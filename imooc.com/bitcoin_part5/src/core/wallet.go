package core

import "crypto/ecdsa"

const versiion = byte(0x00)
const walletFile = "wallet.dat"
const addressChecksumLen = 4

//Wallet stores private and public keys
type Wallet struct {
	PrivateKey ecdsa.PrivateKey
	PublicKey  []byte
}

//NewWallet creates and returns a Wallet
func NewWallet() *Wallet {
	private, public := newKeyPair()
	wallet := Wallet{private, public}

	return &wallet
}
