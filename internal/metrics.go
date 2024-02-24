package internal

import (
	"fmt"
	"net/http"
)

// Middleware to register all the calls that have been made to the server.
func (m *ApiConfig) RegisterHits(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		m.Hits++
		next.ServeHTTP(w, r)
	})
}

// Endpoint used to get the value of the total hits.
func (c *ApiConfig) MetricsHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "text/plain; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(fmt.Sprintf("Hits: %d", c.Hits)))
}
