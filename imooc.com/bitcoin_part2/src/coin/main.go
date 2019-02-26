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
		//fmt.Printf("Hash: %s\n", block.Hash) //乱码
		//fmt.Printf("Hash: %s\n", hex.EncodeToString(block.Hash[:]))

		pow := core.NewProofOfWork(block)                           //直接把block输入证明函数，pow是什么类型
		fmt.Printf("Pow: %s\n", strconv.FormatBool(pow.Validate())) //为什么不直接用pow判断有效性，还要再调用Validate()？
		fmt.Println()
	}
}
