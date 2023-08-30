package svc

import (
	"MicroTikTok/user/model/dao/query"
	"MicroTikTok/user/rpc/internal/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type ServiceContext struct {
	Config  config.Config
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
		BkModel: query.Use(db),
	}
}
