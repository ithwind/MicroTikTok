// Code generated by goctl. DO NOT EDIT.
package handler

import (
	"net/http"

	"MicroTikTok/feed/api/internal/svc"

	"github.com/zeromicro/go-zero/rest"
)

func RegisterHandlers(server *rest.Server, serverCtx *svc.ServiceContext) {
	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodGet,
				Path:    "/douyin/feed",
				Handler: FeedHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/douyin/publish/action",
				Handler: uploadHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/douyin/publish/list",
				Handler: getPublishHandler(serverCtx),
			},
		},
	)
}
