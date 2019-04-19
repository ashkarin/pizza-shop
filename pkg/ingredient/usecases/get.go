package usecases

import (
	"github.com/ashkarin/ashkarin-pizza-shop/pkg/ingredient"
)

// GetIngredientByID get ingredient entry by ID
func GetIngredientByID(s ingredient.StorageGateway, id string) (*ingredient.Ingredient, error) {
	return s.GetByID(id)
}
