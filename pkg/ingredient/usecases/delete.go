package usecases

import "github.com/ashkarin/ashkarin-pizza-shop/pkg/ingredient"

// DeleteIngrdient delete the ingredient entry from the storage
func DeleteIngrdient(s ingredient.StorageGateway, i *ingredient.Ingredient) error {
	return s.Delete(i)
}

// DeleteIngrdientByID delete the ingredient entry from the storage by the ID
func DeleteIngrdientByID(s ingredient.StorageGateway, id string) error {
	return s.DeleteByID(id)
}
