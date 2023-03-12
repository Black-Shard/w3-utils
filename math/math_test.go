package math

import (
	"fmt"
	"math/big"
	"testing"
)

func TestGetBigInt(t *testing.T) {
	a := GetBigInt("10000")
	fmt.Println(a)
}

func TestDecimalUp(t *testing.T) {
	// Test 1
	input := big.NewInt(int64(56))
	DecimalUp(input, 1)
	expected := big.NewInt(int64(560))
	if input.Cmp(expected) != 0 {
		t.Fatalf(`Decimal Up Failed`)
	}

	// Test 2
	input2 := big.NewInt(int64(12))
	DecimalUp(input2, 2)
	expected2 := big.NewInt(int64(1200))
	if input2.Cmp(expected2) != 0 {
		t.Fatalf(`Decimal Up Failed`)
	}
}

func TestDecimalDown(t *testing.T) {
	// Test 1
	input := big.NewInt(int64(5600))
	DecimalDown(input, 2)
	expected := big.NewInt(int64(56))
	if input.Cmp(expected) != 0 {
		t.Fatalf(`Decimal Down Failed`)
	}

	// Test 2
	input2 := big.NewInt(int64(1250))
	DecimalDown(input2, 3)
	expected2 := big.NewInt(int64(1))
	if input2.Cmp(expected2) != 0 {
		t.Fatalf(`Decimal Up Failed`)
	}
}
func TestMulWadDirect(t *testing.T) {
	var x, _ = new(big.Int).SetString("2000000000000000000", 10)
	var y, _ = new(big.Int).SetString("2000000000000000000", 10)
	var res, _ = new(big.Int).SetString("4000000000000000000", 10)
	MulWadDirect(x, y)
	if x.Cmp(res) != 0 {
		t.Fatalf(`MulWadDirect Failed`)
	}
	// Test 2
	var x1, _ = new(big.Int).SetString("2000000000000000000", 10)
	var y1, _ = new(big.Int).SetString("500000000000000000", 10)
	var res1, _ = new(big.Int).SetString("1000000000000000000", 10)
	MulWadDirect(x1, y1)
	if x1.Cmp(res1) != 0 {
		t.Fatalf(`MulWadDirect Failed`)
	}
}

func TestMulWad(t *testing.T) {
	var x, _ = new(big.Int).SetString("2000000000000000000", 10)
	var xStart, _ = new(big.Int).SetString("2000000000000000000", 10)
	var y, _ = new(big.Int).SetString("2000000000000000000", 10)
	var yStart, _ = new(big.Int).SetString("2000000000000000000", 10)

	var res, _ = new(big.Int).SetString("4000000000000000000", 10)
	z := MulWad(x, y)
	if z.Cmp(res) != 0 {
		t.Fatalf(`MulWad Failed`)
	}
	if x.Cmp(xStart) != 0 {
		t.Fatalf(`MulWad Failed - Changed Starting Var`)
	}
	if y.Cmp(yStart) != 0 {
		t.Fatalf(`MulWad Failed - Changed Starting Var`)
	}

	// Test 2
	var x1, _ = new(big.Int).SetString("2000000000000000000", 10)
	var y1, _ = new(big.Int).SetString("500000000000000000", 10)
	var res1, _ = new(big.Int).SetString("1000000000000000000", 10)
	z1 := MulWad(x1, y1)
	if z1.Cmp(res1) != 0 {
		t.Fatalf(`MulWad Failed`)
	}
}

func TestDivWadDirect(t *testing.T) {
	// Test 1
	var x, _ = new(big.Int).SetString("2000000000000000000", 10)
	var y, _ = new(big.Int).SetString("2000000000000000000", 10)
	var res, _ = new(big.Int).SetString("1000000000000000000", 10)
	DivWadDirect(x, y)
	if x.Cmp(res) != 0 {
		t.Fatalf(`DivWadDirect Failed`)
	}
	fmt.Println(x)
	fmt.Println(y)
	// Test 2
	var x1, _ = new(big.Int).SetString("10000000000000000000", 10)
	var y1, _ = new(big.Int).SetString("5000000000000000000", 10)
	var res1, _ = new(big.Int).SetString("2000000000000000000", 10)
	DivWadDirect(x1, y1)
	if x1.Cmp(res1) != 0 {
		t.Fatalf(`DivWadDirect Failed`)
	}
	// Test 3
	var x2, _ = new(big.Int).SetString("10000000000000000000", 10)
	var y2, _ = new(big.Int).SetString("50000000000000000000", 10)
	var res2, _ = new(big.Int).SetString("200000000000000000", 10)
	DivWadDirect(x2, y2)
	if x2.Cmp(res2) != 0 {
		t.Fatalf(`DivWadDirect Failed`)
	}

}

func TestDivWad(t *testing.T) {
	// Test 1
	var x, _ = new(big.Int).SetString("2000000000000000000", 10)
	var xStart, _ = new(big.Int).SetString("2000000000000000000", 10)
	var y, _ = new(big.Int).SetString("2000000000000000000", 10)
	var yStart, _ = new(big.Int).SetString("2000000000000000000", 10)

	var res, _ = new(big.Int).SetString("1000000000000000000", 10)
	z := DivWad(x, y)
	if z.Cmp(res) != 0 {
		t.Fatalf(`DivWad Failed`)
	}
	if x.Cmp(xStart) != 0 {
		t.Fatalf(`DivWad Failed - Changed Starting Var`)
	}
	if y.Cmp(yStart) != 0 {
		t.Fatalf(`DivWad Failed - Changed Starting Var`)
	}
	// Test 2
	var x1, _ = new(big.Int).SetString("10000000000000000000", 10)  // 10
	var y1, _ = new(big.Int).SetString("5000000000000000000", 10)   // 5
	var res1, _ = new(big.Int).SetString("2000000000000000000", 10) // 2
	z1 := DivWad(x1, y1)
	if z1.Cmp(res1) != 0 {
		t.Fatalf(`DivWad Failed`)
	}
	// Test 3
	var x2, _ = new(big.Int).SetString("10000000000000000000", 10) // 10
	var y2, _ = new(big.Int).SetString("50000000000000000000", 10) // 50
	var res2, _ = new(big.Int).SetString("200000000000000000", 10) // 0.2
	z2 := DivWad(x2, y2)
	if z2.Cmp(res2) != 0 {
		t.Fatalf(`DivWad Failed`)
	}

}

func TestMulMany(t *testing.T) {
	a := big.NewInt(6)
	b := big.NewInt(3)
	c := big.NewInt(2)

	y := MulMany(a, b)
	z := MulMany(a, b, c)

	if y.Cmp(new(big.Int).SetInt64(18)) != 0 {
		t.Fatalf(`MulMany Failed`)
	}
	if z.Cmp(new(big.Int).SetInt64(36)) != 0 {
		t.Fatalf(`MulMany Failed`)
	}

	// Check start vars are the same
	if a.Cmp(new(big.Int).SetInt64(6)) != 0 {
		t.Fatalf(`MulMany Failed`)
	}
	if b.Cmp(new(big.Int).SetInt64(3)) != 0 {
		t.Fatalf(`MulMany Failed`)
	}
	if c.Cmp(new(big.Int).SetInt64(2)) != 0 {
		t.Fatalf(`MulMany Failed`)
	}
}

func TestAddMany(t *testing.T) {
	a := big.NewInt(6)
	b := big.NewInt(3)
	c := big.NewInt(2)

	y := AddMany(a, b)
	z := AddMany(a, b, c)

	if y.Cmp(new(big.Int).SetInt64(9)) != 0 {
		t.Fatalf(`AddMany Failed`)
	}
	if z.Cmp(new(big.Int).SetInt64(11)) != 0 {
		t.Fatalf(`AddMany Failed`)
	}
	// Check start vars are the same
	if a.Cmp(new(big.Int).SetInt64(6)) != 0 {
		t.Fatalf(`AddMany Failed`)
	}
	if b.Cmp(new(big.Int).SetInt64(3)) != 0 {
		t.Fatalf(`AddMany Failed`)
	}
	if c.Cmp(new(big.Int).SetInt64(2)) != 0 {
		t.Fatalf(`AddMany Failed`)
	}
}

func TestSubMany(t *testing.T) {
	a := big.NewInt(6)
	b := big.NewInt(3)
	c := big.NewInt(2)

	y := SubMany(a, b)
	z := SubMany(a, b, c)

	if y.Cmp(new(big.Int).SetInt64(3)) != 0 {
		t.Fatalf(`SubMany Failed`)
	}
	if z.Cmp(new(big.Int).SetInt64(1)) != 0 {
		t.Fatalf(`SubMany Failed`)
	}

	// Check start vars are the same
	if a.Cmp(new(big.Int).SetInt64(6)) != 0 {
		t.Fatalf(`SubMany Failed`)
	}

	if b.Cmp(new(big.Int).SetInt64(3)) != 0 {
		t.Fatalf(`SubMany Failed`)
	}

	if c.Cmp(new(big.Int).SetInt64(2)) != 0 {
		t.Fatalf(`SubMany Failed`)
	}

}
func TestDivMany(t *testing.T) {
	a := big.NewInt(6)
	b := big.NewInt(3)
	c := big.NewInt(2)

	y := DivMany(a, b)
	z := DivMany(a, b, c)

	if y.Cmp(new(big.Int).SetInt64(2)) != 0 {
		t.Fatalf(`DivMany Failed`)
	}
	if z.Cmp(new(big.Int).SetInt64(1)) != 0 {
		t.Fatalf(`DivMany Failed`)
	}
	// Check start vars are the same
	if a.Cmp(new(big.Int).SetInt64(6)) != 0 {
		t.Fatalf(`DivMany Failed`)
	}
	if b.Cmp(new(big.Int).SetInt64(3)) != 0 {
		t.Fatalf(`DivMany Failed`)
	}
	if c.Cmp(new(big.Int).SetInt64(2)) != 0 {
		t.Fatalf(`DivMany Failed`)
	}
}

func TestDecimalsToReadable(t *testing.T) {
	x := big.NewInt(0)
	x.SetString("6960694753315980", 10) // 1e18

	r := DecimalsToReadable(x, 18)
	fmt.Printf("Decimal Out: %s\n", r)
}
