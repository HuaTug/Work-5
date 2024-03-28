package mq

import (
	"context"

	"github.com/sirupsen/logrus"
	"github.com/streadway/amqp"

	"Hertz_refactored/biz/dal/db"
)

func ConsumeMessage(ctx context.Context, queueName string) (<-chan amqp.Delivery, error) {
	ch, err := db.RabbitMq.Channel()
	if err != nil {
		logrus.Info(err)
	}
	q, _ := ch.QueueDeclare(queueName, true, false, false, false, nil)
	err = ch.Qos(1, 0, false)
	return ch.Consume(q.Name, "", false, false, false, false, nil)
}
