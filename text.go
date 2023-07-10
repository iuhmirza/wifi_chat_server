package main

import (
	"context"
	"log"
	"net/http"

	"nhooyr.io/websocket"
	"nhooyr.io/websocket/wsjson"
)

var broadcasts = make([]chan textMessage, 0)

func textHandler(w http.ResponseWriter, r *http.Request) {
	wsConn, err := websocket.Accept(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}
	c := make(chan textMessage)
	broadcasts = append(broadcasts, c)
	go clientWriter(wsConn, c)
	go clientReader(wsConn)
}

func clientReader(wsConn *websocket.Conn) {
	for {
		var msg textMessage
		err := wsjson.Read(context.TODO(), wsConn, &msg)
		if err != nil {
			log.Println(err)
			wsConn.Close(websocket.StatusInternalError, "failed to read from ws client")
			return
		}
		for _, v := range broadcasts {
			v <- msg
		}
	}
}

func clientWriter(wsConn *websocket.Conn, broadcast chan textMessage) {
	for {
		message := <-broadcast
		err := wsjson.Write(context.TODO(), wsConn, &message)
		if err != nil {
			log.Println(err)
			wsConn.Close(websocket.StatusInternalError, "failed to write to ws client")
			return
		}
	}
}
