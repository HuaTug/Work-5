package task

import (
	"Hertz_refactored/biz/dal/db/chats/db"
	"Hertz_refactored/biz/dal/db/chats/im/mq"
	"Hertz_refactored/biz/model/chat"
	"context"
	"encoding/json"
	"fmt"
	"github.com/hertz-contrib/websocket"
	"github.com/sirupsen/logrus"
	"log"
)

type SendMsg struct {
	Type    int64  `json:"type"`
	Content string `json:"content"`
}
type ReplyMsg struct {
	From    string `json:"form"`
	Code    int    `json:"code"`
	Content string `json:"content"`
}
type Client struct {
	ID     int64
	ToUid  int64
	Socket *websocket.Conn
	Send   chan []byte
}

type Broadcast struct {
	Client  *Client
	Message []byte
	Type    int64
}

// ClientManager Manager client user
type ClientManager struct {
	Clients   map[int64]*Client //manager
	Broadcast chan *Broadcast
	Reply     chan *Client
	Enter     chan *Client //login
	Leave     chan *Client //exit
}

var Manager = ClientManager{
	Clients:   make(map[int64]*Client),
	Broadcast: make(chan *Broadcast),
	Reply:     make(chan *Client),
	Enter:     make(chan *Client),
	Leave:     make(chan *Client),
}

type SyncTask struct {
}

// ToDo 实现对维度消息的一个存储 不需要将其直接操作数据库
func SendToMessage(message *chat.Message) {
	fmt.Println(message.MessageText)
	err := db.CreateMessage(message)
	if err != nil {
		return
	}
}

func (s *SyncTask) RunTaskCreate(ctx context.Context) error {
	rabbitMqQueue := "create_task"
	msgs, err := mq.ConsumeMessage(ctx, rabbitMqQueue)
	if err != nil {
		logrus.Info(err)
		return err
	}
	var forever chan struct{}
	go func() {
		for d := range msgs {
			//ToDo ：这里还要继续进行完善
			reqRabbitMQ := new(chat.Message)
			err := json.Unmarshal(d.Body, reqRabbitMQ)
			if err != nil {
				log.Printf("Received run Task: %s", err)
			}
			SendToMessage(reqRabbitMQ)
			d.Ack(false)
		}
	}()

	logrus.Info(err)
	<-forever

	return nil
}
