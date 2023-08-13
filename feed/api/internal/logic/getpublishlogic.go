package logic

import (
	"MicroTikTok/feed/api/service"
	"MicroTikTok/pkg/jwt"
	"context"

	"MicroTikTok/feed/api/internal/svc"
	"MicroTikTok/feed/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetPublishLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetPublishLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetPublishLogic {
	return &GetPublishLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetPublishLogic) GetPublish(req *types.PublishListRequest) (resp *types.PublishListResponse, err error) {
	// todo: add your logic here and delete this line\
	/**
	1.获取鉴权token解析获得用户
	2.查找对应视频返回
	*/
	token := req.Token
	claims, err := jwt.ParseToken(token)
	currentUser := claims.User
	videoList := service.NewFeedSerVice().GetPublishList(currentUser.ID)
	var response types.PublishListResponse
	if err != nil {
		response.StatusCode = 400
		response.StatusMsg = "获取失败"
		response.VideoList = nil
		return &response, err
	}
	response.StatusCode = 200
	response.StatusMsg = "获取成功"
	response.VideoList = videoList
	return &response, nil
}
