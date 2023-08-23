package handler

import (
	"net/http"

	"MicroTikTok/chat/api/internal/logic"
	"MicroTikTok/chat/api/internal/svc"
	"MicroTikTok/chat/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func chatMessageHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.ChatMessageRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewChatMessageLogic(r.Context(), svcCtx)
		resp, err := l.ChatMessage(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
