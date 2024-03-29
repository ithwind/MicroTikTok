// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.21.11
// source: video.proto

package video

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	Feed_Feed_FullMethodName           = "/video.Feed/feed"
	Feed_Upload_FullMethodName         = "/video.Feed/upload"
	Feed_GetPublishList_FullMethodName = "/video.Feed/getPublishList"
)

// FeedClient is the client API for Feed service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type FeedClient interface {
	Feed(ctx context.Context, in *FeedRequest, opts ...grpc.CallOption) (*FeedResponse, error)
	Upload(ctx context.Context, in *PublishActionRequest, opts ...grpc.CallOption) (*PublishActionResponse, error)
	GetPublishList(ctx context.Context, in *PublishListRequest, opts ...grpc.CallOption) (*PublishListResponse, error)
}

type feedClient struct {
	cc grpc.ClientConnInterface
}

func NewFeedClient(cc grpc.ClientConnInterface) FeedClient {
	return &feedClient{cc}
}

func (c *feedClient) Feed(ctx context.Context, in *FeedRequest, opts ...grpc.CallOption) (*FeedResponse, error) {
	out := new(FeedResponse)
	err := c.cc.Invoke(ctx, Feed_Feed_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *feedClient) Upload(ctx context.Context, in *PublishActionRequest, opts ...grpc.CallOption) (*PublishActionResponse, error) {
	out := new(PublishActionResponse)
	err := c.cc.Invoke(ctx, Feed_Upload_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *feedClient) GetPublishList(ctx context.Context, in *PublishListRequest, opts ...grpc.CallOption) (*PublishListResponse, error) {
	out := new(PublishListResponse)
	err := c.cc.Invoke(ctx, Feed_GetPublishList_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// FeedServer is the server API for Feed service.
// All implementations must embed UnimplementedFeedServer
// for forward compatibility
type FeedServer interface {
	Feed(context.Context, *FeedRequest) (*FeedResponse, error)
	Upload(context.Context, *PublishActionRequest) (*PublishActionResponse, error)
	GetPublishList(context.Context, *PublishListRequest) (*PublishListResponse, error)
	mustEmbedUnimplementedFeedServer()
}

// UnimplementedFeedServer must be embedded to have forward compatible implementations.
type UnimplementedFeedServer struct {
}

func (UnimplementedFeedServer) Feed(context.Context, *FeedRequest) (*FeedResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Feed not implemented")
}
func (UnimplementedFeedServer) Upload(context.Context, *PublishActionRequest) (*PublishActionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Upload not implemented")
}
func (UnimplementedFeedServer) GetPublishList(context.Context, *PublishListRequest) (*PublishListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPublishList not implemented")
}
func (UnimplementedFeedServer) mustEmbedUnimplementedFeedServer() {}

// UnsafeFeedServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to FeedServer will
// result in compilation errors.
type UnsafeFeedServer interface {
	mustEmbedUnimplementedFeedServer()
}

func RegisterFeedServer(s grpc.ServiceRegistrar, srv FeedServer) {
	s.RegisterService(&Feed_ServiceDesc, srv)
}

func _Feed_Feed_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FeedRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FeedServer).Feed(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Feed_Feed_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FeedServer).Feed(ctx, req.(*FeedRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Feed_Upload_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PublishActionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FeedServer).Upload(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Feed_Upload_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FeedServer).Upload(ctx, req.(*PublishActionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Feed_GetPublishList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PublishListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FeedServer).GetPublishList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Feed_GetPublishList_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FeedServer).GetPublishList(ctx, req.(*PublishListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Feed_ServiceDesc is the grpc.ServiceDesc for Feed service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Feed_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "video.Feed",
	HandlerType: (*FeedServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "feed",
			Handler:    _Feed_Feed_Handler,
		},
		{
			MethodName: "upload",
			Handler:    _Feed_Upload_Handler,
		},
		{
			MethodName: "getPublishList",
			Handler:    _Feed_GetPublishList_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "video.proto",
}
