package im //nolint:gofmt

import (
	"encoding/json"
	"fmt"
	"log"
	"strconv"

	"github.com/hertz-contrib/websocket"
	"github.com/sirupsen/logrus"

	"Hertz_refactored/biz/dal/db"
	"Hertz_refactored/biz/model/chat"
	e "Hertz_refactored/biz/pkg"
)

func (c *Client) Write() {
	defer func() {
		Manager.Leave <- c
		_ = c.Socket.Close()
	}()

	for {
		sendMsg := new(SendMsg)
		err := c.Socket.ReadJSON(&sendMsg)
		if err != nil {
			logrus.Info("数据格式不正确+", err)
			break
		}
		if sendMsg.Type == 1 {
			logrus.Info(c.ID, "：发送消息->", sendMsg.Content)
			Manager.Broadcast <- &Broadcast{
				Client:  c,
				Message: []byte(sendMsg.Content),
			}
		} else if sendMsg.Type == 2 {
			logrus.Info(c.ID, "：正在获取历史未读信息.....")
			results, err := db.GetMessage(c.ID, c.ToUid)
			if err != nil {
				logrus.Info(err)
				return
			}
			for _, result := range results {
				replyMsg := &ReplyMsg{
					From:    strconv.FormatInt(c.ToUid, 10),
					Code:    e.WebsocketSuccessMessage,
					Content: fmt.Sprintf("%s", result.MessageText),
				}
				msg, _ := json.Marshal(replyMsg)
				_ = c.Socket.WriteMessage(websocket.TextMessage, msg)
			}
		} else if sendMsg.Type == 3 {
			logrus.Info(c.ID, "获取全部的信息.....")
			results, err := db.GetAllMessage()
			if err != nil {
				logrus.Info(err)
				return
			}
			for _, result := range results {
				replyMsg := &ReplyMsg{
					From:    strconv.FormatInt(c.ToUid, 10),
					Code:    e.WebsocketSuccess,
					Content: fmt.Sprintf("%s", result.MessageText),
				}
				msg, _ := json.Marshal(replyMsg)
				_ = c.Socket.WriteMessage(websocket.TextMessage, msg)
			}
		}
	}
}

func (c *Client) Read() {
	defer func() {
		_ = c.Socket.Close()
	}()
	for {
		select {
		case message, ok := <-c.Send:
			if !ok {
				_ = c.Socket.WriteMessage(websocket.TextMessage, message)
				return
			}
			log.Println(c.ID, "接受消息:", string(message))
			replyMsg := &chat.ReplyMsg{
				Code:    e.WebsocketSuccessMessage,
				Content: fmt.Sprintf("%s", string(message)),
			}
			msg, _ := json.Marshal(replyMsg)
			_ = c.Socket.WriteMessage(websocket.TextMessage, msg)
		}
	}
}
