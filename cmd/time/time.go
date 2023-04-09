package main

import (
	"math"
)

// Delivery Cost = Base Delivery Cost + (Package Total Weight * 10) + (Distance to Destination * 5)
func (pd *packageDelivery) computeTime(maxSpeed int) float64 {
	// Calculate delivery time
	x := float64(pd.distance) / float64(maxSpeed)
	return math.Floor(x*100) / 100
}

func packagesDeliveryTime(pkgs []string, currentTime float64, maxSpeed int) {
	for _, p := range pkgs {
		myPackage := myPackages[p]
		myPackageTime := myPackage.computeTime(maxSpeed) + currentTime
		finalTime[p] = math.Ceil(myPackageTime*100) / 100
		// log.Println("package ", myPackage.id, " time is ", math.Ceil(myPackageTime*100)/100)
	}
}
