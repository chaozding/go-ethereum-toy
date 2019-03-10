package main

import (
	"core"
)

func main() {
	cli := core.CLI{} //把要操作的对象给命令行
	cli.Run()         //输入数据，运行
}
