package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"github.com/suttapak/cacf/services"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

type ws struct {
	h services.HubServices
}

func NewWsServer(h services.HubServices) Ws {
	return &ws{h}
}
func (ws *ws) Serve(c *gin.Context) {
	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		panic(err)
	}
	token := c.Query("token")
	go ws.h.Run(conn, token)
}
