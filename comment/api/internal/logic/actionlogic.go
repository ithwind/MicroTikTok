package logic

import (
	"MicroTikTok/Constant"
	"MicroTikTok/comment/api/internal/svc"
	"MicroTikTok/comment/api/internal/types"
	"MicroTikTok/comment/rpc/pb/comment"
	"MicroTikTok/pkg/util"
	"context"
	"fmt"
	"github.com/zeromicro/go-zero/core/logx"
	"strconv"
)

type ActionLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewActionLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ActionLogic {
	return &ActionLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ActionLogic) Action(req *types.CommentActionReq) (*types.CommentActionResp, error) {
	var request comment.CommentActionReq
	var response types.CommentActionResp
	commentId, _ := strconv.Atoi(util.GetString(req.CommentID))
	request.Token = req.Token
	request.VideoId = req.VideoId
	request.CommentText = req.CommentText
	request.CommentId = util.Int(int64(commentId))
	fmt.Println("==============================", &request)
	resp, err := l.svcCtx.CommentRpc.CommentAction(l.ctx, &request)
	if err != nil {
		return nil, err
	}

	var r types.Comment
	if req.ActionType == 1 {
		user := resp.GetComment().User
		u := types.User{
			Id:              user.Id,
			Name:            user.Name,
			FollowCount:     util.GetInt(user.FollowCount),
			FollowerCount:   util.GetInt(user.FollowerCount),
			IsFollow:        user.IsFollow,
			Avatar:          util.GetString(user.Avatar),
			BackgroundImage: util.GetString(user.BackgroundImage),
			Signature:       util.GetString(user.Signature),
			TotalFavorited:  util.GetInt(user.TotalFavorited),
			WorkCount:       util.GetInt(user.WorkCount),
			FavoriteCount:   util.GetInt(user.FavoriteCount),
		}

		r = types.Comment{
			Id:       resp.Comment.Id,
			U:        u,
			Content:  resp.Comment.Content,
			CreateAt: resp.Comment.CreateAt,
		}
	}

	response.StatusCode = Constant.StatusHttpOk
	response.StatusMsg = "发布或删除评论成功"
	response.Comment = r
	return &response, nil
}
