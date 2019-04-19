package usecases

import "github.com/ashkarin/ashkarin-pizza-shop/pkg/order"

// PlaceOrder create the order entry in the storage
func PlaceOrder(s order.StorageGateway, o *order.Order) (string, error) {
	return s.Place(o)
}
