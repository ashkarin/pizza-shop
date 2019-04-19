package usecases

import "github.com/ashkarin/ashkarin-pizza-shop/pkg/order"

// UpdateOrder update the order entry in the storage
func UpdateOrder(s order.StorageGateway, o *order.Order) error {
	return s.Update(o)
}
