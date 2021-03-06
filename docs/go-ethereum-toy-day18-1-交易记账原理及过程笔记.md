> 时间流式学习记录
> 原始文档及更新内容在GitHub仓库 [go-ethereum-toy-day18-0-printchain问题解决](https://github.com/chaozding/go-ethereum-toy/blob/master/docs/go-ethereum-toy-day18-0-printchain问题解决.md)
> [](https://github.com/chaozding/go-ethereum-toy/commit/)
> [区块链学习笔记思维导图](https://www.processon.com/mindmap/5c623ed8e4b08a7683be4178)

## 继续学习《私有区块链，我们一起GO-慕课网》
### 5-1 交易记账原理及过程
交易记账和区块是什么关系？
区块创建后数据部分是固定的吗？
比特币和区块的哈希值是什么关系？
比特币和挖矿的关系是什么？
地址和区块是什么关系？
一个地址对应一个区块吗？
比特币总额固定？但是现实里银行里的钱不是固定的啊？
开始整个系统还没有比特币的时候，挖矿是交易记录吗？
**交易及记账**
- 交易机制
  - 区块链能安全可靠地存储交易结果
  - 交易一旦被创建，就没有任何人能够再去修改或者删除
  - 交易由一些输入（input）和输出（output）**组合**而来
    - 输入应该就是谁跟谁交易吧？输出是什么？
系统找零：钱都是系统来消费的
系统找零的币=没有花出去的币
交易输出记录的是余额
余额
如果你的币花光了，则没有系统找零，则没有输出
这个课讲的真清楚，下面的截图均来自于[交易记账原理及过程，私有区块链，我们一起GO教程-慕课网](https://www.imooc.com/video/17558)，理解的话还要看视频
![image](https://user-images.githubusercontent.com/16435896/53558549-2b9d6280-3b83-11e9-865a-bdfebd846d07.png)
**复杂交易情形**
![image](https://user-images.githubusercontent.com/16435896/53558762-a1093300-3b83-11e9-942b-650eed1751b7.png)
有点难理解，最好有个例子。
记账和智能合约是什么关系呢？
- **交易机制**
  - 对于每一笔新的交易，它的输入会引用reference之前一笔交易的输出
  - 交易的输出，也就是比特币实际存放的地方
  - 例外：
    - 有一些输出并没有被关联到某个输入上
    - 一笔交易的输入可以引用之前多笔交易的输出
    - 一个输入必须引用一个输出
没看懂上面的？

## 小结
把笔记拆成单个文档压力小了，和放一个文档里其实也没区别。


## ChangeLog
- 190228 init