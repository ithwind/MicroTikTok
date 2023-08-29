package main

import (
	"MicroTikTok/AcessData/mysql"
	"flag"
	"fmt"

	"MicroTikTok/favoritelist/rpc/internal/config"
	"MicroTikTok/favoritelist/rpc/internal/server"
	"MicroTikTok/favoritelist/rpc/internal/svc"
	"MicroTikTok/favoritelist/rpc/pb/favoritelist"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var configFile = flag.String("f", "etc/favoritelist.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	mysql.Init()
	conf.MustLoad(*configFile, &c)
	ctx := svc.NewServiceContext(c)

	s := zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
		favoritelist.RegisterFavoriteListServer(grpcServer, server.NewFavoriteListServer(ctx))

		if c.Mode == service.DevMode || c.Mode == service.TestMode {
			reflection.Register(grpcServer)
		}
	})
	defer s.Stop()

	fmt.Printf("Starting rpc server at %s...\n", c.ListenOn)
	s.Start()
}
