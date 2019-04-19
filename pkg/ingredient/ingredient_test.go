package ingredient_test

import (
	m "github.com/ashkarin/ashkarin-pizza-shop/pkg/ingredient"
	"github.com/ashkarin/ashkarin-pizza-shop/pkg/ingredient/usecases"

	//log "github.com/sirupsen/logrus"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var (
	ingredientID = ""
)

var _ = Describe("Ingredient", func() {
	It("should add entry to the storage", func() {
		entry := &m.Ingredient{
			Name:       "Cabbage",
			Vegetarian: true,
			Price:      1.2,
		}

		err := usecases.CreateIngredient(storage, entry)
		if err != nil {
			GinkgoWriter.Write([]byte(err.Error()))
			Expect(err).NotTo(HaveOccurred())
		}

	})

	It("should find a single entry in the storage", func() {
		foundIngredients, err := usecases.SearchIngredient(storage, "Cabb")
		if err != nil {
			GinkgoWriter.Write([]byte(err.Error()))
			Expect(err).NotTo(HaveOccurred())
		}
		Expect(foundIngredients).To(HaveLen(1))
		ingredientID = foundIngredients[0].ID.(string)
	})

	It("should update entry to the storage", func() {
		entry := &m.Ingredient{
			ID:         ingredientID,
			Name:       "Bio Cabbage",
			Vegetarian: true,
			Price:      1.2,
		}

		err := usecases.UpdateIngredient(storage, entry)
		if err != nil {
			GinkgoWriter.Write([]byte(err.Error()))
			Expect(err).NotTo(HaveOccurred())
		}
	})

	It("should get entry from the storage by the ID", func() {
		expected := &m.Ingredient{
			ID:         ingredientID,
			Name:       "Bio Cabbage",
			Vegetarian: true,
			Price:      1.2,
		}

		obtained, err := usecases.GetIngredientByID(storage, ingredientID)
		if err != nil {
			GinkgoWriter.Write([]byte(err.Error()))
			Expect(err).NotTo(HaveOccurred())
		}

		Expect(expected.Name).To(Equal(obtained.Name))
		Expect(expected.Vegetarian).To(Equal(obtained.Vegetarian))
		Expect(int(expected.Price)).To(Equal(int(obtained.Price)))
	})

	It("should delete entry from the storage by the ID", func() {
		err := usecases.DeleteIngrdientByID(storage, ingredientID)
		if err != nil {
			GinkgoWriter.Write([]byte(err.Error()))
			Expect(err).NotTo(HaveOccurred())
		}
	})
})
