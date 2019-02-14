## [【区块链和 HyperLedger 微讲堂系列 】第一讲：区块链商用之道（视频回放，讲义下载） (developerWorks 中文社区)](https://www.ibm.com/developerworks/community/blogs/3302cc3b-074e-44da-90b1-5055f1dc0d9c/entry/opentech-blockchain-01?lang=en)
把区块链用作数字ID我还是能理解的，要说用到各行各业还是有点懵逼，真有那个必要吗？
17年区块链落地年

新兴技术成熟度曲线
数字货币交易所
加密货币

感觉就是把数据结构应用到了现实世界的资产，然后用技术进行管理。

## 《用Go语言构建自己的区块链教程-慕课网》
### 组成部分
1. 链式结构
2. 实现一个简单的http server，对外暴露读写接口
前面构造了区块的链式结构，可以理解为数据层。
上面应该就是网络层了，就是通过网络来管理链式结构，我是这么理解的。
数据是很好理解的，所有的难点在于数据的操作上。
### 创建block
觉得这个demo有点无聊啊。
go文件命名大写开头
区块里面存放的是很多交易数据
hash值用字符串表示
hashInBytes[:]是什么意思，原来hashInBytes是[32]bytes
新区块由上一个父区块产生
生成创世区块
区块：
1. 区块编号
2. 时间戳
3. 数据
### 创建Blockchain
把区块串成珍珠
1. 创建Blockchain文件
2. 创建结构体及相关方法，因为Go语言中没有类的说法，只有结构体
撤销 git commit
git log
git reset --hard 4c94c9a 这个操作会导致4c94c9a后的更改全部丢失
git reflog
git reset --hard b0136ee 恢复丢失的commit
git rest HEAD~ 这个是撤销commit的正确操作
参考[git bash - How do I undo the most recent commits in Git? - Stack Overflow](https://stackoverflow.com/questions/927358/how-do-i-undo-the-most-recent-commits-in-git)

Blocks []*Block是什么意思
不应该是链式结构吗？为什是区块数组？