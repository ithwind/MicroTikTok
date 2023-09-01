package svc

import (
	"MicroTikTok/comment/api/internal/config"
	"MicroTikTok/comment/rpc/commentaction"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config     config.Config
	CommentRpc commentaction.CommentAction
}

func NewServiceContext(c config.Config) *ServiceContext {
	commentRpc := commentaction.NewCommentAction(zrpc.MustNewClient(c.CommentRpcConf))
	return &ServiceContext{
		Config:     c,
		CommentRpc: commentRpc,
	}
}
