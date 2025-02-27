package proxy

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"sync"
	"time"
)

// Backend structure to track connections
type Backend struct {
	Address           string
	ActiveConnections int
	mu                sync.Mutex
}

// List of backend servers
var backends = []*Backend{
	{Address: "127.0.0.1:9001"},
	{Address: "127.0.0.1:9002"},
	{Address: "127.0.0.1:9003"},
}

// getLeastConnectionsBackend selects the backend with the least active connections
func getLeastConnectionsBackend() *Backend {
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

// HandleRequest forwards HTTP requests to the least loaded backend
func HandleRequest(w http.ResponseWriter, r *http.Request) {
	backend := getLeastConnectionsBackend()
	if backend == nil {
		http.Error(w, "No available backend servers", http.StatusServiceUnavailable)
		return
	}

	backend.mu.Lock()
	backend.ActiveConnections++
	backend.mu.Unlock()

	start := time.Now()

	// Forward request to backend
	url := fmt.Sprintf("http://%s%s", backend.Address, r.URL.Path)
	log.Printf("Forwarding request to: %s", url)
	resp, err := http.Get(url)
	if err != nil {
		http.Error(w, "Backend unavailable", http.StatusServiceUnavailable)
		return
	}
	defer resp.Body.Close()

	// Copy backend response to client
	w.WriteHeader(resp.StatusCode)
	io.Copy(w, resp.Body)

	duration := time.Since(start).Seconds()
	log.Printf("Request to %s took %f seconds\n", backend.Address, duration)

	backend.mu.Lock()
	backend.ActiveConnections--
	backend.mu.Unlock()
}
