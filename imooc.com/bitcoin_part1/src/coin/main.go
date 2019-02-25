package main

import (
	"fmt"
	"github.com/chaozding/go-ethereum-toy/imooc.com/bitcoin_part1/src/core"
)

func main() {
	//创建一个新的区块链bc
	bc := core.NewBlockchain()

	bc.AddBlcok("Send 1 BTC to Ivan")      //加入一个新区块，发送一个比特币给伊文
	bc.AddBlock("Send 2 more BTC to Ivan") //加入一个新区块，发送更多比特币给伊文

	for _, block := range bc.Blocks {
		fmt.Print("Prev.hash: %x\n", block.PreBlockHash) //打印前一个区块的哈希值
		fmt.Print("Data: %s\n", block.Data)              //以字符串格式显示数据
		fmt.Print("Hash: %x\n", block.Hash)
		fmt.Println()
	}
}
