package realtime

import "sync"

type Hub struct {
	Rooms map[string]map[*Client]bool

	mu sync.RWMutex
}

func NewHub() *Hub {
	return &Hub{
		Rooms: make(map[string]map[*Client]bool),
	}
}