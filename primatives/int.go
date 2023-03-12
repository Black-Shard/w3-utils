package utils

import (
	"math/big"

	"github.com/lmittmann/w3"
)

type EthInt struct {
	Int *big.Int
}

var wad *big.Int = w3.I("1 ether")

func NewEthInt(val string) EthInt {
	b := new(big.Int)
	b.SetString(val, 10)
	return EthInt{
		Int: b,
	}
}

func NewEthWad(val string) EthInt {
	e := NewEthInt(val)
	w := EthInt{
		Int: wad,
	}
	e.MulDirect(w)
	return e
}

func (e *EthInt) Pow(exp string) {
	x := new(big.Int)
	x.SetString(exp, 10)
	e.Int.Exp(e.Int, x, nil)
}

// Returns a new Eth Int
func (e EthInt) Add(y EthInt) EthInt {
	z := big.NewInt(0)
	z.Add(e.Int, y.Int)
	return EthInt{
		Int: z,
	}
}

func (e EthInt) Sub(y EthInt) EthInt {
	z := big.NewInt(0)
	z.Sub(e.Int, y.Int)
	return EthInt{
		Int: z,
	}
}

func (e EthInt) Mul(y EthInt) EthInt {
	z := big.NewInt(0)
	z.Mul(e.Int, y.Int)
	return EthInt{
		Int: z,
	}
}

func (e EthInt) Div(y EthInt) EthInt {
	z := big.NewInt(0)
	z.Div(e.Int, y.Int)
	return EthInt{
		Int: z,
	}
}

// Direct Ops
func (e *EthInt) AddDirect(y EthInt) {
	e.Int.Add(e.Int, y.Int)
}

func (e *EthInt) SubDirect(y EthInt) {
	e.Int.Sub(e.Int, y.Int)
}

func (e *EthInt) MulDirect(y EthInt) {
	e.Int.Mul(e.Int, y.Int)
}

func (e *EthInt) DivDirect(y EthInt) {
	e.Int.Div(e.Int, y.Int)
}

// Equivalency
func (e *EthInt) Equals(y EthInt) bool {
	if e.Int.Cmp(y.Int) == 0 {
		return true
	}
	return false
}

func (e *EthInt) EqualsStr(y string) bool {
	z := new(big.Int)
	z.SetString(y, 10)
	if e.Int.Cmp(z) == 0 {
		return true
	}
	return false
}

func (e *EthInt) Gt(y EthInt) bool {
	if e.Int.Cmp(y.Int) == 1 {
		return true
	}
	return false
}

func (e *EthInt) GtStr(y string) bool {
	z := new(big.Int)
	z.SetString(y, 10)
	if e.Int.Cmp(z) == 1 {
		return true
	}
	return false
}

func (e *EthInt) Lt(y EthInt) bool {
	if e.Int.Cmp(y.Int) == -1 {
		return true
	}
	return false
}

func (e *EthInt) LtStr(y string) bool {
	z := new(big.Int)
	z.SetString(y, 10)
	if e.Int.Cmp(z) == -1 {
		return true
	}
	return false
}
