// Code generated by goctl. DO NOT EDIT.
// Source: video.proto

package server

import (
	"context"

	"MicroTikTok/feed/internal/logic"
	"MicroTikTok/feed/internal/svc"
	"MicroTikTok/pb/video"
)

type FeedServer struct {
	svcCtx *svc.ServiceContext
	video.UnimplementedFeedServer
}

func NewFeedServer(svcCtx *svc.ServiceContext) *FeedServer {
	return &FeedServer{
		svcCtx: svcCtx,
	}
}

func (s *FeedServer) Feed(ctx context.Context, in *video.FeedRequest) (*video.FeedResponse, error) {
	l := logic.NewFeedLogic(ctx, s.svcCtx)
	return l.Feed(in)
}

func (s *FeedServer) Upload(ctx context.Context, in *video.PublishActionRequest) (*video.PublishActionResponse, error) {
	l := logic.NewUploadLogic(ctx, s.svcCtx)
	return l.Upload(in)
}
