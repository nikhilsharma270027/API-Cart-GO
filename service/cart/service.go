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
