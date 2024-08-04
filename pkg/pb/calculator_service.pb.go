// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v4.23.3
// source: service/calculator_service.proto

package pb

import (
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

var File_service_calculator_service_proto protoreflect.FileDescriptor

var file_service_calculator_service_proto_rawDesc = []byte{
	0x0a, 0x20, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x63, 0x61, 0x6c, 0x63, 0x75, 0x6c,
	0x61, 0x74, 0x6f, 0x72, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x07, 0x73, 0x68, 0x6f, 0x70, 0x64, 0x65, 0x76, 0x1a, 0x1e, 0x61, 0x70, 0x70,
	0x2f, 0x69, 0x66, 0x2f, 0x61, 0x70, 0x70, 0x5f, 0x63, 0x61, 0x6c, 0x63, 0x75, 0x6c, 0x61, 0x74,
	0x6f, 0x72, 0x5f, 0x69, 0x66, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x32, 0x59, 0x0a, 0x11, 0x43, 0x61, 0x6c,
	0x63, 0x75, 0x6c, 0x61, 0x74, 0x6f, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x44,
	0x0a, 0x03, 0x53, 0x75, 0x6d, 0x12, 0x13, 0x2e, 0x73, 0x68, 0x6f, 0x70, 0x64, 0x65, 0x76, 0x2e,
	0x53, 0x75, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x14, 0x2e, 0x73, 0x68, 0x6f,
	0x70, 0x64, 0x65, 0x76, 0x2e, 0x53, 0x75, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x12, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0c, 0x3a, 0x01, 0x2a, 0x22, 0x07, 0x2f, 0x76, 0x31,
	0x2f, 0x73, 0x75, 0x6d, 0x42, 0x06, 0x5a, 0x04, 0x2e, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var file_service_calculator_service_proto_goTypes = []any{
	(*SumRequest)(nil),  // 0: shopdev.SumRequest
	(*SumResponse)(nil), // 1: shopdev.SumResponse
}
var file_service_calculator_service_proto_depIdxs = []int32{
	0, // 0: shopdev.CalculatorService.Sum:input_type -> shopdev.SumRequest
	1, // 1: shopdev.CalculatorService.Sum:output_type -> shopdev.SumResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_service_calculator_service_proto_init() }
func file_service_calculator_service_proto_init() {
	if File_service_calculator_service_proto != nil {
		return
	}
	file_app_if_app_calculator_if_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_service_calculator_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_service_calculator_service_proto_goTypes,
		DependencyIndexes: file_service_calculator_service_proto_depIdxs,
	}.Build()
	File_service_calculator_service_proto = out.File
	file_service_calculator_service_proto_rawDesc = nil
	file_service_calculator_service_proto_goTypes = nil
	file_service_calculator_service_proto_depIdxs = nil
}
