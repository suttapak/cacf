package services

import "github.com/gorilla/websocket"

type HubServices interface {
	Run(conn *websocket.Conn, token string)
}
