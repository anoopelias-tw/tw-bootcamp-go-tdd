package main

type CabRides struct {
	cabRides []CabRide
}

type Invoice struct {
	TotalRides int
	TotalFare  int
	Average    float64
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

func (r *CabRides) invoice() Invoice {
	return Invoice{len(r.cabRides), r.TotalFare(), float64(r.TotalFare()) / float64(len(r.cabRides))}
}

func NewCabRides() CabRides {
	return CabRides{
		cabRides: make([]CabRide, 0),
	}
}
