// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.32.0
// 	protoc        v4.25.2
// source: Hello.proto

package __

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

type HelloRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Something string `protobuf:"bytes,1,opt,name=Something,proto3" json:"Something,omitempty"`
}

func (x *HelloRequest) Reset() {
	*x = HelloRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Hello_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HelloRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HelloRequest) ProtoMessage() {}

func (x *HelloRequest) ProtoReflect() protoreflect.Message {
	mi := &file_Hello_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HelloRequest.ProtoReflect.Descriptor instead.
func (*HelloRequest) Descriptor() ([]byte, []int) {
	return file_Hello_proto_rawDescGZIP(), []int{0}
}

func (x *HelloRequest) GetSomething() string {
	if x != nil {
		return x.Something
	}
	return ""
}

type HelloRespone struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Reply string `protobuf:"bytes,1,opt,name=Reply,proto3" json:"Reply,omitempty"`
}

func (x *HelloRespone) Reset() {
	*x = HelloRespone{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Hello_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HelloRespone) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HelloRespone) ProtoMessage() {}

func (x *HelloRespone) ProtoReflect() protoreflect.Message {
	mi := &file_Hello_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HelloRespone.ProtoReflect.Descriptor instead.
func (*HelloRespone) Descriptor() ([]byte, []int) {
	return file_Hello_proto_rawDescGZIP(), []int{1}
}

func (x *HelloRespone) GetReply() string {
	if x != nil {
		return x.Reply
	}
	return ""
}

var File_Hello_proto protoreflect.FileDescriptor

var file_Hello_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x2c, 0x0a,
	0x0c, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1c, 0x0a,
	0x09, 0x53, 0x6f, 0x6d, 0x65, 0x74, 0x68, 0x69, 0x6e, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x09, 0x53, 0x6f, 0x6d, 0x65, 0x74, 0x68, 0x69, 0x6e, 0x67, 0x22, 0x24, 0x0a, 0x0c, 0x48,
	0x65, 0x6c, 0x6c, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x52,
	0x65, 0x70, 0x6c, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x52, 0x65, 0x70, 0x6c,
	0x79, 0x32, 0x38, 0x0a, 0x07, 0x45, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x12, 0x2d, 0x0a, 0x0b,
	0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x0d, 0x2e, 0x48, 0x65,
	0x6c, 0x6c, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0d, 0x2e, 0x48, 0x65, 0x6c,
	0x6c, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x65, 0x30, 0x01, 0x42, 0x04, 0x5a, 0x02, 0x2e,
	0x2f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_Hello_proto_rawDescOnce sync.Once
	file_Hello_proto_rawDescData = file_Hello_proto_rawDesc
)

func file_Hello_proto_rawDescGZIP() []byte {
	file_Hello_proto_rawDescOnce.Do(func() {
		file_Hello_proto_rawDescData = protoimpl.X.CompressGZIP(file_Hello_proto_rawDescData)
	})
	return file_Hello_proto_rawDescData
}

var file_Hello_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_Hello_proto_goTypes = []interface{}{
	(*HelloRequest)(nil), // 0: HelloRequest
	(*HelloRespone)(nil), // 1: HelloRespone
}
var file_Hello_proto_depIdxs = []int32{
	0, // 0: Example.ServerReply:input_type -> HelloRequest
	1, // 1: Example.ServerReply:output_type -> HelloRespone
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_Hello_proto_init() }
func file_Hello_proto_init() {
	if File_Hello_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_Hello_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HelloRequest); i {
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
		file_Hello_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HelloRespone); i {
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
			RawDescriptor: file_Hello_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_Hello_proto_goTypes,
		DependencyIndexes: file_Hello_proto_depIdxs,
		MessageInfos:      file_Hello_proto_msgTypes,
	}.Build()
	File_Hello_proto = out.File
	file_Hello_proto_rawDesc = nil
	file_Hello_proto_goTypes = nil
	file_Hello_proto_depIdxs = nil
}
