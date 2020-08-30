package handlers

import (
	"io"
	"log"
	"net/http"
)

func (s *server) homeHandler(w http.ResponseWriter, r *http.Request) {
	log.Println(r.URL)
	w.Write([]byte("Server stating"))
}
func (s *server) healthCheckHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")

	// In the future we could report back on the status of our DB, or our cache
	// (e.g. Redis) by performing a simple PING, and include them in the response.
	io.WriteString(w, `{"alive": true}`)
}
