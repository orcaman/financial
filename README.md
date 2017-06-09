# financial [![CircleCI](https://circleci.com/gh/orcaman/financial.svg?style=svg)](https://circleci.com/gh/orcaman/financial) [![GoDoc](https://godoc.org/github.com/orcaman/financial?status.svg)](https://godoc.org/github.com/orcaman/financial)

<img src="https://user-images.githubusercontent.com/4884073/26978301-b620e944-4d33-11e7-9879-2b9b3a6aa444.png" width="380">

package `financial` features common financial functions, and is intended to be a golang version of [numpy's financial functions](https://docs.scipy.org/doc/numpy/reference/routines.financial.html).

## Supported Functions:
- NPV(r, values): returns the NPV (Net Present Value) of a cash flow series given an interest rate r, or an error. See [usage example](https://godoc.org/github.com/orcaman/financial#example-NPV).
- IRR(values): returns the Internal Rate of Return (using newton raphson approximation). See [usage example](https://godoc.org/github.com/orcaman/financial#example-IRR).
- FV(pv, r, compoundedAnnually, n): gets the present value, the interest rate (compounded annually if needed) and the number of periods and returns the future value. See [usage example](https://godoc.org/github.com/orcaman/financial#example-fv).

## TODO:
- pv(rate, nper, pmt[, fv, when])	Compute the present value.
- pmt(rate, nper, pv[, fv, when])	Compute the payment against loan principal plus interest.
- ppmt(rate, per, nper, pv[, fv, when])	Compute the payment against loan principal.
- ipmt(rate, per, nper, pv[, fv, when])	Compute the interest portion of a payment.
- mirr(values, finance_rate, reinvest_rate)	Modified internal rate of return.
- nper(rate, pmt, pv[, fv, when])	Compute the number of periodic payments.
- rate(nper, pmt, pv, fv[, when, guess, tol, ...])	Compute the rate of interest per period.