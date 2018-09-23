package main

import "C"
import (
	"github.com/gin-gonic/gin"
	"github.com/gofrs/uuid"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
)

type Client struct {
	ID   string
	Conn *websocket.Conn
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

var ps PubSub

func WsHandlerFunc(c *gin.Context) {
	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)

	if err != nil {
		log.Println(err)
		return
	}

	id := uuid.Must(uuid.NewV4()).String()

	client := Client{
		ID:   id,
		Conn: conn,
	}

	ps.add(client)

	conn.WriteMessage(1, []byte("連線成功:"+client.ID))

	run(client)
}

func run(client Client) {
	defer client.Conn.Close()

	var msg Message

	for {
		err := client.Conn.ReadJSON(&msg)

		if err != nil {
			log.Println("讀取失敗ID:%s", client.ID)
			log.Println("失敗原因", err)
			ps.remove(client)
			return
		}

		ps.HandleReceiveMessage(client, msg)
	}
}
