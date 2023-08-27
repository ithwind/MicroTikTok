package handler

import (
	"MicroTikTok/chat/api/internal/logic"
	"MicroTikTok/chat/api/internal/svc"
	"MicroTikTok/chat/api/internal/types"
	"fmt"
	"github.com/zeromicro/go-zero/rest/httpx"
	"net/http"
	"net/url"
)

func chatActionHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.ChatActionRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}
		fmt.Println(url.QueryEscape(req.Content))
		req.Content = url.QueryEscape(req.Content)
		fmt.Println("ParamContent:", req.Content)
		l := logic.NewChatActionLogic(r.Context(), svcCtx)
		resp, err := l.ChatAction(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
