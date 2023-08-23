package postgres

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

var (
	DB  *gorm.DB
	dsn = "user=postgres password=qb030929 dbname=testTikTok port=5432 sslmode=disable TimeZone=Asia/Shanghai"
)

func Init() {
	var err error
	DB, err = gorm.Open(postgres.New(postgres.Config{
		DSN:                  dsn,
		PreferSimpleProtocol: true, // disables implicit prepared statement usage
	}), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{SingularTable: true},
		PrepareStmt:    true,
		Logger:         logger.Default,
	})
	if err != nil {
		panic(err)
	}
}
