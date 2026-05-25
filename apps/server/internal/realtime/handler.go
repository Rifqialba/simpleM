package realtime

import (
	"encoding/json"
	"log"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"

	"github.com/golang-jwt/jwt/v5"
	"github.com/gofiber/contrib/websocket"
)

type ClientMessage struct {
	Type string `json:"type"`

	Payload any `json:"payload"`
}

type Handler struct {
	manager *Manager

	jwtSecret string
}

func NewHandler(
	manager *Manager,
	jwtSecret string,
) *Handler {

	return &Handler{
		manager: manager,

		jwtSecret: jwtSecret,
	}
}

func (h *Handler) Handle(
	c *websocket.Conn,
) {

	roomID := c.Params("roomId")

	authHeader := c.Headers("Authorization")
	

if authHeader == "" {

	tokenQuery := c.Query("token")

	if tokenQuery != "" {

		authHeader =
			"Bearer " + tokenQuery
	}
}

	if authHeader == "" {

		_ = c.WriteJSON(Event{
			Type: "error",
			Payload: "missing authorization header",
		})

		c.Close()

		return
	}

	tokenString := strings.TrimPrefix(
		authHeader,
		"Bearer ",
	)

	token, err := jwt.Parse(
		tokenString,
		func(token *jwt.Token) (any, error) {
			return []byte(h.jwtSecret), nil
		},
	)

	if err != nil || !token.Valid {

		_ = c.WriteJSON(Event{
			Type: "error",
			Payload: "invalid token",
		})

		c.Close()

		return
	}

	claims, ok := token.Claims.(jwt.MapClaims)

	if !ok {

		c.Close()

		return
	}

	userID, ok := claims["user_id"].(string)

	if !ok {

		c.Close()

		return
	}

	client := &Client{
		ID: userID + ":" + roomID,

		UserID: userID,

		RoomID: roomID,

		Conn: c,
	}

	h.manager.JoinRoom(
		roomID,
		client,
	)

	h.manager.Broadcast(
		roomID,
		Event{
			Type: EventPresenceUpdate,
			Payload: map[string]any{
				"users": h.manager.GetPresence(roomID),
			},
		},
	)

	defer func() {

		h.manager.LeaveRoom(
			roomID,
			client,
		)

		h.manager.Broadcast(
			roomID,
			Event{
				Type: EventPresenceUpdate,
				Payload: map[string]any{
					"users": h.manager.GetPresence(roomID),
				},
			},
		)

		h.manager.Broadcast(
			roomID,
			Event{
				Type: EventRoomLeft,
				Payload: map[string]any{
					"user_id": userID,
				},
			},
		)

		c.Close()
	}()

	h.manager.Broadcast(
		roomID,
		Event{
			Type: EventRoomJoined,
			Payload: map[string]any{
				"user_id": userID,
			},
		},
	)

	c.SetReadDeadline(
		time.Now().Add(60 * time.Second),
	)

	c.SetPongHandler(func(string) error {

		c.SetReadDeadline(
			time.Now().Add(60 * time.Second),
		)

		return nil
	})

	go func() {

		ticker := time.NewTicker(
			30 * time.Second,
		)

		defer ticker.Stop()

		for range ticker.C {

			err := c.WriteMessage(
				websocket.PingMessage,
				[]byte{},
			)

			if err != nil {
				return
			}
		}
	}()

	_ = c.WriteJSON(Event{
		Type: EventPresenceState,
		Payload: map[string]any{
			"users": h.manager.GetPresence(roomID),
		},
	})

	for {

		_, messageBytes, err := c.ReadMessage()

		if err != nil {
			log.Println(err)
			break
		}

		var message ClientMessage

		err = json.Unmarshal(
			messageBytes,
			&message,
		)

		if err != nil {
			continue
		}

		switch message.Type {

		case "whiteboard.update":

			h.manager.Broadcast(
				roomID,
				Event{
					Type: EventWhiteboardUpdated,
					Payload: message.Payload,
				},
			)

		default:

			h.manager.Broadcast(
				roomID,
				Event{
					Type: EventMessageReceived,
					Payload: fiber.Map{
						"user_id": userID,
						"message": string(messageBytes),
					},
				},
			)
		}
	}
}