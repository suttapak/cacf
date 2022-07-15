package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"github.com/suttapak/cacf/services"
)

type ws struct {
	h services.HubServices
}

func NewWsServer(h services.HubServices) Ws {
	return &ws{h}
}
func (ws *ws) Serve(c *gin.Context) {
	conn, err := websocket.Upgrade(c.Writer, c.Request, nil, 1024, 1024)
	if err != nil {
		panic(err)
	}
	token := c.Query("token")
	go ws.h.Run(conn, token)
}
