package financial

import (
	"fmt"
	"testing"
)

var fvTestCases = []struct {
	presentValue       float64
	interestRate       float64
	periods            float64
	compoundedAnnually bool
	want               float64
}{
	{
		presentValue:       100.0,
		interestRate:       0.07,
		periods:            10.0,
		compoundedAnnually: false,
		want:               170.00000000000003,
	},
	{
		presentValue:       100.0,
		interestRate:       0.07,
		periods:            10.0,
		compoundedAnnually: true,
		want:               196.71513572895657,
	},
}

func TestFV(t *testing.T) {
	for _, tt := range fvTestCases {
		got := FV(tt.presentValue, tt.interestRate, tt.compoundedAnnually, tt.periods)
		if got != tt.want {
			t.Errorf("FV(%v, %v, %v) = %v; want %v",
				tt.presentValue,
				tt.interestRate,
				tt.periods,
				got,
				tt.want,
			)
		}
	}
}

var fvBenchResult float64

func BenchmarkFV(b *testing.B) {
	var r float64

	presentValue := 100.0
	interestRate := 0.07
	periods := 10.0
	compoundedAnnually := false

	for n := 0; n < b.N; n++ {
		r = FV(presentValue, interestRate, compoundedAnnually, periods)
	}

	fvBenchResult = r
}

func ExampleFV() {
	presentValue := 100.0
	interestRate := 0.07
	periods := 10.0
	compoundedAnnually := false

	fv := FV(presentValue, interestRate, compoundedAnnually, periods)
	fmt.Printf("FV is: %f", fv)
	// Output:
	// FV is: 170.000000
}
