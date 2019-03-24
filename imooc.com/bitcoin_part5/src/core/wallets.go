package core

import (
	"bytes"
	"crypto/elliptic"
	"encoding/gob"
	"fmt"
	"io/ioutil"
	"log"
)

//Wallets stores a collection of wallets
type Wallets struct {
	Wallets map[string]*Wallets
}

//NewWallets creates Wallets and fills it from a file if it exists
func NewWallets() (*Wallets, error) {
	wallets := Wallets{}
	wallets.Wallets = make(map[string]*Wallets)

	err := wallets.LoadFromFile() //从已有的文件中读取，初始化包含在创建钱包里面

	return &wallets, err //错误返回有什么用
}

func (ws *Wallets) CreateWallet() string {
	wallet := NewWallet() //返回公钥私钥对的
	address := fmt.Sprintf("%s", wallet.GetAddress())

	ws.Wallets[address] = wallet //这个函数是把钱包放入钱包集合里面

	return address //返回地址是为了便于显示
}

//SaveToFile saves wallets to a file
func (ws Wallets) SaveToFile() {
	var content bytes.Buffer

	gob.Register(elliptic.P256())

	encoder := gob.NewEncoder(&content)
	err := encoder.Encode(ws) //把钱包集合编码到content里面
	if err != nil {
		log.Panic(err)
	}

	err = ioutil.WriteFile(walletFile, content.Bytes(), 0644)
	if err != nil {
		log.Panic(err)
	}
}
