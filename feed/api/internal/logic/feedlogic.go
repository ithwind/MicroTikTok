package logic

import (
	"MicroTikTok/feed/api/internal/svc"
	"MicroTikTok/feed/api/internal/types"
	"MicroTikTok/feed/rpc/feed"
	"context"
	"fmt"

	"github.com/zeromicro/go-zero/core/logx"
)

type FeedLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewFeedLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FeedLogic {
	return &FeedLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FeedLogic) Feed(req *types.FeedRequest) (resp *types.FeedResponse, err error) {
	fmt.Println("=========================================")
	request := feed.FeedRequest{
		LatestTime: req.LastTime,
		Token:      req.Token,
	}
	response, err := l.svcCtx.VideoRpc.Feed(l.ctx, &request)
	if err != nil {
		resp.StatusCode = 500
		resp.StatusMsg = "系统错误"
		resp.VideoList = nil
		resp.NextTime = int64(0)
		return resp, err
	}
	// 返回 []*Video 类型的切片
	videoListPtr := response.GetVideoList()
	// 创建一个 []Video 类型的切片
	var videoList []types.Video

	// 将每个指针类型的 Video 转换为实际类型的 Video，并添加到 videoList 中
	for _, v := range videoListPtr {
		videoList = append(videoList, types.Video{
			Id: v.GetId(),
			Author: types.User{
				Id:              v.GetAuthor().GetId(),
				Name:            v.GetAuthor().Name,
				FollowCount:     v.GetAuthor().GetFollowCount(),
				FollowerCount:   v.GetAuthor().GetFollowerCount(),
				IsFollow:        v.GetAuthor().IsFollow,
				Avatar:          v.GetAuthor().GetAvatar(),
				BackgroundImage: v.GetAuthor().GetBackgroundImage(),
				Signature:       v.GetAuthor().GetSignature(),
				TotalFavorited:  v.GetAuthor().GetTotalFavorited(),
				WorkCount:       v.GetAuthor().GetWorkCount(),
				FavoriteCount:   v.GetAuthor().GetFavoriteCount(),
			},
			PlayUrl:       v.PlayUrl,
			CoverUrl:      v.CoverUrl,
			FavoriteCount: v.FavoriteCount,
			CommentCount:  v.CommentCount,
			IsFavorite:    v.IsFavorite,
			Title:         v.Title,
		})
	}

	resp = &types.FeedResponse{
		StatusCode: int64(response.StatusCode),
		StatusMsg:  response.GetStatusMsg(),
		NextTime:   response.GetNextTime(),
		VideoList:  videoList,
	}

	return resp, nil
}
