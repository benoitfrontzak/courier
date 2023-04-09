package main

import (
	"bytes"
	"strings"
	"testing"
)

func TestCostDelivery(t *testing.T) {
	r := strings.NewReader(`100 3
PKG1 5 5 OFR001
PKG2 15 5 OFR002
PKG3 10 100 OFR003
`)
	buf := &bytes.Buffer{}
	if err := process(buf, r); err != nil {
		t.Fatal(err)
	}

	want := `PKG1 0 175
PKG2 0 275
PKG3 35 665
`

	if buf.String() != want {
		t.Fatalf("got:\n%s\nwant:\n%s", buf.String(), want)
	}
}

func Test_computeCost(t *testing.T) {
	tests := []struct {
		// input
		base     int
		weight   int
		distance int
		offer    string

		// output
		wantCost      int
		wantDiscount  int
		wantFinalCost int
	}{
		{
			base: 100, weight: 5, distance: 5, offer: "OFR001",
			wantCost: 175, wantDiscount: 0, wantFinalCost: 175,
		},
		{
			base: 100, weight: 10, distance: 100, offer: "OFR003",
			wantCost: 700, wantDiscount: 35, wantFinalCost: 665,
		},
	}

	for _, tt := range tests {
		pd := packageDelivery{weight: tt.weight, distance: tt.distance, offer: tt.offer}
		pd.computeCost(tt.base)
		if tt.wantCost != pd.cost {
			t.Errorf("cost = %d, want %d", pd.cost, tt.wantCost)
		}
		if tt.wantDiscount != pd.discount {
			t.Errorf("discount = %d, want %d", pd.discount, tt.wantDiscount)
		}
		if tt.wantFinalCost != pd.final {
			t.Errorf("final = %d, want %d", pd.final, tt.wantFinalCost)
		}
	}
}
