package main

const (
	FarePerKm     = 10
	FarePerMinute = 2
)

type CabRide struct {
	distance int
	waitTime int
}

func NewCabRide(distance int, waitTime int) CabRide {
	return CabRide{distance: distance, waitTime: waitTime}
}

func (r *CabRide) Fare() int {
	return r.distance*FarePerKm + r.waitTime*FarePerMinute
}
