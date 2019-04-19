package usecases

import "github.com/ashkarin/ashkarin-pizza-shop/pkg/order"

// DeleteOrder delete the order entry in the storage
func DeleteOrder(s order.StorageGateway, o *order.Order) error {
	return s.Delete(o)
}
