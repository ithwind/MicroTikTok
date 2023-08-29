package svc

import (
	"MicroTikTok/favoritelist/api/internal/config"
	"MicroTikTok/favoritelist/rpc/favoritelistclient"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config          config.Config
	FavoriteListRpc favoritelistclient.FavoriteList
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:          c,
		FavoriteListRpc: favoritelistclient.NewFavoriteList(zrpc.MustNewClient(c.FavoriteListRpc)),
	}
}
