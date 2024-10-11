package common

import (
	"encoding/json"
	pbrelation "github.com/project-kessel/inventory-api/api/kessel/inventory/v1beta1/relationships"
	pbresource "github.com/project-kessel/inventory-api/api/kessel/inventory/v1beta1/resources"
	"github.com/project-kessel/inventory-api/internal/biz/model"
)

func ResourceFromPb(resourceType, reporterId string, resourceData model.JsonObject, metadata *pbresource.Metadata, reporter *pbresource.ReporterData) *model.Resource {
	return &model.Resource{
		ID:           0,
		ResourceData: resourceData,
		ResourceType: resourceType,
		WorkspaceId:  metadata.Workspace,
		Reporter: model.ResourceReporter{
			Reporter: model.Reporter{
				ReporterId:      reporterId,
				ReporterType:    reporter.ReporterType.String(),
				ReporterVersion: reporter.ReporterVersion,
			},
			LocalResourceId: reporter.LocalResourceId,
		},
		ConsoleHref: reporter.ConsoleHref,
		ApiHref:     reporter.ApiHref,
		Labels:      labelsFromPb(metadata.Labels),
	}
}

func ToJsonObject(in interface{}) (model.JsonObject, error) {
	if in == nil {
		return nil, nil
	}

	bytes, err := json.Marshal(in)
	if err != nil {
		return nil, err
	}

	resourceData := model.JsonObject{}
	err = json.Unmarshal(bytes, &resourceData)
	if err != nil {
		return nil, err
	}

	return resourceData, err
}

func labelsFromPb(pbLabels []*pbresource.ResourceLabel) model.Labels {
	labels := model.Labels{}
	for _, pbLabel := range pbLabels {
		labels = append(labels, model.Label{
			Key:   pbLabel.Key,
			Value: pbLabel.Value,
		})
	}
	return labels
}

func RelationshipFromPb(relationshipType, reporterId string, relationshipData model.JsonObject, metadata *pbrelation.Metadata, reporter *pbrelation.ReporterData) *model.Relationship {
	return &model.Relationship{
		ID:               0,
		RelationshipData: relationshipData,
		RelationshipType: relationshipType,
		SubjectId:        0,
		ObjectId:         0,
		Reporter: model.RelationshipReporter{
			Reporter: model.Reporter{
				ReporterId:      reporterId,
				ReporterType:    reporter.ReporterType.String(),
				ReporterVersion: reporter.ReporterVersion,
			},
			SubjectLocalResourceId: reporter.SubjectLocalResourceId,
			ObjectLocalResourceId:  reporter.ObjectLocalResourceId,
		},
	}
}
