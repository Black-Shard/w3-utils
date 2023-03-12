package utils

import (
	"math/big"
)

func AddMany(x ...*big.Int) *big.Int {
	z := new(big.Int).SetInt64(0)
	for _, n := range x {
		z.Add(z, n)
	}
	return z
}

func SubMany(x ...*big.Int) *big.Int {
	z := new(big.Int).Set(x[0])
	for _, n := range x[1:] {
		z.Sub(z, n)
	}
	return z
}

// Multiply many big ints together
func MulMany(x ...*big.Int) (z *big.Int) {
	z = new(big.Int).SetInt64(1)
	for _, n := range x {
		z.Mul(z, n)
	}
	return z
}

func DivMany(x ...*big.Int) *big.Int {
	z := new(big.Int).Set(x[0])
	for i, n := range x {
		if i > 0 {
			z.Div(z, n)
		}
	}
	return z
}

// Always done in 18 decimals
// Directly applies the multiplication result onto x
func MulWadDirect(x *big.Int, y *big.Int) {
	base := new(big.Int).SetInt64(10)
	divisor := base.Exp(base, big.NewInt(int64(18)), nil)
	x.Mul(x, y).Div(x, divisor)
}

// Always done in 18 decimals
// Returns a new int z
func MulWad(x *big.Int, y *big.Int) (z *big.Int) {
	z = new(big.Int).SetInt64(0)
	base := new(big.Int).SetInt64(10)
	divisor := base.Exp(base, big.NewInt(int64(18)), nil)
	z.Mul(x, y)
	z.Div(z, divisor)
	return z
}

// Always done in 18 decimals
func DivWadDirect(x *big.Int, y *big.Int) {
	base := new(big.Int).SetInt64(10)
	multiplier := base.Exp(base, big.NewInt(int64(18)), nil)
	x.Mul(x, multiplier).Div(x, y)
}

// Always done in 18 decimals
func DivWad(x *big.Int, y *big.Int) (z *big.Int) {
	z = new(big.Int).SetInt64(0)
	base := new(big.Int).SetInt64(10)
	multiplier := base.Exp(base, big.NewInt(int64(18)), nil)
	z.Mul(x, multiplier)
	z.Div(z, y)
	return z
}

func DecimalUp(x *big.Int, factor int) {
	base := new(big.Int).SetInt64(10)
	exponent := big.NewInt(int64(factor))
	base.Exp(base, exponent, nil)
	x.Mul(x, base)
}

func DecimalDown(x *big.Int, factor int) {
	base := new(big.Int).SetInt64(10)
	exponent := big.NewInt(int64(factor))
	base.Exp(base, exponent, nil)
	x.Div(x, base)
}

func GetBigInt(val string) *big.Int {
	i := new(big.Int)
	i.SetString(val, 10)
	return i
}

func CompBigInt(x *big.Int, y *big.Int) bool {
	if x.Cmp(y) == 0 {
		return true
	} else {
		return false
	}
}

// Output to 2 decimals
func DecimalsToReadable(x *big.Int, decimals int) string {
	y := big.NewInt(0)
	y.Set(x)
	DecimalDown(y, decimals-2)
	s := y.String()

	var out string
	if len(s) > 2 {
		out = s[:len(s)-2] + "." + s[len(s)-2:]
	} else if len(s) == 2 {
		out = "0" + "." + s[len(s)-2:]
	} else if len(s) == 1 && s != "0" {
		out = "0" + ".0" + s[len(s)-1:]
	} else {
		out = "<0.01"
	}
	return out
}
