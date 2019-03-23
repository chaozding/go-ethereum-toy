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
	//Bc *Blockchain
}

func (cli *CLI) printUsage() {
	fmt.Println("Usage:")
	fmt.Println("  createblockchain - address ADDRESS - Create a blockchain and send genesis block")
	fmt.Println("  createwallet - Generates a new key-pair and saves it into the wallet file")
	fmt.Println("  getbalance - address ADDRESS - Get balance of ADDRESS")
	fmt.Println("  listaddress - Lists all addresses from the wallet file")
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
	bc := NewBlockchain(from) //这个from是传入是什么意思？实际上没用到，作用就是获得数据库指针
	defer bc.Db.Close()

	tx := NewUTXOTransaction(from, to, amount, bc) //创建新的交易
	bc.MineBlock([]*Transaction{tx})               //更新到文件数据库里面去，为新交易挖矿
	fmt.Println("Successs!")
}

//打印区块链中的所有区块
func (cli *CLI) printChain() {
	//bci := cli.Bc.Iterator() //区块链的迭代器
	bc := NewBlockchain("")
	defer bc.Db.Close()

	bci := bc.Iterator()

	for { //这是什么循环
		block := bci.Next()

		//fmt.Printf("in for\n") //test
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

//Run parses command line argument and processes commands
func (cli *CLI) Run() {
	cli.validateArgs() //验证参数合法性

	//每个命令都对应一个主方法调用
	createBlockchainCmd := flag.NewFlagSet("createblockchain", flag.ExitOnError)
	getBalanceCmd := flag.NewFlagSet("getBalance", flag.ExitOnError)
	createWalletCmd := flag.NewFlagSet("createwallet", flag.ExitOnError)
	listAddressCmd := flag.NewFlagSet("listaddresses", flag.ExitOnError)
	sendCmd := flag.NewFlagSet("send", flag.ExitOnError) //用转账交易取代添加区块
	printChainCmd := flag.NewFlagSet("printchain", flag.ExitOnError)

	//配置主命令选项参数
	getBalanceAddress := getBalanceCmd.String("address", "", "The address to get balance")
	createBlockchainAddress := createBlockchainCmd.String("address", "", "The address to coinbase")
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
	case "createwallet": //主命令
		err := createWalletCmd.Parse(os.Args[2:]) //把参数给主命令
		if err != nil {
			log.Panic(err)
		}
	case "listaddresses": //主命令
		err := listAddressCmd.Parse(os.Args[2:])
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

	if createWalletCmd.Parsed() { //需要和参数对应
		cli.createWallet() //走到这里创建钱包
	}

	if listAddressCmd.Parsed() { //需要和参数对应
		cli.listAddresses() //罗列地址
	}

	if printChainCmd.Parsed() {
		cli.printChain()
	}

	if sendCmd.Parsed() { //检查参数有效性
		if *sendFrom == "" || *sendTo == "" || *sendAmount <= 0 {
			sendCmd.Usage()
			os.Exit(1)
		}

		cli.send(*sendFrom, *sendTo, *sendAmount)
	}
}
