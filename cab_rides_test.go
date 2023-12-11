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
	It("should calculate total fare as 110 for rides of 7km and 4km", func() {
		cabRides := NewCabRides()
		cabRides.Add(NewCabRide(7, 0))
		cabRides.Add(NewCabRide(4, 0))

		Expect(cabRides.TotalFare()).To(Equal(110))
	})
	It("should generate Invoice for a ride", func() {
		cabRides := NewCabRides()
		cabRides.Add(NewCabRide(6, 0))

		invoice := cabRides.Invoice()
		Expect(invoice).To(Equal(Invoice{1, 60, 60.0}))
	})
	It("should generate Invoice for multiple rides", func() {
		cabRides := NewCabRides()
		cabRides.Add(NewCabRide(6, 0))
		cabRides.Add(NewCabRide(7, 0))

		invoice := cabRides.Invoice()
		Expect(invoice).To(Equal(Invoice{2, 130, 65.0}))
	})
})
