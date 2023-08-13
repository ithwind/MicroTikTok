package logic

import (
	"MicroTikTok/constant"
	"MicroTikTok/dal/user"
	"MicroTikTok/dal/video"
	"MicroTikTok/feed/api/internal/svc"
	"MicroTikTok/feed/api/internal/types"
	"MicroTikTok/pkg/OSS"
	"MicroTikTok/pkg/ffmpeg"
	"MicroTikTok/pkg/jwt"
	"context"
	"github.com/zeromicro/go-zero/core/logx"
	"time"
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
	// 上传视频形成封面
	/**
	1.解析token获取当前用户信息
	2.将本地视频上传至oss
	3.获取视频封面
	4.上传视频封面
	*/
	claim, err := jwt.ParseToken(req.Token)
	currentUser := claim.User

	var response types.PublishActionResponse
	uploadVideoStatus, err, uploadTime := OSS.Upload("video", savePath)
	if err != nil {
		response.StatusCode = 500
		response.StatusMsg = "发布失败"
		return &response, err
	}
	var coverPath = "feed/uploads/cover/" + uploadTime
	_, err = ffmpeg.GenerateCover(savePath, coverPath, 1)
	uploadCoverStatus, err, _ := OSS.Upload("cover", coverPath+".png")
	if uploadCoverStatus == false || uploadVideoStatus == false {
		response.StatusCode = 400
		response.StatusMsg = "上传视频或封面失败"
		return &response, nil
	}
	if err != nil {
		return nil, err
	}
	if err != nil {
		response.StatusCode = 400
		response.StatusMsg = "上传失败"
		return &response, err
	}
	//todo 通过获取的user更新数据库
	if uploadVideoStatus == true && uploadCoverStatus == true {
		response.StatusCode = 200
		response.StatusMsg = "操作成功"
		/**
		1.更新video表
		2.更新user_video表
		*/
		var addVideo video.Video
		addVideo.Title = req.Title
		addVideo.PublishTime = time.Now()
		addVideo.PlayURL = constant.URLVideoPrefix + "/" + uploadTime
		addVideo.CoverURL = constant.URLCoverPrefix + "/" + uploadTime
		video.UpdateVideo(&addVideo)
		err := user.AddUserVideoTable(currentUser.ID, addVideo.ID)
		if err != nil {
			return nil, err
		}
	}
	return &response, nil
}
