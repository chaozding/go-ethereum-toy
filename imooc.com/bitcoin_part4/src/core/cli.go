package core

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
)

//CLI reponsible for processing command line arguments
type CLI struct {
	Bc *Blockchain
}

func (cli *CLI) printUsage() {
	fmt.Println("Usage:")
	fmt.Println("  getbalance -address ADDRESS - Get balance of ADDRESS")
	//fmt.Println("  addblock -data BLOCK_DATA - add a block to the blockchain")
	fmt.Println("  createblockchain -address ADDRESS - Create a blockchain and send genesis block")
	fmt.Println("  printchain - print all the blocks of the blockchain ")
	fmt.Println("  send -from FROM -to TO -amount AMOUNT - Send AMOUNT of coins from FROM address to TO address")
}

func (cli *CLI) validateArgs() { //验证参数
	if len(os.Args) < 2 {
		cli.printUsage()
		os.Exit(1)
	}
}

func (cli *CLI) createBlockchain(address string) {
	bc := CreateBlockchain(address)
	bc.Db.Close()
	fmt.Println("Done!")
}

func (cli *CLI) getBalance(address string) {
	bc := NewBlockchain(address) //取得区块链
	defer bc.Db.Close()

	balance := 0
	UTXOs := bc.FindUTXO(address) //这个什么意思？

	for _, out := range UTXOs {
		balance += out.Value //把没花出去的钱都累加起来
	}

	fmt.Printf("Balance of '%s': %d\n", address, balance)
}

//Send操作为什么需要创建新的区块链呢？
func (cli *CLI) send(from, to string, amount int) {
	bc := NewBlockchain(from) //这个from是传入是什么意思？
	defer bc.Db.Close()

	tx := NewUTXOTransaction(from, to, amount, bc)
	bc.MineBlock([]*Transaction{tx}) //更新到文件数据库里面去
	fmt.Println("Successs!")
}

//打印区块链中的所有区块
func (cli *CLI) printChain() {
	bci := cli.Bc.Iterator() //区块链的迭代器

	for { //这是什么循环
		block := bci.Next()

		fmt.Printf("in for\n") //test
		fmt.Printf("Pref. hash: %x\n", block.PreBlockHash)
		//fmt.Printf("Data: %s\n", block.Data)
		fmt.Printf("Hash: %x\n", block.Hash)
		pow := NewProofOfWork(block)
		fmt.Printf("Pow: %s\n", strconv.FormatBool(pow.Validate()))
		fmt.Println()

		if len(block.PreBlockHash) == 0 {
			break
		}
	}
}

func (cli *CLI) Run() {
	cli.validateArgs() //验证参数合法性

	//每个命令都对应一个方法调用
	getBalanceCmd := flag.NewFlagSet("getBalance", flag.ExitOnError)
	createBlockchainCmd := flag.NewFlagSet("createblockchain", flag.ExitOnError)
	//addBlockCmd := flag.NewFlagSet("addblock", flag.ExitOnError)
	sendCmd := flag.NewFlagSet("send", flag.ExitOnError) //用转账交易取代添加区块
	printChainCmd := flag.NewFlagSet("printchain", flag.ExitOnError)

	//配置主命令选项参数
	getBalanceAddress := getBalanceCmd.String("address", "", "The address to get balance")
	createBlockchainAddress := createBlockchainCmd.String("address", "", "The address to coinbase")
	//addBlockData := addBlockCmd.String("data", "", "block data")
	sendFrom := sendCmd.String("from", "", "Source wallet address")
	sendTo := sendCmd.String("to", "", "Destination wallet address")
	sendAmount := sendCmd.Int("amount", 0, "Amount to send")

	//提取
	switch os.Args[1] {
	case "getBalance":
		err := getBalanceCmd.Parse(os.Args[2:]) //索引从0开始的，索引0的位置是./coin
		if err != nil {
			log.Panic(err)
		}
	case "createBlockchain":
		err := createBlockchainCmd.Parse(os.Args[2:])
		if err != nil {
			log.Panic(err)
		}
	case "send": //创建完了区块链就开始记录交易啦，交易是由地址发起的
		err := sendCmd.Parse(os.Args[2:])
		if err != nil {
			log.Panic(err)
		}
	case "printchain":
		err := printChainCmd.Parse(os.Args[2:])
		if err != nil {
			log.Panic(err)
		}
	default:
		cli.printUsage()
		os.Exit(1)
	}

	if getBalanceCmd.Parsed() {
		//检查输入参数
		if *getBalanceAddress == "" { //检查输入参数是否有效
			getBalanceCmd.Usage()
			os.Exit(1)
		}
		cli.getBalance(*getBalanceAddress) //输入参数其实就是选项啦
	}

	if createBlockchainCmd.Parsed() {
		if *createBlockchainAddress == "" {
			createBlockchainCmd.Usage()
			os.Exit(1)
		}
		cli.createBlockchain(*createBlockchainAddress)
	}

	if printChainCmd.Parsed() {
		cli.printChain()
	}
}
