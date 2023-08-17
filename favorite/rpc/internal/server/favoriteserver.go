// Code generated by goctl. DO NOT EDIT.
// Source: favorite.proto

package server

import (
	"context"
	"time"

	"MicroTikTok/favorite/rpc/internal/logic"
	"MicroTikTok/favorite/rpc/internal/svc"
	"MicroTikTok/favorite/rpc/pb/favorite"
)

type FavoriteServer struct {
	svcCtx *svc.ServiceContext
	favorite.UnimplementedFavoriteServer
}

func NewFavoriteServer(svcCtx *svc.ServiceContext) *FavoriteServer {
	return &FavoriteServer{
		svcCtx: svcCtx,
	}
}

func (s *FavoriteServer) Favorite(ctx context.Context, in *favorite.FavoriteActionRequest) (*favorite.FavoriteActionResponse, error) {
	c, cancel := context.WithTimeout(ctx, 3*time.Minute)
	defer cancel()
	l := logic.NewFavoriteLogic(c, s.svcCtx)
	return l.Favorite(in)
}