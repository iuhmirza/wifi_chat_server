package main

import (
	"context"
	"log"
	"net/http"

	"nhooyr.io/websocket"
	"nhooyr.io/websocket/wsjson"
)

var clients = make(map[*client]bool)

func textHandler(w http.ResponseWriter, r *http.Request) {
	wsConn, err := websocket.Accept(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}
	c := &client{Conn: wsConn, Send: make(chan *textMessage)}
	clients[c] = true
	go c.clientReader()
	go c.clientWriter()
}

func (c *client) clientReader() {
	for {
		var msg *textMessage = &textMessage{}
		err := wsjson.Read(context.TODO(), c.Conn, msg)
		if err != nil {
			log.Println(err)
			delete(clients, c)
			c.Conn.Close(websocket.StatusInternalError, err.Error())
			return
		}
		for c := range clients {
			c.Send <- msg
		}
	}
}

func (c *client) clientWriter() {
	for {
		msg := <-c.Send
		err := wsjson.Write(context.TODO(), c.Conn, msg)
		if err != nil {
			log.Println(err)
			delete(clients, c)
			c.Conn.Close(websocket.StatusInternalError, err.Error())
			return
		}
	}
}
