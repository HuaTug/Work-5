package mq

import (
	"context"
	"github.com/sirupsen/logrus"
	"github.com/streadway/amqp"
)

func ConsumeMessage(ctx context.Context, queueName string) (<-chan amqp.Delivery, error) {
	ch, err := RabbitMq.Channel()
	if err != nil {
		logrus.Info(err)
	}
	q, _ := ch.QueueDeclare(queueName, true, false, false, false, nil)
	err = ch.Qos(1, 0, false)
	return ch.Consume(q.Name, "", false, false, false, false, nil)
}
