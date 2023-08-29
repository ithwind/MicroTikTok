package handler

import (
	"net/http"

	"MicroTikTok/favoritelist/api/internal/logic"
	"MicroTikTok/favoritelist/api/internal/svc"
	"MicroTikTok/favoritelist/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func FavoritelistHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.FavoritelistRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewFavoritelistLogic(r.Context(), svcCtx)
		resp, err := l.Favoritelist(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
