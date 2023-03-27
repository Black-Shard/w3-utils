package crypto

import (
	"fmt"

	"github.com/ethereum/go-ethereum/common"

	"github.com/tyler-smith/go-bip32"
	"github.com/tyler-smith/go-bip39"
)

// Haven't completely figured how to derive paths from this one - method may not exist.
func Bip() {
	// Generate a mnemonic for memorization or user-friendly seeds
	//entropy, _ := bip39.NewEntropy(256)
	//mnemonic, _ := bip39.NewMnemonic(entropy)

	mnemonic := "tag volcano eight thank tide danger coast health above argue embrace heavy"

	// Generate a Bip32 HD wallet for the mnemonic and a user supplied password
	seed := bip39.NewSeed(mnemonic, "")
	fmt.Println(common.Bytes2Hex(seed))

	masterKey, _ := bip32.NewMasterKey(seed)
	publicKey := masterKey.PublicKey()

	key1, _ := masterKey.NewChildKey(0)
	fmt.Println("Priv Key 1: ", key1)

	sKey, _ := publicKey.Serialize()
	hKey := common.Bytes2Hex(sKey)
	fmt.Println("Pub Key 1: ", hKey)

	// Display mnemonic and keys
	fmt.Println("Mnemonic: ", mnemonic)
	fmt.Println("Master private key: ", masterKey)
	fmt.Println("Master public key: ", publicKey)

}
