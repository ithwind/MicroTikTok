package logic

import (
	"MicroTikTok/pkg/jwt"
	"MicroTikTok/user/api/user/internal/svc"
	"MicroTikTok/user/api/user/internal/types"
	"MicroTikTok/user/rpc/rpc/user"
	"context"
	"fmt"
	"github.com/zeromicro/go-zero/core/logx"
)

type UserInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserInfoLogic {
	return &UserInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserInfoLogic) UserInfo(req *types.UserInfoRequest) (resp *types.UserInfoResponse, err error) {
	// todo: add your logic here and delete this line
	claims, _ := jwt.ParseToken(req.Token)
	fmt.Println(claims.UserId + 1234567890)
	res, err := l.svcCtx.UserRpc.UserInfo(l.ctx, &user.UserInfoRequest{
		UserId: int64(claims.UserId),
		Token:  req.Token,
	})

	if err != nil {
		return nil, err
	}
	return &types.UserInfoResponse{
		StatusMsg:  res.StatusMsg,
		StatusCode: int64(res.StatusCode),
		User: &types.User{
			Avatar:          *res.UserInfo.Avatar,
			BackgroundImage: *res.UserInfo.BackgroundImage,
			FollowCount:     *res.UserInfo.FollowCount,
			FavoriteCount:   *res.UserInfo.FavoriteCount,
			FollowerCount:   *res.UserInfo.FollowerCount,
			ID:              res.UserInfo.Id,
			Name:            res.UserInfo.Name,
			Signature:       *res.UserInfo.Signature,
			TotalFavorited:  *res.UserInfo.TotalFavorited,
			WorkCount:       *res.UserInfo.WorkCount,
		},
	}, nil
}
