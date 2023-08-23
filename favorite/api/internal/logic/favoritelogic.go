package logic

import (
	"MicroTikTok/favorite/rpc/pb/favorite"
	"context"
	"time"

	"MicroTikTok/favorite/api/internal/svc"
	"MicroTikTok/favorite/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type FavoriteLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewFavoriteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FavoriteLogic {
	return &FavoriteLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FavoriteLogic) Favorite(req *types.FavoriteRequest) (*types.FavoriteResponse, error) {
	var request favorite.FavoriteActionRequest
	request.Token = req.Token
	request.VideoId = req.VideoId
	request.ActionType = req.ActionType

	var resp types.FavoriteResponse
	ctx, cancel := context.WithTimeout(l.ctx, 100*time.Minute)
	defer cancel()
	response, err := l.svcCtx.FavoriteRpc.Favorite(ctx, &request)
	if err != nil {
		l.Logger.Error(err)
		resp.StatusCode = 400
		resp.StatusMsg = "点赞或取消点赞操作执行失败"
		return &resp, err
	}
	resp.StatusCode = int64(response.StatusCode)
	resp.StatusMsg = response.GetStatusMsg()
	return &resp, nil
}
