package logic

import (
	"MicroTikTok/AcessData/remark"
	user2 "MicroTikTok/AcessData/user"
	"MicroTikTok/Constant"
	"MicroTikTok/pkg/jwt"
	"MicroTikTok/pkg/util"
	"context"
	"fmt"
	"strconv"

	"MicroTikTok/comment/rpc/internal/svc"
	"MicroTikTok/comment/rpc/pb/comment"

	"github.com/zeromicro/go-zero/core/logx"
)

type CommentActionLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCommentActionLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CommentActionLogic {
	return &CommentActionLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CommentActionLogic) CommentAction(in *comment.CommentActionReq) (*comment.CommentActionResp, error) {
	fmt.Println("========================>>", in)
	var c comment.Comment
	claims, err := jwt.ParseToken(in.Token)
	if err != nil {
		return nil, err
	}
	userId := int64(claims.UserId)
	videoId, _ := strconv.Atoi(in.VideoId)
	user := user2.GetUserById(userId)
	/**
	1.判断新增评论或者删除
	2.在数据执行操作（软删除）
	*/
	if in.CommentText != util.String("Default") {
		err, r := remark.CreateRemark(userId, int64(videoId), util.GetString(in.CommentText))
		if err != nil {
			return nil, err
		}
		u := comment.User{
			Id:              userId,
			Name:            user.UserName,
			FollowCount:     util.Int(user2.GetFollowCountByUserId(userId)),
			FollowerCount:   util.Int(user2.GetFollowerCountByUserId(userId)),
			IsFollow:        user2.GetIsFollowByUserId(user2.GetUserIdByVideoId(int64(videoId)), userId),
			Avatar:          util.String(user.Avatar),
			BackgroundImage: util.String(user.BackgroundImage),
			Signature:       util.String(user.Signature),
			TotalFavorited:  util.Int(user2.GetTotalFavoriteCount(userId)),
			WorkCount:       util.Int(user2.GetWorkCountByUserId(userId)),
			FavoriteCount:   util.Int(user2.GetFavoriteCount(userId)),
		}
		c.User = &u
		c.Id = r.ID
		c.Content = r.CommentText
		c.CreateAt = r.CreateAt.Format("01-02")
	} else if in.CommentId != util.Int(Constant.DefaultCommentIDInt) {
		err := remark.DeleteRemark(util.GetInt(in.CommentId))
		if err != nil {
			return nil, err
		}
	}
	fmt.Println("------------------->>", &c)
	return &comment.CommentActionResp{
		Comment: &c,
	}, nil
}
