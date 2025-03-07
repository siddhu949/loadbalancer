package pool

import (
	"log"
	"sync"
)

// ConnectionPool structure
type ConnectionPool struct {
	mu         sync.Mutex
	connections []string
}

// NewPool creates a connection pool
func NewPool(size int) *ConnectionPool {
	return &ConnectionPool{
		connections: make([]string, size),
	}
}

// Acquire gets a connection
func (p *ConnectionPool) Acquire() string {
	p.mu.Lock()
	defer p.mu.Unlock()

	if len(p.connections) > 0 {
		conn := p.connections[0]
		p.connections = p.connections[1:]
		return conn
	}
	log.Println("No available connections")
	return ""
}

// Release returns a connection
func (p *ConnectionPool) Release(conn string) {
	p.mu.Lock()
	defer p.mu.Unlock()
	p.connections = append(p.connections, conn)
}
