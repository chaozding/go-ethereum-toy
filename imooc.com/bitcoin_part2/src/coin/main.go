package main

import (
	"core"
	"crypto/sha256"
	"fmt"
	"math/big"
	"strconv"
)

const targetBits = 24 //

func proof() {
	data1 := []byte("I like dounts")
	data2 := []byte("I like dountsca07ca")   //这个字符串非常难凑的，反正这次没凑对
	target := big.NewInt(1)                  //初始化为1
	target.Lsh(target, uint(256-targetBits)) //为什么是256 - targetBits，这个为什么对应5个0呢
	fmt.Printf("%x\n", sha256.Sum256(data1))
	fmt.Printf("%064x\n", target)            //某种特殊的目标格式
	fmt.Printf("%x\n", sha256.Sum256(data2)) //这个怎么和视频里面不一样啊？

}

func main() {
	proof()
	//test
	/*
		const targetBits = 24
		target := big.NewInt(1)
		target.Lsh(target, uint(1))
		fmt.Printf("%064x\n", target) //转换为16进制显示
		target = big.NewInt(1)
		target.Lsh(target, uint(7))
		fmt.Printf("%064x\n", target) //转换为16进制显示
		target.Lsh(target, uint(255))
		fmt.Printf("%064x\n", target) //转换为16进制显示
	*/

	//创建一个新的区块链bc
	bc := core.NewBlockchain()

	bc.AddBlock("Send 1 BTC to Ivan")      //加入一个新区块，发送一个比特币给伊文
	bc.AddBlock("Send 2 more BTC to Ivan") //加入一个新区块，发送更多比特币给伊文

	//校验区块链上的区块
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
