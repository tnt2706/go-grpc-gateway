// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v4.23.3
// source: app/model/app_calculator_model.proto

package pb

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

type User struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Number_1 int32 `protobuf:"varint,1,opt,name=number_1,json=number1,proto3" json:"number_1,omitempty"`
	Number_2 int32 `protobuf:"varint,2,opt,name=number_2,json=number2,proto3" json:"number_2,omitempty"`
}

func (x *User) Reset() {
	*x = User{}
	if protoimpl.UnsafeEnabled {
		mi := &file_app_model_app_calculator_model_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *User) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*User) ProtoMessage() {}

func (x *User) ProtoReflect() protoreflect.Message {
	mi := &file_app_model_app_calculator_model_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use User.ProtoReflect.Descriptor instead.
func (*User) Descriptor() ([]byte, []int) {
	return file_app_model_app_calculator_model_proto_rawDescGZIP(), []int{0}
}

func (x *User) GetNumber_1() int32 {
	if x != nil {
		return x.Number_1
	}
	return 0
}

func (x *User) GetNumber_2() int32 {
	if x != nil {
		return x.Number_2
	}
	return 0
}

var File_app_model_app_calculator_model_proto protoreflect.FileDescriptor

var file_app_model_app_calculator_model_proto_rawDesc = []byte{
	0x0a, 0x24, 0x61, 0x70, 0x70, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2f, 0x61, 0x70, 0x70, 0x5f,
	0x63, 0x61, 0x6c, 0x63, 0x75, 0x6c, 0x61, 0x74, 0x6f, 0x72, 0x5f, 0x6d, 0x6f, 0x64, 0x65, 0x6c,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x73, 0x68, 0x6f, 0x70, 0x64, 0x65, 0x76, 0x22,
	0x3c, 0x0a, 0x04, 0x55, 0x73, 0x65, 0x72, 0x12, 0x19, 0x0a, 0x08, 0x6e, 0x75, 0x6d, 0x62, 0x65,
	0x72, 0x5f, 0x31, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x6e, 0x75, 0x6d, 0x62, 0x65,
	0x72, 0x31, 0x12, 0x19, 0x0a, 0x08, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x5f, 0x32, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x32, 0x42, 0x06, 0x5a,
	0x04, 0x2e, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_app_model_app_calculator_model_proto_rawDescOnce sync.Once
	file_app_model_app_calculator_model_proto_rawDescData = file_app_model_app_calculator_model_proto_rawDesc
)

func file_app_model_app_calculator_model_proto_rawDescGZIP() []byte {
	file_app_model_app_calculator_model_proto_rawDescOnce.Do(func() {
		file_app_model_app_calculator_model_proto_rawDescData = protoimpl.X.CompressGZIP(file_app_model_app_calculator_model_proto_rawDescData)
	})
	return file_app_model_app_calculator_model_proto_rawDescData
}

var file_app_model_app_calculator_model_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_app_model_app_calculator_model_proto_goTypes = []any{
	(*User)(nil), // 0: shopdev.User
}
var file_app_model_app_calculator_model_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_app_model_app_calculator_model_proto_init() }
func file_app_model_app_calculator_model_proto_init() {
	if File_app_model_app_calculator_model_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_app_model_app_calculator_model_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*User); i {
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
			RawDescriptor: file_app_model_app_calculator_model_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_app_model_app_calculator_model_proto_goTypes,
		DependencyIndexes: file_app_model_app_calculator_model_proto_depIdxs,
		MessageInfos:      file_app_model_app_calculator_model_proto_msgTypes,
	}.Build()
	File_app_model_app_calculator_model_proto = out.File
	file_app_model_app_calculator_model_proto_rawDesc = nil
	file_app_model_app_calculator_model_proto_goTypes = nil
	file_app_model_app_calculator_model_proto_depIdxs = nil
}
