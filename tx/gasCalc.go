package tx

import (
	"fmt"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/lmittmann/w3"
	"github.com/lmittmann/w3/module/eth"
)

func getGas(walletAddress common.Address, client *w3.Client) (uint64, *big.Int) {
	// Get the gas price
	var gasPrice big.Int
	if err := client.Call(
		eth.GasPrice().Returns(&gasPrice),
	); err != nil {
		log.Printf("Getting Network State Failed: %v\n", err)
	}
	fmt.Printf("Network Gas Price: %v\n", gasPrice)

	gasMultiplier := big.NewInt(0)
	gasMultiplier.SetString("1200000000000000000", 10)

	fmt.Println(gasPrice)
	fmt.Println(gasMultiplier)
	utils.MulWadDirect(&gasPrice, gasMultiplier)
	fmt.Printf("Calculated Gas Price: %v\n", gasPrice)

	var nonce uint64
	client.Call(
		eth.Nonce(walletAddress, nil).Returns(&nonce),
	)
	fmt.Printf("Nonce: %v \n", nonce)

	return nonce, &gasPrice
}

func getChainId(client *w3.Client) *big.Int {
	var chainid uint64
	chainId := big.NewInt(0)
	if err := client.Call(
		eth.ChainID().Returns(&chainid),
	); err != nil {
		log.Printf("Getting Network State Failed: %v\n", err)
	}
	chainId.SetUint64(chainid)
	return chainId
}
