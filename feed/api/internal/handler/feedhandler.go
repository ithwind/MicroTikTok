package handler

import (
	"MicroTikTok/feed/api/internal/logic"
	"MicroTikTok/feed/api/internal/svc"
	"MicroTikTok/feed/api/internal/types"
	"fmt"
	"github.com/zeromicro/go-zero/rest/httpx"
	"net/http"
)

func FeedHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.FeedRequest
		err := httpx.Parse(r, &req)
		if err != nil {
			if req.LastTime == nil {
				var defaultTime int64 = 0
				req.LastTime = &defaultTime
			}
			if req.Token == nil {
				var defaultToken = "Default"
				req.Token = &defaultToken
			}
		}

		l := logic.NewFeedLogic(r.Context(), svcCtx)
		resp, err := l.Feed(&req)
		if err != nil {
			fmt.Printf("Error : %v", err)
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			fmt.Printf("Error : %v", err)
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
