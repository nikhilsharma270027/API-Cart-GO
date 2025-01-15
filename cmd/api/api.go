package api

import (
	"database/sql"
	"net/http"

	"github.com/gorilla/mux"
)

type APIServer struct {
	addr string // address
	db   *sql.DB
}

// we will create new api instance
func NewAPIServer(addr string, db *sql.DB) *APIServer {
	return &APIServer{
		addr: addr,
		db:   db,
	}
}

// to run server in main , we creat ein this file
func (s *APIServer) Run() error {
	// router gorilla mux
	// router := http.NewServeMux()
	router := mux.NewRouter()
	subrouter := router.PathPrefix("api/v1").Subrouter()
	subrouter := router.
	
	return http.ListenAndServe(s.addr, router)
}