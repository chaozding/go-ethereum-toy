> 时间流式学习记录
> 原始文档及更新内容在GitHub仓库 [go-ethereum-toy-day18-0-printchain问题解决](https://github.com/chaozding/go-ethereum-toy/blob/master/docs/go-ethereum-toy-day18-0-printchain问题解决.md)
> https://github.com/chaozding/go-ethereum-toy/commit/
> [区块链学习笔记思维导图](https://www.processon.com/mindmap/5c623ed8e4b08a7683be4178)

## 继续《私有区块链，我们一起GO-慕课网》
### 继续 4-2 工作量证明代码实现
#### 问题：printchain只能打印创世区块的问题
怎么做？
go install coin 在GoLand的Terminal中执行
把blockchain.db删除了，重新运行 ./coin addblock -data "Pay 0.313 BTC for a coffee" 2次，含创世区块，就应该有3个区块在桶bucket里面了。

这个c.Cursor没搞明白，我觉得还是直接把首尾的哈希值存储起来更靠谱。
所以加了一个成员数据tail []byte //末尾区块哈希值

YES!这次对了，用更多的数据换取操作的可理解性，还是值得的。
https://github.com/chaozding/go-ethereum-toy/commit/37d1def6806dcffe1a65943a69c06083e61c52f9

## Go语法
var tail []byte
:= 定义并初始化

## ChangeLog
- 190228 init