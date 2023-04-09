package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

// structure holding a package delivery information
type packageDelivery struct {
	id       string // package id
	weight   int    // package weight (in kg)
	distance int    // package distance (in km)
	offer    string // offer code

	cost     int // delivery cost
	discount int // discount to apply
	final    int // final cost
}

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

// function reading the study case and printing the expected output
func process(w io.Writer, r io.Reader) error {
	var (
		baseDeliveryCost int
		pkgCount         int
		lineCount        = 1
	)

	// Read header (must have only two parameters: baseDeliveryCost & pkgCount)
	n, err := fmt.Fscanf(r, "%d %d\n", &baseDeliveryCost, &pkgCount)
	if n != 2 || err != nil {
		return fmt.Errorf("malformed header at line %d: %v n=%d", lineCount, err, n)
	}

	lineCount++

	// Read packages (must have only four parameters: pkg.id & pkg.weight & pkg.distance & coupon)
	for i := 0; i < pkgCount; i++ {
		pd := packageDelivery{}
		n, err := fmt.Fscanf(r, "%s %d %d %s\n", &pd.id, &pd.weight, &pd.distance, &pd.offer)
		if n != 4 || err != nil {
			return fmt.Errorf("malformed package delivery at line %d: %v", lineCount, err)
		}
		pd.computeCost(baseDeliveryCost)

		fmt.Fprintf(w, "%s %d %d\n", pd.id, pd.discount, pd.final)
		lineCount++
	}

	return nil
}

func main() {
	err := process(os.Stdout, os.Stdin)
	if err != nil {
		log.Fatal(err)
	}
}
