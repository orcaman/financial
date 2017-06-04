package financial

import "testing"

var irrTestCases = []struct {
	ok     bool
	values []float64
	want   float64
}{
	{
		ok:     true,
		values: []float64{-100, 39, 59, 55, 20},
		want:   0.2809484211599531,
	},
	{
		ok:     false,
		values: nil,
		want:   0,
	},
	{
		ok:     false,
		values: []float64{42},
		want:   0,
	},
}

func TestIRR(t *testing.T) {
	for _, tt := range irrTestCases {
		got, err := IRR(tt.values)

		if !tt.ok {
			if err == nil || got != tt.want {
				t.Errorf("IRR(%v) = %v, %v; want %v, %v",
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
			t.Errorf("IRR(%v) = %v, %v; want %v, %v",
				tt.values,
				got,
				err,
				tt.want,
				nil,
			)
		}
	}
}

var irrBenchResult float64

func BenchmarkIRR(b *testing.B) {
	var r float64

	values := []float64{-100, 39, 59, 55, 20}

	for n := 0; n < b.N; n++ {
		r, _ = IRR(values)
	}

	irrBenchResult = r
}
