package tx

import (
	"fmt"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/lmittmann/w3"
	"github.com/lmittmann/w3/module/eth"
)

type rawEnv struct {
	PrivateKey  string
	ArbContract string
}

func executeTransfer(token string, amount *big.Int, to string, rpc string, execute bool) common.Hash {
	// Key Setup
	privateKey, walletAddress := getPk()
	tokenAddress := common.HexToAddress(token) // common.HexToAddress is better than w3
	toAddress := common.HexToAddress(to)

	// w3 client
	client, err := w3.Dial(rpc)
	if err != nil {
		log.Fatal("Failed to connect to network: ", err)
	}
	nonce, gasPrice := getGas(walletAddress, client)
	chainid := getChainId(client)

	value := big.NewInt(0)
	var gasLimit uint64 = 200000
	txInput := erc20Transfer(toAddress, amount)

	tx := types.NewTransaction(nonce, tokenAddress, value, gasLimit, gasPrice, txInput)
	signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainid), privateKey)
	fmt.Printf("Tx Target: %s \n", tx.To())

	var txHash common.Hash
	if execute {
		fmt.Println("Sending Tx...")
		err = client.Call(
			eth.SendTx(signedTx).Returns(&txHash),
		)
		if err != nil {
			fmt.Printf("Error Occured : %s \n", err)
		}
		fmt.Printf("Tx Hash: %s \n", txHash)
	} else {
		fmt.Println("Did Not Send Tx")
	}

	return txHash
}

func erc20Transfer(args ...any) []byte {
	// Tx Setup
	var funcTransfer = w3.MustNewFunc("transfer(address,uint256)", "bool")
	txInput, err := funcTransfer.EncodeArgs(args...)
	if err != nil {
		log.Fatal("Error encoding tx")
	}
	return txInput
}
