package main

import (
	"MicroTikTok/AcessData/mysql"
	"flag"
	"fmt"

	"MicroTikTok/favorite/rpc/internal/config"
	"MicroTikTok/favorite/rpc/internal/server"
	"MicroTikTok/favorite/rpc/internal/svc"
	"MicroTikTok/favorite/rpc/pb/favorite"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var configFile = flag.String("f", "etc/favorite.yaml", "the config file")

func main() {
	flag.Parse()
	mysql.Init()
	var c config.Config
	conf.MustLoad(*configFile, &c)
	ctx := svc.NewServiceContext(c)
	s := zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
		favorite.RegisterFavoriteServer(grpcServer, server.NewFavoriteServer(ctx))

		if c.Mode == service.DevMode || c.Mode == service.TestMode {
			reflection.Register(grpcServer)
		}
	})
	defer s.Stop()

	fmt.Printf("Starting rpc server at %s...\n", c.ListenOn)
	s.Start()
}
