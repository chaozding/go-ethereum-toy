package core

import "fmt"

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
	wallet := NewWallet()
	address := fmt.Sprintf("%s", wallet.GetAddress())

	ws.Wallets[address] = wallet

	return address
}
