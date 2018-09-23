package main

import (
	"github.com/gin-gonic/gin"
	"github.com/gofrs/uuid"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func WsHandlerFunc(c *gin.Context) {
	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)

	if err != nil {
		log.Println(err)
		return
	}

	id := uuid.Must(uuid.NewV4())

	conn.WriteMessage(1, []byte("連線成功:" + id.String()))
}
