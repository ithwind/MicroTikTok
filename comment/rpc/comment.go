package main

import (
	"MicroTikTok/AcessData/mysql"
	"flag"
	"fmt"

	"MicroTikTok/comment/rpc/internal/config"
	"MicroTikTok/comment/rpc/internal/server"
	"MicroTikTok/comment/rpc/internal/svc"
	"MicroTikTok/comment/rpc/pb/comment"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var configFile = flag.String("f", "etc/comment.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	mysql.Init()
	conf.MustLoad(*configFile, &c)
	ctx := svc.NewServiceContext(c)

	s := zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
		comment.RegisterCommentActionServer(grpcServer, server.NewCommentActionServer(ctx))
		if c.Mode == service.DevMode || c.Mode == service.TestMode {
			reflection.Register(grpcServer)
		}
	})
	defer s.Stop()

	fmt.Printf("Starting rpc server at %s...\n", c.ListenOn)
	s.Start()
}
