package chats

import (
	"Hertz_refactored/biz/model/chat"
	chats "Hertz_refactored/biz/service/chats/im"
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/hertz-contrib/websocket"
	"log"
)

type ChatService struct {
	ctx context.Context
}

func NewService(ctx context.Context) *ChatService {
	return &ChatService{ctx: ctx}
}

var upgrader = websocket.HertzUpgrader{
	CheckOrigin: func(ctx *app.RequestContext) bool {
		return true
	},
}

func (s *ChatService) ChatInit(req chat.MessageChatRequest, c *app.RequestContext, userId int64) error {
	client := new(chats.Client)
	err := upgrader.Upgrade(c, func(conn *websocket.Conn) {
		for {
			client = &chats.Client{
				ID:     userId,
				ToUid:  req.ToID,
				Socket: conn,
				Send:   make(chan []byte, 8),
			}
			chats.Manager.Enter <- client
			go client.Read()
			go client.Write()
			forever := make(chan bool)
			<-forever
		}
	})
	if err != nil {
		log.Println("upgrade:", err)
		return err
	}
	return nil
}
