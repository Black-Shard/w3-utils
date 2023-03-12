package primatives

import (
	"testing"
)

func TestEthPow(t *testing.T) {
	a := NewEthInt("2")
	a.Pow("3")
	if !a.EqualsStr("8") {
		t.Fatalf(`EthPow Failed`)
	}
}

func TestEthMulChain(t *testing.T) {
	a := NewEthInt("2")
	b := NewEthInt("3")
	c := NewEthInt("4")

	z := a.Mul(b).Mul(c)
	if !z.EqualsStr("24") {
		t.Fatalf(`MulChain Failed`)
	}
}

func TestNewEthWad(t *testing.T) {
	e := NewEthWad("1")
	if !e.EqualsStr("1000000000000000000") {
		t.Fatalf(`NewEthWad Failed`)
	}
}

func TestEthAdd(t *testing.T) {
	x := NewEthInt("5")
	y := NewEthInt("10")

	z := x.Add(y)
	if !z.EqualsStr("15") {
		t.Fatalf(`EthAdd Failed`)
	}

	x.AddDirect(y)
	if !x.EqualsStr("15") {
		t.Fatalf(`EthAdd Failed`)
	}
}

func TestEthSub(t *testing.T) {
	x := NewEthInt("5")
	y := NewEthInt("10")

	z := x.Sub(y)
	if !z.EqualsStr("-5") {
		t.Fatalf(`EthSub Failed`)
	}

	x.SubDirect(y)
	if !x.EqualsStr("-5") {
		t.Fatalf(`EthSub Failed`)
	}
}

func TestEthMul(t *testing.T) {
	x := NewEthInt("5")
	y := NewEthInt("10")

	z := x.Mul(y)
	if !z.EqualsStr("50") {
		t.Fatalf(`EthMul Failed`)
	}

	x.MulDirect(y)
	if !x.EqualsStr("50") {
		t.Fatalf(`EthMul Failed`)
	}
}

func TestEthDiv(t *testing.T) {
	x := NewEthInt("10")
	y := NewEthInt("5")

	z := x.Div(y)
	if !z.EqualsStr("2") {
		t.Fatalf(`EthDiv Failed`)
	}

	x.DivDirect(y)
	if !x.EqualsStr("2") {
		t.Fatalf(`EthMul Failed`)
	}
}

func TestEthGtLt(t *testing.T) {
	x := NewEthInt("10")
	y := NewEthInt("11")

	if !y.Gt(x) {
		t.Fatalf(`EthGt Failed`)
	}
	if !x.Lt(y) {
		t.Fatalf(`EthLt Failed`)
	}
}
func TestEthGtLtStr(t *testing.T) {
	x := NewEthInt("1")
	y := NewEthInt("-1")

	if !x.GtStr("0") {
		t.Fatalf(`EthGt Failed`)
	}
	if !y.LtStr("0") {
		t.Fatalf(`EthLt Failed`)
	}
}
