package usecases

import (
	"github.com/ashkarin/ashkarin-pizza-shop/pkg/pizza"
)

// CreatePizza create the pizza entry in the storage
func CreatePizza(s pizza.StorageGateway, p *pizza.Pizza) error {
	return s.Store(p)
}
