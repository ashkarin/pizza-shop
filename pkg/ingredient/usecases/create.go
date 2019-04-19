package usecases

import (
	"github.com/ashkarin/ashkarin-pizza-shop/pkg/ingredient"
)

// CreateIngredient create the ingredient entry in the storage
func CreateIngredient(s ingredient.StorageGateway, i *ingredient.Ingredient) error {
	return s.Store(i)
}
