package firewall

import (
	"net"
	"net/http"
)

// List of blacklisted IPs
var blacklistedIPs = map[string]bool{
	"192.168.1.10": true, // Example blocked IP
	"10.0.0.5":     true, // Example blocked IP
}

// Middleware to block blacklisted IPs
func IPBlocker(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ip, _, _ := net.SplitHostPort(r.RemoteAddr)
		if blacklistedIPs[ip] {
			http.Error(w, "403 - Forbidden", http.StatusForbidden)
			return
		}
		next.ServeHTTP(w, r)
	})
}
