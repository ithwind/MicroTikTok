package logic

import (
	"MicroTikTok/favoritelist/rpc/favoritelistclient"
	"context"
	"fmt"
	"strconv"

	"MicroTikTok/favoritelist/api/internal/svc"
	"MicroTikTok/favoritelist/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type FavoriteListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewFavoriteListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FavoriteListLogic {
	return &FavoriteListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FavoriteListLogic) FavoriteList(req *types.FavoriteListRequest) (resp *types.FavoriteListResponse, err error) {
	// todo: add your logic here and delete this line
	userId, _ := strconv.Atoi(req.UserID)
	res, err := l.svcCtx.FavoriteListRpc.FavoriteList(l.ctx, &favoritelistclient.FavoriteListRequest{
		Token:  req.Token,
		UserId: int64(userId),
	})
	if err != nil {
		return nil, err
	}
	videoListPtr := res.GetVideoList()
	fmt.Println(res.GetVideoList())
	var videoList []types.Video
	// 将每个指针类型的 Video 转换为实际类型的 Video，并添加到 videoList 中
	for _, v := range videoListPtr {
		videoList = append(videoList, types.Video{
			ID: v.GetId(),
			Author: types.User{
				ID:              v.GetAuthor().GetId(),
				Name:            v.GetAuthor().Name,
				FollowCount:     v.GetAuthor().GetFollowCount(),
				FollowerCount:   v.GetAuthor().GetFollowerCount(),
				IsFollow:        v.GetAuthor().IsFollow,
				Avatar:          v.GetAuthor().GetAvatar(),
				BackgroundImage: v.GetAuthor().GetBackgroundImage(),
				Signature:       v.GetAuthor().GetSignature(),
				TotalFavorited:  strconv.FormatInt(v.GetAuthor().TotalFavorited, 10),
				WorkCount:       v.GetAuthor().GetWorkCount(),
				FavoriteCount:   v.GetAuthor().GetFavoriteCount(),
			},
			PlayURL:       v.PlayUrl,
			CoverURL:      v.CoverUrl,
			FavoriteCount: v.FavoriteCount,
			CommentCount:  v.CommentCount,
			IsFavorite:    v.IsFavorite,
			Title:         v.Title,
		})
	}
	return &types.FavoriteListResponse{
		StatusCode: strconv.Itoa(int(res.StatusCode)),
		StatusMsg:  &res.StatusMsg,
		VideoList:  videoList,
	}, nil
}
