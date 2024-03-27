package db

import (
	"Hertz_refactored/biz/config"
	"Hertz_refactored/biz/pkg/utils"
	"context"
	"github.com/sirupsen/logrus"
	"github.com/streadway/amqp"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var Db *gorm.DB
var (
	MongoDBClient *mongo.Client
	RabbitMq      *amqp.Connection
)

func Init() {
	var err error
	dsn := utils.GetMysqlDsn()
	Db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		SkipDefaultTransaction: true,
		PrepareStmt:            true,
		Logger:                 logger.Default,
	})
	if err != nil {
		panic(err)
	}
	logrus.Info("数据库连接成功")
	MongoDB()
}
func MongoDB() {
	// 设置mongoDB客户端连接信息
	clientOptions := options.Client().ApplyURI("mongodb://" + config.ConfigInfo.MongoDb.Addr)
	var err error
	MongoDBClient, err = mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		logrus.Info(err)
	}
	err = MongoDBClient.Ping(context.TODO(), nil)
	if err != nil {
		logrus.Info(err)
	}
	logrus.Info("MongoDB Connect")
	InitRabbitMQ()
}
func InitRabbitMQ() {
	conn, err := amqp.Dial("amqp://" + config.ConfigInfo.RabbitMq.Username + ":" + config.ConfigInfo.RabbitMq.Password + "@" + config.ConfigInfo.RabbitMq.Addr + "/")
	if err != nil {
		logrus.Info(err)
	}
	if err != nil {
		logrus.Info(err)
	}
	RabbitMq = conn
	logrus.Info("Mq连接成功")
}
