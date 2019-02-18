//package rpc
package main

import (
	"encoding/json"
	"github.com/chaozding/go-ethereum-toy/imooc.com/demochain/core"
	"io"
	"net/http"
)

var blockchain *core.Blockchain //全局变量，这里定义了一个指向结构体的指针
// 这个变量是空的呀

func run() {
	http.HandleFunc("/blockchain/get", blockchainGetHandler)
	http.HandleFunc("/blockchain/write", blockchainWriteHandler)
	//启动对端口的监听并启动服务器
	http.ListenAndServe("localhost:8888", nil) //不需要错误处理函数
}

//访问URL的时候，会调用这个函数
func blockchainGetHandler(w http.ResponseWriter, r *http.Request) {
	bytes, error := json.Marshal(blockchain) //转换为json格式
	if error != nil {
		//错误处理函数
		http.Error(w, error.Error(), http.StatusInternalServerError)
		return
	}
	//并没有用到r
	io.WriteString(w, string(bytes)) //把处理后的内容写到响应里面
}

//Request是发送方的请求对象
//Response是接收方的响应对象
func blockchainWriteHandler(w http.ResponseWriter, r *http.Request) {
	blockData := r.URL.Query().Get("Data") //字段
	blockchain.SendData(blockData)
	blockchainGetHandler(w, r) //把最新的区块链回写给接受者
}

func main() {
	blockchain = core.NewBlockchain()
	run()
	//localhost:8888/blockchain/get
}
