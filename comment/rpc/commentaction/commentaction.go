// Code generated by goctl. DO NOT EDIT.
// Source: comment.proto

package commentaction

import (
	"context"

	"MicroTikTok/comment/rpc/pb/comment"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	Comment           = comment.Comment
	CommentActionReq  = comment.CommentActionReq
	CommentActionResp = comment.CommentActionResp
	CommentListReq    = comment.CommentListReq
	CommentListResp   = comment.CommentListResp
	User              = comment.User

	CommentAction interface {
		CommentList(ctx context.Context, in *CommentListReq, opts ...grpc.CallOption) (*CommentListResp, error)
		CommentAction(ctx context.Context, in *CommentActionReq, opts ...grpc.CallOption) (*CommentActionResp, error)
	}

	defaultCommentAction struct {
		cli zrpc.Client
	}
)

func NewCommentAction(cli zrpc.Client) CommentAction {
	return &defaultCommentAction{
		cli: cli,
	}
}

func (m *defaultCommentAction) CommentList(ctx context.Context, in *CommentListReq, opts ...grpc.CallOption) (*CommentListResp, error) {
	client := comment.NewCommentActionClient(m.cli.Conn())
	return client.CommentList(ctx, in, opts...)
}

func (m *defaultCommentAction) CommentAction(ctx context.Context, in *CommentActionReq, opts ...grpc.CallOption) (*CommentActionResp, error) {
	client := comment.NewCommentActionClient(m.cli.Conn())
	return client.CommentAction(ctx, in, opts...)
}