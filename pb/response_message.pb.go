// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.23.0
// 	protoc        v3.6.1
// source: response_message.proto

package pb

import (
	proto "github.com/golang/protobuf/proto"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type Response_ReturnCode int32

const (
	Response_unknown Response_ReturnCode = 0
	Response_ok      Response_ReturnCode = 1
	Response_skip    Response_ReturnCode = 2
	Response_error   Response_ReturnCode = 3
)

// Enum value maps for Response_ReturnCode.
var (
	Response_ReturnCode_name = map[int32]string{
		0: "unknown",
		1: "ok",
		2: "skip",
		3: "error",
	}
	Response_ReturnCode_value = map[string]int32{
		"unknown": 0,
		"ok":      1,
		"skip":    2,
		"error":   3,
	}
)

func (x Response_ReturnCode) Enum() *Response_ReturnCode {
	p := new(Response_ReturnCode)
	*p = x
	return p
}

func (x Response_ReturnCode) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Response_ReturnCode) Descriptor() protoreflect.EnumDescriptor {
	return file_response_message_proto_enumTypes[0].Descriptor()
}

func (Response_ReturnCode) Type() protoreflect.EnumType {
	return &file_response_message_proto_enumTypes[0]
}

func (x Response_ReturnCode) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Response_ReturnCode.Descriptor instead.
func (Response_ReturnCode) EnumDescriptor() ([]byte, []int) {
	return file_response_message_proto_rawDescGZIP(), []int{0, 0}
}

type Response struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ReturnCode    Response_ReturnCode `protobuf:"varint,1,opt,name=return_code,json=returnCode,proto3,enum=dsipe.transfer.Response_ReturnCode" json:"return_code,omitempty"`
	ReturnMessage string              `protobuf:"bytes,2,opt,name=return_message,json=returnMessage,proto3" json:"return_message,omitempty"`
	Context       []int32             `protobuf:"varint,3,rep,packed,name=context,proto3" json:"context,omitempty"`
}

func (x *Response) Reset() {
	*x = Response{}
	if protoimpl.UnsafeEnabled {
		mi := &file_response_message_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Response) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Response) ProtoMessage() {}

func (x *Response) ProtoReflect() protoreflect.Message {
	mi := &file_response_message_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Response.ProtoReflect.Descriptor instead.
func (*Response) Descriptor() ([]byte, []int) {
	return file_response_message_proto_rawDescGZIP(), []int{0}
}

func (x *Response) GetReturnCode() Response_ReturnCode {
	if x != nil {
		return x.ReturnCode
	}
	return Response_unknown
}

func (x *Response) GetReturnMessage() string {
	if x != nil {
		return x.ReturnMessage
	}
	return ""
}

func (x *Response) GetContext() []int32 {
	if x != nil {
		return x.Context
	}
	return nil
}

var File_response_message_proto protoreflect.FileDescriptor

var file_response_message_proto_rawDesc = []byte{
	0x0a, 0x16, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x12, 0x70, 0x6c, 0x75, 0x67, 0x67, 0x61,
	0x62, 0x6c, 0x65, 0x2e, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x22, 0xcd, 0x01, 0x0a,
	0x08, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x48, 0x0a, 0x0b, 0x72, 0x65, 0x74,
	0x75, 0x72, 0x6e, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x27,
	0x2e, 0x70, 0x6c, 0x75, 0x67, 0x67, 0x61, 0x62, 0x6c, 0x65, 0x2e, 0x74, 0x72, 0x61, 0x6e, 0x73,
	0x66, 0x65, 0x72, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x52, 0x65, 0x74,
	0x75, 0x72, 0x6e, 0x43, 0x6f, 0x64, 0x65, 0x52, 0x0a, 0x72, 0x65, 0x74, 0x75, 0x72, 0x6e, 0x43,
	0x6f, 0x64, 0x65, 0x12, 0x25, 0x0a, 0x0e, 0x72, 0x65, 0x74, 0x75, 0x72, 0x6e, 0x5f, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x72, 0x65, 0x74,
	0x75, 0x72, 0x6e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f,
	0x6e, 0x74, 0x65, 0x78, 0x74, 0x18, 0x03, 0x20, 0x03, 0x28, 0x05, 0x52, 0x07, 0x63, 0x6f, 0x6e,
	0x74, 0x65, 0x78, 0x74, 0x22, 0x36, 0x0a, 0x0a, 0x52, 0x65, 0x74, 0x75, 0x72, 0x6e, 0x43, 0x6f,
	0x64, 0x65, 0x12, 0x0b, 0x0a, 0x07, 0x75, 0x6e, 0x6b, 0x6e, 0x6f, 0x77, 0x6e, 0x10, 0x00, 0x12,
	0x06, 0x0a, 0x02, 0x6f, 0x6b, 0x10, 0x01, 0x12, 0x08, 0x0a, 0x04, 0x73, 0x6b, 0x69, 0x70, 0x10,
	0x02, 0x12, 0x09, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x10, 0x03, 0x42, 0x06, 0x5a, 0x04,
	0x2e, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_response_message_proto_rawDescOnce sync.Once
	file_response_message_proto_rawDescData = file_response_message_proto_rawDesc
)

func file_response_message_proto_rawDescGZIP() []byte {
	file_response_message_proto_rawDescOnce.Do(func() {
		file_response_message_proto_rawDescData = protoimpl.X.CompressGZIP(file_response_message_proto_rawDescData)
	})
	return file_response_message_proto_rawDescData
}

var file_response_message_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_response_message_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_response_message_proto_goTypes = []interface{}{
	(Response_ReturnCode)(0), // 0: dsipe.transfer.Response.ReturnCode
	(*Response)(nil),         // 1: dsipe.transfer.Response
}
var file_response_message_proto_depIdxs = []int32{
	0, // 0: dsipe.transfer.Response.return_code:type_name -> dsipe.transfer.Response.ReturnCode
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_response_message_proto_init() }
func file_response_message_proto_init() {
	if File_response_message_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_response_message_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Response); i {
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
			RawDescriptor: file_response_message_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_response_message_proto_goTypes,
		DependencyIndexes: file_response_message_proto_depIdxs,
		EnumInfos:         file_response_message_proto_enumTypes,
		MessageInfos:      file_response_message_proto_msgTypes,
	}.Build()
	File_response_message_proto = out.File
	file_response_message_proto_rawDesc = nil
	file_response_message_proto_goTypes = nil
	file_response_message_proto_depIdxs = nil
}
