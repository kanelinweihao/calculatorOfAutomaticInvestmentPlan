package calc

import (
	// "fmt"
	// "math/big"
	"github.com/shopspring/decimal"
	"go.lwh.com/linweihao/calculatorOfAutomaticInvestmentPlan/utils/err"
)

func Add(x string, y string) (z string) {
	X, errDeciaml := decimal.NewFromString(x)
	err.ErrCheck(errDeciaml)
	Y, errDeciaml := decimal.NewFromString(y)
	err.ErrCheck(errDeciaml)
	Z := X.Add(Y)
	z = Z.String()
	return z
}

func Sub(x string, y string) (z string) {
	X, errDeciaml := decimal.NewFromString(x)
	err.ErrCheck(errDeciaml)
	Y, errDeciaml := decimal.NewFromString(y)
	err.ErrCheck(errDeciaml)
	Z := X.Sub(Y)
	z = Z.String()
	return z
}

func Mul(x string, y string) (z string) {
	X, errDeciaml := decimal.NewFromString(x)
	err.ErrCheck(errDeciaml)
	Y, errDeciaml := decimal.NewFromString(y)
	err.ErrCheck(errDeciaml)
	Z := X.Mul(Y)
	z = Z.String()
	return z
}

func Div(x string, y string) (z string) {
	X, errDeciaml := decimal.NewFromString(x)
	err.ErrCheck(errDeciaml)
	Y, errDeciaml := decimal.NewFromString(y)
	err.ErrCheck(errDeciaml)
	Z := X.Div(Y)
	z = Z.String()
	return z
}
