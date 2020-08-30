package handlers

import (
	"net/http"

	"github.com/gorilla/mux"
)

type server struct {
	Router *mux.Router
}

func (s *server) configureRouter() *mux.Router {
	s.Router.Use(mux.CORSMethodMiddleware(s.Router))
	s.Router.HandleFunc("/", s.homeHandler)
	s.Router.HandleFunc("/health-check", s.homeHandler).Methods(http.MethodGet)
	return s.Router
}

// StartServer ...
func StartServer() *mux.Router {
	server := server{
		Router: mux.NewRouter(),
	}
	router := server.configureRouter()
	return router
}
