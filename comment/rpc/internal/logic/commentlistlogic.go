package logic

import (
	"MicroTikTok/AcessData/remark"
	"MicroTikTok/AcessData/user"
	"MicroTikTok/pkg/jwt"
	"context"
	"fmt"
	"strconv"

	"MicroTikTok/comment/rpc/internal/svc"
	"MicroTikTok/comment/rpc/pb/comment"

	"github.com/zeromicro/go-zero/core/logx"
)

type CommentListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCommentListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CommentListLogic {
	return &CommentListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CommentListLogic) CommentList(in *comment.CommentListReq) (*comment.CommentListResp, error) {
	fmt.Println("Param:", in.Token, in.VideoId)
	/**
	1.解析token获取当前用户Id
	2.查询comment
	*/
	claims, err := jwt.ParseToken(in.Token)
	if err != nil {
		fmt.Println("Error:", err)
		return nil, err
	}
	userId := int64(claims.UserId)
	videoId, _ := strconv.Atoi(in.VideoId)
	fmt.Println("========================>>Param:", userId, videoId)
	remarks := remark.GetRemarkListByVideoIdAndUserId(int64(videoId), userId)
	fmt.Println("=====================>>remarks:", remarks)
	var commentList = make([]*comment.Comment, 0)

	for _, r := range remarks {
		var com comment.Comment
		u := user.GetUserById(userId)
		followCount := user.GetFollowCountByUserId(userId)
		followerCount := user.GetFollowerCountByUserId(userId)
		avatar := u.Avatar
		backgroundImg := u.BackgroundImage
		signature := u.Signature
		totalFavorited := user.GetTotalFavoriteCount(userId)
		workOut := user.GetWorkCountByUserId(userId)
		favoriteCount := user.GetFavoriteCount(userId)

		com.Id = r.ID
		com.Content = r.CommentText
		com.CreateAt = r.CreateAt.Format("01-02")
		com.User = &comment.User{
			Id:              userId,
			Name:            u.UserName,
			FollowCount:     &followCount,
			FollowerCount:   &followerCount,
			IsFollow:        user.GetIsFollowByUserId(userId, int64(videoId)),
			Avatar:          &avatar,
			BackgroundImage: &backgroundImg,
			Signature:       &signature,
			TotalFavorited:  &totalFavorited,
			WorkCount:       &workOut,
			FavoriteCount:   &favoriteCount,
		}
		commentList = append(commentList, &com)
	}
	fmt.Println("========================>>List: ", commentList)
	return &comment.CommentListResp{
		CommentList: commentList,
	}, nil
}
