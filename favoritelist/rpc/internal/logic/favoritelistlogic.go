package logic

import (
	"MicroTikTok/favoritelist/rpc/internal/svc"
	"MicroTikTok/favoritelist/rpc/pb/favoritelist"
	_ "MicroTikTok/pkg/corn"
	"context"

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
	// todo: add your logic here and delete this line

	return &favoritelist.FavoriteListResponse{}, nil
}
