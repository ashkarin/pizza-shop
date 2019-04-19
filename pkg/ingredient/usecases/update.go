package usecases

import "github.com/ashkarin/ashkarin-pizza-shop/pkg/ingredient"

// UpdateIngredient update the ingredient entry in the storage
func UpdateIngredient(s ingredient.StorageGateway, p *ingredient.Ingredient) error {
	return s.Update(p)
}
