package logic

import (
	"MicroTikTok/common/cryptx"
	"MicroTikTok/user/rpc/internal/svc"
	"MicroTikTok/user/rpc/rpc/user"
	"context"
	"github.com/zeromicro/go-zero/core/logx"
	"google.golang.org/grpc/status"
	"gorm.io/gorm"
)

type LoginLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *LoginLogic) Login(in *user.LoginRequest) (*user.LoginResponse, error) {
	// todo: add your logic here and delete this line
	table := l.svcCtx.BkModel.User
	res, err := table.WithContext(l.ctx).Where(table.Username.Eq(in.Username)).Debug().First()
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, status.Error(100, "用户不存在")
		}
		return nil, status.Error(500, err.Error())
	}
	password := cryptx.PasswordEncrypt(l.svcCtx.Config.Salt, in.Password)
	if password != res.Password {
		return nil, status.Error(100, "密码错误")
	}
	statusMsg := "登陆成功"
	return &user.LoginResponse{
		UserId:     int64(res.ID),
		StatusCode: 200,
		StatusMsg:  &statusMsg,
	}, nil
}
