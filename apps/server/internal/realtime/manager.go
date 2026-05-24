package realtime

import (
	"encoding/json"
)

type Manager struct {
	hub *Hub
}

func NewManager() *Manager {
	return &Manager{
		hub: NewHub(),
	}
}

func (m *Manager) JoinRoom(
	roomID string,
	client *Client,
) {

	m.hub.mu.Lock()
	defer m.hub.mu.Unlock()

	if _, exists := m.hub.Rooms[roomID]; !exists {

		m.hub.Rooms[roomID] = make(map[*Client]bool)
	}

	m.hub.Rooms[roomID][client] = true
}

func (m *Manager) LeaveRoom(
	roomID string,
	client *Client,
) {

	m.hub.mu.Lock()
	defer m.hub.mu.Unlock()

	if _, exists := m.hub.Rooms[roomID]; !exists {
		return
	}

	delete(
		m.hub.Rooms[roomID],
		client,
	)

	if len(m.hub.Rooms[roomID]) == 0 {

		delete(m.hub.Rooms, roomID)
	}
}

func (m *Manager) Broadcast(
	roomID string,
	event Event,
) {

	m.hub.mu.RLock()

	clients := m.hub.Rooms[roomID]

	m.hub.mu.RUnlock()

	payload, err := json.Marshal(event)

	if err != nil {
		return
	}

	for client := range clients {

		err := client.Conn.WriteMessage(
			1,
			payload,
		)

		if err != nil {

			m.LeaveRoom(
				roomID,
				client,
			)
		}
	}
}