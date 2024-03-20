package mysql

import (
	"github.com/sirupsen/logrus"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var dsn = "root:123456@tcp(localhost:3306)/Hertz?charset=utf8&parseTime=True&loc=Local"

var Db *gorm.DB

func Init() {
	var err error
	Db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		SkipDefaultTransaction: true,
		PrepareStmt:            true,
		Logger:                 logger.Default,
	})
	if err != nil {
		panic(err)
	}
	logrus.Info("数据库连接成功")
}
