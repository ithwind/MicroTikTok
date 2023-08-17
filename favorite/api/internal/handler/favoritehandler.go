package handler

import (
	"context"
	"net/http"
	"time"

	"MicroTikTok/favorite/api/internal/logic"
	"MicroTikTok/favorite/api/internal/svc"
	"MicroTikTok/favorite/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func favoriteHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.FavoriteRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}
		ctx, cancel := context.WithTimeout(r.Context(), 10*time.Minute)
		defer cancel()
		l := logic.NewFavoriteLogic(ctx, svcCtx)
		resp, err := l.Favorite(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
