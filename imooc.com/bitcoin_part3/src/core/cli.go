package core

import (
	"fmt"
	"os"
)

//CLI reponsible for processing command line arguments
type CLI struct {
	Bc *Blockchain
}

func (cli *CLI) printUsage() {
	fmt.Println("Usage:")
	fmt.Println("  addblock -data BLOCK_DATA - add a block to the blockchain")
	fmt.Println("  printchain - print all the blocks of the blockchain ")
}

func (cli *CLI) validateArgs() { //验证参数
	if len(os.Args) < 2 {
		cli.printUsage()
		os.Exit(1)
	}
}
