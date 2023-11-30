package main

type CabRides struct {
	cabRides []CabRide
}

func (r *CabRides) Add(ride CabRide) {
	r.cabRides = append(r.cabRides, ride)
}

func (r *CabRides) TotalFare() int {
	total := 0
	for i := 0; i < len(r.cabRides); i++ {
		total += r.cabRides[i].Fare()
	}
	return total
}

func NewCabRides() CabRides {
	return CabRides{
		cabRides: make([]CabRide, 0),
	}
}
