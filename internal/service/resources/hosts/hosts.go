package hosts

import (
	"context"
	pb "github.com/project-kessel/inventory-api/api/kessel/inventory/v1beta1/resources"
	authnapi "github.com/project-kessel/inventory-api/internal/authn/api"
	"github.com/project-kessel/inventory-api/internal/biz/model"
	"github.com/project-kessel/inventory-api/internal/biz/resources"
	"github.com/project-kessel/inventory-api/internal/middleware"
	conv "github.com/project-kessel/inventory-api/internal/service/common"
	"golang.org/x/exp/slices"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/types/known/fieldmaskpb"
)

const (
	ResourceType = "rhel_host"
)

var (
	PatchFieldMaskPaths = []string{
		"rhel_host.metadata.workspace_id",
		"rhel_host.reporter_data.api_href",
		"rhel_host.reporter_data.console_href",
	}
)

// HostsService handles requests for Rhel hosts
type HostsService struct {
	pb.UnimplementedKesselRhelHostServiceServer

	Ctl *resources.Usecase
}

// New creates a new HostsService to handle requests for Rhel hosts
func New(c *resources.Usecase) *HostsService {
	return &HostsService{
		Ctl: c,
	}
}

func (c *HostsService) CreateRhelHost(ctx context.Context, r *pb.CreateRhelHostRequest) (*pb.CreateRhelHostResponse, error) {
	identity, err := middleware.GetIdentity(ctx)
	if err != nil {
		return nil, err
	}

	if h, err := hostFromCreateRequest(r, identity); err == nil {
		if resp, err := c.Ctl.Create(ctx, h); err == nil {
			return createResponseFromHost(resp), nil
		} else {
			return nil, err
		}
	} else {
		return nil, err
	}
}

func (c *HostsService) UpdateRhelHost(ctx context.Context, r *pb.UpdateRhelHostRequest) (*pb.UpdateRhelHostResponse, error) {
	identity, err := middleware.GetIdentity(ctx)
	if err != nil {
		return nil, err
	}

	if h, err := hostFromUpdateRequest(r, identity); err == nil {
		if resp, err := c.Ctl.Update(ctx, h, model.ReporterResourceIdFromResource(h)); err == nil {
			return updateResponseFromHost(resp), nil
		} else {
			return nil, err
		}
	} else {
		return nil, err
	}
}

func (c *HostsService) PatchRhelHost(ctx context.Context, r *pb.PatchRhelHostRequest) (*pb.PatchRhelHostResponse, error) {
	identity, err := middleware.GetIdentity(ctx)
	if err != nil {
		return nil, err
	}

	r.UpdateMask, err = fieldmaskpb.New(r, PatchFieldMaskPaths...)
	if err != nil {
		return nil, err
	}

	// The idea of filter is to take the request and clear any values set that are not defined in the mask
	Filter(r, r.UpdateMask)

	h, err := hostFromPatchRequest(r, identity)
	if err != nil {
		return nil, err
	}
	resp, err := c.Ctl.Patch(ctx, h, model.ReporterResourceIdFromResource(h))
	if err != nil {
		return nil, err
	}
	return patchResponseFromHost(resp), nil
}

func (c *HostsService) DeleteRhelHost(ctx context.Context, r *pb.DeleteRhelHostRequest) (*pb.DeleteRhelHostResponse, error) {
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

func hostFromCreateRequest(r *pb.CreateRhelHostRequest, identity *authnapi.Identity) (*model.Resource, error) {
	return conv.ResourceFromPb(ResourceType, identity.Principal, nil, r.RhelHost.Metadata, r.RhelHost.ReporterData), nil
}

func createResponseFromHost(resource *model.Resource) *pb.CreateRhelHostResponse {
	return &pb.CreateRhelHostResponse{}
}

func hostFromUpdateRequest(r *pb.UpdateRhelHostRequest, identity *authnapi.Identity) (*model.Resource, error) {
	return conv.ResourceFromPb(ResourceType, identity.Principal, nil, r.RhelHost.Metadata, r.RhelHost.ReporterData), nil
}

func updateResponseFromHost(resource *model.Resource) *pb.UpdateRhelHostResponse {
	return &pb.UpdateRhelHostResponse{}
}

func hostFromPatchRequest(r *pb.PatchRhelHostRequest, identity *authnapi.Identity) (*model.Resource, error) {
	return conv.ResourceFromPb(ResourceType, identity.Principal, nil, r.RhelHost.Metadata, r.RhelHost.ReporterData), nil
}

func patchResponseFromHost(resource *model.Resource) *pb.PatchRhelHostResponse {
	return &pb.PatchRhelHostResponse{}
}

func fromDeleteRequest(r *pb.DeleteRhelHostRequest, identity *authnapi.Identity) (model.ReporterResourceId, error) {
	return conv.ReporterResourceIdFromPb(ResourceType, identity.Principal, r.ReporterData), nil
}

func toDeleteResponse() *pb.DeleteRhelHostResponse {
	return &pb.DeleteRhelHostResponse{}
}

// Filter removes any values set in the request not permitted by the field mask
// currently this doesnt work as expected because its using the top object value "rhel_host" and not nested values
// like rhel_host.metadata.workspace_id -- it needs to further parse into fields to ensure it doesnt remove the expected values
func Filter(msg proto.Message, mask *fieldmaskpb.FieldMask) {
	if mask == nil || len(mask.Paths) == 0 {
		return
	}
	rft := msg.ProtoReflect()
	rft.Range(func(fd protoreflect.FieldDescriptor, _ protoreflect.Value) bool {
		if !slices.Contains(mask.Paths, string(fd.Name())) {
			rft.Clear(fd)
		}
		return true
	})
}
