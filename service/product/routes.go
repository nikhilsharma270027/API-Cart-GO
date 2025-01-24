package product

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/nikhilsharma270027/API-Cart-GO/types"
	"github.com/nikhilsharma270027/API-Cart-GO/utils"
)

type Handler struct {
	store types.ProductStore
}

func NewHandler(store types.ProductStore) *Handler {
	return &Handler{store: store}
}

func (h *Handler) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/products", h.handleCreateProduct).Methods(http.MethodGet)
	router.HandleFunc("/products", h.handleCreateProduct).Methods(http.MethodPost)
}

func (h *Handler) handleCreateProduct(w http.ResponseWriter, r *http.Request) {
	// var product types.ProductStore
	// if err := utils.ParseJSON(r, &product); err != nil {
	// 	utils.WriteError(w, http.StatusBadRequest, err)
	// 	return
	// }

	// if err := utils.Validate.Struct(product); err != nil {
	// 	errors := err.(validator.ValidationErrors)
	// 	utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("invalid payload: %v", errors))
	// 	return
	// }

	ps, err := h.store.GetProducts()
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	utils.WriteJSon(w, http.StatusCreated, ps)
}
