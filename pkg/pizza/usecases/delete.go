package usecases

import (
	"github.com/ashkarin/ashkarin-pizza-shop/pkg/pizza"
)

// DeletePizza delete the pizza entry in the storage
func DeletePizza(s pizza.StorageGateway, p *pizza.Pizza) error {
	return s.Delete(p)
}
