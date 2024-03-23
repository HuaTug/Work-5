package im

import (
	"Hertz_refactored/biz/dal/db/chats/db"
	"Hertz_refactored/biz/dal/db/chats/im/mq"
	"Hertz_refactored/biz/model/chat"
	e "Hertz_refactored/biz/pkg"
	"encoding/json"
	"github.com/hertz-contrib/websocket"
	"github.com/sirupsen/logrus"
	"log"
	"time"
)

func (manager *ClientManager) Listen() {
	for {
		log.Println("监听管道通信")

		select {

		case client := <-Manager.Enter:
			log.Printf("%v:online\n", client.ID)
			Manager.Clients[client.ID] = client //把连接放到用户管理上
			/*
				baseResp := pack.BuildChatBaseResp(errno.WebSocketSuccess)
				resp, _ := json.Marshal(baseResp)
				_ = client.Socket.WriteMessage(websocket.TextMessage, resp)*/
			replyMsg := &chat.BaseResp{
				Code: 200,
				Msg:  "成功连接至服务器",
			}
			msg, _ := json.Marshal(replyMsg)
			_ = client.Socket.WriteMessage(websocket.TextMessage, msg)
		case client := <-Manager.Leave:
			log.Printf("%v:offline\n", client.ID)
			replyMsg := &chat.BaseResp{
				Code: 200,
				Msg:  "成功断掉服务器",
			}
			msg, _ := json.Marshal(replyMsg)
			_ = client.Socket.WriteMessage(websocket.TextMessage, msg)
			close(client.Send)                 //close chan
			delete(Manager.Clients, client.ID) //delete map
		case broadcast := <-Manager.Broadcast:
			//ToDo:对这里进行完善
			message := broadcast.Message
			touid := broadcast.Client.ToUid
			//每一个用户都有自己id和to_id，这便是能够使得两个用户进行通信的关键因素
			flag := false
			//这是一个广播，广播的内容就是要去将一个消息发送给特定的用户，所以有了id==touid的判断逻辑
			for id, conn := range Manager.Clients {
				if id != touid {
					continue
				}
				//这时候表示在进行循环过后，发送消息方找到了接收消息方，此时接受方可以打开其消息通道进行消息的收发
				select {
				case conn.Send <- message:
					flag = true
					//找到接收方 有用户在线 改flag标志
				default:
					close(conn.Send)
					delete(Manager.Clients, conn.ID)
				}
			}
			id := broadcast.Client.ID
			if flag {
				logrus.Info("对方在线应答")
				replyMsg := &ReplyMsg{
					Code:    e.WebsocketOnlineReply,
					Content: "对方在线应答",
				}
				msg, err := json.Marshal(replyMsg)
				_ = broadcast.Client.Socket.WriteMessage(websocket.TextMessage, msg)
				messages := &chat.Message{
					SenderID:    id,
					ReceiverID:  touid,
					MessageText: string(message),
					SendTime:    time.Now().Format(time.DateTime),
					State:       1,
				}
				err = db.CreateMessage(messages)
				if err != nil {
					logrus.Info(err)
				}
			} else {
				logrus.Info("对方不在线")
				replyMsg := &ReplyMsg{
					Code:    e.WebsocketOfflineReply,
					Content: "对方不在线",
				}
				msg, err := json.Marshal(replyMsg)
				_ = broadcast.Client.Socket.WriteMessage(websocket.TextMessage, msg)
				messages := &chat.Message{
					SenderID:    id,
					ReceiverID:  touid,
					MessageText: string(message),
					SendTime:    time.Now().Format(time.DateTime),
					State:       0,
				}
				Msg, err := json.Marshal(messages)
				if err != nil {
					logrus.Info("Error: ", err)
				}
				err = mq.SendMessageMQ(Msg)
				if err != nil {
					logrus.Info(err)
				}
				if err != nil {
					logrus.Info(err)
				}
			}
		}
	}
}
