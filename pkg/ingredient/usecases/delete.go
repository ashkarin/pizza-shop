package usecases

import "github.com/ashkarin/ashkarin-pizza-shop/pkg/ingredient"

// DeleteIngrdient delete the ingredient entry in the storage
func DeleteIngrdient(s ingredient.StorageGateway, i *ingredient.Ingredient) error {
	return s.Delete(i)
}
