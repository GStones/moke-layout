// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.3
// 	protoc        (unknown)
// source: demo/demo.proto

package pb

import (
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

type HiRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Uid           string                 `protobuf:"bytes,1,opt,name=uid,proto3" json:"uid,omitempty"`
	Message       string                 `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	Topic         string                 `protobuf:"bytes,3,opt,name=topic,proto3" json:"topic,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *HiRequest) Reset() {
	*x = HiRequest{}
	mi := &file_demo_demo_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *HiRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HiRequest) ProtoMessage() {}

func (x *HiRequest) ProtoReflect() protoreflect.Message {
	mi := &file_demo_demo_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HiRequest.ProtoReflect.Descriptor instead.
func (*HiRequest) Descriptor() ([]byte, []int) {
	return file_demo_demo_proto_rawDescGZIP(), []int{0}
}

func (x *HiRequest) GetUid() string {
	if x != nil {
		return x.Uid
	}
	return ""
}

func (x *HiRequest) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *HiRequest) GetTopic() string {
	if x != nil {
		return x.Topic
	}
	return ""
}

type HiResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Message       string                 `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *HiResponse) Reset() {
	*x = HiResponse{}
	mi := &file_demo_demo_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *HiResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HiResponse) ProtoMessage() {}

func (x *HiResponse) ProtoReflect() protoreflect.Message {
	mi := &file_demo_demo_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HiResponse.ProtoReflect.Descriptor instead.
func (*HiResponse) Descriptor() ([]byte, []int) {
	return file_demo_demo_proto_rawDescGZIP(), []int{1}
}

func (x *HiResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type WatchRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Uid           string                 `protobuf:"bytes,1,opt,name=uid,proto3" json:"uid,omitempty"`
	Topic         string                 `protobuf:"bytes,2,opt,name=topic,proto3" json:"topic,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *WatchRequest) Reset() {
	*x = WatchRequest{}
	mi := &file_demo_demo_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *WatchRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WatchRequest) ProtoMessage() {}

func (x *WatchRequest) ProtoReflect() protoreflect.Message {
	mi := &file_demo_demo_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WatchRequest.ProtoReflect.Descriptor instead.
func (*WatchRequest) Descriptor() ([]byte, []int) {
	return file_demo_demo_proto_rawDescGZIP(), []int{2}
}

func (x *WatchRequest) GetUid() string {
	if x != nil {
		return x.Uid
	}
	return ""
}

func (x *WatchRequest) GetTopic() string {
	if x != nil {
		return x.Topic
	}
	return ""
}

type WatchResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Message       string                 `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *WatchResponse) Reset() {
	*x = WatchResponse{}
	mi := &file_demo_demo_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *WatchResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WatchResponse) ProtoMessage() {}

func (x *WatchResponse) ProtoReflect() protoreflect.Message {
	mi := &file_demo_demo_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WatchResponse.ProtoReflect.Descriptor instead.
func (*WatchResponse) Descriptor() ([]byte, []int) {
	return file_demo_demo_proto_rawDescGZIP(), []int{3}
}

func (x *WatchResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

var File_demo_demo_proto protoreflect.FileDescriptor

var file_demo_demo_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x64, 0x65, 0x6d, 0x6f, 0x2f, 0x64, 0x65, 0x6d, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x07, 0x64, 0x65, 0x6d, 0x6f, 0x2e, 0x70, 0x62, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x4d, 0x0a, 0x09, 0x48, 0x69, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x03, 0x75, 0x69, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x70, 0x69, 0x63, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x74, 0x6f, 0x70, 0x69, 0x63, 0x22, 0x26, 0x0a, 0x0a, 0x48, 0x69, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22,
	0x36, 0x0a, 0x0c, 0x57, 0x61, 0x74, 0x63, 0x68, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x10, 0x0a, 0x03, 0x75, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x69,
	0x64, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x70, 0x69, 0x63, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x74, 0x6f, 0x70, 0x69, 0x63, 0x22, 0x29, 0x0a, 0x0d, 0x57, 0x61, 0x74, 0x63, 0x68,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x32, 0xa8, 0x01, 0x0a, 0x0b, 0x44, 0x65, 0x6d, 0x6f, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x12, 0x46, 0x0a, 0x02, 0x48, 0x69, 0x12, 0x12, 0x2e, 0x64, 0x65, 0x6d, 0x6f, 0x2e,
	0x70, 0x62, 0x2e, 0x48, 0x69, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x13, 0x2e, 0x64,
	0x65, 0x6d, 0x6f, 0x2e, 0x70, 0x62, 0x2e, 0x48, 0x69, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x17, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x11, 0x3a, 0x01, 0x2a, 0x22, 0x0c, 0x2f, 0x76,
	0x31, 0x2f, 0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x2f, 0x68, 0x69, 0x12, 0x51, 0x0a, 0x05, 0x57, 0x61,
	0x74, 0x63, 0x68, 0x12, 0x15, 0x2e, 0x64, 0x65, 0x6d, 0x6f, 0x2e, 0x70, 0x62, 0x2e, 0x57, 0x61,
	0x74, 0x63, 0x68, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x64, 0x65, 0x6d,
	0x6f, 0x2e, 0x70, 0x62, 0x2e, 0x57, 0x61, 0x74, 0x63, 0x68, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x17, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x11, 0x12, 0x0f, 0x2f, 0x76, 0x31, 0x2f,
	0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x2f, 0x77, 0x61, 0x74, 0x63, 0x68, 0x30, 0x01, 0x42, 0x0d, 0x5a,
	0x0b, 0x64, 0x65, 0x6d, 0x6f, 0x2f, 0x61, 0x70, 0x69, 0x3b, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_demo_demo_proto_rawDescOnce sync.Once
	file_demo_demo_proto_rawDescData = file_demo_demo_proto_rawDesc
)

func file_demo_demo_proto_rawDescGZIP() []byte {
	file_demo_demo_proto_rawDescOnce.Do(func() {
		file_demo_demo_proto_rawDescData = protoimpl.X.CompressGZIP(file_demo_demo_proto_rawDescData)
	})
	return file_demo_demo_proto_rawDescData
}

var file_demo_demo_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_demo_demo_proto_goTypes = []any{
	(*HiRequest)(nil),     // 0: demo.pb.HiRequest
	(*HiResponse)(nil),    // 1: demo.pb.HiResponse
	(*WatchRequest)(nil),  // 2: demo.pb.WatchRequest
	(*WatchResponse)(nil), // 3: demo.pb.WatchResponse
}
var file_demo_demo_proto_depIdxs = []int32{
	0, // 0: demo.pb.DemoService.Hi:input_type -> demo.pb.HiRequest
	2, // 1: demo.pb.DemoService.Watch:input_type -> demo.pb.WatchRequest
	1, // 2: demo.pb.DemoService.Hi:output_type -> demo.pb.HiResponse
	3, // 3: demo.pb.DemoService.Watch:output_type -> demo.pb.WatchResponse
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_demo_demo_proto_init() }
func file_demo_demo_proto_init() {
	if File_demo_demo_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_demo_demo_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_demo_demo_proto_goTypes,
		DependencyIndexes: file_demo_demo_proto_depIdxs,
		MessageInfos:      file_demo_demo_proto_msgTypes,
	}.Build()
	File_demo_demo_proto = out.File
	file_demo_demo_proto_rawDesc = nil
	file_demo_demo_proto_goTypes = nil
	file_demo_demo_proto_depIdxs = nil
}
