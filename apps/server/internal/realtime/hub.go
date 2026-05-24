package realtime

import "sync"

type Hub struct {
	Rooms map[string]map[*Client]bool

	Presence map[string]map[string]bool

	mu sync.RWMutex
}

func NewHub() *Hub {
	return &Hub{
		Rooms: make(map[string]map[*Client]bool),

		Presence: make(map[string]map[string]bool),
	}
}