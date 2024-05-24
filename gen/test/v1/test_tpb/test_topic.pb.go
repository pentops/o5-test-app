// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        (unknown)
// source: test/v1/topic/test_topic.proto

package test_tpb

import (
	_ "buf.build/gen/go/bufbuild/protovalidate/protocolbuffers/go/buf/validate"
	_ "github.com/pentops/o5-go/messaging/v1/messaging_pb"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type GreetingMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	GreetingId string `protobuf:"bytes,1,opt,name=greeting_id,json=greetingId,proto3" json:"greeting_id,omitempty"`
	Name       string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *GreetingMessage) Reset() {
	*x = GreetingMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_test_v1_topic_test_topic_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GreetingMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GreetingMessage) ProtoMessage() {}

func (x *GreetingMessage) ProtoReflect() protoreflect.Message {
	mi := &file_test_v1_topic_test_topic_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GreetingMessage.ProtoReflect.Descriptor instead.
func (*GreetingMessage) Descriptor() ([]byte, []int) {
	return file_test_v1_topic_test_topic_proto_rawDescGZIP(), []int{0}
}

func (x *GreetingMessage) GetGreetingId() string {
	if x != nil {
		return x.GreetingId
	}
	return ""
}

func (x *GreetingMessage) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

var File_test_v1_topic_test_topic_proto protoreflect.FileDescriptor

var file_test_v1_topic_test_topic_proto_rawDesc = []byte{
	0x0a, 0x1e, 0x74, 0x65, 0x73, 0x74, 0x2f, 0x76, 0x31, 0x2f, 0x74, 0x6f, 0x70, 0x69, 0x63, 0x2f,
	0x74, 0x65, 0x73, 0x74, 0x5f, 0x74, 0x6f, 0x70, 0x69, 0x63, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x0d, 0x74, 0x65, 0x73, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x74, 0x6f, 0x70, 0x69, 0x63, 0x1a,
	0x1b, 0x62, 0x75, 0x66, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61,
	0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d,
	0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x21, 0x6f, 0x35, 0x2f, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x69, 0x6e, 0x67, 0x2f, 0x76, 0x31, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x58, 0x0a, 0x0f,
	0x47, 0x72, 0x65, 0x65, 0x74, 0x69, 0x6e, 0x67, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12,
	0x29, 0x0a, 0x0b, 0x67, 0x72, 0x65, 0x65, 0x74, 0x69, 0x6e, 0x67, 0x5f, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x42, 0x08, 0xba, 0x48, 0x05, 0x72, 0x03, 0xb0, 0x01, 0x01, 0x52, 0x0a,
	0x67, 0x72, 0x65, 0x65, 0x74, 0x69, 0x6e, 0x67, 0x49, 0x64, 0x12, 0x1a, 0x0a, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x06, 0xba, 0x48, 0x03, 0xc8, 0x01, 0x01,
	0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x32, 0x6a, 0x0a, 0x09, 0x54, 0x65, 0x73, 0x74, 0x54, 0x6f,
	0x70, 0x69, 0x63, 0x12, 0x44, 0x0a, 0x08, 0x47, 0x72, 0x65, 0x65, 0x74, 0x69, 0x6e, 0x67, 0x12,
	0x1e, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x74, 0x6f, 0x70, 0x69, 0x63, 0x2e,
	0x47, 0x72, 0x65, 0x65, 0x74, 0x69, 0x6e, 0x67, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x1a,
	0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x00, 0x1a, 0x17, 0xd2, 0xa2, 0xf5, 0xe4, 0x02,
	0x11, 0x0a, 0x0f, 0x0a, 0x0d, 0x6f, 0x35, 0x2d, 0x74, 0x65, 0x73, 0x74, 0x2d, 0x74, 0x6f, 0x70,
	0x69, 0x63, 0x42, 0x35, 0x5a, 0x33, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x70, 0x65, 0x6e, 0x74, 0x6f, 0x70, 0x73, 0x2f, 0x6f, 0x35, 0x2d, 0x74, 0x65, 0x73, 0x74,
	0x2d, 0x61, 0x70, 0x70, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x74, 0x65, 0x73, 0x74, 0x2f, 0x76, 0x31,
	0x2f, 0x74, 0x65, 0x73, 0x74, 0x5f, 0x74, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_test_v1_topic_test_topic_proto_rawDescOnce sync.Once
	file_test_v1_topic_test_topic_proto_rawDescData = file_test_v1_topic_test_topic_proto_rawDesc
)

func file_test_v1_topic_test_topic_proto_rawDescGZIP() []byte {
	file_test_v1_topic_test_topic_proto_rawDescOnce.Do(func() {
		file_test_v1_topic_test_topic_proto_rawDescData = protoimpl.X.CompressGZIP(file_test_v1_topic_test_topic_proto_rawDescData)
	})
	return file_test_v1_topic_test_topic_proto_rawDescData
}

var file_test_v1_topic_test_topic_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_test_v1_topic_test_topic_proto_goTypes = []interface{}{
	(*GreetingMessage)(nil), // 0: test.v1.topic.GreetingMessage
	(*emptypb.Empty)(nil),   // 1: google.protobuf.Empty
}
var file_test_v1_topic_test_topic_proto_depIdxs = []int32{
	0, // 0: test.v1.topic.TestTopic.Greeting:input_type -> test.v1.topic.GreetingMessage
	1, // 1: test.v1.topic.TestTopic.Greeting:output_type -> google.protobuf.Empty
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_test_v1_topic_test_topic_proto_init() }
func file_test_v1_topic_test_topic_proto_init() {
	if File_test_v1_topic_test_topic_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_test_v1_topic_test_topic_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GreetingMessage); i {
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
			RawDescriptor: file_test_v1_topic_test_topic_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_test_v1_topic_test_topic_proto_goTypes,
		DependencyIndexes: file_test_v1_topic_test_topic_proto_depIdxs,
		MessageInfos:      file_test_v1_topic_test_topic_proto_msgTypes,
	}.Build()
	File_test_v1_topic_test_topic_proto = out.File
	file_test_v1_topic_test_topic_proto_rawDesc = nil
	file_test_v1_topic_test_topic_proto_goTypes = nil
	file_test_v1_topic_test_topic_proto_depIdxs = nil
}
