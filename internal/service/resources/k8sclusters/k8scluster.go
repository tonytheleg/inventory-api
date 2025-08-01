package k8sclusters

import (
	"context"

	"github.com/project-kessel/inventory-api/internal/biz/usecase/resources"

	pb "github.com/project-kessel/inventory-api/api/kessel/inventory/v1beta1/resources"
	authnapi "github.com/project-kessel/inventory-api/internal/authn/api"
	"github.com/project-kessel/inventory-api/internal/biz/model_legacy"
	"github.com/project-kessel/inventory-api/internal/middleware"
	conv "github.com/project-kessel/inventory-api/internal/service/common"
)

const (
	ResourceType = "k8s_cluster"
)

// K8sClustersService handles requests for k8s clusters
type K8sClustersService struct {
	pb.UnimplementedKesselK8SClusterServiceServer

	Ctl *resources.Usecase
}

// NewKesselK8SClusterServiceV1beta1 creates a new K8sClusterService to handle requests for k8s clusters
func NewKesselK8SClusterServiceV1beta1(c *resources.Usecase) *K8sClustersService {
	return &K8sClustersService{
		Ctl: c,
	}
}

func (c *K8sClustersService) CreateK8SCluster(ctx context.Context, r *pb.CreateK8SClusterRequest) (*pb.CreateK8SClusterResponse, error) {
	identity, err := middleware.GetIdentity(ctx)
	if err != nil {
		return nil, err
	}

	if k, err := k8sClusterFromCreateRequest(r, identity); err == nil {
		if resp, err := c.Ctl.Create(ctx, k); err == nil { //nolint:staticcheck
			return createResponseFromK8sCluster(resp), nil
		} else {
			return nil, err
		}
	} else {
		return nil, err
	}
}

func (c *K8sClustersService) UpdateK8SCluster(ctx context.Context, r *pb.UpdateK8SClusterRequest) (*pb.UpdateK8SClusterResponse, error) {
	identity, err := middleware.GetIdentity(ctx)
	if err != nil {
		return nil, err
	}

	if k, err := k8sClusterFromUpdateRequest(r, identity); err == nil {
		// Todo: Update to use the right ID
		if resp, err := c.Ctl.Update(ctx, k, model_legacy.ReporterResourceIdFromResource(k)); err == nil { //nolint:staticcheck
			return updateResponseFromK8sCluster(resp), nil
		} else {
			return nil, err
		}
	} else {
		return nil, err
	}
}

func (c *K8sClustersService) DeleteK8SCluster(ctx context.Context, r *pb.DeleteK8SClusterRequest) (*pb.DeleteK8SClusterResponse, error) {
	identity, err := middleware.GetIdentity(ctx)
	if err != nil {
		return nil, err
	}

	if resourceId, err := fromDeleteRequest(r, identity); err == nil {
		if err := c.Ctl.Delete(ctx, resourceId); err == nil {
			return toDeleteResponse(), nil
		} else {
			return nil, err
		}
	} else {
		return nil, err
	}
}

func k8sClusterFromCreateRequest(r *pb.CreateK8SClusterRequest, identity *authnapi.Identity) (*model_legacy.Resource, error) {
	resourceData, err := conv.ToJsonObject(r.K8SCluster.ResourceData)
	if err != nil {
		return nil, err
	}

	return conv.ResourceFromPbv1beta1(ResourceType, identity.Principal, resourceData, r.K8SCluster.Metadata, r.K8SCluster.ReporterData), nil
}

func createResponseFromK8sCluster(c *model_legacy.Resource) *pb.CreateK8SClusterResponse {
	return &pb.CreateK8SClusterResponse{}
}

func k8sClusterFromUpdateRequest(r *pb.UpdateK8SClusterRequest, identity *authnapi.Identity) (*model_legacy.Resource, error) {
	resourceData, err := conv.ToJsonObject(r.K8SCluster.ResourceData)
	if err != nil {
		return nil, err
	}

	return conv.ResourceFromPbv1beta1(ResourceType, identity.Principal, resourceData, r.K8SCluster.Metadata, r.K8SCluster.ReporterData), nil
}

func updateResponseFromK8sCluster(c *model_legacy.Resource) *pb.UpdateK8SClusterResponse {
	return &pb.UpdateK8SClusterResponse{}
}

func fromDeleteRequest(r *pb.DeleteK8SClusterRequest, identity *authnapi.Identity) (model_legacy.ReporterResourceId, error) {
	return conv.ReporterResourceIdFromPb(ResourceType, identity.Principal, r.ReporterData), nil
}

func toDeleteResponse() *pb.DeleteK8SClusterResponse {
	return &pb.DeleteK8SClusterResponse{}
}
