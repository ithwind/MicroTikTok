package logic

import (
	"MicroTikTok/feed/rpc/pb/video"
	"MicroTikTok/pkg/util"
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

func (l *GetPublishLogic) GetPublish(req *types.PublishListRequest) (*types.PublishListResponse, error) {
	// todo: add your logic here and delete this line\

	var request video.PublishListRequest
	var resp types.PublishListResponse
	request.UserId = req.UserId
	request.Token = req.Token
	publishListResponse, err := l.svcCtx.VideoRpc.GetPublishList(l.ctx, &request)
	if err != nil {
		return nil, err
	}
	videoPointers := publishListResponse.GetVideoList()

	// 创建一个 []Video 类型的切片，遍历原始切片并解引用指针元素，将其添加到新切片中
	var videos []types.Video
	for _, videoPointer := range videoPointers {
		author := videoPointer.Author
		v := types.Video{
			Id: videoPointer.Id,
			Author: types.User{
				Id:              author.Id,
				Name:            author.Name,
				FollowCount:     *author.FollowCount,
				FollowerCount:   *author.FollowerCount,
				IsFollow:        author.IsFollow,
				Avatar:          util.GetString(author.Avatar),
				BackgroundImage: util.GetString(author.BackgroundImage),
				Signature:       util.GetString(author.Signature),
				TotalFavorited:  *author.TotalFavorited,
				WorkCount:       *author.WorkCount,
				FavoriteCount:   *author.FavoriteCount,
			},
			PlayUrl:       videoPointer.PlayUrl,
			CoverUrl:      videoPointer.CoverUrl,
			FavoriteCount: videoPointer.FavoriteCount,
			CommentCount:  videoPointer.CommentCount,
			IsFavorite:    videoPointer.IsFavorite,
			Title:         videoPointer.Title,
		}
		videos = append(videos, v)
	}
	resp.StatusCode = int64(publishListResponse.StatusCode)
	resp.StatusMsg = publishListResponse.GetStatusMsg()
	resp.VideoList = videos
	return &resp, nil
}
