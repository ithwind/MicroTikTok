package logic

import (
	"MicroTikTok/favoritelist/rpc/favoritelistclient"
	"context"

	"MicroTikTok/favoritelist/api/internal/svc"
	"MicroTikTok/favoritelist/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type FavoritelistLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewFavoritelistLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FavoritelistLogic {
	return &FavoritelistLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FavoritelistLogic) Favoritelist(req *types.FavoritelistRequest) (resp *types.FavoritelistResponse, err error) {
	// todo: add your logic here and delete this line
	res, err := l.svcCtx.FavoriteListRpc.FavoriteList(l.ctx, &favoritelistclient.FavoriteListRequest{
		Token:  req.Token,
		UserId: req.UserID,
	})
	if err != nil {
		return nil, err
	}
	videoListPtr := res.GetVideoList()
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
				TotalFavorited:  v.GetAuthor().TotalFavorited,
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
	return &types.FavoritelistResponse{
		StatusCode: string(res.StatusCode),
		StatusMsg:  &res.StatusMsg,
		VideoList:  videoList,
	}, nil
}
