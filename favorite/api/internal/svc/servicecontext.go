package svc

import (
	"MicroTikTok/favorite/api/internal/config"
	"MicroTikTok/favorite/rpc/favoriteclient"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config      config.Config
	FavoriteRpc favoriteclient.Favorite
}

func NewServiceContext(c config.Config) *ServiceContext {
	favoriteRpc := favoriteclient.NewFavorite(zrpc.MustNewClient(c.FavoriteRpcConf))
	return &ServiceContext{
		Config:      c,
		FavoriteRpc: favoriteRpc,
	}
}
