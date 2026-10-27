package model_test

import (
	"context"
	"testing"

	"github.com/project-kessel/inventory-api/internal/biz/model"
	"github.com/project-kessel/inventory-api/internal/data"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// TestUnifiedReporterRelationsRegression guards against the unified schema
// repository registering reporter schema instances without their relations.
// The reporter-aware SchemaService fetches reporter schemas and filters
// representations to reporter-only data before calculating tuples; if the
// reporter instance carries no relations (as happened when the #1503 combined
// wiring moved them to the common instance only), every reporter-side
// relationship produces zero tuples and replication silently stops.
func TestUnifiedReporterRelationsRegression(t *testing.T) {
	ctx := context.Background()

	repository, err := data.NewInMemorySchemaRepositoryFromUnifiedYAMLDir(ctx, "../../../development/unified-schema/resources")
	require.NoError(t, err)

	schemaService := model.NewSchemaService(repository, nil)

	t.Run("billing account services", func(t *testing.T) {
		key := unifiedTestKey(t, "3055cc32-b398-11f1-bdd0-6ae246604903", "billing_account", "features", "1b9aa216-b39a-11f1-8e5d-6ae246604903")

		reporterVersion := model.NewVersion(0)
		current, err := model.NewRepresentations(
			nil, nil,
			model.Representation(map[string]interface{}{
				"services": []interface{}{"4f829a2c-b398-11f1-ba8e-6ae246604903", "54c3d5a0-b398-11f1-9332-6ae246604903"},
			}),
			&reporterVersion,
		)
		require.NoError(t, err)

		tuples, err := schemaService.CalculateTuplesForResource(ctx, current, nil, key)
		require.NoError(t, err)
		assert.True(t, tuples.HasTuplesToCreate(), "expected services tuples to be created for the billing account")
	})

	t.Run("workspace direct service preferences", func(t *testing.T) {
		key := unifiedTestKey(t, "9554d7a6-b396-11f1-8fa2-6ae246604903", "workspace", "features", "3ba27996-b398-11f1-bf17-6ae246604903")

		reporterVersion := model.NewVersion(0)
		current, err := model.NewRepresentations(
			nil, nil,
			model.Representation(map[string]interface{}{
				"direct_service_preferences": []interface{}{"4f829a2c-b398-11f1-ba8e-6ae246604903"},
			}),
			&reporterVersion,
		)
		require.NoError(t, err)

		tuples, err := schemaService.CalculateTuplesForResource(ctx, current, nil, key)
		require.NoError(t, err)
		assert.True(t, tuples.HasTuplesToCreate(), "expected direct_service_preferences tuples to be created for the workspace")
	})

	t.Run("workspace direct billing account", func(t *testing.T) {
		key := unifiedTestKey(t, "31a6d1a6-b395-11f1-adbd-6ae246604903", "workspace", "features", "3ba27996-b398-11f1-bf17-6ae246604903")

		reporterVersion := model.NewVersion(0)
		current, err := model.NewRepresentations(
			nil, nil,
			model.Representation(map[string]interface{}{
				"direct_billing_account": "3055cc32-b398-11f1-bdd0-6ae246604903",
			}),
			&reporterVersion,
		)
		require.NoError(t, err)

		tuples, err := schemaService.CalculateTuplesForResource(ctx, current, nil, key)
		require.NoError(t, err)
		assert.True(t, tuples.HasTuplesToCreate(), "expected direct_billing_account tuple to be created for the workspace")
	})
}

// unifiedTestKey builds a ReporterResourceKey for the unified schema tests.
func unifiedTestKey(t *testing.T, localResourceId, resourceType, reporterType, reporterInstanceId string) model.ReporterResourceKey {
	t.Helper()
	localId, err := model.NewLocalResourceId(localResourceId)
	require.NoError(t, err)
	resType, err := model.NewResourceType(resourceType)
	require.NoError(t, err)
	repType, err := model.NewReporterType(reporterType)
	require.NoError(t, err)
	instanceId, err := model.NewReporterInstanceId(reporterInstanceId)
	require.NoError(t, err)
	key, err := model.NewReporterResourceKey(localId, resType, repType, instanceId)
	require.NoError(t, err)
	return key
}
