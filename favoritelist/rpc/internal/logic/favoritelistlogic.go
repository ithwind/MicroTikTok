package logic

import (
	"MicroTikTok/AcessData/user"
	"MicroTikTok/AcessData/video"
	"MicroTikTok/favoritelist/rpc/internal/svc"
	"MicroTikTok/favoritelist/rpc/pb/favoritelist"
	_ "MicroTikTok/pkg/corn"
	"MicroTikTok/pkg/jwt"
	"context"
	"fmt"

	"github.com/zeromicro/go-zero/core/logx"
)

type FavoriteListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewFavoriteListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FavoriteListLogic {
	return &FavoriteListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *FavoriteListLogic) FavoriteList(in *favoritelist.FavoriteListRequest) (*favoritelist.FavoriteListResponse, error) {
	/**
	0.通过token获取当前登录用户Id
	1.获取用户点赞视频Id
	2.获取点赞视频的信息
	*/
	claims, err := jwt.ParseToken(in.Token)
	if err != nil {
		fmt.Println("TokenError:", err)
		return nil, err
	}
	currentUserId := int64(claims.UserId)
	videoIds, _ := video.GetFavoriteVideoIdsByUserId(in.UserId)
	videoList := make([]*favoritelist.Video, 0)
	for _, videoId := range videoIds {
		v := video.GetVideoById(videoId)
		u := user.GetUserById(user.GetUserIdByVideoId(v.ID))
		var video favoritelist.Video
		var author favoritelist.User
		author.Id = u.ID
		author.Name = u.UserName
		author.FavoriteCount = user.GetFavoriteCount(u.ID)
		author.TotalFavorited = user.GetTotalFavoriteCount(u.ID)
		author.Avatar = u.Avatar
		author.Signature = u.Signature
		author.BackgroundImage = u.BackgroundImage
		author.WorkCount = user.GetWorkCountByUserId(u.ID)
		author.FollowerCount = user.GetFollowCountByUserId(u.ID)
		author.FollowerCount = user.GetFollowCountByUserId(u.ID)
		author.IsFollow = user.GetIsFollowByUserId(u.ID, user.GetUserIdByVideoId(v.ID))

		video.Id = v.ID
		video.Title = v.Title
		video.PlayUrl = v.PlayURL
		video.CoverUrl = v.CoverURL
		video.Author = &author
		video.IsFavorite = user.GetIsFavoriteByUserId(currentUserId, v.ID)
		video.FavoriteCount = video.GetFavoriteCount()
		video.CommentCount = video.GetCommentCount()
		videoList = append(videoList, &video)
	}
	fmt.Println("Param:", in.Token, in.UserId)
	return &favoritelist.FavoriteListResponse{
		StatusCode: 0,
		StatusMsg:  "获取成功",
		VideoList:  videoList,
	}, nil
}
