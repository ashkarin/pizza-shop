package ingredient_test

import (
	"testing"

	m "github.com/ashkarin/ashkarin-pizza-shop/pkg/ingredient"
	"github.com/ashkarin/ashkarin-pizza-shop/pkg/ingredient/gateways"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var storage m.StorageGateway = nil

func TestIngredient(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Ingredient Suite")
}

var _ = BeforeSuite(func() {
	// Open a gateway to the MongoDB storage
	dbHost := "mongodb"
	dbPort := "27017"
	dbName := "pizzashop_test"
	dbUser := ""
	dbPass := ""
	dbCollection := "ingredients"

	var err error
	storage, err = gateways.NewMongoDbGateway(dbHost, dbPort, dbUser, dbPass, dbName, dbCollection)
	Expect(err).NotTo(HaveOccurred())
})

var _ = AfterSuite(func() {
	//err := storage.Clean()
	//Expect(err).NotTo(HaveOccurred())
})
