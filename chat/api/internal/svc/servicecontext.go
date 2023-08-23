package svc

import (
	"MicroTikTok/chat/api/internal/config"
	"MicroTikTok/chat/rpc/chatclient"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config  config.Config
	ChatRpc chatclient.Chat
}

func NewServiceContext(c config.Config) *ServiceContext {
	chatRpc := chatclient.NewChat(zrpc.MustNewClient(c.ChatRpcConf))
	return &ServiceContext{
		Config:  c,
		ChatRpc: chatRpc,
	}
}
