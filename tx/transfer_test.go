package tx

import (
	"fmt"
	"log"
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/lmittmann/w3"
	"github.com/lmittmann/w3/module/eth"
)

// @todo add a switch to send entire balance
func TestTransfer(t *testing.T) {
	amount := big.NewInt(0)
	//amount.SetString("4651362", 10)
	//amount = w3.I("1.2 ether")
	token := "0xe3f5a90f9cb311505cd691a46596599aa1a0ad7d"
	//rpc := "https://polygon-rpc.com/"
	rpc := "https://rpc01.dogechain.dog/"
	to := "0x55fabd00693962d017564876d8c390dcf78f47e0"

	account := "0x1783Ad78f0FEFCfa7e1d964D22a7c6e46c6EB4C7"
	amount = getTokenBalance(common.HexToAddress(token), common.HexToAddress(account), rpc)
	fmt.Println(amount)

	executeTransfer(token, amount, to, rpc, true)
}

func TestReturnTokens(t *testing.T) {
	to := "0x55fabd00693962d017564876d8c390dcf78f47e0"
	token := "0x7b4328c127b85369d9f82ca0503b000d09cf9180"
	rpc := "https://rpc01.dogechain.dog/"

	// w3 client
	client, err := w3.Dial(rpc)
	if err != nil {
		log.Fatal("Failed to connect to network: ", err)
	}

	txInput := returnTokens(common.HexToAddress(token))
	signedTx := createTx(to, token, txInput, client)
	sendTx(signedTx, client)

}

func TestBalance(t *testing.T) {
	token := common.HexToAddress("0x765277eebeca2e31912c9946eae1021199b39c61")
	account := common.HexToAddress("0x1783Ad78f0FEFCfa7e1d964D22a7c6e46c6EB4C7")
	rpc := "https://rpc01.dogechain.dog/"
	bal := getTokenBalance(token, account, rpc)
	fmt.Println(bal)
}

func getTokenBalance(token common.Address, account common.Address, rpc string) *big.Int {
	// w3 client
	client, err := w3.Dial(rpc)
	if err != nil {
		log.Fatal("Failed to connect to network: ", err)
	}

	balanceFunc := w3.MustNewFunc("balanceOf(address)", "uint256")
	var bal *big.Int

	err = client.Call(
		eth.CallFunc(balanceFunc, token, account).Returns(&bal),
	)
	if err != nil {
		fmt.Println(err)
	}

	return bal
}
