package main //这里为什么是main而不是cmd呢？

import "github.com/chaozding/go-ethereum-toy/imooc.com/demochain/core"

func main() {
	bc := core.NewBlockchain()
	bc.SendData("Send 1 BTC to Jacky")
	bc.SendData("Send 1 EOC to Jack")
	bc.Print()
}
