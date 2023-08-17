package logic

import (
	"MicroTikTok/dal/user"
	video2 "MicroTikTok/dal/video"
	"MicroTikTok/pkg/jwt"
	"MicroTikTok/pkg/util"
	"context"

	"MicroTikTok/feed/rpc/internal/svc"
	"MicroTikTok/feed/rpc/pb/video"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetPublishListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetPublishListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetPublishListLogic {
	return &GetPublishListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetPublishListLogic) GetPublishList(in *video.PublishListRequest) (*video.PublishListResponse, error) {
	// todo: commentCount未实现
	var response video.PublishListResponse
	/**
	1.解析token
	*/
	token := in.Token
	claims, err := jwt.ParseToken(token)
	currentUser := claims.User
	if err != nil {
		response.StatusCode = 400
		response.StatusMsg = util.String("获取失败")
		response.VideoList = nil
		return &response, err
	}
	videoList := make([]*video.Video, 0, 30)
	RawVideos, _ := video2.GetPublishList(currentUser.ID)
	for _, v := range RawVideos {
		author := user.GetUserById(user.GetUserIdByVideoId(v.ID))
		FollowCount := user.GetFollowCountByUserId(author.ID)
		FollowerCount := user.GetFollowerCountByUserId(author.ID)
		TotalFavoriteCount := user.GetTotalFavoriteCount(author.ID)
		WorkCount := user.GetWorkCountByUserId(author.ID)
		FavoriteCount := user.GetFavoriteCount(author.ID)
		backVideo := video.Video{
			Id: v.ID,
			Author: &video.User{
				Id:              author.ID,
				Name:            author.UserName,
				FollowCount:     &FollowCount,
				FollowerCount:   &FollowerCount,
				IsFollow:        user.GetIsFavoriteByUserId(author.ID, v.ID),
				Avatar:          util.String(author.Avatar),
				BackgroundImage: util.String(author.BackgroundImage),
				Signature:       util.String(author.Signature),
				TotalFavorited:  &TotalFavoriteCount,
				WorkCount:       &WorkCount,
				FavoriteCount:   &FavoriteCount,
			},
			PlayUrl:       v.PlayURL,
			CoverUrl:      v.CoverURL,
			FavoriteCount: video2.GetFavoriteCountByVideoId(v.ID),
			CommentCount:  0,
			IsFavorite:    user.GetIsFavoriteByUserId(currentUser.ID, v.ID),
			Title:         v.Title,
		}
		videoList = append(videoList, &backVideo)
	}
	response.StatusCode = 200
	response.StatusMsg = util.String("获取成功")
	response.VideoList = videoList
	return &response, nil
}
