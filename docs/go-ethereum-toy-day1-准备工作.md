## 缘起
前段时间看了《区块链之新》，解答了很多疑惑，所以对区块链有了新的看法，看过纪录片之后，我对区块链的理解是去中间人，但是后来又看了各种新闻，感觉所谓的去中间人好像也不是那么回事。还有一点看好的就是用于数字ID，不过这东西为什么一定要等区块链出来了才有呢。
至于各种币，我觉得更多的是金钱游戏，反正钱是不可能凭空生出来。
当然，这些认识是很肤浅的，看过纪录片谁都会又一堆看似合理的想法和冲动。
想起之前听过的茶电台[How to build blockchain from scratch? - Teahour.fm](http://teahour.fm/2018/07/08/how-to-build-blockchain-from-scratch.html)，一直没时间搞，主要是觉得难吧，现在区块链应用又多起来了，所以想把这个再重新捡起来，好像在投机。 

## 方案
以[太坊客户端](https://github.com/ethereum/go-ethereum)的工程已经很大了，感觉根本无从下手。
所以，参考[cryptape/ruby-ethereum: A Ruby implementation of Ethereum.](https://github.com/cryptape/ruby-ethereum)的commit历史来做，使用go语言。
主要的障碍是：
1. 对区块链和以太坊具体实现原理并不了解->参考论坛贡献的[技术文档](http://8btc.com/doc-doc.html)以及官方的文档。
2. ruby和go语言没学过，把语言学好了再来写软件显然是不现实的->ruby只要求能看懂代码是做什么的即可，语法细节不管，go语言参考官方文档的入门教程，实现功能即可。
3. 工作量大->实现ruby-ethereum的功能即可，如果超过45分钟没有解决，考虑需要更换方案以及到论坛寻找帮助。
4. 放弃半途而废->详细记录驱动开发，帖子/记录之后会放在代码的[docs](https://github.com/chaozding/go-ethereum-toy/docs)文件夹下。

## 步骤
### 编译运行go-ethereum
[以太坊是什么？](https://ethereum.github.io/go-ethereum/)
以太坊是一个去中心化的平台，运行智能合约，智能合约其实是一种应用，以程序的方式运行，不停机、无审查（没懂，那想审查不还得审查）、诈骗、第三方接口。
Go Ethereum可以作为独立的客户端或者嵌入库。

参考[#1 明说(01)：从0开始搭建区块链开发环境 - EthCast](http://ethcast.com/v1)
以太坊客户端
编程语言Solidity
开发框架Truffle是什么意思？
geth
账户是什么？
账户列表
账户余额
当前区块号是什么？
启动挖矿
充值到第一个账户？
转账
交易确认需要挖矿写入区块链？
编译参考GitHub的[README.md](https://github.com/ethereum/go-ethereum)
Mac参考[Installation Instructions for Mac · ethereum/go-ethereum Wiki](https://github.com/ethereum/go-ethereum/wiki/Installation-Instructions-for-Mac)
git clone https://github.com/ethereum/go-ethereum
brew install go
cd go-ethereum
make geth
运行  build/bin/geth 启动节点
vim .zshrc
export PATH=$HOME/bin:$HOME/git/go-ethereum/build/bin:/usr/local/bin:$PATH
然后就可以运行geth就有效果了
可执行命令
没看懂
geth实例
以太坊网络
用户和以太坊网络的交互：
1. 创建账户
2. 交易
3. 部署合约，和合约交互，智能合约是什么
geth console
exit
geth --testnet console
可编程的Geth节点接口
私人矿机
读我大致看了下，没看懂。
[Developers' Guide · ethereum/go-ethereum Wiki](https://github.com/ethereum/go-ethereum/wiki/Developers'-Guide)这个是如何贡献代码的教程。
#### 配置Go环境
vi .zshrc
export GOPATH=$HOME/go
suorce .zshrc
echo $GOPATH
mkdir -p $GOPATH/src/github.com/chaozding
rm -rf go-ethereum
git clone git@github.com:chaozding/go-ethereum.git $GOPATH/src/github.com/chaozding/go-ethereum
#### 构建可执行程序
cd $GOPATH/src/github.com/chaozding/go-ethereum
go install -v ./cmd/geth
奇怪只能在$GOPATH/src/github.com/ethereum/go-ethereum目录下运行，晕
geth会被安装到 $GOPATH/bin/geth
#### Git flow
#### 测试
Command + Shift + . Mac显示隐藏文件
go test -v -cpu 4 ./eth -run TestMethod 这是做什么的？
#### [Metrics and Monitoring · ethereum/go-ethereum Wiki](https://github.com/ethereum/go-ethereum/wiki/Metrics-and-Monitoring)
这是什么？
#### 获取栈跟踪
退出 geth 后台驻留程序
参考[go ethereum - How do you stop a running geth node? - Ethereum Stack Exchange](https://ethereum.stackexchange.com/questions/709/how-do-you-stop-a-running-geth-node)
killall -HUP geth
或者
ps ax | grep geth
kill -HUP <pid>
好像没效果，杀不死进程，不知道为什么？
geth --pprof console
然后可以访问 http://localhost:6060/debug/pprof 查看堆
访问 http://localhost:6060/debug/pprof/goroutine?debug=2 产生跟踪，用于调试 
geth -port=30300 -verbosity 5 --pprof --pprofport 6060 2>> /tmp/00.glog
killall -QUIT geth
#### 贡献
鼓励'in progress'PRs
### 编译运行ruby-ethereum
一下午都没搞这个，看电影看视频，然后就发散出去了。
git clone https://github.com/chaozding/ruby-ethereum
->《用Go语言构建自己的区块链教程》