package logic

import (
	"MicroTikTok/Constant"
	"MicroTikTok/feed/api/internal/svc"
	"MicroTikTok/feed/api/internal/types"
	"MicroTikTok/feed/rpc/pb/video"
	"MicroTikTok/pkg/OSS"
	"MicroTikTok/pkg/ffmpeg"
	"MicroTikTok/pkg/util"
	"context"
	"fmt"
	"github.com/zeromicro/go-zero/core/logx"
	"sync"
)

type UploadLogic struct {
	logx.Logger
	ctx               context.Context
	svcCtx            *svc.ServiceContext
	mu                sync.Mutex
	uploadVideoStatus bool
	uploadCoverStatus bool
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
	var uploadTime string
	var wg sync.WaitGroup
	wg.Add(1) // 设置两个协程
	//上传视频
	go func() {
		defer wg.Done()

		status, _, ut := OSS.Upload("video", savePath)
		l.mu.Lock()
		l.uploadVideoStatus = status
		uploadTime = ut
		fmt.Println("video:", l.uploadVideoStatus)
		fmt.Println("VideoTime:", uploadTime)
		l.mu.Unlock()

	}()
	if err != nil {
		r.StatusCode = 500
		r.StatusMsg = "发布失败"
		return &r, err
	}
	wg.Wait()
	wg.Add(1)
	//上传封面
	go func() {
		defer wg.Done()
		var coverPath = "uploads/cover/" + uploadTime
		fmt.Println("CoverTime:", uploadTime)
		_, err = ffmpeg.GenerateCover(savePath, coverPath, 1)
		status, _, _ := OSS.Upload("cover", coverPath+".png")
		l.mu.Lock()
		l.uploadCoverStatus = status
		fmt.Println("cover:", l.uploadCoverStatus)
		l.mu.Unlock()
	}()
	wg.Wait()

	request := video.PublishActionRequest{}
	request.Token = req.Token
	request.Title = req.Title
	request.UploadTime = util.String(uploadTime)
	l.Logger.Info(request.GetToken(), request.Title, request.UploadTime)
	fmt.Printf("ReqToken:%v, ReqTitle:%v, ReqTime:%v", request.GetToken(), request.Title, request.UploadTime)
	response, err := l.svcCtx.VideoRpc.Upload(l.ctx, &request)
	if err != nil {
		r.StatusCode = Constant.StatusHttpFail
		r.StatusMsg = "上传失败"
		return &r, err
	}
	if l.uploadCoverStatus == false || l.uploadVideoStatus == false {
		fmt.Println("Cover:", l.uploadCoverStatus)
		fmt.Println("Video:", l.uploadVideoStatus)
		r.StatusCode = Constant.StatusHttpFail
		r.StatusMsg = "上传视频或封面失败"
		return &r, nil
	}
	fmt.Println("========上传服务Api==========")
	r.StatusCode = response.StatusCode
	r.StatusMsg = response.GetStatusMsg()
	return &r, nil
}
