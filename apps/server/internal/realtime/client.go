package realtime

import (
	"github.com/gofiber/contrib/websocket"
)

type Client struct {
	ID string

	UserID string

	RoomID string

	Conn *websocket.Conn
}