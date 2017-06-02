package financial

import (
	"fmt"
	"math"
)

const irrMaxInterations = 20
const irrAccuracy = 1E-7
const irrInitialGuess = 0

// NPV returns the NPV (Net Present Value) of a cash flow series.
func NPV(rate float64, values []float64) (float64, error) {
	if len(values) == 0 {
		return -1, fmt.Errorf("values must include the initial investment (usually negative number) and period cash flows")
	}
	npv := values[0]
	values = values[1:]
	for i, v := range values {
		npv += (v / math.Pow(1+rate, float64(i+1)))
	}
	return npv, nil
}

// IRR return the Internal Rate of Return (IRR).
func IRR(values []float64) (float64, error) {
	if len(values) == 0 {
		return -1, fmt.Errorf("values must include the initial investment (usually negative number) and period cash flows")
	}
	x0 := float64(irrInitialGuess)
	var x1 float64
	for i := 0; i < irrMaxInterations; i++ {
		fValue := float64(0)
		fDerivative := float64(0)
		for k := 0; k < len(values); k++ {
			fk := float64(k)
			fValue += values[k] / math.Pow(1.0+x0, fk)
			fDerivative += -fk * values[k] / math.Pow(1.0+x0, fk+1.0)
		}
		x1 = x0 - fValue/fDerivative

		if math.Abs(x1-x0) <= irrAccuracy {
			return x1, nil
		}
		x0 = x1
	}
	return -1, fmt.Errorf("could not find irr for the provided values")
}
