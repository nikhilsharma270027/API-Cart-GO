package cart

import (
	"fmt"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
	"github.com/nikhilsharma270027/API-Cart-GO/service/auth"
	"github.com/nikhilsharma270027/API-Cart-GO/types"
	"github.com/nikhilsharma270027/API-Cart-GO/utils"
)

type Handler struct {
	store      types.ProductStore
	orderStore types.OrderStore
	userStore  types.UserStore
	// productStore types.ProductStore
}

func NewHandler(store types.ProductStore, orderStore types.OrderStore, userStore types.UserStore) *Handler {
	return &Handler{
		store:      store,
		orderStore: orderStore,
		userStore:  userStore,
	}
}

func (h *Handler) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/cart/checkout", auth.WithJWTAuth(h.handleCheckout, h.userStore)).Methods(http.MethodPost)
}

// we need to know if product exists in Cart , so we use products dependency
func (h *Handler) handleCheckout(w http.ResponseWriter, r *http.Request) {
	// the user id will be coming fro m jwt token
	userID := auth.GetUserIDFromContext(r.Context())

	// check if product exists using order and productStore
	var cart types.CartCheckoutPayload
	if err := utils.ParseJSON(r, &cart); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	if err := utils.Validate.Struct(cart); err != nil {
		errors := err.(validator.ValidationErrors)
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("invalid payload: %v", errors))
		return
	}

	//get products
	// we need to iterate over the cart as it has ites
	// we create a func to do so
	productIDs, err := getCartItemsIDs(cart.Items)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	// get products
	products, err := h.store.GetProductsByID(productIDs)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	orderID, totalPrice, err := h.createOrder(products, cart.Items, userID)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	utils.WriteJSon(w, http.StatusOK, map[string]interface{}{
		"total_price": totalPrice,
		"order_id":    orderID,
	})

}
