// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v5.27.2
// source: kessel/inventory/v1beta1/common_attributes_reporters_inner.proto

package v1beta1

import (
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

type CommonAttributesReportersInner struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The reporter type, for example, Satellite, ACM, OCM The reporter type is
	// authenticated during resource registration/update.
	Type string `protobuf:"bytes,2622298,opt,name=Type,proto3" json:"Type,omitempty"`
	// Identifies the reporter. This is meant to uniquely identify reporters that
	// can have multiple instances running, such as ACM. The instance ID is
	// authenticated during resource registration/update.
	ReporterInstanceId string `protobuf:"bytes,513872551,opt,name=reporter_instance_id,json=reporterInstanceId,proto3" json:"reporter_instance_id,omitempty"`
	// The time when this reporter record was created
	Created string `protobuf:"bytes,491683561,opt,name=created,proto3" json:"created,omitempty"`
	// The time when this reporter record was updated
	LastUpdated string `protobuf:"bytes,338699282,opt,name=last_updated,json=lastUpdated,proto3" json:"last_updated,omitempty"`
}

func (x *CommonAttributesReportersInner) Reset() {
	*x = CommonAttributesReportersInner{}
	if protoimpl.UnsafeEnabled {
		mi := &file_kessel_inventory_v1beta1_common_attributes_reporters_inner_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CommonAttributesReportersInner) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CommonAttributesReportersInner) ProtoMessage() {}

func (x *CommonAttributesReportersInner) ProtoReflect() protoreflect.Message {
	mi := &file_kessel_inventory_v1beta1_common_attributes_reporters_inner_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CommonAttributesReportersInner.ProtoReflect.Descriptor instead.
func (*CommonAttributesReportersInner) Descriptor() ([]byte, []int) {
	return file_kessel_inventory_v1beta1_common_attributes_reporters_inner_proto_rawDescGZIP(), []int{0}
}

func (x *CommonAttributesReportersInner) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *CommonAttributesReportersInner) GetReporterInstanceId() string {
	if x != nil {
		return x.ReporterInstanceId
	}
	return ""
}

func (x *CommonAttributesReportersInner) GetCreated() string {
	if x != nil {
		return x.Created
	}
	return ""
}

func (x *CommonAttributesReportersInner) GetLastUpdated() string {
	if x != nil {
		return x.LastUpdated
	}
	return ""
}

var File_kessel_inventory_v1beta1_common_attributes_reporters_inner_proto protoreflect.FileDescriptor

var file_kessel_inventory_v1beta1_common_attributes_reporters_inner_proto_rawDesc = []byte{
	0x0a, 0x40, 0x6b, 0x65, 0x73, 0x73, 0x65, 0x6c, 0x2f, 0x69, 0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f,
	0x72, 0x79, 0x2f, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f,
	0x6e, 0x5f, 0x61, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x73, 0x5f, 0x72, 0x65, 0x70,
	0x6f, 0x72, 0x74, 0x65, 0x72, 0x73, 0x5f, 0x69, 0x6e, 0x6e, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x1c, 0x61, 0x70, 0x69, 0x2e, 0x6b, 0x65, 0x73, 0x73, 0x65, 0x6c, 0x2e, 0x69,
	0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f, 0x72, 0x79, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31,
	0x22, 0xb2, 0x01, 0x0a, 0x1e, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x41, 0x74, 0x74, 0x72, 0x69,
	0x62, 0x75, 0x74, 0x65, 0x73, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x65, 0x72, 0x73, 0x49, 0x6e,
	0x6e, 0x65, 0x72, 0x12, 0x15, 0x0a, 0x04, 0x54, 0x79, 0x70, 0x65, 0x18, 0xda, 0x86, 0xa0, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x54, 0x79, 0x70, 0x65, 0x12, 0x34, 0x0a, 0x14, 0x72, 0x65,
	0x70, 0x6f, 0x72, 0x74, 0x65, 0x72, 0x5f, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x5f,
	0x69, 0x64, 0x18, 0xa7, 0xa5, 0x84, 0xf5, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x12, 0x72, 0x65,
	0x70, 0x6f, 0x72, 0x74, 0x65, 0x72, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x49, 0x64,
	0x12, 0x1c, 0x0a, 0x07, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x18, 0xe9, 0xfd, 0xb9, 0xea,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x12, 0x25,
	0x0a, 0x0c, 0x6c, 0x61, 0x73, 0x74, 0x5f, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x18, 0x92,
	0xc8, 0xc0, 0xa1, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x6c, 0x61, 0x73, 0x74, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x64, 0x42, 0x46, 0x0a, 0x1c, 0x61, 0x70, 0x69, 0x2e, 0x6b, 0x65, 0x73,
	0x73, 0x65, 0x6c, 0x2e, 0x69, 0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f, 0x72, 0x79, 0x2e, 0x76, 0x31,
	0x62, 0x65, 0x74, 0x61, 0x31, 0x50, 0x01, 0x5a, 0x24, 0x61, 0x70, 0x69, 0x2f, 0x6b, 0x65, 0x73,
	0x73, 0x65, 0x6c, 0x2f, 0x69, 0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f, 0x72, 0x79, 0x2f, 0x76, 0x31,
	0x62, 0x65, 0x74, 0x61, 0x31, 0x3b, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_kessel_inventory_v1beta1_common_attributes_reporters_inner_proto_rawDescOnce sync.Once
	file_kessel_inventory_v1beta1_common_attributes_reporters_inner_proto_rawDescData = file_kessel_inventory_v1beta1_common_attributes_reporters_inner_proto_rawDesc
)

func file_kessel_inventory_v1beta1_common_attributes_reporters_inner_proto_rawDescGZIP() []byte {
	file_kessel_inventory_v1beta1_common_attributes_reporters_inner_proto_rawDescOnce.Do(func() {
		file_kessel_inventory_v1beta1_common_attributes_reporters_inner_proto_rawDescData = protoimpl.X.CompressGZIP(file_kessel_inventory_v1beta1_common_attributes_reporters_inner_proto_rawDescData)
	})
	return file_kessel_inventory_v1beta1_common_attributes_reporters_inner_proto_rawDescData
}

var file_kessel_inventory_v1beta1_common_attributes_reporters_inner_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_kessel_inventory_v1beta1_common_attributes_reporters_inner_proto_goTypes = []any{
	(*CommonAttributesReportersInner)(nil), // 0: api.kessel.inventory.v1beta1.CommonAttributesReportersInner
}
var file_kessel_inventory_v1beta1_common_attributes_reporters_inner_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_kessel_inventory_v1beta1_common_attributes_reporters_inner_proto_init() }
func file_kessel_inventory_v1beta1_common_attributes_reporters_inner_proto_init() {
	if File_kessel_inventory_v1beta1_common_attributes_reporters_inner_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_kessel_inventory_v1beta1_common_attributes_reporters_inner_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*CommonAttributesReportersInner); i {
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
			RawDescriptor: file_kessel_inventory_v1beta1_common_attributes_reporters_inner_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_kessel_inventory_v1beta1_common_attributes_reporters_inner_proto_goTypes,
		DependencyIndexes: file_kessel_inventory_v1beta1_common_attributes_reporters_inner_proto_depIdxs,
		MessageInfos:      file_kessel_inventory_v1beta1_common_attributes_reporters_inner_proto_msgTypes,
	}.Build()
	File_kessel_inventory_v1beta1_common_attributes_reporters_inner_proto = out.File
	file_kessel_inventory_v1beta1_common_attributes_reporters_inner_proto_rawDesc = nil
	file_kessel_inventory_v1beta1_common_attributes_reporters_inner_proto_goTypes = nil
	file_kessel_inventory_v1beta1_common_attributes_reporters_inner_proto_depIdxs = nil
}