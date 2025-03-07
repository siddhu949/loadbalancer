package algorithm

import "sync"

// Backend structure
type Backend struct {
	Address           string
	ActiveConnections int
	mu                sync.Mutex
}

// Round Robin Load Balancing
var rrIndex int

func GetRoundRobinBackend(backends []*Backend) *Backend {
	backend := backends[rrIndex%len(backends)]
	rrIndex++
	return backend
}

// Least Connections Load Balancing
func GetLeastConnectionsBackend(backends []*Backend) *Backend {
	var selected *Backend
	for _, backend := range backends {
		backend.mu.Lock()
		if selected == nil || backend.ActiveConnections < selected.ActiveConnections {
			selected = backend
		}
		backend.mu.Unlock()
	}
	return selected
}
