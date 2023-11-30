package main

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("CabRide", func() {
	It("should calculate fare as 50 when distance is 5", func() {
		cabRide := NewCabRide(5)
		fare := cabRide.Fare()
		Expect(fare).To(Equal(50))
	})
	It("should calculate fare as 60 when distance is 6", func() {
		cabRide := NewCabRide(6)
		fare := cabRide.Fare()
		Expect(fare).To(Equal(60))
	})
})
