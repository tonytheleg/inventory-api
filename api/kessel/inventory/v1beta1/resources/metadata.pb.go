// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        (unknown)
// source: kessel/inventory/v1beta1/resources/metadata.proto

package resources

import (
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Metadata struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Kessel Asset Inventory generated identifier.
	Id string `protobuf:"bytes,3355,opt,name=id,proto3" json:"id,omitempty"`
	// The type of the Resource
	ResourceType string `protobuf:"bytes,442752204,opt,name=resource_type,proto3" json:"resource_type,omitempty"`
	// Date and time when the inventory item was first reported.
	CreatedAt *timestamppb.Timestamp `protobuf:"bytes,3400,opt,name=created_at,proto3" json:"created_at,omitempty"`
	// Date and time when the inventory item was last updated.
	UpdatedAt *timestamppb.Timestamp `protobuf:"bytes,3401,opt,name=updated_at,proto3" json:"updated_at,omitempty"`
	// Date and time when the inventory item was deleted.
	DeletedAt *timestamppb.Timestamp `protobuf:"bytes,3402,opt,name=deleted_at,proto3" json:"deleted_at,omitempty"`
	// The org id in which this resource is a member for access control.  A
	// resource can only be a member of one org.
	OrgId string `protobuf:"bytes,3500,opt,name=org_id,proto3" json:"org_id,omitempty"`
	// The workspace id in which this resource is a member for access control.  A
	// resource can only be a member of one workspace.
	WorkspaceId string           `protobuf:"bytes,3501,opt,name=workspace_id,proto3" json:"workspace_id,omitempty"`
	Labels      []*ResourceLabel `protobuf:"bytes,3552281,rep,name=labels,proto3" json:"labels,omitempty"`
}

func (x *Metadata) Reset() {
	*x = Metadata{}
	if protoimpl.UnsafeEnabled {
		mi := &file_kessel_inventory_v1beta1_resources_metadata_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Metadata) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Metadata) ProtoMessage() {}

func (x *Metadata) ProtoReflect() protoreflect.Message {
	mi := &file_kessel_inventory_v1beta1_resources_metadata_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Metadata.ProtoReflect.Descriptor instead.
func (*Metadata) Descriptor() ([]byte, []int) {
	return file_kessel_inventory_v1beta1_resources_metadata_proto_rawDescGZIP(), []int{0}
}

func (x *Metadata) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Metadata) GetResourceType() string {
	if x != nil {
		return x.ResourceType
	}
	return ""
}

func (x *Metadata) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *Metadata) GetUpdatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdatedAt
	}
	return nil
}

func (x *Metadata) GetDeletedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.DeletedAt
	}
	return nil
}

func (x *Metadata) GetOrgId() string {
	if x != nil {
		return x.OrgId
	}
	return ""
}

func (x *Metadata) GetWorkspaceId() string {
	if x != nil {
		return x.WorkspaceId
	}
	return ""
}

func (x *Metadata) GetLabels() []*ResourceLabel {
	if x != nil {
		return x.Labels
	}
	return nil
}

var File_kessel_inventory_v1beta1_resources_metadata_proto protoreflect.FileDescriptor

var file_kessel_inventory_v1beta1_resources_metadata_proto_rawDesc = []byte{
	0x0a, 0x31, 0x6b, 0x65, 0x73, 0x73, 0x65, 0x6c, 0x2f, 0x69, 0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f,
	0x72, 0x79, 0x2f, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x73, 0x2f, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x22, 0x6b, 0x65, 0x73, 0x73, 0x65, 0x6c, 0x2e, 0x69, 0x6e, 0x76, 0x65,
	0x6e, 0x74, 0x6f, 0x72, 0x79, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x72, 0x65,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f,
	0x61, 0x70, 0x69, 0x2f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x62, 0x65, 0x68, 0x61, 0x76, 0x69,
	0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x37, 0x6b, 0x65, 0x73, 0x73, 0x65,
	0x6c, 0x2f, 0x69, 0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f, 0x72, 0x79, 0x2f, 0x76, 0x31, 0x62, 0x65,
	0x74, 0x61, 0x31, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2f, 0x72, 0x65,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0xa1, 0x03, 0x0a, 0x08, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x12,
	0x14, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x9b, 0x1a, 0x20, 0x01, 0x28, 0x09, 0x42, 0x03, 0xe0, 0x41,
	0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x2d, 0x0a, 0x0d, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0xcc, 0xb9, 0x8f, 0xd3, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x42, 0x03, 0xe0, 0x41, 0x03, 0x52, 0x0d, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f,
	0x74, 0x79, 0x70, 0x65, 0x12, 0x40, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f,
	0x61, 0x74, 0x18, 0xc8, 0x1a, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x42, 0x03, 0xe0, 0x41, 0x03, 0x52, 0x0a, 0x63, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x12, 0x40, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x64, 0x5f, 0x61, 0x74, 0x18, 0xc9, 0x1a, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x42, 0x03, 0xe0, 0x41, 0x03, 0x52, 0x0a, 0x75, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x12, 0x40, 0x0a, 0x0a, 0x64, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0xca, 0x1a, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x42, 0x03, 0xe0, 0x41, 0x03, 0x52, 0x0a,
	0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x12, 0x17, 0x0a, 0x06, 0x6f, 0x72,
	0x67, 0x5f, 0x69, 0x64, 0x18, 0xac, 0x1b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6f, 0x72, 0x67,
	0x5f, 0x69, 0x64, 0x12, 0x23, 0x0a, 0x0c, 0x77, 0x6f, 0x72, 0x6b, 0x73, 0x70, 0x61, 0x63, 0x65,
	0x5f, 0x69, 0x64, 0x18, 0xad, 0x1b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x77, 0x6f, 0x72, 0x6b,
	0x73, 0x70, 0x61, 0x63, 0x65, 0x5f, 0x69, 0x64, 0x12, 0x4c, 0x0a, 0x06, 0x6c, 0x61, 0x62, 0x65,
	0x6c, 0x73, 0x18, 0x99, 0xe8, 0xd8, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x31, 0x2e, 0x6b, 0x65,
	0x73, 0x73, 0x65, 0x6c, 0x2e, 0x69, 0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f, 0x72, 0x79, 0x2e, 0x76,
	0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73,
	0x2e, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x52, 0x06,
	0x6c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x42, 0x86, 0x01, 0x0a, 0x32, 0x6f, 0x72, 0x67, 0x2e, 0x70,
	0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x5f, 0x6b, 0x65, 0x73, 0x73, 0x65, 0x6c, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x69, 0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f, 0x72, 0x79, 0x2e, 0x76, 0x31, 0x62, 0x65,
	0x74, 0x61, 0x31, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x50, 0x01, 0x5a,
	0x4e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x70, 0x72, 0x6f, 0x6a,
	0x65, 0x63, 0x74, 0x2d, 0x6b, 0x65, 0x73, 0x73, 0x65, 0x6c, 0x2f, 0x69, 0x6e, 0x76, 0x65, 0x6e,
	0x74, 0x6f, 0x72, 0x79, 0x2d, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x6b, 0x65, 0x73,
	0x73, 0x65, 0x6c, 0x2f, 0x69, 0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f, 0x72, 0x79, 0x2f, 0x76, 0x31,
	0x62, 0x65, 0x74, 0x61, 0x31, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_kessel_inventory_v1beta1_resources_metadata_proto_rawDescOnce sync.Once
	file_kessel_inventory_v1beta1_resources_metadata_proto_rawDescData = file_kessel_inventory_v1beta1_resources_metadata_proto_rawDesc
)

func file_kessel_inventory_v1beta1_resources_metadata_proto_rawDescGZIP() []byte {
	file_kessel_inventory_v1beta1_resources_metadata_proto_rawDescOnce.Do(func() {
		file_kessel_inventory_v1beta1_resources_metadata_proto_rawDescData = protoimpl.X.CompressGZIP(file_kessel_inventory_v1beta1_resources_metadata_proto_rawDescData)
	})
	return file_kessel_inventory_v1beta1_resources_metadata_proto_rawDescData
}

var file_kessel_inventory_v1beta1_resources_metadata_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_kessel_inventory_v1beta1_resources_metadata_proto_goTypes = []any{
	(*Metadata)(nil),              // 0: kessel.inventory.v1beta1.resources.Metadata
	(*timestamppb.Timestamp)(nil), // 1: google.protobuf.Timestamp
	(*ResourceLabel)(nil),         // 2: kessel.inventory.v1beta1.resources.ResourceLabel
}
var file_kessel_inventory_v1beta1_resources_metadata_proto_depIdxs = []int32{
	1, // 0: kessel.inventory.v1beta1.resources.Metadata.created_at:type_name -> google.protobuf.Timestamp
	1, // 1: kessel.inventory.v1beta1.resources.Metadata.updated_at:type_name -> google.protobuf.Timestamp
	1, // 2: kessel.inventory.v1beta1.resources.Metadata.deleted_at:type_name -> google.protobuf.Timestamp
	2, // 3: kessel.inventory.v1beta1.resources.Metadata.labels:type_name -> kessel.inventory.v1beta1.resources.ResourceLabel
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_kessel_inventory_v1beta1_resources_metadata_proto_init() }
func file_kessel_inventory_v1beta1_resources_metadata_proto_init() {
	if File_kessel_inventory_v1beta1_resources_metadata_proto != nil {
		return
	}
	file_kessel_inventory_v1beta1_resources_resource_label_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_kessel_inventory_v1beta1_resources_metadata_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*Metadata); i {
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
			RawDescriptor: file_kessel_inventory_v1beta1_resources_metadata_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_kessel_inventory_v1beta1_resources_metadata_proto_goTypes,
		DependencyIndexes: file_kessel_inventory_v1beta1_resources_metadata_proto_depIdxs,
		MessageInfos:      file_kessel_inventory_v1beta1_resources_metadata_proto_msgTypes,
	}.Build()
	File_kessel_inventory_v1beta1_resources_metadata_proto = out.File
	file_kessel_inventory_v1beta1_resources_metadata_proto_rawDesc = nil
	file_kessel_inventory_v1beta1_resources_metadata_proto_goTypes = nil
	file_kessel_inventory_v1beta1_resources_metadata_proto_depIdxs = nil
}
