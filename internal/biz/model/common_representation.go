package model

import (
	"fmt"
	"strings"
	"time"

	"github.com/google/uuid"
)

type CommonRepresentation struct {
	Representation
	resourceId     ResourceId
	version        Version
	reporter       ReporterId
	toBeDetermined ToBeDetermined
}

func NewCommonRepresentation(
	resourceId ResourceId,
	data Representation,
	version Version,
	reporterType ReporterType,
	reporterInstanceId ReporterInstanceId,
	toBeDetermined ToBeDetermined,
) (CommonRepresentation, error) {
	if resourceId.UUID() == uuid.Nil {
		return CommonRepresentation{}, fmt.Errorf("%w: ResourceId", ErrInvalidUUID)
	}

	if strings.TrimSpace(string(reporterType)) == "" {
		return CommonRepresentation{}, fmt.Errorf("%w: ReporterType", ErrEmpty)
	}

	if strings.TrimSpace(string(reporterInstanceId)) == "" {
		return CommonRepresentation{}, fmt.Errorf("%w: ReporterInstanceId", ErrEmpty)
	}
	if len(data) == 0 {
		return CommonRepresentation{}, fmt.Errorf("%w: CommonRepresentation data", ErrInvalidData)
	}

	if data.Data() == nil {
		return CommonRepresentation{}, fmt.Errorf("CommonRepresentation requires non-empty data")
	}

	reporter := NewReporterId(reporterType, reporterInstanceId)

	return CommonRepresentation{
		Representation: data,
		resourceId:     resourceId,
		version:        version,
		reporter:       reporter,
		toBeDetermined: toBeDetermined,
	}, nil
}

func (cr CommonRepresentation) Serialize() CommonRepresentationSnapshot {
	reporterType, reporterInstanceId := cr.reporter.Serialize()

	// Create representation snapshot
	representationSnapshot := RepresentationSnapshot{
		Data: cr.Data(),
	}

	// Create CommonRepresentation snapshot - direct initialization without validation
	return CommonRepresentationSnapshot{
		Representation:             representationSnapshot,
		ResourceId:                 cr.resourceId.UUID(),
		Version:                    cr.version.Serialize(),
		ReportedByReporterType:     reporterType,
		ReportedByReporterInstance: reporterInstanceId,
		ToBeDetermined:             cr.toBeDetermined.Serialize(),
		CreatedAt:                  time.Now(), // TODO: Add proper timestamp from domain entity if available
	}
}

func DeserializeCommonRepresentation(snapshot *CommonRepresentationSnapshot) CommonRepresentation {
	// Create domain tiny types directly from snapshot values - no validation
	resourceId := ResourceId(snapshot.ResourceId)
	representation := Representation(snapshot.Representation.Data)
	version := DeserializeVersion(snapshot.Version)
	reporterType := ReporterType(snapshot.ReportedByReporterType)
	reporterInstanceId := ReporterInstanceId(snapshot.ReportedByReporterInstance)
	toBeDetermined := ToBeDetermined(snapshot.ToBeDetermined)

	// Create reporter ID
	reporterId := ReporterId{
		reporterType:       reporterType,
		reporterInstanceId: reporterInstanceId,
	}

	return CommonRepresentation{
		Representation: representation,
		resourceId:     resourceId,
		version:        version,
		reporter:       reporterId,
		toBeDetermined: toBeDetermined,
	}
}

// CreateSnapshot creates a snapshot of the CommonRepresentation
func (cr CommonRepresentation) CreateSnapshot() (CommonRepresentationSnapshot, error) {
	return cr.Serialize(), nil
}
