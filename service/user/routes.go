package user

import (
	"fmt"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
	"github.com/nikhilsharma270027/API-Cart-GO/config"
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
	// get jSON payload
	var payload types.LoginUserPayload

	if err := utils.ParseJSON(r, &payload); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	// validate payload
	if err := utils.Validate.Struct(payload); err != nil {
		errors := err.(validator.ValidationErrors)
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("invalid palyload %v", errors))
		return
	}

	// define user/ we gonna find using email
	u, err := h.store.GetUserByEmail(payload.Email)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("Not found, invalid email or password"))
		return
	}

	if !auth.ComparePasswords(u.Password, []byte(payload.Password)) {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("not found, invalid email or password"))
		return
	}

	//after we create jwt
	secret := []byte(config.Envs.JWTSecret)
	token, err := auth.CreateJWT(secret, u.ID)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	// we will create jwt auth
	// when the passwords match we send the jwt token in return
	utils.WriteJSon(w, http.StatusOK, map[string]string{"token": token})

}

func (h *Handler) handlerRegister(w http.ResponseWriter, r *http.Request) {
	// get jSON payload
	var payload types.RegisterUserPayload

	if err := utils.ParseJSON(r, &payload); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	// validate payload
	if err := utils.Validate.Struct(payload); err != nil {
		errors := err.(validator.ValidationErrors)
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("invalid palyload %v", errors))
		return
	}
	// check if the user exists
	_, err := h.store.GetUserByEmail(payload.Email)
	if err == nil {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("user with email %s already exists", payload.Email))
		return
	}

	// hashing password
	hashedPassword, err := auth.HashPassword(payload.Password)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
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
