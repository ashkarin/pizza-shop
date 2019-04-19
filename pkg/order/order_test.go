package order_test

import (
	m "github.com/ashkarin/ashkarin-pizza-shop/pkg/order"
	"github.com/ashkarin/ashkarin-pizza-shop/pkg/order/usecases"

	//log "github.com/sirupsen/logrus"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var (
	orderID = ""
)

var _ = Describe("Order", func() {
	It("should add entry to the storage", func() {
		entry := &m.Order{
			AcceptedAt:  accepted,
			CompletedAt: completed,
			Status:      m.StatusAccepted,
		}

		var err error
		orderID, err = usecases.PlaceOrder(storage, entry)
		if err != nil {
			GinkgoWriter.Write([]byte(err.Error()))
			Expect(err).NotTo(HaveOccurred())
		}

	})

	It("should update entry to the storage", func() {
		entry := &m.Order{
			ID:          orderID,
			AcceptedAt:  accepted,
			CompletedAt: completed,
			Status:      m.StatusInProcess,
		}

		err := usecases.UpdateOrder(storage, entry)
		if err != nil {
			GinkgoWriter.Write([]byte(err.Error()))
			Expect(err).NotTo(HaveOccurred())
		}
	})

	It("should get entry from the storage by the ID", func() {
		expected := &m.Order{
			ID:          orderID,
			AcceptedAt:  accepted,
			CompletedAt: completed,
			Status:      m.StatusInProcess,
		}

		obtained, err := usecases.GetOrderByID(storage, orderID)
		if err != nil {
			GinkgoWriter.Write([]byte(err.Error()))
			Expect(err).NotTo(HaveOccurred())
		}

		//Expect(expected.AcceptedAt).To(Equal(obtained.AcceptedAt))
		//Expect(expected.CompletedAt).To(Equal(obtained.CompletedAt))
		Expect(int(expected.Status)).To(Equal(int(obtained.Status)))
	})

	It("should delete entry from the storage by the ID", func() {
		err := usecases.DeleteOrderByID(storage, orderID)
		if err != nil {
			GinkgoWriter.Write([]byte(err.Error()))
			Expect(err).NotTo(HaveOccurred())
		}
	})
})
