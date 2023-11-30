package main

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("CabRides", func() {
	It("should calculate total fare as 100 for rides of 6km and 4km", func() {
		cabRides := NewCabRides()
		cabRides.Add(NewCabRide(6, 0))
		cabRides.Add(NewCabRide(4, 0))

		Expect(cabRides.TotalFare()).To(Equal(100))
	})
})
