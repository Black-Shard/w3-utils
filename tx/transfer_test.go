package tx

import (
	"fmt"
	"math/big"
	"testing"

	"github.com/lmittmann/w3"
)

func TestTransfer(t *testing.T) {
	amount := big.NewInt(0)
	//amount.SetString("10000000", 10)
	amount = w3.I("1.2 ether")
	fmt.Println(amount)

	token := "0xc03345448969Dd8C00e9E4A85d2d9722d093aF8E"
	//rpc := "https://polygon-rpc.com/"
	rpc := "https://eth.bd.evmos.org:8545"
	to := "0xd0b3e774b1f2c3cafda3f018ee9cd6f08c91c471"

	executeTransfer(token, amount, to, rpc, false)
}
