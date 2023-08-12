package logic

import (
	"MicroTikTok/pkg/util"
	"context"
	"fmt"
	"github.com/qiniu/go-sdk/v7/sms/bytes"

	"MicroTikTok/feed/internal/svc"
	"MicroTikTok/pb/video"

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
	reader := bytes.NewReader(in.Data)

	fmt.Println(reader)

	return &video.PublishActionResponse{
		StatusCode: 200,
		StatusMsg:  util.String("发布成功"),
	}, nil
}
