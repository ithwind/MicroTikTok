package logic

import (
	"MicroTikTok/feed/model"
	jwt2 "MicroTikTok/pkg/jwt"
	"MicroTikTok/user/api/user/internal/svc"
	"MicroTikTok/user/api/user/internal/types"
	"MicroTikTok/user/rpc/rpc/user"
	"context"
	"github.com/zeromicro/go-zero/core/logx"
)

type RegisterLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RegisterLogic {
	return &RegisterLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}
func (l *RegisterLogic) Register(req *types.RegisterRequest) (resp *types.RegisterResponse, err error) {
	res, err := l.svcCtx.UserRpc.Register(l.ctx, &user.RegisterRequest{
		Username: req.Username,
		Password: req.Password,
	})
	if err != nil {
		return nil, err
	}
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
	return &types.RegisterResponse{
		StatusCode: int64(res.StatusCode),
		UserID:     res.UserId,
		Token:      token,
		StatusMsg:  res.StatusMsg,
	}, nil
}
