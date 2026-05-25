package realtime

const (
	EventRoomJoined = "room.joined"

	EventRoomLeft = "room.left"

	EventMessageReceived = "message.received"

	EventPresenceUpdate = "presence.update"

	EventPresenceState = "presence.state"

	EventTabCreated = "tab.created"

	EventTabActivated = "tab.activated"

	EventWhiteboardUpdated = "whiteboard.updated"
)

type Event struct {
	Type string `json:"type"`

	Payload any `json:"payload"`
}