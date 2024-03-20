// Code generated by hertz generator.

package chat

import (
	chats "Hertz_refactored/biz/dal/db/chats/im"
	"Hertz_refactored/biz/model/chat"
	"context"
	"fmt"
	"github.com/cloudwego/hertz/pkg/app"
	_ "github.com/cloudwego/hertz/pkg/app/server"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"github.com/hertz-contrib/websocket"
	"github.com/sirupsen/logrus"

	"log"
)

var upgrader = websocket.HertzUpgrader{
	CheckOrigin: func(ctx *app.RequestContext) bool {
		return true
	},
}

// Chat .
// @router /chats [GET]
func Chat(_ context.Context, c *app.RequestContext) {
	var err error
	var req chat.MessageChatRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}
	userId, _ := c.Get("user_id")
	var id int64
	if v, ok := userId.(float64); ok {
		id = int64(v)
	} else {
		logrus.Info("数据转换出错")
		return
	}
	fmt.Println(id)
	client := new(chats.Client)
	err = upgrader.Upgrade(c, func(conn *websocket.Conn) {
		for {
			client = &chats.Client{
				ID:     id,
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
		return
	}
	/*
		c.JSON(consts.StatusBadRequest, "Msg")
		不要引入这段代码 否则无法进行websocket通信
		这段代码执行后会退出函数。在 Go 中，当函数调用了 return 语句后，函数的执行将立即终止，并且将返回值（如果有的话）传递给调用者。
	*/
}
