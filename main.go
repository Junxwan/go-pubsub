package main

import "github.com/gin-gonic/gin"

func main() {
	route := gin.Default()

	route.StaticFile("/", "./index.html")
	route.GET("/ws", WsHandlerFunc)

	route.Run()
}
