package usecases

import "github.com/ashkarin/ashkarin-pizza-shop/pkg/ingredient"

// SearchIngredient find ingredients by the name in the storage
func SearchIngredient(s ingredient.StorageGateway, name string) ([]*ingredient.Ingredient, error) {
	return s.Search(name)
}
