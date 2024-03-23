package mq

import (
	"github.com/sirupsen/logrus"
	"github.com/streadway/amqp"
)

var RabbitMq *amqp.Connection

func InitRabbitMQ() {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		logrus.Info(err)
	}
	if err != nil {
		logrus.Info(err)
	}
	RabbitMq = conn
	logrus.Info("Mq连接成功")
}
