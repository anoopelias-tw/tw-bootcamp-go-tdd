package main

type CabRide struct {
	distance int
	waitTime int
}

func NewCabRide(distance int, waitTime int) CabRide {
	return CabRide{distance: distance, waitTime: waitTime}
}

func (r *CabRide) Fare() int {
	return r.distance*10 + r.waitTime*2
}
