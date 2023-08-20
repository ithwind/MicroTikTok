package logic

import (
	"MicroTikTok/feed/model"
	_ "MicroTikTok/feed/model"
	jwt2 "MicroTikTok/pkg/jwt"
	"MicroTikTok/user/api/user/internal/svc"
	"MicroTikTok/user/api/user/internal/types"
	_ "MicroTikTok/user/model/dao/model"
	"MicroTikTok/user/rpc/rpc/user"
	"context"
	"github.com/zeromicro/go-zero/core/logx"
)

type LoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LoginLogic) Login(req *types.LoginRequest) (resp *types.LoginResponse, err error) {
	// todo: add your logic here and delete this line
	res, err := l.svcCtx.UserRpc.Login(l.ctx, &user.LoginRequest{
		Username: req.Username,
		Password: req.Password,
	})

	table := l.svcCtx.BkModel.User
	userinfo, _ := table.WithContext(l.ctx).Where(table.Username.Eq(req.Username)).First()
	u := model.User{
		ID:              int64(userinfo.ID),
		UserName:        userinfo.Username,
		Avatar:          userinfo.Avatar,
		BackgroundImage: userinfo.BackgroundImage,
		Signature:       userinfo.Signature,
	}
	token := jwt2.GenerateToken(int(userinfo.ID), u)
	return &types.LoginResponse{
		StatusCode: int64(res.StatusCode),
		UserID:     &res.UserId,
		Token:      &token,
		StatusMsg:  res.StatusMsg,
	}, nil
}
