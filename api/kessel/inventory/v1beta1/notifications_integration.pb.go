// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v5.27.2
// source: kessel/inventory/v1beta1/notifications_integration.proto

package v1beta1

import (
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
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

type NotificationsIntegration struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Metadata about this resource
	Metadata *Metadata `protobuf:"bytes,1,opt,name=metadata,proto3" json:"metadata,omitempty"`
	// Write only reporter specific data
	ReporterData *ReporterData `protobuf:"bytes,245278792,opt,name=reporter_data,json=reporterData,proto3" json:"reporter_data,omitempty"`
	// The entities that registered this item in the Kessel Asset Inventory. The
	// same resource may be registered by multiple reporters.
	Reporters []*ReporterData `protobuf:"bytes,353323086,rep,name=reporters,proto3" json:"reporters,omitempty"`
}

func (x *NotificationsIntegration) Reset() {
	*x = NotificationsIntegration{}
	if protoimpl.UnsafeEnabled {
		mi := &file_kessel_inventory_v1beta1_notifications_integration_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NotificationsIntegration) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NotificationsIntegration) ProtoMessage() {}

func (x *NotificationsIntegration) ProtoReflect() protoreflect.Message {
	mi := &file_kessel_inventory_v1beta1_notifications_integration_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NotificationsIntegration.ProtoReflect.Descriptor instead.
func (*NotificationsIntegration) Descriptor() ([]byte, []int) {
	return file_kessel_inventory_v1beta1_notifications_integration_proto_rawDescGZIP(), []int{0}
}

func (x *NotificationsIntegration) GetMetadata() *Metadata {
	if x != nil {
		return x.Metadata
	}
	return nil
}

func (x *NotificationsIntegration) GetReporterData() *ReporterData {
	if x != nil {
		return x.ReporterData
	}
	return nil
}

func (x *NotificationsIntegration) GetReporters() []*ReporterData {
	if x != nil {
		return x.Reporters
	}
	return nil
}

var File_kessel_inventory_v1beta1_notifications_integration_proto protoreflect.FileDescriptor

var file_kessel_inventory_v1beta1_notifications_integration_proto_rawDesc = []byte{
	0x0a, 0x38, 0x6b, 0x65, 0x73, 0x73, 0x65, 0x6c, 0x2f, 0x69, 0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f,
	0x72, 0x79, 0x2f, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2f, 0x6e, 0x6f, 0x74, 0x69, 0x66,
	0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x5f, 0x69, 0x6e, 0x74, 0x65, 0x67, 0x72, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1c, 0x61, 0x70, 0x69, 0x2e,
	0x6b, 0x65, 0x73, 0x73, 0x65, 0x6c, 0x2e, 0x69, 0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f, 0x72, 0x79,
	0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x61, 0x70, 0x69, 0x2f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x62, 0x65, 0x68, 0x61, 0x76,
	0x69, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x76, 0x61, 0x6c, 0x69, 0x64,
	0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x27, 0x6b, 0x65, 0x73, 0x73, 0x65, 0x6c, 0x2f, 0x69, 0x6e, 0x76, 0x65, 0x6e,
	0x74, 0x6f, 0x72, 0x79, 0x2f, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2f, 0x6d, 0x65, 0x74,
	0x61, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2c, 0x6b, 0x65, 0x73,
	0x73, 0x65, 0x6c, 0x2f, 0x69, 0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f, 0x72, 0x79, 0x2f, 0x76, 0x31,
	0x62, 0x65, 0x74, 0x61, 0x31, 0x2f, 0x72, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x65, 0x72, 0x5f, 0x64,
	0x61, 0x74, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xa2, 0x02, 0x0a, 0x18, 0x4e, 0x6f,
	0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x49, 0x6e, 0x74, 0x65, 0x67,
	0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x50, 0x0a, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61,
	0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x26, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x6b,
	0x65, 0x73, 0x73, 0x65, 0x6c, 0x2e, 0x69, 0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f, 0x72, 0x79, 0x2e,
	0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61,
	0x42, 0x0c, 0xe2, 0x41, 0x01, 0x02, 0xfa, 0x42, 0x05, 0xa2, 0x01, 0x02, 0x08, 0x01, 0x52, 0x08,
	0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x12, 0x60, 0x0a, 0x0d, 0x72, 0x65, 0x70, 0x6f,
	0x72, 0x74, 0x65, 0x72, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x18, 0xc8, 0xd0, 0xfa, 0x74, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x2a, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x6b, 0x65, 0x73, 0x73, 0x65, 0x6c, 0x2e,
	0x69, 0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f, 0x72, 0x79, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61,
	0x31, 0x2e, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x65, 0x72, 0x44, 0x61, 0x74, 0x61, 0x42, 0x0c,
	0xe2, 0x41, 0x01, 0x04, 0xfa, 0x42, 0x05, 0xa2, 0x01, 0x02, 0x08, 0x01, 0x52, 0x0c, 0x72, 0x65,
	0x70, 0x6f, 0x72, 0x74, 0x65, 0x72, 0x44, 0x61, 0x74, 0x61, 0x12, 0x52, 0x0a, 0x09, 0x72, 0x65,
	0x70, 0x6f, 0x72, 0x74, 0x65, 0x72, 0x73, 0x18, 0xce, 0x90, 0xbd, 0xa8, 0x01, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x2a, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x6b, 0x65, 0x73, 0x73, 0x65, 0x6c, 0x2e, 0x69,
	0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f, 0x72, 0x79, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31,
	0x2e, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x65, 0x72, 0x44, 0x61, 0x74, 0x61, 0x42, 0x04, 0xe2,
	0x41, 0x01, 0x03, 0x52, 0x09, 0x72, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x65, 0x72, 0x73, 0x42, 0x46,
	0x0a, 0x1c, 0x61, 0x70, 0x69, 0x2e, 0x6b, 0x65, 0x73, 0x73, 0x65, 0x6c, 0x2e, 0x69, 0x6e, 0x76,
	0x65, 0x6e, 0x74, 0x6f, 0x72, 0x79, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x50, 0x01,
	0x5a, 0x24, 0x61, 0x70, 0x69, 0x2f, 0x6b, 0x65, 0x73, 0x73, 0x65, 0x6c, 0x2f, 0x69, 0x6e, 0x76,
	0x65, 0x6e, 0x74, 0x6f, 0x72, 0x79, 0x2f, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x3b, 0x76,
	0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x50, 0x02, 0x50, 0x03, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_kessel_inventory_v1beta1_notifications_integration_proto_rawDescOnce sync.Once
	file_kessel_inventory_v1beta1_notifications_integration_proto_rawDescData = file_kessel_inventory_v1beta1_notifications_integration_proto_rawDesc
)

func file_kessel_inventory_v1beta1_notifications_integration_proto_rawDescGZIP() []byte {
	file_kessel_inventory_v1beta1_notifications_integration_proto_rawDescOnce.Do(func() {
		file_kessel_inventory_v1beta1_notifications_integration_proto_rawDescData = protoimpl.X.CompressGZIP(file_kessel_inventory_v1beta1_notifications_integration_proto_rawDescData)
	})
	return file_kessel_inventory_v1beta1_notifications_integration_proto_rawDescData
}

var file_kessel_inventory_v1beta1_notifications_integration_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_kessel_inventory_v1beta1_notifications_integration_proto_goTypes = []any{
	(*NotificationsIntegration)(nil), // 0: api.kessel.inventory.v1beta1.NotificationsIntegration
	(*Metadata)(nil),                 // 1: api.kessel.inventory.v1beta1.Metadata
	(*ReporterData)(nil),             // 2: api.kessel.inventory.v1beta1.ReporterData
}
var file_kessel_inventory_v1beta1_notifications_integration_proto_depIdxs = []int32{
	1, // 0: api.kessel.inventory.v1beta1.NotificationsIntegration.metadata:type_name -> api.kessel.inventory.v1beta1.Metadata
	2, // 1: api.kessel.inventory.v1beta1.NotificationsIntegration.reporter_data:type_name -> api.kessel.inventory.v1beta1.ReporterData
	2, // 2: api.kessel.inventory.v1beta1.NotificationsIntegration.reporters:type_name -> api.kessel.inventory.v1beta1.ReporterData
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_kessel_inventory_v1beta1_notifications_integration_proto_init() }
func file_kessel_inventory_v1beta1_notifications_integration_proto_init() {
	if File_kessel_inventory_v1beta1_notifications_integration_proto != nil {
		return
	}
	file_kessel_inventory_v1beta1_metadata_proto_init()
	file_kessel_inventory_v1beta1_reporter_data_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_kessel_inventory_v1beta1_notifications_integration_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*NotificationsIntegration); i {
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
			RawDescriptor: file_kessel_inventory_v1beta1_notifications_integration_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_kessel_inventory_v1beta1_notifications_integration_proto_goTypes,
		DependencyIndexes: file_kessel_inventory_v1beta1_notifications_integration_proto_depIdxs,
		MessageInfos:      file_kessel_inventory_v1beta1_notifications_integration_proto_msgTypes,
	}.Build()
	File_kessel_inventory_v1beta1_notifications_integration_proto = out.File
	file_kessel_inventory_v1beta1_notifications_integration_proto_rawDesc = nil
	file_kessel_inventory_v1beta1_notifications_integration_proto_goTypes = nil
	file_kessel_inventory_v1beta1_notifications_integration_proto_depIdxs = nil
}
