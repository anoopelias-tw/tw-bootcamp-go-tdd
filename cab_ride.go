package main

const (
	farePerKm     = 10
	farePerMinute = 2
	minimumFare   = 40
)

type CabRide struct {
	distance int
	waitTime int
}

func NewCabRide(distance int, waitTime int) CabRide {
	return CabRide{distance: distance, waitTime: waitTime}
}

func (r *CabRide) Fare() int {
	fare := r.distance*farePerKm + r.waitTime*farePerMinute
	if fare < minimumFare {
		return minimumFare
	}
	return fare
}
