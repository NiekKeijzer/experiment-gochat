package controllers

import (
	"chat/service/chat"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

const anonymousName = "anonymous"

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func ServeWS(hub *chat.Hub, ctx *gin.Context) {
	conn, err := upgrader.Upgrade(ctx.Writer, ctx.Request, nil)
	if err != nil {
		log.Printf("error: %v", err)

		return
	}

	client := hub.NewClient(conn)

	go client.ReadPump()
	go client.WritePump()
}
