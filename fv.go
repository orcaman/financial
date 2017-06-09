package financial

import "math"

// FV gets the present value, the interest rate (compounded annually if needed) and the number of periods and returns the future value.
func FV(pv, r float64, compoundedAnnually bool, n float64) float64 {
	if !compoundedAnnually {
		return pv * (1 + (r * n))
	}
	return pv * math.Pow(1+r, n)
}
