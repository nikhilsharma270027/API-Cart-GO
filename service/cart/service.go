package cart

import (
	"fmt"

	"github.com/nikhilsharma270027/API-Cart-GO/types"
)

func getCartItemsIDs(items []types.CartItem) ([]int, error) {
	// create a slice and storing b y making its len of product in items
	productIds := make([]int, len(items))
	for i, item := range items {
		if item.Quantity <= 0 {
			return nil, fmt.Errorf("invalid quantity for the product 5 %d", item.ProductID)
		}

		productIds[i] = item.ProductID
	}

	return productIds, nil
}

func (h *Handler) createOrder(ps []types.Product, cartItems []types.CartItem, userID int) (int, float64, error) {
	// to quickly access all products  we create product map
	productMap := make(map[int]types.Product)
	for _, product := range ps { // populating
		productMap[product.ID] = product
	}

	// check if all products are actually in stock | if aviable
	if err := checkIfCartIsInStock(cartItems, productMap); err != nil {
		return 0, 0, nil
	}

	// calculate the total price
	totalPrice := calculateTotalPrice(cartItems, productMap)

	// reduce quantity of products in our db
	for _, item := range cartItems {
		product := productMap[item.ProductID]
		product.Quantity -= item.Quantity

		h.store.UpdateProduct(product)
	}
	// create the order - creating order table
	orderID, err := h.orderStore.CreateOrder(types.Order{
		UserID:  userID,
		Total:   totalPrice,
		Status:  "pending",
		Address: "some address", // could fetch address from a user addresses table
	})
	if err != nil {
		return 0, 0, err
	}
	// create order items
	for _, item := range cartItems {
		h.orderStore.CreateOrderItem(types.OrderItem{
			OrderID:   orderID,
			ProductID: item.ProductID,
			Quantity:  item.Quantity,
			Price:     productMap[item.ProductID].Price,
		})
	}
	// return 0,0,nil
	// return 0,totalPrice,nil
	return orderID, totalPrice, nil
}

func checkIfCartIsInStock(cartItems []types.CartItem, products map[int]types.Product) error {
	if len(cartItems) == 0 {
		return fmt.Errorf("cart is empty")
	}

	for _, item := range cartItems {
		product, ok := products[item.ProductID]
		if !ok {
			return fmt.Errorf("product %d is not available in the store, please refresh your cart", item.ProductID)
		}

		if product.Quantity < item.Quantity {
			return fmt.Errorf("product %s is not available in the quantity requested", product.Name)
		}
	}

	return nil

}

func calculateTotalPrice(cartItems []types.CartItem, products map[int]types.Product) float64 {
	var total float64

	for _, item := range cartItems {
		product := products[item.ProductID]
		total += product.Price * float64(item.Quantity)
	}

	return total
}
