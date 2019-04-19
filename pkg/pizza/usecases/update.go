package usecases

import "github.com/ashkarin/ashkarin-pizza-shop/pkg/pizza"

// UpdatePizza update the pizza entry in the storage
func UpdatePizza(s pizza.StorageGateway, p *pizza.Pizza) error {
	return s.Update(p)
}
