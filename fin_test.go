package financial

import (
	"fmt"
	"log"
	"testing"
)

func TestNPV(t *testing.T) {
	npv, err := NPV(0.281, []float64{-100, 39, 59, 55, 20})
	if err != nil {
		t.Fatal(err.Error())
	}
	s := fmt.Sprintf("%.5f", npv)
	log.Printf("npv: %s", s)
	if s != "-0.00848" {
		t.Fail()
	}
}

func TestIRR(t *testing.T) {
	irr, err := IRR([]float64{-100, 39, 59, 55, 20})
	if err != nil {
		t.Fatal(err.Error())
	}
	s := fmt.Sprintf("%.5f", irr)
	log.Printf("irr: %s", s)
	if s != "0.28095" {
		t.Fail()
	}
}
