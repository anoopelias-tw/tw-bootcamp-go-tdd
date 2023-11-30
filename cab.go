package main

const (
	FarePerKm     = 10
	FarePerMinute = 2
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
	fare := r.distance*FarePerKm + r.waitTime*FarePerMinute
	if fare < minimumFare {
		return minimumFare
	}
	return fare
}
