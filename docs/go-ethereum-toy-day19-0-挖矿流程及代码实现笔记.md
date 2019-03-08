> 时间流式学习记录
> 原始记录及更新内容在GitHub仓库 [go-ethereum-toy-day19-0-挖矿流程及代码实现笔记](https://github.com/chaozding/go-ethereum-toy/blob/master/docs/go-ethereum-toy-day19-0-挖矿流程及代码实现笔记.md)
> [区块链学习笔记思维导图](https://www.processon.com/mindmap/5c623ed8e4b08a7683be4178)
> [区块链应用笔记思维导图](https://www.processon.com/view/link/5c7a0699e4b0628e820ab8a2)
> [区块链技术相关书签分享](https://www.gettoby.com/p/swgw7g8sayvw)

[](https://github.com/chaozding/go-ethereum-toy/commit/)

## 继续学习《私有区块链，我们一起GO-慕课网》
### [5-3 挖矿流程及代码实现](https://www.imooc.com/video/17560)
**交易记账示意图**
![image](https://user-images.githubusercontent.com/16435896/53558549-2b9d6280-3b83-11e9-865a-bdfebd846d07.png)
来源：[5-1 交易记账原理及过程](https://www.imooc.com/video/17558)
transaction.go 文件引入了一个新的结构体 Transaction，一个交易包含很多输入输出。
发币交易[coinbase](https://www.coinbase.com/)：用于奖励挖矿成功的交易
交易记录对象的属性：
交易记录
交易输入
交易输出
把比特币从一个地址发转到另一个地址（转账交易/普通交易）是什么意思？
1. 存在区块链
2. 查询是否有足够的余额
输出
系统找零
比特币要么你不花，要么全花掉，然后系统找零。
**交易记录**这张图是理解的关键
![image](https://user-images.githubusercontent.com/16435896/53680406-1dcc1680-3d16-11e9-83db-e134a7832719.png)
来源：[5-1 交易记账原理及过程](https://www.imooc.com/video/17558)
没花掉是什么意思？
首先到区块链中遍历所有的区块，把余额/没有花出去的Output找出来，然后把余额累加出来，先看看够不够资格转账，如果够资格，则按照Output一笔笔往外转就可以了，如果需要找零则在新的交易里面做一个Output来找零
数字地址（必须唯一）：叫做比特币的**钱包**。
为什么比特币的数量是有限？一直挖下去不是无限吗？
为什么交易的输入和输出不可能相同?
初始化 bitcoin_part4 [839d69e](https://github.com/chaozding/go-ethereum-toy/commit/839d69e673f59ea5b445f205ef1c7ed665569740)
添加命令行主命令 [d40bf3e](https://github.com/chaozding/go-ethereum-toy/commit/d40bf3ef4388a27fb1bedf4a79505cad686f74a6)
添加命令行主命令的选项参数 [b4caf5](https://github.com/chaozding/go-ethereum-toy/commit/b4caf577b6ddd274ddf33490e2033839796fd408)
解析主命令并获得全部选项参数到命令对象中去 [54637fd](https://github.com/chaozding/go-ethereum-toy/commit/54637fd05753fe8698ae8cac0ebee05c3b8ececf)
根据解析结果调用函数,调用了查询余额的函数,调用了创建区块链的函数(挖矿) 
添加CLI对象的成员函数 createBlockchain 和 getBalance ,成员函数再调用核心库函数 [744f01a](https://github.com/chaozding/go-ethereum-toy/commit/744f01a53e13b46fd9f9ad635f65e70f46df7c19)
创建发币交易函数 NewCoinbaseTX(to, data string) *Transaction [6c7bb05](https://github.com/chaozding/go-ethereum-toy/commit/6c7bb05ce1fea0101268679ac8ae3dfa1cb4fc65)
交易对象的成员函数设置交易序号 (tx *Transaction) SetID() [4a5bab6](https://github.com/chaozding/go-ethereum-toy/commit/4a5bab6058cd045733d5843d8b3ab6008fcf9898)
交易序号使用哈希值可以验证输入输出有没有被改变过
完成命令行函数 (cli *CLI) printUsage() [b58dcbb](https://github.com/chaozding/go-ethereum-toy/commit/b58dcbbbae6d37e71a320be7621b5e5a826dfe52)
NewUTXOTransaction(from, to string, amount int, bc *Blockchain) *Transaction [3f8be71](https://github.com/chaozding/go-ethereum-toy/commit/3f8be71757532d592573386e182bc104e37754ab)

## Go语法
log.Panic("...") 
结构体编码为字节数组的方法
```go
var encoded bytes.Buffer //字符数组缓冲器对象,encoded.Bytes()
enc : gob.NewEncoder(&encoded) //传递进指针
err := gob.Encode(tx)
if err != nil {
    log.Panic(err)
}
hash = sha256.Sum256(encoded.Bytes())
tx.ID = hash[:]
```
结构体的初始化 []TXInput{}
const subsidy = 10
## 算法

## 疑问
- 挖矿成功奖励的10个比特币来自哪里？
- 挖矿成功是什么意思？
- 区块的Data部分存放的是真实的Transaction结构体？
  - 现在使用真实的交易记录结构体取代描述占位符
- 每次进行交易就要创建区块？
- 创建区块链和发币是什么关系？
- 每个区块和钱包地址是什么关系?
- 交易不是应该记录在每个区块上,还是说区块链存在在每个区块上?
- 怎么知道有交易的?有交易了就开始挖矿?原来的区块都不能在有任何改动了?
- 奖励交易是最简单的交易,只要挖矿就会存储这么一个交易?
- 多个人同时挖矿不是资源浪费吗?
- NewBlockchain 和 CreateBlockchain 有什么区别呢, 注释都是一样的?
  - CreateBlockchain 这个要先于 NewBlockchain 存在, CreateBlockchain 用于创建创世区块区块链, NewBlockchain 的作用应该就是获取区块链信息.
- 为什么send里面调用的是NewBlockchain(from)?
- NewBlockchain 和 MineBlock 有什么区别?
  - MineBlock 意思更明确
- 为什么交易输入里面有交易序号Txid,交易输出TXOutput没有交易序号?做完了?haha


## 小结
- 理解的话还是要看那个视频课程。
- 有代码的情况下，一定要忍住不要找其他资料，没有任何资料会比写代码的人更懂过程原理。
- 拿个笔和纸随便写写，有利于启发思路。

## 参考资料
- [Gavin Wood: 我与以太坊的二三事](https://mp.weixin.qq.com/s?__biz=MzU2MDE2MDU3Mg==&mid=2247487009&idx=1&sn=516b02f3b37d746f82a065201b2a295f&chksm=fc0d0472cb7a8d644c644e7834d7f1e1c89fc836f626b3c4a422ee97eea5dd6b357e4463b625&mpshare=1&scene=1&srcid=0710l7Q3ubkDQ6xPmpGNIeOc%23rd)
  - 这个访谈不错，其谈到了对比特币投资不感兴趣、对比特币背后的区块链应用技术很感兴趣、EOS项目比较水
- - [Bitcoin Block #0](https://www.blockchain.com/btc/block/000000000019d6689c085ae165831e934ff763ae46a2a6c172b3f1b60a8ce26f)
- 《区块链革命》对比特币和区块链的过程讲解的比较清楚

## ChangeLog
- 190308 ...
- 190307 代码理解，拖了几天了
- 190304 代码理解
- 190303 代码理解
- 190301 init