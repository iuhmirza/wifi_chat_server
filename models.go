package main

import "nhooyr.io/websocket"

type client struct {
	Conn *websocket.Conn
	Send chan *textMessage
}

type textMessage struct {
	Username string `json:"username"`
	Content  string `json:"content"`
}
