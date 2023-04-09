package main

// Delivery Cost = Base Delivery Cost + (Package Total Weight * 10) + (Distance to Destination * 5)
func (pd *packageDelivery) computeCost(base int) {
	// First calculate cost without discount
	pd.cost = base + (pd.weight * 10) + (pd.distance * 5)

	// Check offer code is valid
	if off, ok := isValidOffer(pd.offer, pd.distance, pd.weight); ok {
		pd.discount = int(float64(pd.cost) * (float64(off.discount) / 100))
	}

	pd.final = pd.cost - pd.discount
}
