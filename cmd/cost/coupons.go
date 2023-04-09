package main

// stucture holding interval (min & max)
type interval struct {
	min int
	max int
}

// structure holding one coupon offer
type offer struct {
	discount int      // offer discount (in %)
	distance interval // distance required (in km)
	weight   interval // weight required (in kg)
}

// variable storing all available coupon offers
var offers = map[string]offer{
	"OFR001": {discount: 10, distance: interval{0, 200}, weight: interval{70, 200}},
	"OFR002": {discount: 7, distance: interval{50, 150}, weight: interval{150, 250}},
	"OFR003": {discount: 5, distance: interval{50, 250}, weight: interval{10, 150}},
	//
	// add additional offer here
	//
}

// function checking if a coupon offer is valid for the package delivery
func isValidOffer(code string, distance, weight int) (offer, bool) {
	off, ok := offers[code]
	if !ok {
		return offer{}, false
	}
	if distance < off.distance.min || distance > off.distance.max {
		return offer{}, false
	}
	if weight < off.weight.min || weight > off.weight.max {
		return offer{}, false
	}
	return off, true
}
