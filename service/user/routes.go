package user

import (
	"net/http"

	"github.com/gorilla/mux"
)

type Handler struct {
}

func NewHandler() *Handler {
	return &Handler{}
}

func (h *Handler) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/login", h.handlerLogin).Methods("POST")
	router.HandleFunc("/Register", h.handlerRegister).Methods("POST")
}

func (h *Handler) handlerLogin(w http.ResponseWriter, r *http.Request) {

}

func (h *Handler) handlerRegister(w http.ResponseWriter, r *http.Request) {

}
