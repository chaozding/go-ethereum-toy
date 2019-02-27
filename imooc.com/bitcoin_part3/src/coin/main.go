package main

import (
	"core"
)

func main() {
	//创建一个新的区块链bc
	bc := core.NewBlockchain() //这里的bc结构不同于以往的了
	defer bc.Db.Close()        //main函数结束后才能关闭

	//bc.AddBlock("Send 1 BTC to Ivan")      //加入一个新区块，发送一个比特币给伊文
	//bc.AddBlock("Send 2 more BTC to Ivan") //加入一个新区块，发送更多比特币给伊文

	cli := core.CLI(bc) //把要操作的对象给命令行
	cli.Run()           //输入数据，运行

	////校验区块链上的区块
	//for _, block := range bc.Blocks {
	//	fmt.Printf("Prev.hash: %x\n", block.PreBlockHash) //打印前一个区块的哈希值
	//	fmt.Printf("Data: %s\n", block.Data)              //以字符串格式显示数据
	//	fmt.Printf("Hash: %x\n", block.Hash)              //%x适合[]byte
	//	//fmt.Printf("Hash: %s\n", block.Hash) //乱码
	//	//fmt.Printf("Hash: %s\n", hex.EncodeToString(block.Hash[:]))
	//
	//	pow := core.NewProofOfWork(block)                           //直接把block输入证明函数，pow是什么类型
	//	fmt.Printf("Pow: %s\n", strconv.FormatBool(pow.Validate())) //为什么不直接用pow判断有效性，还要再调用Validate()？
	//	fmt.Println()
	//}
	//bc.DeleteBlockchain()
}
