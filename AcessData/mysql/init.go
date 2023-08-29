package mysql

import (
	"gorm.io/driver/mysql"
	//"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

var (
	DB *gorm.DB
	//mysql
	dsn = "root:qb030929@tcp(localhost:3306)/tiktok?parseTime=true&loc=Asia%2FShanghai"
	//dsn = "user=mysql password=qb030929 dbname=testTikTok port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	/*host     = os.Getenv("MYSQL_HOST")
	name     = os.Getenv("MYSQL_USER")
	password = os.Getenv("MYSQL_PASSWORD")
	port     = os.Getenv("MYSQL_PORT")
	dbName   = "tiktok"
	dsn      = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true&loc=Asia%%2FShanghai", name, password, host, port, dbName)*/
)

func Init() {
	var err error
	DB, err = gorm.Open(mysql.New(mysql.Config{
		DSN: dsn,
	}), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{SingularTable: true},
		PrepareStmt:    true,
		Logger:         logger.Default,
	})
	if err != nil {
		panic(err)
	}
}
