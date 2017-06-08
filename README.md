# financial [![CircleCI](https://circleci.com/gh/orcaman/financial.svg?style=svg)](https://circleci.com/gh/orcaman/financial) [![GoDoc](https://godoc.org/github.com/orcaman/financial?status.svg)](https://godoc.org/github.com/orcaman/financial)

package `financial` features common financial functions, and is intended to be a golang version of [numpy's financial functions](https://docs.scipy.org/doc/numpy/reference/routines.financial.html).

## Supported Functions:
- npv(rate, values)	Returns the NPV (Net Present Value) of a cash flow series. See [usage example](https://godoc.org/github.com/orcaman/financial#ex-NPV).
- irr(values)	Return the Internal Rate of Return (using newton raphson approximation). See [usage example](https://godoc.org/github.com/orcaman/financial#ex-IRR).


## TODO:
- fv(rate, nper, pmt, pv[, when])	Compute the future value.
- pv(rate, nper, pmt[, fv, when])	Compute the present value.
- pmt(rate, nper, pv[, fv, when])	Compute the payment against loan principal plus interest.
- ppmt(rate, per, nper, pv[, fv, when])	Compute the payment against loan principal.
- ipmt(rate, per, nper, pv[, fv, when])	Compute the interest portion of a payment.
- mirr(values, finance_rate, reinvest_rate)	Modified internal rate of return.
- nper(rate, pmt, pv[, fv, when])	Compute the number of periodic payments.
- rate(nper, pmt, pv, fv[, when, guess, tol, ...])	Compute the rate of interest per period.