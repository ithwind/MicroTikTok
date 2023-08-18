package logic

import (
	"MicroTikTok/feed/rpc/internal/svc"
	"MicroTikTok/feed/rpc/pb/video"
	service2 "MicroTikTok/feed/rpc/service"
	"MicroTikTok/pkg/util"
	"context"
	"fmt"

	"github.com/zeromicro/go-zero/core/logx"
)

type FeedLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewFeedLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FeedLogic {
	return &FeedLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *FeedLogic) Feed(in *video.FeedRequest) (*video.FeedResponse, error) {
	var resp video.FeedResponse
	var feedResp *video.FeedResponse
	fmt.Println("+++++++++++++++++++++++++++++++++")
	if in.GetToken() != "Default" || in.GetLatestTime() != 0 {
		feedResp, err := service2.NewFeedSerVice().Feed(in)
		if err != nil {
			resp.StatusCode = 500
			resp.StatusMsg = util.String("系统错误")
			resp.VideoList = nil
			resp.NextTime = nil
			return feedResp, err
		}
	}
	fmt.Printf("token:%v  time:%v", &in.Token, &in.LatestTime)
	feedResp, _ = service2.NewFeedSerVice().Feed(in)

	return feedResp, nil
}
