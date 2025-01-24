package api

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/nikhilsharma270027/API-Cart-GO/service/product"
	"github.com/nikhilsharma270027/API-Cart-GO/service/user"
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
	subrouter := router.PathPrefix("/api/v1").Subrouter()

	// userService
	userStore := user.NewStore(s.db)
	userHanlder := user.NewHandler(userStore) // adding user store from store.go
	userHanlder.RegisterRoutes(subrouter)     // add subrouter now /api/vi/login

	//Product service
	productStore := product.NewStore(s.db)
	productHandler := product.NewHandler(productStore)
	productHandler.RegisterRoutes(subrouter)

	log.Println("Listening on", s.addr)

	return http.ListenAndServe(s.addr, router)
}

// You have an API with routes like:

// /api/v1/users
// /api/v1/products
// /api/v1/orders
// Instead of defining the prefix (/api/v1) repeatedly for every route, you can use PathPrefix("api/v1") to group them.
