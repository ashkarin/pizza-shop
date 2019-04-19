package usecases

import "github.com/ashkarin/ashkarin-pizza-shop/pkg/pizza"

// SearchPizza find pizzas by the name
func SearchPizza(s pizza.StorageGateway, name string) ([]*pizza.Pizza, error) {
	return s.Search(name)
}
