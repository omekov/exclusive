package handlers

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

type server struct {
	router *mux.Router
	logger *logrus.Logger
}

// StartServer ...
func Router() *mux.Router {
	server := server{
		router: mux.NewRouter(),
		logger: logrus.New(),
	}
	router := server.configureRouter()
	return router
}

func (s *server) configureRouter() *mux.Router {
	s.router.Use(mux.CORSMethodMiddleware(s.router))
	s.router.Use(s.logRequest)
	s.router.HandleFunc("/", s.homeHandler).Methods(http.MethodGet)
	s.router.HandleFunc("/signin", s.signInHandler).Methods(http.MethodPost)
	s.router.HandleFunc("/signup", s.signUpHandler).Methods(http.MethodPost)
	s.router.HandleFunc("/signout", s.homeHandler)
	s.router.HandleFunc("/mail", s.sendMailHandler)
	s.router.HandleFunc("/health-check", s.homeHandler).Methods(http.MethodGet)
	return s.router
}
