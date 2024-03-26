package mq

import (
	"Hertz_refactored/biz/dal/db"
	"github.com/sirupsen/logrus"
	"github.com/streadway/amqp"
)

func SendMessageMQ(body []byte) (err error) {
	ch, err := db.RabbitMq.Channel()
	if err != nil {
		logrus.Info(err)
	}
	q, _ := ch.QueueDeclare("create_task", true, false, false, false, nil)
	err = ch.Publish("", q.Name, false, false, amqp.Publishing{
		DeliveryMode: amqp.Persistent,
		ContentType:  "application/json",
		Body:         body,
	})
	if err != nil {
		logrus.Info(err)
		return err
	}
	logrus.Info("发送MQ成功...")
	return
}
