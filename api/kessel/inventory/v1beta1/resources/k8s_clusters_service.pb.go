// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        (unknown)
// source: kessel/inventory/v1beta1/resources/k8s_clusters_service.proto

package resources

import (
	_ "buf.build/gen/go/bufbuild/protovalidate/protocolbuffers/go/buf/validate"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type CreateK8SClusterRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The k8s cluster to create in Kessel Asset Inventory
	K8SCluster *K8SCluster `protobuf:"bytes,1,opt,name=k8s_cluster,proto3" json:"k8s_cluster,omitempty"`
	// Will wait for synchronous processes to complete before returning
	WaitForSync bool `protobuf:"varint,2,opt,name=wait_for_sync,json=waitForSync,proto3" json:"wait_for_sync,omitempty"`
}

func (x *CreateK8SClusterRequest) Reset() {
	*x = CreateK8SClusterRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_kessel_inventory_v1beta1_resources_k8s_clusters_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateK8SClusterRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateK8SClusterRequest) ProtoMessage() {}

func (x *CreateK8SClusterRequest) ProtoReflect() protoreflect.Message {
	mi := &file_kessel_inventory_v1beta1_resources_k8s_clusters_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateK8SClusterRequest.ProtoReflect.Descriptor instead.
func (*CreateK8SClusterRequest) Descriptor() ([]byte, []int) {
	return file_kessel_inventory_v1beta1_resources_k8s_clusters_service_proto_rawDescGZIP(), []int{0}
}

func (x *CreateK8SClusterRequest) GetK8SCluster() *K8SCluster {
	if x != nil {
		return x.K8SCluster
	}
	return nil
}

func (x *CreateK8SClusterRequest) GetWaitForSync() bool {
	if x != nil {
		return x.WaitForSync
	}
	return false
}

type CreateK8SClusterResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *CreateK8SClusterResponse) Reset() {
	*x = CreateK8SClusterResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_kessel_inventory_v1beta1_resources_k8s_clusters_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateK8SClusterResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateK8SClusterResponse) ProtoMessage() {}

func (x *CreateK8SClusterResponse) ProtoReflect() protoreflect.Message {
	mi := &file_kessel_inventory_v1beta1_resources_k8s_clusters_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateK8SClusterResponse.ProtoReflect.Descriptor instead.
func (*CreateK8SClusterResponse) Descriptor() ([]byte, []int) {
	return file_kessel_inventory_v1beta1_resources_k8s_clusters_service_proto_rawDescGZIP(), []int{1}
}

type UpdateK8SClusterRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The resource to be updated will be defined by
	// \"<reporter_data.reporter_type>:<reporter_instaance_id>:<reporter_data.local_resource_id>\"
	// from the request body.
	K8SCluster *K8SCluster `protobuf:"bytes,1,opt,name=k8s_cluster,proto3" json:"k8s_cluster,omitempty"`
	// Will wait for synchronous processes to complete before returning
	WaitForSync bool `protobuf:"varint,2,opt,name=wait_for_sync,json=waitForSync,proto3" json:"wait_for_sync,omitempty"`
}

func (x *UpdateK8SClusterRequest) Reset() {
	*x = UpdateK8SClusterRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_kessel_inventory_v1beta1_resources_k8s_clusters_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateK8SClusterRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateK8SClusterRequest) ProtoMessage() {}

func (x *UpdateK8SClusterRequest) ProtoReflect() protoreflect.Message {
	mi := &file_kessel_inventory_v1beta1_resources_k8s_clusters_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateK8SClusterRequest.ProtoReflect.Descriptor instead.
func (*UpdateK8SClusterRequest) Descriptor() ([]byte, []int) {
	return file_kessel_inventory_v1beta1_resources_k8s_clusters_service_proto_rawDescGZIP(), []int{2}
}

func (x *UpdateK8SClusterRequest) GetK8SCluster() *K8SCluster {
	if x != nil {
		return x.K8SCluster
	}
	return nil
}

func (x *UpdateK8SClusterRequest) GetWaitForSync() bool {
	if x != nil {
		return x.WaitForSync
	}
	return false
}

type UpdateK8SClusterResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *UpdateK8SClusterResponse) Reset() {
	*x = UpdateK8SClusterResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_kessel_inventory_v1beta1_resources_k8s_clusters_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateK8SClusterResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateK8SClusterResponse) ProtoMessage() {}

func (x *UpdateK8SClusterResponse) ProtoReflect() protoreflect.Message {
	mi := &file_kessel_inventory_v1beta1_resources_k8s_clusters_service_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateK8SClusterResponse.ProtoReflect.Descriptor instead.
func (*UpdateK8SClusterResponse) Descriptor() ([]byte, []int) {
	return file_kessel_inventory_v1beta1_resources_k8s_clusters_service_proto_rawDescGZIP(), []int{3}
}

type DeleteK8SClusterRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The resource to be deleted will be defined by
	// \"<reporter_data.reporter_type>:<reporter_instaance_id>:<reporter_data.local_resource_id>\"
	// from the request body.
	ReporterData *ReporterData `protobuf:"bytes,1,opt,name=reporter_data,proto3" json:"reporter_data,omitempty"`
}

func (x *DeleteK8SClusterRequest) Reset() {
	*x = DeleteK8SClusterRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_kessel_inventory_v1beta1_resources_k8s_clusters_service_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteK8SClusterRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteK8SClusterRequest) ProtoMessage() {}

func (x *DeleteK8SClusterRequest) ProtoReflect() protoreflect.Message {
	mi := &file_kessel_inventory_v1beta1_resources_k8s_clusters_service_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteK8SClusterRequest.ProtoReflect.Descriptor instead.
func (*DeleteK8SClusterRequest) Descriptor() ([]byte, []int) {
	return file_kessel_inventory_v1beta1_resources_k8s_clusters_service_proto_rawDescGZIP(), []int{4}
}

func (x *DeleteK8SClusterRequest) GetReporterData() *ReporterData {
	if x != nil {
		return x.ReporterData
	}
	return nil
}

type DeleteK8SClusterResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *DeleteK8SClusterResponse) Reset() {
	*x = DeleteK8SClusterResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_kessel_inventory_v1beta1_resources_k8s_clusters_service_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteK8SClusterResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteK8SClusterResponse) ProtoMessage() {}

func (x *DeleteK8SClusterResponse) ProtoReflect() protoreflect.Message {
	mi := &file_kessel_inventory_v1beta1_resources_k8s_clusters_service_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteK8SClusterResponse.ProtoReflect.Descriptor instead.
func (*DeleteK8SClusterResponse) Descriptor() ([]byte, []int) {
	return file_kessel_inventory_v1beta1_resources_k8s_clusters_service_proto_rawDescGZIP(), []int{5}
}

var File_kessel_inventory_v1beta1_resources_k8s_clusters_service_proto protoreflect.FileDescriptor

var file_kessel_inventory_v1beta1_resources_k8s_clusters_service_proto_rawDesc = []byte{
	0x0a, 0x3d, 0x6b, 0x65, 0x73, 0x73, 0x65, 0x6c, 0x2f, 0x69, 0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f,
	0x72, 0x79, 0x2f, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x73, 0x2f, 0x6b, 0x38, 0x73, 0x5f, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72,
	0x73, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x22, 0x6b, 0x65, 0x73, 0x73, 0x65, 0x6c, 0x2e, 0x69, 0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f, 0x72,
	0x79, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x73, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f,
	0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x1b, 0x62, 0x75, 0x66, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f,
	0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x34,
	0x6b, 0x65, 0x73, 0x73, 0x65, 0x6c, 0x2f, 0x69, 0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f, 0x72, 0x79,
	0x2f, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x73, 0x2f, 0x6b, 0x38, 0x73, 0x5f, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x36, 0x6b, 0x65, 0x73, 0x73, 0x65, 0x6c, 0x2f, 0x69, 0x6e, 0x76,
	0x65, 0x6e, 0x74, 0x6f, 0x72, 0x79, 0x2f, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2f, 0x72,
	0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2f, 0x72, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x65,
	0x72, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x8f, 0x01, 0x0a,
	0x17, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4b, 0x38, 0x73, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65,
	0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x50, 0x0a, 0x0b, 0x6b, 0x38, 0x73, 0x5f,
	0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2e, 0x2e,
	0x6b, 0x65, 0x73, 0x73, 0x65, 0x6c, 0x2e, 0x69, 0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f, 0x72, 0x79,
	0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x73, 0x2e, 0x4b, 0x38, 0x73, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x52, 0x0b, 0x6b,
	0x38, 0x73, 0x5f, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x12, 0x22, 0x0a, 0x0d, 0x77, 0x61,
	0x69, 0x74, 0x5f, 0x66, 0x6f, 0x72, 0x5f, 0x73, 0x79, 0x6e, 0x63, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x0b, 0x77, 0x61, 0x69, 0x74, 0x46, 0x6f, 0x72, 0x53, 0x79, 0x6e, 0x63, 0x22, 0x1a,
	0x0a, 0x18, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4b, 0x38, 0x73, 0x43, 0x6c, 0x75, 0x73, 0x74,
	0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x8f, 0x01, 0x0a, 0x17, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x4b, 0x38, 0x73, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x50, 0x0a, 0x0b, 0x6b, 0x38, 0x73, 0x5f, 0x63, 0x6c,
	0x75, 0x73, 0x74, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2e, 0x2e, 0x6b, 0x65,
	0x73, 0x73, 0x65, 0x6c, 0x2e, 0x69, 0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f, 0x72, 0x79, 0x2e, 0x76,
	0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73,
	0x2e, 0x4b, 0x38, 0x73, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x52, 0x0b, 0x6b, 0x38, 0x73,
	0x5f, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x12, 0x22, 0x0a, 0x0d, 0x77, 0x61, 0x69, 0x74,
	0x5f, 0x66, 0x6f, 0x72, 0x5f, 0x73, 0x79, 0x6e, 0x63, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x0b, 0x77, 0x61, 0x69, 0x74, 0x46, 0x6f, 0x72, 0x53, 0x79, 0x6e, 0x63, 0x22, 0x1a, 0x0a, 0x18,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4b, 0x38, 0x73, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x79, 0x0a, 0x17, 0x44, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x4b, 0x38, 0x73, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x5e, 0x0a, 0x0d, 0x72, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x65, 0x72, 0x5f,
	0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x30, 0x2e, 0x6b, 0x65, 0x73,
	0x73, 0x65, 0x6c, 0x2e, 0x69, 0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f, 0x72, 0x79, 0x2e, 0x76, 0x31,
	0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2e,
	0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x65, 0x72, 0x44, 0x61, 0x74, 0x61, 0x42, 0x06, 0xba, 0x48,
	0x03, 0xc8, 0x01, 0x01, 0x52, 0x0d, 0x72, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x65, 0x72, 0x5f, 0x64,
	0x61, 0x74, 0x61, 0x22, 0x1a, 0x0a, 0x18, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x4b, 0x38, 0x73,
	0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x32,
	0xf7, 0x04, 0x0a, 0x17, 0x4b, 0x65, 0x73, 0x73, 0x65, 0x6c, 0x4b, 0x38, 0x73, 0x43, 0x6c, 0x75,
	0x73, 0x74, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0xc7, 0x01, 0x0a, 0x10,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4b, 0x38, 0x73, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72,
	0x12, 0x3b, 0x2e, 0x6b, 0x65, 0x73, 0x73, 0x65, 0x6c, 0x2e, 0x69, 0x6e, 0x76, 0x65, 0x6e, 0x74,
	0x6f, 0x72, 0x79, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x72, 0x65, 0x73, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x73, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4b, 0x38, 0x73, 0x43,
	0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x3c, 0x2e,
	0x6b, 0x65, 0x73, 0x73, 0x65, 0x6c, 0x2e, 0x69, 0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f, 0x72, 0x79,
	0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x73, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4b, 0x38, 0x73, 0x43, 0x6c, 0x75, 0x73,
	0x74, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x38, 0x82, 0xd3, 0xe4,
	0x93, 0x02, 0x32, 0x3a, 0x01, 0x2a, 0x22, 0x2d, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x69, 0x6e, 0x76,
	0x65, 0x6e, 0x74, 0x6f, 0x72, 0x79, 0x2f, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2f, 0x72,
	0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2f, 0x6b, 0x38, 0x73, 0x2d, 0x63, 0x6c, 0x75,
	0x73, 0x74, 0x65, 0x72, 0x73, 0x12, 0xc7, 0x01, 0x0a, 0x10, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x4b, 0x38, 0x73, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x12, 0x3b, 0x2e, 0x6b, 0x65, 0x73,
	0x73, 0x65, 0x6c, 0x2e, 0x69, 0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f, 0x72, 0x79, 0x2e, 0x76, 0x31,
	0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2e,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4b, 0x38, 0x73, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x3c, 0x2e, 0x6b, 0x65, 0x73, 0x73, 0x65, 0x6c,
	0x2e, 0x69, 0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f, 0x72, 0x79, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74,
	0x61, 0x31, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2e, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x4b, 0x38, 0x73, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x38, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x32, 0x3a, 0x01, 0x2a,
	0x1a, 0x2d, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x69, 0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f, 0x72, 0x79,
	0x2f, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x73, 0x2f, 0x6b, 0x38, 0x73, 0x2d, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x73, 0x12,
	0xc7, 0x01, 0x0a, 0x10, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x4b, 0x38, 0x73, 0x43, 0x6c, 0x75,
	0x73, 0x74, 0x65, 0x72, 0x12, 0x3b, 0x2e, 0x6b, 0x65, 0x73, 0x73, 0x65, 0x6c, 0x2e, 0x69, 0x6e,
	0x76, 0x65, 0x6e, 0x74, 0x6f, 0x72, 0x79, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e,
	0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65,
	0x4b, 0x38, 0x73, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x3c, 0x2e, 0x6b, 0x65, 0x73, 0x73, 0x65, 0x6c, 0x2e, 0x69, 0x6e, 0x76, 0x65, 0x6e,
	0x74, 0x6f, 0x72, 0x79, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x72, 0x65, 0x73,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x4b, 0x38, 0x73,
	0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x38, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x32, 0x3a, 0x01, 0x2a, 0x2a, 0x2d, 0x2f, 0x61, 0x70, 0x69,
	0x2f, 0x69, 0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f, 0x72, 0x79, 0x2f, 0x76, 0x31, 0x62, 0x65, 0x74,
	0x61, 0x31, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2f, 0x6b, 0x38, 0x73,
	0x2d, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x73, 0x42, 0x86, 0x01, 0x0a, 0x32, 0x6f, 0x72,
	0x67, 0x2e, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x5f, 0x6b, 0x65, 0x73, 0x73, 0x65, 0x6c,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x69, 0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f, 0x72, 0x79, 0x2e, 0x76,
	0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73,
	0x50, 0x01, 0x5a, 0x4e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x70,
	0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x2d, 0x6b, 0x65, 0x73, 0x73, 0x65, 0x6c, 0x2f, 0x69, 0x6e,
	0x76, 0x65, 0x6e, 0x74, 0x6f, 0x72, 0x79, 0x2d, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x70, 0x69, 0x2f,
	0x6b, 0x65, 0x73, 0x73, 0x65, 0x6c, 0x2f, 0x69, 0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f, 0x72, 0x79,
	0x2f, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_kessel_inventory_v1beta1_resources_k8s_clusters_service_proto_rawDescOnce sync.Once
	file_kessel_inventory_v1beta1_resources_k8s_clusters_service_proto_rawDescData = file_kessel_inventory_v1beta1_resources_k8s_clusters_service_proto_rawDesc
)

func file_kessel_inventory_v1beta1_resources_k8s_clusters_service_proto_rawDescGZIP() []byte {
	file_kessel_inventory_v1beta1_resources_k8s_clusters_service_proto_rawDescOnce.Do(func() {
		file_kessel_inventory_v1beta1_resources_k8s_clusters_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_kessel_inventory_v1beta1_resources_k8s_clusters_service_proto_rawDescData)
	})
	return file_kessel_inventory_v1beta1_resources_k8s_clusters_service_proto_rawDescData
}

var file_kessel_inventory_v1beta1_resources_k8s_clusters_service_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_kessel_inventory_v1beta1_resources_k8s_clusters_service_proto_goTypes = []any{
	(*CreateK8SClusterRequest)(nil),  // 0: kessel.inventory.v1beta1.resources.CreateK8sClusterRequest
	(*CreateK8SClusterResponse)(nil), // 1: kessel.inventory.v1beta1.resources.CreateK8sClusterResponse
	(*UpdateK8SClusterRequest)(nil),  // 2: kessel.inventory.v1beta1.resources.UpdateK8sClusterRequest
	(*UpdateK8SClusterResponse)(nil), // 3: kessel.inventory.v1beta1.resources.UpdateK8sClusterResponse
	(*DeleteK8SClusterRequest)(nil),  // 4: kessel.inventory.v1beta1.resources.DeleteK8sClusterRequest
	(*DeleteK8SClusterResponse)(nil), // 5: kessel.inventory.v1beta1.resources.DeleteK8sClusterResponse
	(*K8SCluster)(nil),               // 6: kessel.inventory.v1beta1.resources.K8sCluster
	(*ReporterData)(nil),             // 7: kessel.inventory.v1beta1.resources.ReporterData
}
var file_kessel_inventory_v1beta1_resources_k8s_clusters_service_proto_depIdxs = []int32{
	6, // 0: kessel.inventory.v1beta1.resources.CreateK8sClusterRequest.k8s_cluster:type_name -> kessel.inventory.v1beta1.resources.K8sCluster
	6, // 1: kessel.inventory.v1beta1.resources.UpdateK8sClusterRequest.k8s_cluster:type_name -> kessel.inventory.v1beta1.resources.K8sCluster
	7, // 2: kessel.inventory.v1beta1.resources.DeleteK8sClusterRequest.reporter_data:type_name -> kessel.inventory.v1beta1.resources.ReporterData
	0, // 3: kessel.inventory.v1beta1.resources.KesselK8sClusterService.CreateK8sCluster:input_type -> kessel.inventory.v1beta1.resources.CreateK8sClusterRequest
	2, // 4: kessel.inventory.v1beta1.resources.KesselK8sClusterService.UpdateK8sCluster:input_type -> kessel.inventory.v1beta1.resources.UpdateK8sClusterRequest
	4, // 5: kessel.inventory.v1beta1.resources.KesselK8sClusterService.DeleteK8sCluster:input_type -> kessel.inventory.v1beta1.resources.DeleteK8sClusterRequest
	1, // 6: kessel.inventory.v1beta1.resources.KesselK8sClusterService.CreateK8sCluster:output_type -> kessel.inventory.v1beta1.resources.CreateK8sClusterResponse
	3, // 7: kessel.inventory.v1beta1.resources.KesselK8sClusterService.UpdateK8sCluster:output_type -> kessel.inventory.v1beta1.resources.UpdateK8sClusterResponse
	5, // 8: kessel.inventory.v1beta1.resources.KesselK8sClusterService.DeleteK8sCluster:output_type -> kessel.inventory.v1beta1.resources.DeleteK8sClusterResponse
	6, // [6:9] is the sub-list for method output_type
	3, // [3:6] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_kessel_inventory_v1beta1_resources_k8s_clusters_service_proto_init() }
func file_kessel_inventory_v1beta1_resources_k8s_clusters_service_proto_init() {
	if File_kessel_inventory_v1beta1_resources_k8s_clusters_service_proto != nil {
		return
	}
	file_kessel_inventory_v1beta1_resources_k8s_cluster_proto_init()
	file_kessel_inventory_v1beta1_resources_reporter_data_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_kessel_inventory_v1beta1_resources_k8s_clusters_service_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*CreateK8SClusterRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_kessel_inventory_v1beta1_resources_k8s_clusters_service_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*CreateK8SClusterResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_kessel_inventory_v1beta1_resources_k8s_clusters_service_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*UpdateK8SClusterRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_kessel_inventory_v1beta1_resources_k8s_clusters_service_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*UpdateK8SClusterResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_kessel_inventory_v1beta1_resources_k8s_clusters_service_proto_msgTypes[4].Exporter = func(v any, i int) any {
			switch v := v.(*DeleteK8SClusterRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_kessel_inventory_v1beta1_resources_k8s_clusters_service_proto_msgTypes[5].Exporter = func(v any, i int) any {
			switch v := v.(*DeleteK8SClusterResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_kessel_inventory_v1beta1_resources_k8s_clusters_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_kessel_inventory_v1beta1_resources_k8s_clusters_service_proto_goTypes,
		DependencyIndexes: file_kessel_inventory_v1beta1_resources_k8s_clusters_service_proto_depIdxs,
		MessageInfos:      file_kessel_inventory_v1beta1_resources_k8s_clusters_service_proto_msgTypes,
	}.Build()
	File_kessel_inventory_v1beta1_resources_k8s_clusters_service_proto = out.File
	file_kessel_inventory_v1beta1_resources_k8s_clusters_service_proto_rawDesc = nil
	file_kessel_inventory_v1beta1_resources_k8s_clusters_service_proto_goTypes = nil
	file_kessel_inventory_v1beta1_resources_k8s_clusters_service_proto_depIdxs = nil
}
