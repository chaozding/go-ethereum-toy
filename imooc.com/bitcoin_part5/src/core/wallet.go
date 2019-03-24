package core

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"log"
)

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
	private, public := newKeyPair() //这样就生成了公钥和私钥对了
	wallet := Wallet{private, public}

	return &wallet
}

//GetAddress returns wallet address
func (w Wallet) GetAddress() []byte {
	pubKeyHash := HashPubKey(w.PublicKey)

	//func append(slice []Type, elems ...Type) []Type
	versionedPayload := append([]byte{version}, pubKeyHash...)
	checksum := checksum(versionedPayload) //这是干嘛的？

	fullPayload := append(versionPayload, checksum...) //为什么把checksum加在后面？
	address := Base58Encode(fullPayload)               //这个函数也要自己写？

	return address
}

func newKeyPair() (ecdsa.PrivateKey, []byte) {
	curve := elliptic.P256()
	private, err := ecdsa.GenerateKey(curve, rand.Reader)
	if err != nil {
		log.Panic(err)
	}
	pubKey := append(private.PublicKey.X.Bytes(), private.PublicKey.Y.Bytes()...)
	return *private, pubKey
}
