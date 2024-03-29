// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.21.11
// source: comment.proto

package comment

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
	CommentAction_CommentList_FullMethodName   = "/comment.CommentAction/CommentList"
	CommentAction_CommentAction_FullMethodName = "/comment.CommentAction/CommentAction"
)

// CommentActionClient is the client API for CommentAction service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CommentActionClient interface {
	CommentList(ctx context.Context, in *CommentListReq, opts ...grpc.CallOption) (*CommentListResp, error)
	CommentAction(ctx context.Context, in *CommentActionReq, opts ...grpc.CallOption) (*CommentActionResp, error)
}

type commentActionClient struct {
	cc grpc.ClientConnInterface
}

func NewCommentActionClient(cc grpc.ClientConnInterface) CommentActionClient {
	return &commentActionClient{cc}
}

func (c *commentActionClient) CommentList(ctx context.Context, in *CommentListReq, opts ...grpc.CallOption) (*CommentListResp, error) {
	out := new(CommentListResp)
	err := c.cc.Invoke(ctx, CommentAction_CommentList_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *commentActionClient) CommentAction(ctx context.Context, in *CommentActionReq, opts ...grpc.CallOption) (*CommentActionResp, error) {
	out := new(CommentActionResp)
	err := c.cc.Invoke(ctx, CommentAction_CommentAction_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CommentActionServer is the server API for CommentAction service.
// All implementations must embed UnimplementedCommentActionServer
// for forward compatibility
type CommentActionServer interface {
	CommentList(context.Context, *CommentListReq) (*CommentListResp, error)
	CommentAction(context.Context, *CommentActionReq) (*CommentActionResp, error)
	mustEmbedUnimplementedCommentActionServer()
}

// UnimplementedCommentActionServer must be embedded to have forward compatible implementations.
type UnimplementedCommentActionServer struct {
}

func (UnimplementedCommentActionServer) CommentList(context.Context, *CommentListReq) (*CommentListResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CommentList not implemented")
}
func (UnimplementedCommentActionServer) CommentAction(context.Context, *CommentActionReq) (*CommentActionResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CommentAction not implemented")
}
func (UnimplementedCommentActionServer) mustEmbedUnimplementedCommentActionServer() {}

// UnsafeCommentActionServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CommentActionServer will
// result in compilation errors.
type UnsafeCommentActionServer interface {
	mustEmbedUnimplementedCommentActionServer()
}

func RegisterCommentActionServer(s grpc.ServiceRegistrar, srv CommentActionServer) {
	s.RegisterService(&CommentAction_ServiceDesc, srv)
}

func _CommentAction_CommentList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CommentListReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CommentActionServer).CommentList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CommentAction_CommentList_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CommentActionServer).CommentList(ctx, req.(*CommentListReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _CommentAction_CommentAction_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CommentActionReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CommentActionServer).CommentAction(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CommentAction_CommentAction_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CommentActionServer).CommentAction(ctx, req.(*CommentActionReq))
	}
	return interceptor(ctx, in, info, handler)
}

// CommentAction_ServiceDesc is the grpc.ServiceDesc for CommentAction service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CommentAction_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "comment.CommentAction",
	HandlerType: (*CommentActionServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CommentList",
			Handler:    _CommentAction_CommentList_Handler,
		},
		{
			MethodName: "CommentAction",
			Handler:    _CommentAction_CommentAction_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "comment.proto",
}
