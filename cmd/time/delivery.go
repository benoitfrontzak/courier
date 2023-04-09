package main

import (
	"golang.org/x/exp/slices"
)

type vehicle struct {
	returnTime float64
}

type delivery struct {
	id       int
	vehicle  int
	packages []string
}

// What's the best delivery packaging for a list a packages for a specific max capacity
// return [PKG2 PKG4] 185kg **
func oneDeliveryPackaging(capacity int, packages []packageDelivery) ([]string, int) {
	n := len(packages)

	// initialize the dp table
	dp := make([][]int, n+1)
	for i := range dp {
		dp[i] = make([]int, capacity+1)
	}

	// fill the dp table
	for i := 1; i <= n; i++ {
		for w := 1; w <= capacity; w++ {
			if packages[i-1].weight <= w {
				dp[i][w] = max(dp[i-1][w], dp[i-1][w-packages[i-1].weight]+packages[i-1].weight)
			} else {
				dp[i][w] = dp[i-1][w]
			}
		}
	}

	// find the indices of the selected packages
	indices := make([]string, 0)
	w := capacity
	for i := n; i >= 1; i-- {
		if dp[i][w] != dp[i-1][w] {
			indices = append(indices, packages[i-1].id)
			w -= packages[i-1].weight
		}
	}

	// reverse the indices and return them along with the total weight
	reverse(indices)
	return indices, dp[n][capacity]
}

// What's the max number **
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// Reverse the indices **
func reverse(a []string) {
	n := len(a)
	for i := 0; i < n/2; i++ {
		j := n - i - 1
		a[i], a[j] = a[j], a[i]
	}
}

// What are the best deliveries packaging for a list a packages for a specific max capacity
// returns [[PKG2 PKG4] [PKG3] [PKG5] [PKG1]]
func allDeliveriesPackaging(capacity int, packages []packageDelivery) [][]string {
	// while there is still packages to be packed
	for len(packages) > 0 {
		oneDelivery, _ := oneDeliveryPackaging(capacity, packages)
		deliveries = append(deliveries, oneDelivery)
		// for each package
		for _, d := range oneDelivery {
			// find index of pkg thanks to myPackages dictionary
			myIndex := slices.IndexFunc(packages, func(pd packageDelivery) bool {
				return pd.id == d
			})
			// remove onedeliveries from packages
			packages = append(packages[:myIndex], packages[myIndex+1:]...)
		}
	}
	return deliveries
}

// What's the delivery distribution per vehicle (by returning time)
// returns [
// {deliveryID:0 vehicleID:0 packages:[PKG2 PKG4]}
// {deliveryID:1 vehicleID:1 packages[PKG3]}
// {deliveryID:2 vehicleID:1 packages[PKG5]}
// {deliveryID:3 vehicleID:0 packages[PKG1]}
// ]
func allDeliveriesOrder(vehicleCount, maxSpeed int, deliveries [][]string) []delivery {
	vehicles := make([]vehicle, vehicleCount)

	// store all deliveries by priority order and vehicle
	var allDeliveries = []delivery{}

	for i, dlv := range deliveries {
		mt := maxTimeDelivery(dlv, maxSpeed)
		// what's the next vehicle available
		nextv := minReturnTimes(vehicles)
		myDelivery := delivery{
			id:       i,
			vehicle:  nextv,
			packages: dlv,
		}
		allDeliveries = append(allDeliveries, myDelivery)
		// increment next vehicle return time
		vehicles[nextv].returnTime += mt * 2
	}

	return allDeliveries
}

// What's the maximum time of the delivery
// [PKG2 PKG4], 70  returns 1.78 (125km / 70km/h)
// so to have the returned time must times 2
func maxTimeDelivery(deliveryIDs []string, maxSpeed int) float64 {
	var mt float64
	for _, id := range deliveryIDs {
		dp := myPackages[id]
		time := dp.computeTime(maxSpeed)

		if time > mt {
			mt = time
		}
	}

	return mt
}

// What's the next vehicle to be back?
func minReturnTimes(vehicles []vehicle) int {
	minidx := 0
	for i, v := range vehicles {
		if v.returnTime < vehicles[minidx].returnTime {
			minidx = i
		}
	}
	return minidx
}
