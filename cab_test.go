package main

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("CabRide", func() {
	It("should calculate fare as 50 when distance is 5", func() {
		cabRide := NewCabRide(5, 0)
		fare := cabRide.Fare()
		Expect(fare).To(Equal(50))
	})
	It("should calculate fare as 60 when distance is 6", func() {
		cabRide := NewCabRide(6, 0)
		fare := cabRide.Fare()
		Expect(fare).To(Equal(60))
	})
	It("should calculate fare as 6 when waiting time is 3 minutes", func() {
		cabRide := NewCabRide(0, 30)
		Expect(cabRide.Fare()).To(Equal(60))
	})
	It("should calculate minimum fare as 40", func() {
		cabRide := NewCabRide(1, 1)
		Expect(cabRide.Fare()).To(Equal(40))
	})
})
