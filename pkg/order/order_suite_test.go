package order_test

import (
	"testing"
	"time"

	m "github.com/ashkarin/ashkarin-pizza-shop/pkg/order"
	"github.com/ashkarin/ashkarin-pizza-shop/pkg/order/gateways"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var (
	storage   m.StorageGateway = nil
	accepted  time.Time
	completed time.Time
)

func TestOrder(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Order Suite")
}

var _ = BeforeSuite(func() {
	// Open a gateway to the MongoDB storage
	dbHost := "mongodb"
	dbPort := "27017"
	dbName := "pizzashop_test"
	dbUser := ""
	dbPass := ""
	dbCollection := "orders"

	var err error
	storage, err = gateways.NewMongoDbGateway(dbHost, dbPort, dbUser, dbPass, dbName, dbCollection)
	Expect(err).NotTo(HaveOccurred())

	loc, err := time.LoadLocation("UTC")
	Expect(err).NotTo(HaveOccurred())

	accepted = time.Date(2019, 04, 11, 9, 00, 20, 0, loc)
	completed = time.Date(2019, 04, 11, 10, 15, 00, 0, loc)
})

var _ = AfterSuite(func() {
	//err := storage.Clean()
	//Expect(err).NotTo(HaveOccurred())
})
