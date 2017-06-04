package financial

import (
	"errors"
	"math"
)

// NPV returns the NPV (Net Present Value) of a cash flow series.
func NPV(rate float64, values []float64) (float64, error) {
	if len(values) < 2 {
		return 0, errors.New("values must include the initial investment (usually negative number) and period cash flows")
	}
	npv := values[0]
	values = values[1:]
	for i, v := range values {
		npv += (v / math.Pow(1+rate, float64(i+1)))
	}
	return npv, nil
}
