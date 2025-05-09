// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        (unknown)
// source: kessel/inventory/v1beta2/allowed.proto

package v1beta2

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

type Allowed int32

const (
	Allowed_ALLOWED_UNSPECIFIED Allowed = 0
	Allowed_ALLOWED_TRUE        Allowed = 1
	Allowed_ALLOWED_FALSE       Allowed = 2 // ALLOWED_CONDITIONAL = 3; // uncomment when used
)

// Enum value maps for Allowed.
var (
	Allowed_name = map[int32]string{
		0: "ALLOWED_UNSPECIFIED",
		1: "ALLOWED_TRUE",
		2: "ALLOWED_FALSE",
	}
	Allowed_value = map[string]int32{
		"ALLOWED_UNSPECIFIED": 0,
		"ALLOWED_TRUE":        1,
		"ALLOWED_FALSE":       2,
	}
)

func (x Allowed) Enum() *Allowed {
	p := new(Allowed)
	*p = x
	return p
}

func (x Allowed) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Allowed) Descriptor() protoreflect.EnumDescriptor {
	return file_kessel_inventory_v1beta2_allowed_proto_enumTypes[0].Descriptor()
}

func (Allowed) Type() protoreflect.EnumType {
	return &file_kessel_inventory_v1beta2_allowed_proto_enumTypes[0]
}

func (x Allowed) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Allowed.Descriptor instead.
func (Allowed) EnumDescriptor() ([]byte, []int) {
	return file_kessel_inventory_v1beta2_allowed_proto_rawDescGZIP(), []int{0}
}

var File_kessel_inventory_v1beta2_allowed_proto protoreflect.FileDescriptor

var file_kessel_inventory_v1beta2_allowed_proto_rawDesc = []byte{
	0x0a, 0x26, 0x6b, 0x65, 0x73, 0x73, 0x65, 0x6c, 0x2f, 0x69, 0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f,
	0x72, 0x79, 0x2f, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x32, 0x2f, 0x61, 0x6c, 0x6c, 0x6f, 0x77,
	0x65, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x18, 0x6b, 0x65, 0x73, 0x73, 0x65, 0x6c,
	0x2e, 0x69, 0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f, 0x72, 0x79, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74,
	0x61, 0x32, 0x2a, 0x47, 0x0a, 0x07, 0x41, 0x6c, 0x6c, 0x6f, 0x77, 0x65, 0x64, 0x12, 0x17, 0x0a,
	0x13, 0x41, 0x4c, 0x4c, 0x4f, 0x57, 0x45, 0x44, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49,
	0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x10, 0x0a, 0x0c, 0x41, 0x4c, 0x4c, 0x4f, 0x57, 0x45,
	0x44, 0x5f, 0x54, 0x52, 0x55, 0x45, 0x10, 0x01, 0x12, 0x11, 0x0a, 0x0d, 0x41, 0x4c, 0x4c, 0x4f,
	0x57, 0x45, 0x44, 0x5f, 0x46, 0x41, 0x4c, 0x53, 0x45, 0x10, 0x02, 0x42, 0x72, 0x0a, 0x28, 0x6f,
	0x72, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x5f, 0x6b, 0x65, 0x73, 0x73, 0x65,
	0x6c, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x69, 0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f, 0x72, 0x79, 0x2e,
	0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x32, 0x50, 0x01, 0x5a, 0x44, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x2d, 0x6b, 0x65,
	0x73, 0x73, 0x65, 0x6c, 0x2f, 0x69, 0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f, 0x72, 0x79, 0x2d, 0x61,
	0x70, 0x69, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x6b, 0x65, 0x73, 0x73, 0x65, 0x6c, 0x2f, 0x69, 0x6e,
	0x76, 0x65, 0x6e, 0x74, 0x6f, 0x72, 0x79, 0x2f, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x32, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_kessel_inventory_v1beta2_allowed_proto_rawDescOnce sync.Once
	file_kessel_inventory_v1beta2_allowed_proto_rawDescData = file_kessel_inventory_v1beta2_allowed_proto_rawDesc
)

func file_kessel_inventory_v1beta2_allowed_proto_rawDescGZIP() []byte {
	file_kessel_inventory_v1beta2_allowed_proto_rawDescOnce.Do(func() {
		file_kessel_inventory_v1beta2_allowed_proto_rawDescData = protoimpl.X.CompressGZIP(file_kessel_inventory_v1beta2_allowed_proto_rawDescData)
	})
	return file_kessel_inventory_v1beta2_allowed_proto_rawDescData
}

var file_kessel_inventory_v1beta2_allowed_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_kessel_inventory_v1beta2_allowed_proto_goTypes = []any{
	(Allowed)(0), // 0: kessel.inventory.v1beta2.Allowed
}
var file_kessel_inventory_v1beta2_allowed_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_kessel_inventory_v1beta2_allowed_proto_init() }
func file_kessel_inventory_v1beta2_allowed_proto_init() {
	if File_kessel_inventory_v1beta2_allowed_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_kessel_inventory_v1beta2_allowed_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_kessel_inventory_v1beta2_allowed_proto_goTypes,
		DependencyIndexes: file_kessel_inventory_v1beta2_allowed_proto_depIdxs,
		EnumInfos:         file_kessel_inventory_v1beta2_allowed_proto_enumTypes,
	}.Build()
	File_kessel_inventory_v1beta2_allowed_proto = out.File
	file_kessel_inventory_v1beta2_allowed_proto_rawDesc = nil
	file_kessel_inventory_v1beta2_allowed_proto_goTypes = nil
	file_kessel_inventory_v1beta2_allowed_proto_depIdxs = nil
}
