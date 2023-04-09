package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"sort"
)

// all deliveries packaging by max vehicle capacity
// returns [[PKG2 PKG4] [PKG3] [PKG5] [PKG1]]
var deliveries [][]string

// dictionary to store package delivery information per package id
var myPackages = make(map[string]*packageDelivery)

// store package final delivery time
var finalTime = make(map[string]float64)

func process(w io.Writer, r io.Reader) error {
	// decode io.Reader to baseDelivery structure,
	// slice of packageDelivery structure,
	// vehicle structure,
	// handle potential errors
	bd, pkgs, vd, err := decodeReader(r)
	if err != nil {
		fmt.Println(err)
	}

	// Define a copy of pkgs
	toBePacked := make([]packageDelivery, len(pkgs))
	copy(toBePacked, pkgs)
	// Sort packages: by lightest, then by closest, then by unique id (so we get a stable sort).
	sortPkgs(toBePacked)

	// all deliveries packaging by max vehicle capacity
	// returns [[PKG2 PKG4] [PKG3] [PKG5] [PKG1]]
	deliveries = allDeliveriesPackaging(vd.maxWeight, toBePacked)

	// distribution of deliveries to each vehicle (based on vehicle returned time)
	// returns [
	// {deliveryID:0 vehicleID:0 packages:[PKG2 PKG4]}
	// {deliveryID:1 vehicleID:1 packages[PKG3]}
	// {deliveryID:2 vehicleID:1 packages[PKG5]}
	// {deliveryID:3 vehicleID:0 packages[PKG1]}
	// ]
	orderedDeliveries := allDeliveriesOrder(vd.vehicleCount, vd.maxSpeed, deliveries)

	// calculate final time of each package
	for i := 0; i < len(orderedDeliveries); i++ {
		currentTime := 0.0
		for j := i + 1; j < len(orderedDeliveries); j++ {
			if orderedDeliveries[i].vehicle == orderedDeliveries[j].vehicle {
				packagesDeliveryTime(orderedDeliveries[i].packages, currentTime, vd.maxSpeed)
				currentTime += maxTimeDelivery(orderedDeliveries[i].packages, vd.maxSpeed) * 2
				packagesDeliveryTime(orderedDeliveries[j].packages, currentTime, vd.maxSpeed)
			}
		}

	}

	// populate pkgs
	for _, p := range pkgs {
		p.computeCost(bd.baseCost)
		p.time = finalTime[p.id]
		fmt.Fprintf(w, "%s %d %d %.2f\n", p.id, p.discount, p.final, p.time)
	}

	return nil
}

// Sort packages: by lightest, then by closest, then by unique id (so we get a stable sort).
func sortPkgs(pkgs []packageDelivery) []packageDelivery {
	sort.Slice(pkgs, func(i, j int) bool {
		if pkgs[i].weight != pkgs[j].weight {
			return pkgs[i].weight < pkgs[j].weight
		}
		if pkgs[i].distance != pkgs[j].distance {
			return pkgs[i].distance < pkgs[j].distance
		}
		return pkgs[i].id < pkgs[j].id
	})

	return pkgs
}

func main() {
	err := process(os.Stdout, os.Stdin)
	if err != nil {
		log.Fatal(err)
	}
}
