> 时间流式学习记录
> 原始文档及更新内容在GitHub仓库，文档里的40位哈希值需要在GitHub仓库里面才能生成commit链接。
> [区块链学习笔记思维导图](https://www.processon.com/mindmap/5c623ed8e4b08a7683be4178)
> [go-ethereum-toy-day17-工作量证明实现代码笔记](https://github.com/chaozding/go-ethereum-toy/blob/master/docs/go-ethereum-toy-day17-工作量证明实现代码笔记.md)

## 继续《私有区块链，我们一起GO-慕课网》
### 继续 4-2 工作量证明代码实现
把bitcoin_part2复制为bitcoin_part3 4008fd9e13c3b8632a545c14b96d69f0e7964ad7 ，便于迭代代码。
之前4-1是原型代码，代码不实用，main函数太死板了，已经有工作量证明了，应该就可以发币了。
发币是什么意思？
#### 问题一：区块链是在内存上的，现在要搬到本地数据库/文件存储
改进方法：引入特殊的存储机制
报错了，搞不明白，明明一模一样的，错误截图如下：
![image](https://user-images.githubusercontent.com/16435896/53483173-e2361000-3abb-11e9-91f9-7b9091728d20.png)
哈哈，通过在bolt库路径下搜Update找到用例了，发现是语法错了，这里是匿名函数啊！

一个block区块里面可以有多笔交易是什么意思？
合约和区块链什么关系？

完成数据库替换内存数据 f383d000df2c68b875f5cdbc233fb2ee20facd78
但是数据库内容还不知道对不对。
#### 问题二：main函数太不优雅了，不灵活
把交易记录写入区块放到命令行做。
交易记录和区块是什么关系？
难道没次发生交易就要创建区块，那不是太麻烦了？
main函数操作部分交给CLI执行 3521b35776fcad749fa8830494e68b6a50d5be3b
CLI可以理解为一个类，先是初始化/构造函数，然后就是和数据相关的操作了。
定义CLI结构体 0c64883897b70aa42ffc6582e3350a4971ead5b8 
func (cli *CLI) printUsage() 1383575d21e1c8c9dcd57f753877829835e0e155 
验证参数 6611b9309378d50ee40ca0c684fe400b0605a4a7 
添加新区块到区块链 c59a70a888b3d572245668f89beed2672f273d62 
这些方法又都是供Run()选择调用的
添加blockchain的迭代器，并配置Next方法 5c024d134cbbc60c4cc9ed87a9a980734d8dc54d 

测试命令如下，此时的版本是 297618742b808e0b759a3481cf454dbfaa5f1f99 
./coin
./coin addblock -data "Pay 0.313 BTC for a coffee" 区块链被存储到本地数据库了
./coin printchain
效果截图如下：
![image](https://user-images.githubusercontent.com/16435896/53509398-a9fff300-3af6-11e9-8166-f8c572fb7a9e.png)
但是打印后发现只有创世区块，不知道是存储问题还是打印问题？
于是我不断运行 ./coin addblock -data "Pay 0.313 BTC for a coffee" 发现文件大小是增加的，说明添加数据是没问题
就是打印问题。
问题还没解决，明天再说吧。

## 语法
GO语言真的很强大，包管理直接依托github.com，一个包就是一个文件夹，文件夹里面直接就是代码。
bytes.Join()和genesis.Serialize()的区别？
[]byte转换为string：hex.EncodeToString(block.Hash[:]) -> %s
block.Hash -> %x
第一次用:=，之后不需要冒号:了
defer bc.Db.Close()表示函数结束后再次执行这个语句。
成员函数都用小写开头。
成员数据都用大写开头。
strconf.FormatBool(pow.Validate())
for {}
cli := core.CLI{注意结构体初始化是花括号} b2459e54cea7c4e8a6a8f7bc11eb81b65ed50fbb 


## 技巧
复制仓库全路径https://github.com/boltdb/bolt，GoLand会自动询问是否安装。
相当于执行 go get -t github.com/boltdb/bolt，安装的库在当前项目的src/github.com/boltdb/bolt下面，不会被git识别。

## 其他
有难度应该高兴啊。
培训班的老师讲课声音很大但不影响他会代码。
语法不重要，重要的是逻辑。

## 参考资料
[Golang公链开发003-数据持久化存储 - Go语言中文网 - Golang中文社区](https://studygolang.com/articles/13798)
[bolt - GoDoc](https://godoc.org/github.com/boltdb/bolt#Tx)
[gob - The Go Programming Language](https://golang.org/pkg/encoding/gob/)

## 变更
- 190227 init