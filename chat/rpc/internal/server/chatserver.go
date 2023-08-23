// Code generated by goctl. DO NOT EDIT.
// Source: chat.proto

package server

import (
	"context"

	"MicroTikTok/chat/rpc/internal/logic"
	"MicroTikTok/chat/rpc/internal/svc"
	"MicroTikTok/chat/rpc/pb/chat"
)

type ChatServer struct {
	svcCtx *svc.ServiceContext
	chat.UnimplementedChatServer
}

func NewChatServer(svcCtx *svc.ServiceContext) *ChatServer {
	return &ChatServer{
		svcCtx: svcCtx,
	}
}

func (s *ChatServer) ChatAction(ctx context.Context, in *chat.ChatActionRequest) (*chat.ChatActionResponse, error) {
	l := logic.NewChatActionLogic(ctx, s.svcCtx)
	return l.ChatAction(in)
}

func (s *ChatServer) ChatMessage(ctx context.Context, in *chat.ChatMessageRequest) (*chat.ChatMessageResponse, error) {
	l := logic.NewChatMessageLogic(ctx, s.svcCtx)
	return l.ChatMessage(in)
}
