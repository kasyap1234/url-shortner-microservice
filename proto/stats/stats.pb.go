// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.2
// 	protoc        v3.21.12
// source: proto/stats.proto

package stats

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

type IncrementAccessCountRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ShortUrl string `protobuf:"bytes,1,opt,name=short_url,json=shortUrl,proto3" json:"short_url,omitempty"`
}

func (x *IncrementAccessCountRequest) Reset() {
	*x = IncrementAccessCountRequest{}
	mi := &file_proto_stats_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *IncrementAccessCountRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IncrementAccessCountRequest) ProtoMessage() {}

func (x *IncrementAccessCountRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_stats_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IncrementAccessCountRequest.ProtoReflect.Descriptor instead.
func (*IncrementAccessCountRequest) Descriptor() ([]byte, []int) {
	return file_proto_stats_proto_rawDescGZIP(), []int{0}
}

func (x *IncrementAccessCountRequest) GetShortUrl() string {
	if x != nil {
		return x.ShortUrl
	}
	return ""
}

type IncrementAccessCountResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success bool `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
}

func (x *IncrementAccessCountResponse) Reset() {
	*x = IncrementAccessCountResponse{}
	mi := &file_proto_stats_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *IncrementAccessCountResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IncrementAccessCountResponse) ProtoMessage() {}

func (x *IncrementAccessCountResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_stats_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IncrementAccessCountResponse.ProtoReflect.Descriptor instead.
func (*IncrementAccessCountResponse) Descriptor() ([]byte, []int) {
	return file_proto_stats_proto_rawDescGZIP(), []int{1}
}

func (x *IncrementAccessCountResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

type GetAccessCountRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ShortUrl string `protobuf:"bytes,1,opt,name=short_url,json=shortUrl,proto3" json:"short_url,omitempty"`
}

func (x *GetAccessCountRequest) Reset() {
	*x = GetAccessCountRequest{}
	mi := &file_proto_stats_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetAccessCountRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAccessCountRequest) ProtoMessage() {}

func (x *GetAccessCountRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_stats_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAccessCountRequest.ProtoReflect.Descriptor instead.
func (*GetAccessCountRequest) Descriptor() ([]byte, []int) {
	return file_proto_stats_proto_rawDescGZIP(), []int{2}
}

func (x *GetAccessCountRequest) GetShortUrl() string {
	if x != nil {
		return x.ShortUrl
	}
	return ""
}

type GetAccessCountResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AccessCount int64 `protobuf:"varint,1,opt,name=access_count,json=accessCount,proto3" json:"access_count,omitempty"`
}

func (x *GetAccessCountResponse) Reset() {
	*x = GetAccessCountResponse{}
	mi := &file_proto_stats_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetAccessCountResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAccessCountResponse) ProtoMessage() {}

func (x *GetAccessCountResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_stats_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAccessCountResponse.ProtoReflect.Descriptor instead.
func (*GetAccessCountResponse) Descriptor() ([]byte, []int) {
	return file_proto_stats_proto_rawDescGZIP(), []int{3}
}

func (x *GetAccessCountResponse) GetAccessCount() int64 {
	if x != nil {
		return x.AccessCount
	}
	return 0
}

var File_proto_stats_proto protoreflect.FileDescriptor

var file_proto_stats_proto_rawDesc = []byte{
	0x0a, 0x11, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x73, 0x74, 0x61, 0x74, 0x73, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x05, 0x73, 0x74, 0x61, 0x74, 0x73, 0x22, 0x3a, 0x0a, 0x1b, 0x49, 0x6e,
	0x63, 0x72, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x43, 0x6f, 0x75,
	0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x73, 0x68, 0x6f,
	0x72, 0x74, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x73, 0x68,
	0x6f, 0x72, 0x74, 0x55, 0x72, 0x6c, 0x22, 0x38, 0x0a, 0x1c, 0x49, 0x6e, 0x63, 0x72, 0x65, 0x6d,
	0x65, 0x6e, 0x74, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73,
	0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73,
	0x22, 0x34, 0x0a, 0x15, 0x47, 0x65, 0x74, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x43, 0x6f, 0x75,
	0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x73, 0x68, 0x6f,
	0x72, 0x74, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x73, 0x68,
	0x6f, 0x72, 0x74, 0x55, 0x72, 0x6c, 0x22, 0x3b, 0x0a, 0x16, 0x47, 0x65, 0x74, 0x41, 0x63, 0x63,
	0x65, 0x73, 0x73, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x21, 0x0a, 0x0c, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0b, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x43, 0x6f,
	0x75, 0x6e, 0x74, 0x32, 0xbe, 0x01, 0x0a, 0x0c, 0x53, 0x74, 0x61, 0x74, 0x73, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x12, 0x5f, 0x0a, 0x14, 0x49, 0x6e, 0x63, 0x72, 0x65, 0x6d, 0x65, 0x6e,
	0x74, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x22, 0x2e, 0x73,
	0x74, 0x61, 0x74, 0x73, 0x2e, 0x49, 0x6e, 0x63, 0x72, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x41, 0x63,
	0x63, 0x65, 0x73, 0x73, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x23, 0x2e, 0x73, 0x74, 0x61, 0x74, 0x73, 0x2e, 0x49, 0x6e, 0x63, 0x72, 0x65, 0x6d, 0x65,
	0x6e, 0x74, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x4d, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x41, 0x63, 0x63, 0x65,
	0x73, 0x73, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x1c, 0x2e, 0x73, 0x74, 0x61, 0x74, 0x73, 0x2e,
	0x47, 0x65, 0x74, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1d, 0x2e, 0x73, 0x74, 0x61, 0x74, 0x73, 0x2e, 0x47, 0x65,
	0x74, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x42, 0x0d, 0x5a, 0x0b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x73, 0x74,
	0x61, 0x74, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_stats_proto_rawDescOnce sync.Once
	file_proto_stats_proto_rawDescData = file_proto_stats_proto_rawDesc
)

func file_proto_stats_proto_rawDescGZIP() []byte {
	file_proto_stats_proto_rawDescOnce.Do(func() {
		file_proto_stats_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_stats_proto_rawDescData)
	})
	return file_proto_stats_proto_rawDescData
}

var file_proto_stats_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_proto_stats_proto_goTypes = []any{
	(*IncrementAccessCountRequest)(nil),  // 0: stats.IncrementAccessCountRequest
	(*IncrementAccessCountResponse)(nil), // 1: stats.IncrementAccessCountResponse
	(*GetAccessCountRequest)(nil),        // 2: stats.GetAccessCountRequest
	(*GetAccessCountResponse)(nil),       // 3: stats.GetAccessCountResponse
}
var file_proto_stats_proto_depIdxs = []int32{
	0, // 0: stats.StatsService.IncrementAccessCount:input_type -> stats.IncrementAccessCountRequest
	2, // 1: stats.StatsService.GetAccessCount:input_type -> stats.GetAccessCountRequest
	1, // 2: stats.StatsService.IncrementAccessCount:output_type -> stats.IncrementAccessCountResponse
	3, // 3: stats.StatsService.GetAccessCount:output_type -> stats.GetAccessCountResponse
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_proto_stats_proto_init() }
func file_proto_stats_proto_init() {
	if File_proto_stats_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_proto_stats_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_stats_proto_goTypes,
		DependencyIndexes: file_proto_stats_proto_depIdxs,
		MessageInfos:      file_proto_stats_proto_msgTypes,
	}.Build()
	File_proto_stats_proto = out.File
	file_proto_stats_proto_rawDesc = nil
	file_proto_stats_proto_goTypes = nil
	file_proto_stats_proto_depIdxs = nil
}
