package logic

import (
	"MicroTikTok/AcessData/user"
	video2 "MicroTikTok/AcessData/video"
	"MicroTikTok/Constant"
	"MicroTikTok/feed/rpc/internal/svc"
	"MicroTikTok/feed/rpc/pb/video"
	"MicroTikTok/pkg/jwt"
	"MicroTikTok/pkg/util"
	"context"
	"fmt"
	"github.com/zeromicro/go-zero/core/logx"
	"time"
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
	var resp video.PublishActionResponse
	/**
	1.解析token获取当前用户信息
	*/
	claim, err := jwt.ParseToken(in.Token)
	currentUser := claim.UserVo
	l.Logger.Infof("currentUser:", currentUser)
	/**
	1.更新video表
	2.更新user_video表
	*/

	var addVideo video2.Video
	addVideo.Title = in.Title
	addVideo.PublishTime = time.Now()
	addVideo.PlayURL = Constant.URLVideoPrefix + "/" + util.GetString(in.UploadTime)
	addVideo.CoverURL = Constant.URLCoverPrefix + "/" + util.GetString(in.UploadTime)
	video2.UpdateVideo(&addVideo)
	err = user.AddUserVideoTable(currentUser.ID, addVideo.ID)

	if err != nil {
		resp.StatusCode = Constant.StatusHttpFail
		resp.StatusMsg = util.String("发布失败")
		return &resp, err
	}
	resp.StatusCode = Constant.StatusHttpOk
	resp.StatusMsg = util.String("发布成功")
	fmt.Printf("RPCResp:%v", &resp)
	return &resp, nil
}
