package user

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/nikhilsharma270027/API-Cart-GO/service/auth"
	"github.com/nikhilsharma270027/API-Cart-GO/types"
	"github.com/nikhilsharma270027/API-Cart-GO/utils"
)

type Handler struct {
	store types.UserStore
}

func NewHandler(store types.UserStore) *Handler {
	return &Handler{store: store}
}

func (h *Handler) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/login", h.handlerLogin).Methods("POST")
	router.HandleFunc("/register", h.handlerRegister).Methods("POST")
}

// we need to json payload and implement jwt

func (h *Handler) handlerLogin(w http.ResponseWriter, r *http.Request) {

}

func (h *Handler) handlerRegister(w http.ResponseWriter, r *http.Request) {
	// get jSON payload
	var payload types.RegisterUserPayload

	if err := utils.ParseJSON(r, &payload); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
	}
	// check if the user exists
	_, err := h.store.GetUserByEmail(payload.Email)
	if err == nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	// hashing password
	hashedPassword, err := auth.HashPassword(payload.Password)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
	}
	// if it doesnt we create the new user
	// we gonna create it using
	err = h.store.CreateUser(types.User{
		FirstName: payload.FirstName,
		LastName:  payload.LastName,
		Email:     payload.Email,
		Password:  hashedPassword,
	})

	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	utils.WriteJSon(w, http.StatusCreated, nil)

}
