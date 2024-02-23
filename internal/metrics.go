package internal

import (
	"fmt"
	"net/http"
)

func (m *ApiConfig) RegisterHits(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		m.Hits++
		next.ServeHTTP(w, r)
	})
}

func (c *ApiConfig) MetricsHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "text/plain; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(fmt.Sprintf("Hits: %d", c.Hits)))
}
