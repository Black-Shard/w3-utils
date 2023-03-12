package tx

import (
	"crypto/ecdsa"
	"log"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/spf13/viper"
)

func getPk() (*ecdsa.PrivateKey, common.Address) {
	pk := readEnv()
	privateKey, _ := crypto.HexToECDSA(pk)
	pubKey := privateKey.Public()
	publicKeyECDSA, ok := pubKey.(*ecdsa.PublicKey)
	if !ok {
		panic("error casting public key to ECDSA")
	}
	walletAddress := crypto.PubkeyToAddress(*publicKeyECDSA)

	return privateKey, walletAddress
}

func readEnv() string {
	viper.SetConfigFile("../.env")
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal("Cannot read dotenv config")
	}
	//env := viper.Get("privateKey").(string)
	var rawEnv rawEnv
	viper.Unmarshal(&rawEnv)
	return rawEnv.PrivateKey
}
