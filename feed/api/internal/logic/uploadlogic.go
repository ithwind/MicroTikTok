package logic

import (
	"MicroTikTok/feed/api/internal/svc"
	"MicroTikTok/feed/api/internal/types"
	"MicroTikTok/feed/rpc/pb/video"
	"MicroTikTok/pkg/OSS"
	"MicroTikTok/pkg/ffmpeg"
	"MicroTikTok/pkg/util"
	"context"
	"fmt"
	"github.com/zeromicro/go-zero/core/logx"
)

type UploadLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUploadLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UploadLogic {
	return &UploadLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UploadLogic) Upload(req *types.PublishActionRequest, savePath string) (resp *types.PublishActionResponse, err error) {
	var r types.PublishActionResponse
	//上传视频
	uploadVideoStatus, err, uploadTime := OSS.Upload("video", savePath)
	if err != nil {
		resp.StatusCode = 500
		resp.StatusMsg = "发布失败"
		return resp, err
	}
	//上传封面
	var coverPath = "feed/uploads/cover/" + uploadTime
	_, err = ffmpeg.GenerateCover(savePath, coverPath, 1)
	uploadCoverStatus, err, _ := OSS.Upload("cover", coverPath+".png")

	request := video.PublishActionRequest{}
	request.Token = req.Token
	request.Title = req.Title
	request.UploadTime = util.String(uploadTime)
	fmt.Printf("ReqToken:%v, ReqTitle:%v, ReqTime:%v", request.GetToken(), request.Title, request.UploadTime)
	response, err := l.svcCtx.VideoRpc.Upload(l.ctx, &request)
	fmt.Printf("ApiResp: %v", response)
	if err != nil {
		r.StatusCode = 400
		r.StatusMsg = "上传失败"
		return &r, err
	}
	if uploadCoverStatus == false || uploadVideoStatus == false {
		r.StatusCode = 400
		r.StatusMsg = "上传视频或封面失败"
		return &r, nil
	}
	r.StatusCode = response.StatusCode
	r.StatusMsg = response.GetStatusMsg()
	return &r, nil
}
