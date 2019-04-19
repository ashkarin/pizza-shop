package usecases

import (
	"github.com/ashkarin/ashkarin-pizza-shop/pkg/pizza"
)

// GetPizzaByID get pizza entry by ID
func GetPizzaByID(s pizza.StorageGateway, id string) (*pizza.Pizza, error) {
	return s.GetByID(id)
}
