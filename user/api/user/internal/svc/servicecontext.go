package svc

import (
	"MicroTikTok/user/api/user/internal/config"
	"MicroTikTok/user/model/dao/query"
	"MicroTikTok/user/rpc/rpc/user"
	"MicroTikTok/user/rpc/userinfo"
	"github.com/zeromicro/go-zero/zrpc"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type ServiceContext struct {
	Config  config.Config
	UserRpc user.UserInfoClient
	BkModel *query.Query
}

func NewServiceContext(c config.Config) *ServiceContext {
	db, err := gorm.Open(mysql.Open(c.DataSource), &gorm.Config{})
	//如果出错就GameOver了
	if err != nil {
		panic(err)
	}
	return &ServiceContext{
		Config:  c,
		UserRpc: userinfo.NewUserInfo(zrpc.MustNewClient(c.UserRpc)),
		BkModel: query.Use(db),
	}
}
