package usecases

import (
	"github.com/ashkarin/ashkarin-pizza-shop/pkg/order"
)

// GetOrderByID get order entry from the storage by the ID
func GetOrderByID(s order.StorageGateway, id string) (*order.Order, error) {
	return s.GetByID(id)
}
