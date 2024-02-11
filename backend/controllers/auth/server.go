package auth

import (
	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
)

// Server is a container for controllers used in authentification
type Server struct {
	db *sqlx.DB
}

func NewServer(db *sqlx.DB) *Server {
	return &Server{db: db}
}

func (s *Server) Route(subrouter *mux.Router) {
	subrouter.HandleFunc("/new", s.CreateUserHandler).Methods("POST")
}
