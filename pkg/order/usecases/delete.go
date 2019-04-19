package usecases

import "github.com/ashkarin/ashkarin-pizza-shop/pkg/order"

// DeleteOrder delete the order entry in the storage
func DeleteOrder(s order.StorageGateway, o *order.Order) error {
	return s.Delete(o)
}

// DeleteOrderByID delete the order entry from the storage by the ID
func DeleteOrderByID(s order.StorageGateway, id string) error {
	return s.DeleteByID(id)
}
