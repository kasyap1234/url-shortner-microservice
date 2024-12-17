// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.2
// 	protoc        v3.21.12
// source: proto/url.proto

package url

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

type ShortenURLRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	LongUrl string `protobuf:"bytes,1,opt,name=long_url,json=longUrl,proto3" json:"long_url,omitempty"`
}

func (x *ShortenURLRequest) Reset() {
	*x = ShortenURLRequest{}
	mi := &file_proto_url_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ShortenURLRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ShortenURLRequest) ProtoMessage() {}

func (x *ShortenURLRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_url_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ShortenURLRequest.ProtoReflect.Descriptor instead.
func (*ShortenURLRequest) Descriptor() ([]byte, []int) {
	return file_proto_url_proto_rawDescGZIP(), []int{0}
}

func (x *ShortenURLRequest) GetLongUrl() string {
	if x != nil {
		return x.LongUrl
	}
	return ""
}

type ShortenURLResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ShortUrl string `protobuf:"bytes,1,opt,name=short_url,json=shortUrl,proto3" json:"short_url,omitempty"`
}

func (x *ShortenURLResponse) Reset() {
	*x = ShortenURLResponse{}
	mi := &file_proto_url_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ShortenURLResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ShortenURLResponse) ProtoMessage() {}

func (x *ShortenURLResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_url_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ShortenURLResponse.ProtoReflect.Descriptor instead.
func (*ShortenURLResponse) Descriptor() ([]byte, []int) {
	return file_proto_url_proto_rawDescGZIP(), []int{1}
}

func (x *ShortenURLResponse) GetShortUrl() string {
	if x != nil {
		return x.ShortUrl
	}
	return ""
}

type ResolveURLRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ShortUrl string `protobuf:"bytes,1,opt,name=short_url,json=shortUrl,proto3" json:"short_url,omitempty"`
}

func (x *ResolveURLRequest) Reset() {
	*x = ResolveURLRequest{}
	mi := &file_proto_url_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ResolveURLRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResolveURLRequest) ProtoMessage() {}

func (x *ResolveURLRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_url_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResolveURLRequest.ProtoReflect.Descriptor instead.
func (*ResolveURLRequest) Descriptor() ([]byte, []int) {
	return file_proto_url_proto_rawDescGZIP(), []int{2}
}

func (x *ResolveURLRequest) GetShortUrl() string {
	if x != nil {
		return x.ShortUrl
	}
	return ""
}

type ResolveURLResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	LongUrl string `protobuf:"bytes,1,opt,name=long_url,json=longUrl,proto3" json:"long_url,omitempty"`
}

func (x *ResolveURLResponse) Reset() {
	*x = ResolveURLResponse{}
	mi := &file_proto_url_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ResolveURLResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResolveURLResponse) ProtoMessage() {}

func (x *ResolveURLResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_url_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResolveURLResponse.ProtoReflect.Descriptor instead.
func (*ResolveURLResponse) Descriptor() ([]byte, []int) {
	return file_proto_url_proto_rawDescGZIP(), []int{3}
}

func (x *ResolveURLResponse) GetLongUrl() string {
	if x != nil {
		return x.LongUrl
	}
	return ""
}

var File_proto_url_proto protoreflect.FileDescriptor

var file_proto_url_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x75, 0x72, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x0c, 0x75, 0x72, 0x6c, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x65, 0x6e, 0x65, 0x72, 0x22,
	0x2e, 0x0a, 0x11, 0x53, 0x68, 0x6f, 0x72, 0x74, 0x65, 0x6e, 0x55, 0x52, 0x4c, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x19, 0x0a, 0x08, 0x6c, 0x6f, 0x6e, 0x67, 0x5f, 0x75, 0x72, 0x6c,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6c, 0x6f, 0x6e, 0x67, 0x55, 0x72, 0x6c, 0x22,
	0x31, 0x0a, 0x12, 0x53, 0x68, 0x6f, 0x72, 0x74, 0x65, 0x6e, 0x55, 0x52, 0x4c, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x5f, 0x75,
	0x72, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x55,
	0x72, 0x6c, 0x22, 0x30, 0x0a, 0x11, 0x52, 0x65, 0x73, 0x6f, 0x6c, 0x76, 0x65, 0x55, 0x52, 0x4c,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x73, 0x68, 0x6f, 0x72, 0x74,
	0x5f, 0x75, 0x72, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x73, 0x68, 0x6f, 0x72,
	0x74, 0x55, 0x72, 0x6c, 0x22, 0x2f, 0x0a, 0x12, 0x52, 0x65, 0x73, 0x6f, 0x6c, 0x76, 0x65, 0x55,
	0x52, 0x4c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x19, 0x0a, 0x08, 0x6c, 0x6f,
	0x6e, 0x67, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6c, 0x6f,
	0x6e, 0x67, 0x55, 0x72, 0x6c, 0x32, 0xb7, 0x01, 0x0a, 0x13, 0x55, 0x52, 0x4c, 0x53, 0x68, 0x6f,
	0x72, 0x74, 0x65, 0x6e, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x4f, 0x0a,
	0x0a, 0x53, 0x68, 0x6f, 0x72, 0x74, 0x65, 0x6e, 0x55, 0x52, 0x4c, 0x12, 0x1f, 0x2e, 0x75, 0x72,
	0x6c, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x65, 0x6e, 0x65, 0x72, 0x2e, 0x53, 0x68, 0x6f, 0x72, 0x74,
	0x65, 0x6e, 0x55, 0x52, 0x4c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x20, 0x2e, 0x75,
	0x72, 0x6c, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x65, 0x6e, 0x65, 0x72, 0x2e, 0x53, 0x68, 0x6f, 0x72,
	0x74, 0x65, 0x6e, 0x55, 0x52, 0x4c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x4f,
	0x0a, 0x0a, 0x52, 0x65, 0x73, 0x6f, 0x6c, 0x76, 0x65, 0x55, 0x52, 0x4c, 0x12, 0x1f, 0x2e, 0x75,
	0x72, 0x6c, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x65, 0x6e, 0x65, 0x72, 0x2e, 0x52, 0x65, 0x73, 0x6f,
	0x6c, 0x76, 0x65, 0x55, 0x52, 0x4c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x20, 0x2e,
	0x75, 0x72, 0x6c, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x65, 0x6e, 0x65, 0x72, 0x2e, 0x52, 0x65, 0x73,
	0x6f, 0x6c, 0x76, 0x65, 0x55, 0x52, 0x4c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42,
	0x0b, 0x5a, 0x09, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x75, 0x72, 0x6c, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_url_proto_rawDescOnce sync.Once
	file_proto_url_proto_rawDescData = file_proto_url_proto_rawDesc
)

func file_proto_url_proto_rawDescGZIP() []byte {
	file_proto_url_proto_rawDescOnce.Do(func() {
		file_proto_url_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_url_proto_rawDescData)
	})
	return file_proto_url_proto_rawDescData
}

var file_proto_url_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_proto_url_proto_goTypes = []any{
	(*ShortenURLRequest)(nil),  // 0: urlshortener.ShortenURLRequest
	(*ShortenURLResponse)(nil), // 1: urlshortener.ShortenURLResponse
	(*ResolveURLRequest)(nil),  // 2: urlshortener.ResolveURLRequest
	(*ResolveURLResponse)(nil), // 3: urlshortener.ResolveURLResponse
}
var file_proto_url_proto_depIdxs = []int32{
	0, // 0: urlshortener.URLShortenerService.ShortenURL:input_type -> urlshortener.ShortenURLRequest
	2, // 1: urlshortener.URLShortenerService.ResolveURL:input_type -> urlshortener.ResolveURLRequest
	1, // 2: urlshortener.URLShortenerService.ShortenURL:output_type -> urlshortener.ShortenURLResponse
	3, // 3: urlshortener.URLShortenerService.ResolveURL:output_type -> urlshortener.ResolveURLResponse
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_proto_url_proto_init() }
func file_proto_url_proto_init() {
	if File_proto_url_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_proto_url_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_url_proto_goTypes,
		DependencyIndexes: file_proto_url_proto_depIdxs,
		MessageInfos:      file_proto_url_proto_msgTypes,
	}.Build()
	File_proto_url_proto = out.File
	file_proto_url_proto_rawDesc = nil
	file_proto_url_proto_goTypes = nil
	file_proto_url_proto_depIdxs = nil
}
