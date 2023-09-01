package handler

import (
	"MicroTikTok/user/api/user/internal/logic"
	"MicroTikTok/user/api/user/internal/svc"
	"MicroTikTok/user/api/user/internal/types"
	"net/http"
	"time"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func LoginHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.LoginRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewLoginLogic(r.Context(), svcCtx)
		resp, err := l.Login(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
		token := resp.Token
		expires := time.Now().Add(24 * time.Hour) // 一天后的时间
		http.SetCookie(w, &http.Cookie{
			Name:    "token",
			Value:   *token,
			Expires: expires,
		})
		httpx.OkJsonCtx(r.Context(), w, resp)
	}
}
