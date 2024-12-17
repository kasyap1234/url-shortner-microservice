// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v3.21.12
// source: proto/url.proto

package url

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	URLShortenerService_ShortenURL_FullMethodName = "/urlshortener.URLShortenerService/ShortenURL"
	URLShortenerService_ResolveURL_FullMethodName = "/urlshortener.URLShortenerService/ResolveURL"
)

// URLShortenerServiceClient is the client API for URLShortenerService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type URLShortenerServiceClient interface {
	ShortenURL(ctx context.Context, in *ShortenURLRequest, opts ...grpc.CallOption) (*ShortenURLResponse, error)
	ResolveURL(ctx context.Context, in *ResolveURLRequest, opts ...grpc.CallOption) (*ResolveURLResponse, error)
}

type uRLShortenerServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewURLShortenerServiceClient(cc grpc.ClientConnInterface) URLShortenerServiceClient {
	return &uRLShortenerServiceClient{cc}
}

func (c *uRLShortenerServiceClient) ShortenURL(ctx context.Context, in *ShortenURLRequest, opts ...grpc.CallOption) (*ShortenURLResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ShortenURLResponse)
	err := c.cc.Invoke(ctx, URLShortenerService_ShortenURL_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *uRLShortenerServiceClient) ResolveURL(ctx context.Context, in *ResolveURLRequest, opts ...grpc.CallOption) (*ResolveURLResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ResolveURLResponse)
	err := c.cc.Invoke(ctx, URLShortenerService_ResolveURL_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// URLShortenerServiceServer is the server API for URLShortenerService service.
// All implementations must embed UnimplementedURLShortenerServiceServer
// for forward compatibility.
type URLShortenerServiceServer interface {
	ShortenURL(context.Context, *ShortenURLRequest) (*ShortenURLResponse, error)
	ResolveURL(context.Context, *ResolveURLRequest) (*ResolveURLResponse, error)
	mustEmbedUnimplementedURLShortenerServiceServer()
}

// UnimplementedURLShortenerServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedURLShortenerServiceServer struct{}

func (UnimplementedURLShortenerServiceServer) ShortenURL(context.Context, *ShortenURLRequest) (*ShortenURLResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ShortenURL not implemented")
}
func (UnimplementedURLShortenerServiceServer) ResolveURL(context.Context, *ResolveURLRequest) (*ResolveURLResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ResolveURL not implemented")
}
func (UnimplementedURLShortenerServiceServer) mustEmbedUnimplementedURLShortenerServiceServer() {}
func (UnimplementedURLShortenerServiceServer) testEmbeddedByValue()                             {}

// UnsafeURLShortenerServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to URLShortenerServiceServer will
// result in compilation errors.
type UnsafeURLShortenerServiceServer interface {
	mustEmbedUnimplementedURLShortenerServiceServer()
}

func RegisterURLShortenerServiceServer(s grpc.ServiceRegistrar, srv URLShortenerServiceServer) {
	// If the following call pancis, it indicates UnimplementedURLShortenerServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&URLShortenerService_ServiceDesc, srv)
}

func _URLShortenerService_ShortenURL_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ShortenURLRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(URLShortenerServiceServer).ShortenURL(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: URLShortenerService_ShortenURL_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(URLShortenerServiceServer).ShortenURL(ctx, req.(*ShortenURLRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _URLShortenerService_ResolveURL_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ResolveURLRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(URLShortenerServiceServer).ResolveURL(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: URLShortenerService_ResolveURL_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(URLShortenerServiceServer).ResolveURL(ctx, req.(*ResolveURLRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// URLShortenerService_ServiceDesc is the grpc.ServiceDesc for URLShortenerService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var URLShortenerService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "urlshortener.URLShortenerService",
	HandlerType: (*URLShortenerServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ShortenURL",
			Handler:    _URLShortenerService_ShortenURL_Handler,
		},
		{
			MethodName: "ResolveURL",
			Handler:    _URLShortenerService_ResolveURL_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/url.proto",
}
