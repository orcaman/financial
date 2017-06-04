package financial

import "testing"

var npvTestCases = []struct {
	ok     bool
	rate   float64
	values []float64
	want   float64
}{
	{
		ok:     true,
		rate:   0.281,
		values: []float64{-100, 39, 59, 55, 20},
		want:   -0.00847859163845488,
	},
	{
		ok:     false,
		rate:   0.281,
		values: nil,
		want:   0,
	},
	{
		ok:     false,
		rate:   0.281,
		values: []float64{42},
		want:   0,
	},
}

func TestNPV(t *testing.T) {
	for _, tt := range npvTestCases {
		got, err := NPV(tt.rate, tt.values)

		if !tt.ok {
			if err == nil || got != tt.want {
				t.Errorf("NPV(%v, %v) = %v, %v; want %v, %v",
					tt.rate,
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
				tt.rate,
				tt.values,
				got,
				err,
				tt.want,
				nil,
			)
		}
	}
}
