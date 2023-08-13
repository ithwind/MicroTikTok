package logic

import (
	"MicroTikTok/feed/rpc/internal/svc"
	"MicroTikTok/feed/rpc/pb/video"
	"MicroTikTok/pkg/util"
	"context"
	"fmt"
	"github.com/zeromicro/go-zero/core/logx"
)

type UploadLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUploadLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UploadLogic {
	return &UploadLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UploadLogic) Upload(in *video.PublishActionRequest) (*video.PublishActionResponse, error) {
	// todo: add your logic here and delete this line
	value := l.ctx.Value("data")
	fmt.Println(value)
	return &video.PublishActionResponse{
		StatusCode: 200,
		StatusMsg:  util.String("发布成功"),
	}, nil
}
