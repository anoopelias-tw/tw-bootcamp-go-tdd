package main

type CabRide struct {
	distance int
}

func NewCabRide(distance int) CabRide {
	return CabRide{distance: distance}
}

func (r *CabRide) Fare() int {
	return r.distance * 10
}
