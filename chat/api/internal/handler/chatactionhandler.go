package handler

import (
	"net/http"

	"MicroTikTok/chat/api/internal/logic"
	"MicroTikTok/chat/api/internal/svc"
	"MicroTikTok/chat/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func chatActionHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.ChatActionRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewChatActionLogic(r.Context(), svcCtx)
		resp, err := l.ChatAction(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
