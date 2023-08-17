package svc

import (
	"MicroTikTok/feed/api/internal/config"
	"MicroTikTok/feed/rpc/feed"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config   config.Config
	VideoRpc feed.Feed
}

func NewServiceContext(c config.Config) *ServiceContext {
	videoRpc := feed.NewFeed(zrpc.MustNewClient(c.VideoRpcConf))
	return &ServiceContext{
		Config:   c,
		VideoRpc: videoRpc,
	}
}
