> 时间流式学习记录
> 原始文档及更新内容在GitHub仓库，文档里的40位哈希值需要在仓库里面才生成链接
> [go-ethereum-toy-day16-私有区块链课程笔记](https://github.com/chaozding/go-ethereum-toy/blob/master/docs/go-ethereum-toy-day16-私有区块链课程笔记.md)

上午荒废了半天，主要是想到昨天那个课程示例代码比较多，看明白了挺费时，就潜意识里拖着，拖到下午还是得看啊，还不如先开个logging文档开始呢。

## 继续《私有区块链，我们一起GO-慕课网》
### 继续 4-1 工作量证明及哈希算法
初始化工程，把bitcoin_part1复制为bitcoin_part2 bef695fbcd8aa6ad2c1314226063ec1d8f82bc5a
配置项目GOPATH路径。
什么是工作量证明？
1. 工作的结果Nonce作为数据的一部分加入区块链从而成为一个区块 2727a9989124232041bb2ffb37765c7c716fa336 。
计算当前区块哈希值的时候，包含了工作的结果数据吗？
工作的结果是证明吗？
还是说的工作的结果需要额外的证明？
2. 完成这个工作的人也会获得奖励，完成这个工作量就是指的是挖矿。
3. 整个“努力工作（消耗电能）并进行证明（证明你确实消耗了电能）”的机制，就叫做工作量（费电（数学难度））证明 6c5d6b15e18bf6719cd35bde36f497054a0daf40 ，独一无二的苦力是应该得到报酬的，如果这个苦力足够苦的话。
代码可以从结果往回逆推，这里代码更新是通过往旧的实现里添加新代码实现的。
区块里面的数据部分是什么东西？

给区块数据结构 block.go 添加工作量证明部分 06ce2bf3c82f1e08187683ca26bd5faf13591994 
如何完成一定的工作（苦力活俗称挖矿），从而生成一个很特殊的哈希值，并返回工作量？
因为这个步骤比较复杂，所以新建一个proofofwork.go的文件来负责这件事，这个proofofwork.go可以理解为一个类文件，还是用面向对象习惯思维好理解，类名就是ProofOfWork，然后又一个函数NewProofOfWork()作为构造函数，初始化ProofOfWork的数据，后面还有类的方法了。
0439130144ca65669fa650455e878a02c5760dd8

target.Lsh(target, uint(256 - targetBits))是什么意思？为什么 const targetBits = 20 表示要求满足哈希值前面5个全是0？256 - 20 = 236，接下来呢？

fix block.go 75f666a4a2be133398d5e7a1a0209319016676dd
添加挖矿和工作量证明 
创世区块也要进行工作量证明
算法本身很简单，但是计算满足要求的难，有点类似暴力算法。
挖矿难度对所有人一视同仁，所以出现了矿机。
努力工作经过证明，则奖励比特币，比特币怎么来的？

## 语法
结构体初始化是{}，例如 ProofOfWork{b, target}。
“类”的方法的接收器一般使用指针形式，例如 
err := binary.Write(buff, binary.BigEndian, num)这个二进制读写什么意思，不是字节读写吗？
big.Int 是整数
if hashInt.Cmp(pow.target) == -1s表示相等吗？看来-1就是表示big.Int相等了
\r表示什么类型
hashInt整数
strconv.FormatBool(pow.Validate())转换布尔类型为字符串类型
## 变更
- 190226 init