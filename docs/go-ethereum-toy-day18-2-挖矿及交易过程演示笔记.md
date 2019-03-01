> 时间流式学习记录
> 原始文档及更新内容在GitHub仓库 [go-ethereum-toy-day18-2-挖矿及交易过程演示笔记](https://github.com/chaozding/go-ethereum-toy/blob/master/docs/go-ethereum-toy-day18-2-挖矿及交易过程演示笔记.md)
> [区块链学习笔记思维导图](https://www.processon.com/mindmap/5c623ed8e4b08a7683be4178)
> [](https://github.com/chaozding/go-ethereum-toy/commit/)

## 继续《私有区块链，我们一起GO-慕课网》
### 5-2 挖矿及交易过程演示
哈希值是16进制的数值吗吗？还是仅仅是64位长度的字符串？
为什么目标是用`target.Lsh(target, uint(256 - targetBits))`
挖矿难度示意代码 [205e154](https://github.com/chaozding/go-ethereum-toy/commit/205e15455053a9cdd053749c662c85b110cdba90)
`./coin` 显示用法，内部没有执行任何命令
`./coin createblockchain -address Ivan` 这个Ivan此处是地址，表示挖矿的奖励到账10个比特币
`./coin getbalance -address Ivan` 获知余额
没挖矿就没有余额，地址和区块是什么关系？
`./coin send -from Ivan -to Pedro -amount 3` 发3个比特币作为生日礼物
系统很健壮，想随便发钱是不可以的。
`./coin printchain`


## Go语法
`target.Lsh(target, uint(1))`移位运算是什么意思？
这里的移位是指target转换为二进制情况下的左移运算，比如1左移1位得(10)2，转换位十进制则是(2)10，转换为16进制则是(2)16。
比如1左移7位得(1000 0000)2，转换为十进制则是(128)10，转换为16进制则是(80)16。
测试代码截图如下：
![image](https://user-images.githubusercontent.com/16435896/53574362-6e266580-3baa-11e9-859e-fcfe0987d295.png)
`rm *.db`

## 疑问
- 每个地址可以看作一个账户（可以理解为一个人/企业的银行账户），地址是如何产生的，如何保证是有效的地址？
- 挖矿是否可以理解为银行呢？
- 第一个比特币是哪里来的，那时不是还没区块链吗？
  - 这个问题最好还是看白皮书清楚，否则都是只言片语。
- 最初的50个比特币存在哪里？
- 奖励的币和区块产生的币是什么关系？
- 比特币和金本位的黄金有什么区别？
- 为什么比特币需要用现金买？而且价格有波动，既然是为了取代货币，不是应该价格不变互换吗？
- 感觉比特币的最大的价值在于区块链概念，投资价值完全是泡沫，但是搞不懂比特币多出来的现实金钱是从哪里的？
  - 感觉还是要从数字货币的历史来理解。
- 比特币投资的积极意义是什么？
- 数字化货币、数字货币、虚拟货币、电子货币、加密货币有什么区别和关系？
- 钱包和地址什么区别？
- 钱包和区块链的关系是什么？
- 小数个比特币是什么意思？
- 交易记录和区块是什么关系？

## 小结

## 参考资料
- [世界上第一枚比特币是怎么来的？ - 知乎](https://www.zhihu.com/question/265573962)
- [Bitcoin Block #0](https://www.blockchain.com/btc/block/000000000019d6689c085ae165831e934ff763ae46a2a6c172b3f1b60a8ce26f)
- [区块链浏览器——搜索以太坊区块链 | BTC | ETH](https://www.blockchain.com/zh-cn/explorer)
- [第一个比特币诞生啦\_百度经验](https://jingyan.baidu.com/article/6b97984def64b91ca2b0bfb9.html)
- [liuchengxu/blockchain-tutorial: A step-by-step blockchain tutorial in simplified Chinese](https://github.com/liuchengxu/blockchain-tutorial)

## ChangeLog
- 190301 init