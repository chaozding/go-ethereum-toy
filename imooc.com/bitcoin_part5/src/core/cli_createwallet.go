package core

import "fmt"

func (cli *CLI) createWallet() {
	wallets, _ := NewWallets()        //创建了一个钱包集合对象， 如果钱包集合已经存在则不需要再重复创建了
	address := wallets.CreateWallet() //在钱包集合里面创建一个钱包对象
	wallets.SaveToFile()              //把钱包对象存储到钱包集合文件

	fmt.Printf("Your new address: %s\n", address)
}
