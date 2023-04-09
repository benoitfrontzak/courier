package main

import (
	"fmt"
	"io"
)

type baseDelivery struct {
	baseCost int
	pkgCount int
}

type packageDelivery struct {
	id       string // package id
	weight   int    // package weight (in kg)
	distance int    // package distance (in km)
	offer    string // offer code

	cost     int     // delivery cost
	discount int     // discount to apply
	final    int     // final cost
	time     float64 // delivery time
}

type vehicleDelivery struct {
	vehicleCount int
	maxSpeed     int
	maxWeight    int
}

func decodeReader(r io.Reader) (baseDelivery, []packageDelivery, vehicleDelivery, error) {
	lineCount := 1

	var bd baseDelivery
	// Read header
	n, err := fmt.Fscanf(r, "%d %d\n", &bd.baseCost, &bd.pkgCount)
	if n != 2 || err != nil {
		return bd, []packageDelivery{}, vehicleDelivery{}, fmt.Errorf("malformed header at line %d: %v", lineCount, err)
	}
	lineCount++

	pkgs := make([]packageDelivery, 0, bd.pkgCount)
	// Read packages
	for i := 0; i < bd.pkgCount; i++ {
		pd := packageDelivery{}
		n, err := fmt.Fscanf(r, "%s %d %d %s\n", &pd.id, &pd.weight, &pd.distance, &pd.offer)
		if n != 4 || err != nil {
			return bd, pkgs, vehicleDelivery{}, fmt.Errorf("malformed package delivery at line %d: %v", lineCount, err)
		}

		pkgs = append(pkgs, pd)
		myPackages[pd.id] = &pd

		lineCount++
	}

	var vd vehicleDelivery
	// Read footer
	n, err = fmt.Fscanf(r, "%d %d %d\n", &vd.vehicleCount, &vd.maxSpeed, &vd.maxWeight)
	if n != 3 || err != nil {
		return bd, pkgs, vd, fmt.Errorf("malformed footer at line %d: %v", lineCount, err)
	}

	return bd, pkgs, vd, nil
}
