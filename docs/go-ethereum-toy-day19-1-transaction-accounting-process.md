> 时间流式学习记录
> 原始记录及更新内容在GitHub仓库 [go-ethereum-toy-day19-1-交易记账流程笔记](https://github.com/chaozding/go-ethereum-toy/blob/master/docs/go-ethereum-toy-day19-1-transaction-accounting-process.md)
> [区块链技术笔记思维导图](https://www.processon.com/mindmap/5c623ed8e4b08a7683be4178)
> [区块链应用笔记思维导图](https://www.processon.com/view/link/5c7a0699e4b0628e820ab8a2)
> [区块链技术相关书签分享](https://www.gettoby.com/p/swgw7g8sayvw)

[](https://github.com/chaozding/go-ethereum-toy/commit/)

## 继续学习《私有区块链,我们一起GO-慕课网》
### [5-4 交易记账及代码实现](https://www.imooc.com/video/17561)
**交易示例**
![image](https://user-images.githubusercontent.com/16435896/53558549-2b9d6280-3b83-11e9-865a-bdfebd846d07.png)
来源: [5-1 交易记账原理及过程]()
上图中交易Transaction就是区块里面存放的东西.
交易输入TXInput有输入序号(其实作用就是类似下图中的输出到输入的箭头,标明是哪个交易的交易输出作为新交易的交易输入)
交易输入TXInput有来源地址的钱.
交易输入TXInput有来源地址的值.

交易输出TXOutput其实并不需要序号,因为不像交易输入那样需要标明来源.
交易输出TXOutput需要的有去向地址的余额.  
交易输出TXOutput需要的有去向地址的值.  

图中没有箭头出来的的OutputX所在的交Transactions就构成了unspentTXs []Transaction->过滤出其中的目标地址  

#### 创建一个转账交易的过程
创建一个转账交易 func NewUTXOTransaction(from, to string, amount int, bc *Blockchain) *Transaction [a77fde1](https://github.com/chaozding/go-ethereum-toy/commit/a77fde1c1e9a060df8b00ae630b2b8c6adfc42bc)
在没有消费掉的交易输出中再找到可以消费(相对于新交易,余额足够的)的OutputX,用于形成新的OutputX到InputX的箭头 func (bc *Blockchain) FindSpendableOutputs(address string, amount int) (int, map[string][]int) [30bbaef](https://github.com/chaozding/go-ethereum-toy/commit/30bbaefe7636119f7006c4b01172471ae9cc2187)
找到没有消费掉(相对于已经产生的交易)的OutputX,也就是还没有箭头出来的OutputX,不需要余额足够 func (bc *Blockchain) FindUnspentTransactions(address string) []Transaction [9b2c82d](https://github.com/chaozding/go-ethereum-toy/commit/9b2c82d549edb7eae4d18688c83454f9ece207f4)

**交易机制**
![image](https://user-images.githubusercontent.com/16435896/54068038-bc122c00-4282-11e9-91bf-819d3c022018.png)
来源: [5-4 交易记账及代码实现](https://www.imooc.com/video/17561)  
可以使用map[string][]int的数据结构存储整个交易的交易输出部分.
凡是有输入的地方都表示这个比特币被花出去了,如果没有全花掉,则会转变为找零成为交易输出部分.  
上图中凡是OutputX和InputX之间有连线的,他们其实就是一回事,因为交易输入必须花光来源地址的钱,所以OutputX的钱必须全部转移到InputX中.  
上图中的有的OutputX上没有线引出来,说明就是有余额的,然后找出钱包地址. 

添加(tx Transaction) IsCoinbase() bool [5e180e7](https://github.com/chaozding/go-ethereum-toy/commit/5e180e756325597301c80843f2b82b892e931cb5)  
能实用解锁数据解锁交易输出TXOutput, (in *TXInput) CanUnlockOutputWith(unlockingData string) bool [6ba63e1](https://github.com/chaozding/go-ethereum-toy/commit/6ba63e13ec3fe0fbd5275f31130c78147b678226)
能被解锁数据解锁交易输出TXOutput, (out *TXOutput) CanBeUnlockedWidth [f7c1877](https://github.com/chaozding/go-ethereum-toy/commit/f7c1877540133f8ae2ad7b16f730a07fc9296867)
这两个的唯一区别调用主题不同,一个是交易输入(主动),一个是交易输出(被动)

获取某个账户地址的余额
查找未花掉余额的输出的所有交易/找到属于我的那些没有花出去的交易 func (bc *Blockchain) FindUTXO(address string) []TXOutput [66ccf5d](https://github.com/chaozding/go-ethereum-toy/commit/66ccf5d7b94f463a5a8075ac91886c83f295eb78)

数字地址:比特币钱包
## Go语法
- hex.EncodeToString([]byte) 
- func append(slice []Type, elems ...Type) []Type

## 疑问
- 没有花出去的交易是什么意思?交易怎么能花费的?
  - FindUnspentTransaction returns a list of transactions containing unspent outputs
  - 比特币要么你没花,要么你花出去了然后找零(找零就变成交易输出了),所以有箭头连着说明就是花掉了
  - 交易输出类型(比特币都是对应特定地址的)
    - 找零剩下
    - 别人转账给我的
  - 交易输出类型但是花掉的(全花掉)
    - 找零剩下花掉了
    - 别人转账给我的花掉了
  - 把我自己的找出来
  - ![image](https://user-images.githubusercontent.com/16435896/54044054-048d0380-4209-11e9-9798-c4a21dd3b33a.png)
    - 
- 交易ID和交易输入的ID什么区别?
  - 交易输出应该不需要序号吧,输出就是对应账户地址的(排除掉花掉的输出/就是有输出到输入箭头的那种)
  - 交易输入的ID应该就是交易ID吧,没有在合并具体的交易输出ID
- 交易输入ID是怎么生成的?
  - 交易输入序号/ID是之前某个交易的某个输出:对应图中的输出到输入的箭头
- 每一个交易的OutputX只能对应一个地址吗,有没有可能同时出现多个OutputX对应同一个地址?
  - 按照代码,好像不是这么回事啊,可能一个地址可以解锁多个交易输出,从而同时用于支付(交易输入)
  - 这样子一来的话地址的余额就是分散在多个交易的交易输出之中了,这不合理,虽然可以运行
  - 获取余额是通过遍历未花出去的交易输出实现的
- ScriptSig 和 ScriptPubKey 什么区别?
- 交易是多个账户地址向多个账户地址转账吗?

代码问题:
- unspentOutputs[txID] = append(unspentOutputs[txID], out.Value) [aa5b687](https://github.com/chaozding/go-ethereum-toy/commit/aa5b6877bd790638061b3b323af3e47cd4611fa0)

## 小结
- 看三遍

## 参考资料
- 

## ChangeLog
- 190301 init