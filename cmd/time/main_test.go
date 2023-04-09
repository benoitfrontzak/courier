package main

import (
	"bytes"
	"strings"
	"testing"
)

func TestTimeDelivery(t *testing.T) {
	r := strings.NewReader(`100 5
PKG1 50 30 OFR001
PKG2 75 125 OFR008
PKG3 175 100 OFR003
PKG4 110 60 OFR003
PKG5 155 95 NA
2 70 200
`)
	buf := &bytes.Buffer{}
	if err := process(buf, r); err != nil {
		t.Fatal(err)
	}

	want := `PKG1 0 750 3.98
PKG2 0 1475 1.78
PKG3 0 2350 1.42
PKG4 105 1395 0.85
PKG5 0 2125 4.19
`

	if buf.String() != want {
		t.Fatalf("got:\n%s\nwant:\n%s", buf.String(), want)
	}
}
