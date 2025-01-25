package api

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/nikhilsharma270027/API-Cart-GO/service/cart"
	"github.com/nikhilsharma270027/API-Cart-GO/service/order"
	"github.com/nikhilsharma270027/API-Cart-GO/service/product"
	"github.com/nikhilsharma270027/API-Cart-GO/service/user"
	"github.com/rs/cors"
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
	productHandler := product.NewHandler(productStore, userStore)
	productHandler.RegisterRoutes(subrouter)

	orderStore := order.NewStore(s.db)

	cartHandler := cart.NewHandler(productStore, orderStore, userStore)
	cartHandler.RegisterRoutes(subrouter)

	// Configure CORS
	corsMiddleware := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:3000", "http://example.com"}, // Replace with your frontend domains
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Authorization", "Content-Type"},
		AllowCredentials: true,
	})

	// Wrap the router with CORS middleware
	handler := corsMiddleware.Handler(router)

	log.Println("Listening on", s.addr)

	return http.ListenAndServe(s.addr, handler)
}

// You have an API with routes like:

// /api/v1/users
// /api/v1/products
// /api/v1/orders
// Instead of defining the prefix (/api/v1) repeatedly for every route, you can use PathPrefix("api/v1") to group them.
