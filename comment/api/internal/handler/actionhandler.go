package handler

import (
	"MicroTikTok/Constant"
	"MicroTikTok/pkg/util"
	"net/http"

	"MicroTikTok/comment/api/internal/logic"
	"MicroTikTok/comment/api/internal/svc"
	"MicroTikTok/comment/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func ActionHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.CommentActionReq

		err := httpx.Parse(r, &req)

		if err != nil {
			if req.CommentID == nil {
				req.CommentID = util.String(Constant.DefaultCommentIDString)
			}
			if req.CommentText == nil {
				req.CommentText = util.String("Default")
			}
		}

		l := logic.NewActionLogic(r.Context(), svcCtx)
		resp, err := l.Action(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
