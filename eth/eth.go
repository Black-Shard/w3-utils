package eth

import (
	"context"
	"encoding/hex"
	"fmt"
	"log"
	"math/big"
	"strings"
	"sync"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/params"
)

// has0xPrefix validates hexStr begins with '0x' or '0X'.
func has0xPrefix(hexStr string) bool {
	return len(hexStr) >= 2 && hexStr[0] == '0' && hexStr[1] == 'x'
}

// isHexCharacter returns bool of c being a valid hexadecimal.
func isHexCharacter(c rune) bool {
	return ('0' <= c && c <= '9') || ('a' <= c && c <= 'f') || ('A' <= c && c <= 'F')
}

// isHex validates whether each byte is valid hexadecimal string.
func isHex(str string) bool {
	for _, c := range str {
		if !isHexCharacter(c) {
			return false
		}
	}
	return true
}

// A returns an address from a hexstring or panics if the hexstring does not
// represent a valid checksum encoded address.
//
// Use [common.HexToAddress] to get the address from a hexstring without
// panicking.
func A(hexAddress string) common.Address {
	if !has0xPrefix(hexAddress) {
		panic(fmt.Sprintf("hex address %q must have 0x prefix", hexAddress))
	}
	if !isHex(hexAddress[2:]) {
		panic(fmt.Sprintf("hex address %q must be hex", hexAddress))
	}
	if len(hexAddress) != 42 {
		panic(fmt.Sprintf("hex address %q must have 20 bytes", hexAddress))
	}

	var addr common.Address
	b, _ := hex.DecodeString(hexAddress[2:])
	copy(addr[:], b)

	if hexAddress != addr.Hex() {
		panic(fmt.Sprintf("hex address %q must be checksum encoded", hexAddress))
	}
	return addr
}

func WeiToEther(wei *big.Int) *big.Float {
	f := new(big.Float)
	f.SetPrec(236)
	f.SetMode(big.ToNearestEven)
	fWei := new(big.Float)
	fWei.SetPrec(236)
	fWei.SetMode(big.ToNearestEven)
	return f.Quo(fWei.SetInt(wei), big.NewFloat(params.Ether))
}

func EtherToWei(eth *big.Float) *big.Int {
	truncInt, _ := eth.Int(nil)
	truncInt = new(big.Int).Mul(truncInt, big.NewInt(params.Ether))
	fracStr := strings.Split(fmt.Sprintf("%.18f", eth), ".")[1]
	fracStr += strings.Repeat("0", 18-len(fracStr))
	fracInt, _ := new(big.Int).SetString(fracStr, 10)
	wei := new(big.Int).Add(truncInt, fracInt)
	return wei
}

func WeiToGwei(wei *big.Int) *big.Float {
	f := new(big.Float)
	f.SetPrec(236)
	f.SetMode(big.ToNearestEven)
	fWei := new(big.Float)
	fWei.SetPrec(236)
	fWei.SetMode(big.ToNearestEven)
	return f.Quo(fWei.SetInt(wei), big.NewFloat(1e9))
}

func BigIntToNumb(i *big.Int, decimals int) string {
	return BigFloatToNumb(new(big.Float).SetInt(i), decimals)
}

func BigFloatToNumb(f *big.Float, decimals int) string {
	output := f.Text('f', decimals)
	dotIndex := strings.Index(output, ".")
	if dotIndex == -1 {
		dotIndex = len(output)
	}
	for outputIndex := dotIndex; outputIndex > 3; {
		outputIndex -= 3
		output = output[:outputIndex] + "," + output[outputIndex:]
	}
	return output
}

func PrintBlock(block *types.Block) {
	t := time.Unix(int64(block.Header().Time), 0).UTC()
	fmt.Printf("%d \t %s \t tx=%-4d \t gas=%d\n", block.Header().Number, t, len(block.Transactions()), block.GasUsed())
}

// GetBlocks is a fast block query pipeline. It queries blocks concurrently and pushes it into a channel for processing.
func GetBlocks(blockChan chan<- *types.Block, client *ethclient.Client, startBlock int64, endBlock int64, concurrency int) {
	var blockWorkerWg sync.WaitGroup         // for waiting for all workers to finish
	blockHeightChan := make(chan int64, 100) // channel for workers to know which heights to download

	// Start eth client thread pool
	for w := 1; w <= concurrency; w++ {
		blockWorkerWg.Add(1)

		// Worker gets a block height from blockHeightChan, downloads it, and puts it in the blockChan
		go func() {
			defer blockWorkerWg.Done()
			for blockHeight := range blockHeightChan {
				// fmt.Println(blockHeight)
				block, err := client.BlockByNumber(context.Background(), big.NewInt(blockHeight))
				if err != nil {
					log.Println("Error getting block:", blockHeight, err)
					continue
				}
				blockChan <- block
			}
		}()
	}

	// Push blockheights into channel, for workers to pick up
	for currentBlockNumber := startBlock; currentBlockNumber <= endBlock; currentBlockNumber++ {
		blockHeightChan <- currentBlockNumber
	}

	// Close worker channel and wait for workers to finish
	close(blockHeightChan)
	blockWorkerWg.Wait()
}
