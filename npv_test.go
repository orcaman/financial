package financial

import (
	"fmt"
	"testing"
)

var npvTestCases = []struct {
	ok     bool
	r      float64
	values []float64
	want   float64
}{
	{
		ok:     true,
		r:      0.281,
		values: []float64{-100, 39, 59, 55, 20},
		want:   -0.00847859163845488,
	},
	{
		ok:     false,
		r:      0.281,
		values: nil,
		want:   0,
	},
	{
		ok:     false,
		r:      0.281,
		values: []float64{42},
		want:   0,
	},
}

func TestNPV(t *testing.T) {
	for _, tt := range npvTestCases {
		got, err := NPV(tt.r, tt.values)

		if !tt.ok {
			if err == nil || got != tt.want {
				t.Errorf("NPV(%v, %v) = %v, %v; want %v, %v",
					tt.r,
					tt.values,
					got,
					err,
					tt.want,
					"<error>",
				)
			}
			continue
		}

		if got != tt.want {
			t.Errorf("NPV(%v, %v) = %v, %v; want %v, %v",
				tt.r,
				tt.values,
				got,
				err,
				tt.want,
				nil,
			)
		}
	}
}

var npvBenchResult float64

func BenchmarkNPV(b *testing.B) {
	var res float64

	r := 0.281
	values := []float64{-100, 39, 59, 55, 20}

	for n := 0; n < b.N; n++ {
		res, _ = NPV(r, values)
	}

	npvBenchResult = res
}

func ExampleNPV() {
	initialInvestment := 100.0

	cashFlowYear1 := 39.0
	cashFlowYear2 := 59.0
	cashFlowYear3 := 55.0
	cashFlowYear4 := 20.0

	discountRate := 0.281

	npv, err := NPV(discountRate, []float64{-initialInvestment, cashFlowYear1, cashFlowYear2, cashFlowYear3, cashFlowYear4})
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Printf("NPV is: %f", npv)
	// Output:
	// NPV is: -0.008479
}
