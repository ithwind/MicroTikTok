package main

import (
	"MicroTikTok/AcessData/mysql"
	"flag"
	"fmt"

	"MicroTikTok/chat/rpc/internal/config"
	"MicroTikTok/chat/rpc/internal/server"
	"MicroTikTok/chat/rpc/internal/svc"
	"MicroTikTok/chat/rpc/pb/chat"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var configFile = flag.String("f", "etc/chat.yaml", "the config file")

func main() {
	//go service2.StartPolling()
	flag.Parse()
	mysql.Init()
	var c config.Config
	conf.MustLoad(*configFile, &c)
	ctx := svc.NewServiceContext(c)

	s := zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
		chat.RegisterChatServer(grpcServer, server.NewChatServer(ctx))

		if c.Mode == service.DevMode || c.Mode == service.TestMode {
			reflection.Register(grpcServer)
		}
	})
	defer s.Stop()

	fmt.Printf("Starting rpc server at %s...\n", c.ListenOn)
	s.Start()
}
