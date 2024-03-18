package wallet_manager

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"os"

	"github.com/tyler-smith/go-bip32"
	"github.com/tyler-smith/go-bip39"
)

type CurrencyWallet struct {
}

func CreateCurrencyWallet() *CurrencyWallet {
	privateKey, _ := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	privateKeyBytes, err := x509.MarshalECPrivateKey(privateKey)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	privateKeyPem := &pem.Block{
		Type:  "EC PRIVATE KEY",
		Bytes: privateKeyBytes,
	}
	pemFile, _ := os.Create("privateKey.pem")
	pem.Encode(pemFile, privateKeyPem)
	pemFile.Close()

	publicKey := privateKey.PublicKey
	publicKeyBytes, err := x509.MarshalPKIXPublicKey(&publicKey)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	publicKeyPem := &pem.Block{
		Type:  "PUBLIC KEY",
		Bytes: publicKeyBytes,
	}
	pemFile, _ = os.Create("publicKey.pem")
	pem.Encode(pemFile, publicKeyPem)
	pemFile.Close()

	fmt.Println("密钥对生成成功！")
	return nil
}

func CreateCurrencyWallet2() {
	// Generate a mnemonic for memorization or user-friendly seeds
	entropy, _ := bip39.NewEntropy(128)
	mnemonic, _ := bip39.NewMnemonic(entropy)

	// Generate a Bip32 HD wallet for the mnemonic and a user supplied password
	seed := bip39.NewSeed("scissors innocent bulk present time affair parade fuel thought scrub marriage ice cause laundry stereo spider home dish know repeat flip rude bonus seed", "Secret Passphrase")

	masterKey, _ := bip32.NewMasterKey(seed)
	publicKey := masterKey.PublicKey()

	// Display mnemonic and keys
	fmt.Println("Mnemonic: ", mnemonic)
	fmt.Println("Master private key: ", masterKey)
	fmt.Println("Master public key: ", publicKey)

	//  Mnemonic:  scissors innocent bulk present time affair parade fuel thought scrub marriage ice cause laundry stereo spider home dish know repeat flip rude bonus seed
	//  Master private key:  xprv9s21ZrQH143K3VpQJFVKLeNDV3XjphmPa6UmBHr89ESwZijKoj1dfQy3mNSn4DiRggrgJJJEeE6YjZZzZC7AvKNmztmy9isshgpcy9taDR7
	//  Master public key:  xpub661MyMwAqRbcFytsQH2KhnJx35NEEAVEwKQMygFjhZyvSX4UMGKtDDHXcf5pnaq5YoJJtiYBWBeRfaYsKu9m7Ajgo1D4BzcE3oFQpcTymU8

}
