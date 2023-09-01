package logic

import (
	"MicroTikTok/comment/api/internal/svc"
	"MicroTikTok/comment/api/internal/types"
	"MicroTikTok/comment/rpc/pb/comment"
	"MicroTikTok/pkg/util"
	"context"
	"github.com/zeromicro/go-zero/core/logx"
)

type CommentLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCommentLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CommentLogic {
	return &CommentLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CommentLogic) Comment(req *types.CommentListReq) (*types.CommentListResp, error) {
	// todo: add your logic here and delete this line
	var request comment.CommentListReq
	var resp types.CommentListResp
	request.Token = req.Token
	request.VideoId = req.VideoId
	response, err := l.svcCtx.CommentRpc.CommentList(l.ctx, &request)
	if err != nil {
		return nil, err
	}
	//转换评论传输形式
	var remarkList = make([]types.Comment, 0)
	for _, c := range response.CommentList {
		u := c.User
		r := types.Comment{
			Id: c.Id,
			U: types.User{
				Id:              u.Id,
				Name:            u.Name,
				FollowCount:     util.GetInt(u.FollowCount),
				FollowerCount:   util.GetInt(u.FollowerCount),
				IsFollow:        u.IsFollow,
				Avatar:          util.GetString(u.Avatar),
				BackgroundImage: util.GetString(u.BackgroundImage),
				Signature:       util.GetString(u.Signature),
				TotalFavorited:  util.GetInt(u.TotalFavorited),
				WorkCount:       util.GetInt(u.WorkCount),
				FavoriteCount:   util.GetInt(u.FavoriteCount),
			},
			Content:  c.Content,
			CreateAt: c.CreateAt,
		}
		remarkList = append(remarkList, r)
	}

	resp.StatusCode = 0
	resp.StatusMsg = "获取评论成功"
	resp.CommentList = remarkList
	return &resp, nil
}
