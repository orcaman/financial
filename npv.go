package financial

import (
	"errors"
	"math"
)

// NPV returns the NPV (Net Present Value) of a cash flow series given an interest rate r, or an error.
func NPV(r float64, values []float64) (float64, error) {
	if len(values) < 2 {
		return 0, errors.New("values must include the initial investment (usually negative number) and period cash flows")
	}
	npv := values[0]
	values = values[1:]
	for i, v := range values {
		npv += (v / math.Pow(1+r, float64(i+1)))
	}
	return npv, nil
}
