package handlers

import (
	"encoding/json"
	"io"
	"log"
	"net/http"

	"github.com/omekov/exclusive/internal/apiserver/models"
	"github.com/omekov/exclusive/internal/apiserver/stores"
)

func (s *server) homeHandler(w http.ResponseWriter, r *http.Request) {
	log.Println(r.URL)
	w.Write([]byte("Server stating"))
}

func (s *server) sendMailHandler(w http.ResponseWriter, r *http.Request) {
	log.Println(r.URL)
	w.Write([]byte("Send mail"))
}
func (s *server) healthCheckHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")

	// In the future we could report back on the status of our DB, or our cache
	// (e.g. Redis) by performing a simple PING, and include them in the response.
	io.WriteString(w, `{"alive": true}`)
}

func (s *server) signInHandler(w http.ResponseWriter, r *http.Request) {
	auth := models.Auth{}
	if err := json.NewDecoder(r.Body).Decode(&auth); err != nil {
		s.error(w, r, http.StatusBadRequest, err)
	}
	store := stores.Store{}
	store.ConfigMongo("mongodb://ds261716.mlab.com:61716/")
	store.FindCustomer(auth.Username)
	s.respond(w, r, http.StatusOK, auth)
}

func (s *server) signUpHandler(w http.ResponseWriter, r *http.Request) {
	custormer := models.Customer{}
	if err := json.NewDecoder(r.Body).Decode(&custormer); err != nil {
		s.error(w, r, http.StatusBadRequest, err)
	}
	store := stores.Store{}
	store.RegisterCustomer(&custormer)
	s.respond(w, r, http.StatusOK, custormer)
}
