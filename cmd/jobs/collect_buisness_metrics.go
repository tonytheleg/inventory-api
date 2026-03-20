package jobs

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/metrics"
	"github.com/project-kessel/inventory-api/cmd/common"
	"github.com/project-kessel/inventory-api/internal/storage"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/spf13/cobra"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/exporters/prometheus"
	"go.opentelemetry.io/otel/metric"
	sdkmetric "go.opentelemetry.io/otel/sdk/metric"
	"go.opentelemetry.io/otel/sdk/resource"
	semconv "go.opentelemetry.io/otel/semconv/v1.4.0"
	"gorm.io/gorm"
)

const (
	prefix       = "kessel_inventory_"
	queryTimeout = 5 * time.Minute
)

func NewCollectBizMetricsJobCommand(storageOptions *storage.Options, loggerOptions common.LoggerOptions, ctx context.Context) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "collect-buisness-metrics",
		Short: "Captures business-related metrics and exposes them via metrics endpoint",
		Long:  "Captures business-related metrics and exposes them via metrics endpoint",
		RunE: func(cmd *cobra.Command, args []string) error {
			bc := &bizMetricsCollector{}

			meterProvider, err := newMeterProvider()
			if err != nil {
				return err
			}
			meter, err := newMeter(meterProvider)
			if err != nil {
				return err
			}
			err = bc.New(meter)
			if err != nil {
				return fmt.Errorf("failed to initialize business metricscollector: %w", err)
			}
			return startBusinessMetricsCollector(storageOptions, loggerOptions, ctx, bc)
		},
	}
	return cmd
}

type workspaceResourceCount struct {
	ResourceType string
	Count        float64
}

type reporterResourceCount struct {
	ResourceType       string
	ReporterType       string
	ReporterInstanceID string
	Count              int64
}

func newMeterProvider() (metric.MeterProvider, error) {
	exporter, err := prometheus.New()
	if err != nil {
		return nil, fmt.Errorf("failed to setup exporter for meter provider: %w", err)
	}

	provider := sdkmetric.NewMeterProvider(
		sdkmetric.WithResource(
			resource.NewWithAttributes(
				semconv.SchemaURL,
				semconv.ServiceNameKey.String("biz-metricscolector"),
			),
		),
		sdkmetric.WithReader(exporter),
		sdkmetric.WithView(
			metrics.DefaultSecondsHistogramView(metrics.DefaultServerSecondsHistogramName),
		),
	)
	otel.SetMeterProvider(provider)
	return provider, nil
}

func newMeter(provider metric.MeterProvider) (metric.Meter, error) {
	return provider.Meter("biz-metricscollector"), nil
}

// bizMetricsCollector queries the database for specific metrics and publishes them to Prometheus
type bizMetricsCollector struct {
	ResourcesPerWorkspace metric.Float64Histogram
	ResourceCount         metric.Int64Gauge
}

func (b *bizMetricsCollector) New(meter metric.Meter) error {
	var err error
	b.ResourcesPerWorkspace, err = meter.Float64Histogram(
		prefix+"resources_per_workspace",
		metric.WithDescription("Distribution of resource counts per workspace, grouped by resource type"),
		metric.WithExplicitBucketBoundaries(1, 2, 5, 10, 20, 50, 100, 200, 500, 1000),
	)
	if err != nil {
		return fmt.Errorf("initiating meter failed: %w", err)
	}

	b.ResourceCount, err = meter.Int64Gauge(
		prefix+"resource_count",
		metric.WithDescription("Number of resources grouped by resource type, reporter name, and reporter ID"),
	)
	if err != nil {
		return fmt.Errorf("initiating meter failed: %w", err)
	}

	return nil
}

func startBusinessMetricsCollector(storageOptions *storage.Options, loggerOptions common.LoggerOptions, ctx context.Context, bc *bizMetricsCollector) error {
	_, logger := common.InitLogger(common.GetLogLevel(), loggerOptions)
	logHelper := log.NewHelper(log.With(logger, "job", "collect_business_metrics"))

	storageConfig := storage.NewConfig(storageOptions).Complete()
	db, err := storage.New(storageConfig, logHelper)
	if err != nil {
		return err
	}

	go func() {
		collectBusinessMetrics(ctx, db, bc, logHelper)

		ticker := time.NewTicker(24 * time.Hour)
		defer ticker.Stop()

		for {
			select {
			case <-ctx.Done():
				return
			case <-ticker.C:
				collectBusinessMetrics(ctx, db, bc, logHelper)
			}
		}
	}()

	http.Handle("/metrics", promhttp.Handler())
	err = http.ListenAndServe(":9090", nil)
	if err != nil {
		fmt.Printf("error serving metrics: %v", err)
		return err
	}

	return nil
}

func collectBusinessMetrics(ctx context.Context, db *gorm.DB, bc *bizMetricsCollector, logger *log.Helper) {
	queryCtx, cancel := context.WithTimeout(ctx, queryTimeout)
	defer cancel()

	collectResourcesPerWorkspace(queryCtx, db, bc, logger)
	collectResourceCount(queryCtx, db, bc, logger)
}

func collectResourcesPerWorkspace(ctx context.Context, db *gorm.DB, bc *bizMetricsCollector, logger *log.Helper) {
	var results []workspaceResourceCount

	// For each resource, get the latest common representation (max version),
	// extract workspace_id from the JSONB data, and count resources per
	// workspace_id and resource_type. The result is one row per workspace per
	// resource_type with the count of resources in that workspace.
	err := db.WithContext(ctx).Raw(`
		SELECT r.type AS resource_type,
		       COUNT(*)::float AS count
		FROM resource r
		JOIN LATERAL (
			SELECT data
			FROM common_representations
			WHERE resource_id = r.id
			ORDER BY version DESC
			LIMIT 1
		) cr ON true
		WHERE cr.data->>'workspace_id' IS NOT NULL
		GROUP BY r.type, cr.data->>'workspace_id'
	`).Scan(&results).Error

	if err != nil {
		logger.Errorf("failed to collect resources_per_workspace metric: %v", err)
		return
	}

	for _, r := range results {
		bc.ResourcesPerWorkspace.Record(ctx, r.Count,
			metric.WithAttributes(
				attribute.String("resource_type", r.ResourceType),
			),
		)
	}

	logger.Infof("recorded resources_per_workspace metric for %d (workspace, resource_type) groups", len(results))
}

func collectResourceCount(ctx context.Context, db *gorm.DB, bc *bizMetricsCollector, logger *log.Helper) {
	var results []reporterResourceCount

	err := db.WithContext(ctx).Raw(`
		SELECT resource_type,
		       reporter_type,
		       reporter_instance_id,
		       COUNT(*) AS count
		FROM reporter_resources
		WHERE NOT tombstone
		GROUP BY resource_type, reporter_type, reporter_instance_id
	`).Scan(&results).Error

	if err != nil {
		logger.Errorf("failed to collect resource_count metric: %v", err)
		return
	}

	for _, r := range results {
		bc.ResourceCount.Record(ctx, r.Count,
			metric.WithAttributes(
				attribute.String("resource_type", r.ResourceType),
				attribute.String("reporter_name", r.ReporterType),
				attribute.String("reporter_id", r.ReporterInstanceID),
			),
		)
	}
	logger.Infof("recorded resource_count metric for %d (resource_type, reporter_name, reporter_id) groups", len(results))
}
