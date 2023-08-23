package logic

import (
	"MicroTikTok/favorite/rpc/internal/service"
	"MicroTikTok/favorite/rpc/internal/svc"
	"MicroTikTok/favorite/rpc/pb/favorite"
	"MicroTikTok/pkg/jwt"
	"MicroTikTok/pkg/redis"
	"MicroTikTok/pkg/util"
	"context"
	"github.com/zeromicro/go-zero/core/logx"
	"strconv"
)

type FavoriteLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewFavoriteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FavoriteLogic {
	return &FavoriteLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *FavoriteLogic) Favorite(in *favorite.FavoriteActionRequest) (*favorite.FavoriteActionResponse, error) {
	redisService, err := redis.NewRedisCacheService()
	var resp favorite.FavoriteActionResponse
	/**
	1.解析token
	2.将用户id和视频id存入redis
	*/
	claims, err := jwt.ParseToken(in.Token)
	if err != nil {
		resp.StatusCode = 400
		resp.StatusMsg = util.String("当前用户不存在")
		return &resp, err
	}
	currentUser := claims.UserVo
	currentUserId := currentUser.ID
	//插入redis
	l.Logger.Infof("currentUserId", currentUserId)
	l.Logger.Infof("videoId:", in.VideoId)
	redisService.HashSetRedis(strconv.FormatInt(currentUserId, 10), strconv.FormatInt(in.VideoId, 10), in.ActionType)

	go service.FavoriteService()

	if in.ActionType == "1" {
		resp.StatusCode = 200
		resp.StatusMsg = util.String("点赞成功")
	} else if in.ActionType == "0" {
		resp.StatusCode = 200
		resp.StatusMsg = util.String("取消点赞成功")
	}
	return &resp, nil
}
