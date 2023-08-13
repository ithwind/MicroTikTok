package logic

import (
	"MicroTikTok/feed/api/internal/svc"
	"MicroTikTok/feed/api/internal/types"
	"MicroTikTok/feed/api/service"
	"context"

	"github.com/zeromicro/go-zero/core/logx"
)

type FeedLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewFeedLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FeedLogic {
	return &FeedLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FeedLogic) Feed(req *types.FeedRequest) (resp *types.FeedResponse, err error) {
	if req.Token != nil || req.LastTime != nil {
		resp, err = service.NewFeedSerVice().Feed(req)
		if err != nil {
			resp.StatusCode = 500
			resp.StatusMsg = "系统错误"
			resp.VideoList = nil
			resp.NextTime = int64(0)
			return resp, err
		}
		return resp, err
	}
	resp, _ = service.NewFeedSerVice().Feed(req)

	return resp, nil
}
