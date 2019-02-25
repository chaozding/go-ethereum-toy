package main

import (
	"core"
	"fmt"
	"strconv"
)

func main() {
	//创建一个新的区块链bc
	bc := core.NewBlockchain()

	bc.AddBlock("Send 1 BTC to Ivan")      //加入一个新区块，发送一个比特币给伊文
	bc.AddBlock("Send 2 more BTC to Ivan") //加入一个新区块，发送更多比特币给伊文

	for _, block := range bc.Blocks {
		fmt.Printf("Prev.hash: %x\n", block.PreBlockHash) //打印前一个区块的哈希值
		fmt.Printf("Data: %s\n", block.Data)              //以字符串格式显示数据
		fmt.Printf("Hash: %x\n", block.Hash)              //%x适合[]byte

		pow := core.NewProofofWork(block)
		fmt.Printf("Pow: %s\n", strconv.FormatBool(pow.Validate()))
		fmt.Println()
	}
}
