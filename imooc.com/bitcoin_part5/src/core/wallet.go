package core

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/sha256"
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

	return &wallet //这个存储地址有什么用啊，到时候保存的是地址值?
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

//HashPubKey hashes public key
func HashPubKey(pubKey []byte) []byte {
	publicSHA256 := sha256.Sum256(pubKey)

	RIPEMD160Hasher := ripemd160.New() //这是什么？
	_, err := RIPEMD160Hasher.Write(publicSHA256[:])
	if err != nil {
		log.Panic(err)
	}
	publicRIPEMD160 := RIPEMD160Hasher.Sum(nil)

	return publicRIPEMD160
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
