package tx

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/lmittmann/w3"
	"github.com/lmittmann/w3/module/eth"
)

type Tokens struct {
	Tokens []Token `json:"tokens"`
}
type Token struct {
	ChainId int    `json:"chainID"`
	Address string `json:"address"`
	Name    string `json:"name"`
	Symbol  string `json:"symbol"`
}

func TestGetPoolAddressQuick(t *testing.T) {

	rpc := "https://rpc-us.dogechain.dog/"
	//pk, _ := getPk()
	client, _ := w3.Dial(rpc)
	factory := common.HexToAddress("0xaaa04462e35f3e40d798331657ca015169e005d7") //yodeswap factory
	//read tokens json file
	content, err := ioutil.ReadFile("C:/Users/tim_i/freya/config/dogechainTokens.json")
	if err != nil {
		log.Fatal(err)
		return
	}
	var tokens Tokens
	err = json.Unmarshal(content, &tokens)
	if err != nil {
		log.Fatal("Error during Unmarshal(): ", err)
	}
	for i := 0; i < len(tokens.Tokens); i++ {
		for j := 0; j < len(tokens.Tokens); j++ {
			token0 := common.HexToAddress(tokens.Tokens[i].Address)
			token1 := common.HexToAddress(tokens.Tokens[j].Address)
			var pairAddr common.Address

			var getPair = w3.MustNewFunc("getPair(address,address)", "address")
			client.Call(
				eth.CallFunc(getPair, factory, token0, token1).Returns(&pairAddr),
			)
			if pairAddr.String() == "0x0000000000000000000000000000000000000000" {
				continue
			}
			/*fmt.Printf("The Pair Address is: %v\n", pairAddr)
			fmt.Printf("Pair: %v and %v\n", tokens.Tokens[i].Symbol, tokens.Tokens[j].Symbol)
			fmt.Printf("Token0: %v\n", token0)
			fmt.Printf("Token1: %v\n", token1)
			*/
			var (
				getToken0 = w3.MustNewFunc("token0()", "address")
				getToken1 = w3.MustNewFunc("token1()", "address")
			)
			client.Call(
				eth.CallFunc(getToken0, pairAddr).Returns(&token0),
				eth.CallFunc(getToken1, pairAddr).Returns(&token1),
			)

			fmt.Printf("{\n")
			fmt.Printf("\"address\":\"%v\",\n", pairAddr)
			if token0.String() == tokens.Tokens[i].Address {
				fmt.Printf("\"token0\":\"%v\",\n", tokens.Tokens[i].Symbol)
				fmt.Printf("\"token1\":\"%v\",\n", tokens.Tokens[j].Symbol)
				fmt.Printf("\"name\":\"%v-%v-YODE\",\n", tokens.Tokens[i].Symbol, tokens.Tokens[j].Symbol)
			} else {
				fmt.Printf("\"token0\":\"%v\",\n", tokens.Tokens[i].Symbol)
				fmt.Printf("\"token1\":\"%v\",\n", tokens.Tokens[j].Symbol)
				fmt.Printf("\"name\":\"%v-%v-YODE\",\n", tokens.Tokens[i].Symbol, tokens.Tokens[j].Symbol)
				/*	fmt.Printf("address of token 0\":\"%v\"\n", tokens.Tokens[i].Address)
					fmt.Printf("address of tokens.Tokens[i] 0\":\"%v\"\n", token0.String())
					fmt.Printf("Branch 2")
					fmt.Printf("token0\":\"%v\"\n", tokens.Tokens[j].Symbol)
					fmt.Printf("token1\":\"%v\"\n", tokens.Tokens[i].Symbol)
					fmt.Printf("name\":\"%v-%v-YODE\"\n", tokens.Tokens[j].Symbol, tokens.Tokens[i].Symbol)*/
			}
			fmt.Printf("\"dex\":\"Yodeswap\",\n")
			fmt.Printf("\"type\":\"UniV2\"\n")
			fmt.Printf("},\n")
		}
	}
}
