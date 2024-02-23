package internal

import (
	"log"
	"net/http"
)

func (cfg *ApiConfig) Log(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		//log.Printf("Method: %s Path: %s", r.Method, r.URL.Path)
		log.Printf(`
Method: %s
Path: %s
		`, r.Method, r.URL.Path)
		next.ServeHTTP(w, r)
	})
}
