package realtime

const (
	EventRoomJoined = "room.joined"

	EventRoomLeft = "room.left"

	EventMessageReceived = "message.received"

	EventPresenceUpdate = "presence.update"
)

type Event struct {
	Type string `json:"type"`

	Payload any `json:"payload"`
}