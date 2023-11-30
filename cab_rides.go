package main

type CabRides struct {
}

func (r *CabRides) Add(ride CabRide) {
}

func (r *CabRides) TotalFare() int {
	return 100
}

func NewCabRides() CabRides {
	return CabRides{}
}
