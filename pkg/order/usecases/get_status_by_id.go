package usecases

import "github.com/ashkarin/ashkarin-pizza-shop/pkg/order"

// GetStatusByID get order status by the ID
func GetStatusByID(s order.StorageGateway, id string) (order.Status, error) {
	o, err := s.GetByID(id)
	if err != nil {
		return order.StatusUnknown, err
	}
	return o.Status, nil
}
