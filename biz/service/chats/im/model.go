package im //nolint:gofmt

import "github.com/hertz-contrib/websocket"

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
