package tx

import (
	"fmt"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/lmittmann/w3"
	"github.com/lmittmann/w3/module/eth"
)

func TestGetPoolAddress(t *testing.T) {
	rpc := "https://rpc-us.dogechain.dog/"
	//pk, _ := getPk()
	client, _ := w3.Dial(rpc)
	factory := common.HexToAddress("0xF4bc79D32A7dEfd87c8A9C100FD83206bbF19Af5") // Kibbleswap factory

	token0 := common.HexToAddress("0xB3306f03595490e5cC3a1b1704a5a158D3436ffC")
	token1 := common.HexToAddress("0xB7ddC6414bf4F5515b52D8BdD69973Ae205ff101")
	var pairAddr common.Address

	var getPair = w3.MustNewFunc("getPair(address,address)", "address")
	client.Call(
		eth.CallFunc(getPair, factory, token0, token1).Returns(&pairAddr),
	)
	fmt.Printf("The Pair Address is: %v\n", pairAddr)

	var (
		getToken0 = w3.MustNewFunc("token0()", "address")
		getToken1 = w3.MustNewFunc("token1()", "address")
	)

	client.Call(
		eth.CallFunc(getToken0, pairAddr).Returns(&token0),
		eth.CallFunc(getToken1, pairAddr).Returns(&token1),
	)

	fmt.Printf("Token0: %v\n", token0)
	fmt.Printf("Token1: %v\n", token1)

}
