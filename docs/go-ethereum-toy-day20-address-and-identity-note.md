> 时间流式学习记录  
> 原始记录及更新内容在GitHub仓库[go-ethereum-toy-day20-地址及身份标识笔记](https://github.com/chaozding/go-ethereum-toy/blob/master/docs/go-ethereum-toy-day20-address-and-identity-note)  
> [区块链技术笔记思维导图](https://www.processon.com/mindmap/5c623ed8e4b08a7683be4178)  
> [区块链应用笔记思维导图](https://www.processon.com/view/link/5c7a0699e4b0628e820ab8a2)  
> [区块链技术相关书签分享](https://www.gettoby.com/p/swgw7g8sayvw)  

[](https://github.com/chaozding/go-ethereum-toy/commit/)  

## 继续《私有区块链,我们一起GO-慕课网》
### 第5章 交易及记账
bitcoin_part4可执行调试,初始状态 [b201273](https://github.com/chaozding/go-ethereum-toy/commit/b201273a3f1a5ce83b32aca85155f4e5dbc0c217)
`cd src`
`go install coin` 安装命令行程序
报错了:
- 简化main.go代码,初始化命令行对象不再需要传递区块链对象进去了 [54bc3cf](https://github.com/chaozding/go-ethereum-toy/commit/54bc3cf38b0f68874ffa5819dea7ee41361fb578) 
- 添加 func dbExits() bool [bc23a32](https://github.com/chaozding/go-ethereum-toy/commit/bc23a3216a2af93b1e0fc18194430b14ff1c90fc)
- func (bc *Blockchain) MineBlock(transactions []*Transaction) 代替 func (bc *Blockchain) AddBlock(data string) [05ef498](https://github.com/chaozding/go-ethereum-toy/commit/05ef498c6e90d69253ddfeb911e73b2a463184d6)
- 添加 func (b *Block) HashTransactions() []byte [484dc35](https://github.com/chaozding/go-ethereum-toy/commit/484dc3595c827032c4dda2becdd1d383cf1e035d)
  - 调用 HashTransactions() [846d8d5](https://github.com/chaozding/go-ethereum-toy/commit/846d8d5c5355c8e6e8eb81aa6563f6f889899347)
- blockchain.go use outIdx [83543f5](https://github.com/chaozding/go-ethereum-toy/commit/83543f59b594a19d4f157bf84cebb2c588cfc743)
- cli.go use sendFrom sendTo sendAmount [89c8a6e](https://github.com/chaozding/go-ethereum-toy/commit/89c8a6eca37706611286452b0e163f86218a4fc7)
`cd ../bin`
`rm *.db`
`ls -l`
`coin`: 会打印使用方法
`createBlockchain -address ADDRESS`: 创建某个区块链(挖矿)针对某个地址, 需要给这个地址发奖励
`./coin createBlockchain -address Ivan`
![image](https://user-images.githubusercontent.com/16435896/54485379-6d801580-48b2-11e9-957d-121d72987b83.png)
然后可以发现区块链的存储文件已经存在了:
![image](https://user-images.githubusercontent.com/16435896/54485386-91dbf200-48b2-11e9-89f8-28d58daeb83e.png)
`getBalance -address ADDRESS`: 得到某个地址的余额(发币奖励/转账交易)
`./coin getBalance -address Ivan`
![image](https://user-images.githubusercontent.com/16435896/54485453-bb494d80-48b3-11e9-82de-04f23c1a39ca.png)
可以发现报错了
运行`printchain`, 发现也报错了
![image](https://user-images.githubusercontent.com/16435896/54485726-c0f56200-48b8-11e9-8060-b07d1551e626.png)
解决错误: [b7013b7](https://github.com/chaozding/go-ethereum-toy/commit/b7013b750296ebc5ce4149ff926c8194cbb0c7ef) 
`./coin getBalance -address Pedro`
![image](https://user-images.githubusercontent.com/16435896/54486175-598ee080-48bf-11e9-861c-addabf062574.png)
`./coin getBalance -address Helen`
![image](https://user-images.githubusercontent.com/16435896/54486178-76c3af00-48bf-11e9-9f6b-90d74fd7902a.png)
上面就是第一笔奖励交易的余额情况了.
`send -from FROM -to TO -amount AMOUNT`: 从FROM向TO转账AMOUNT个比特币
`./coin send -from Ivan -to Pedro -amount 3`
![image](https://user-images.githubusercontent.com/16435896/54486187-affc1f00-48bf-11e9-8ddc-e517c2534211.png)
`./coin printchain`
![image](https://user-images.githubusercontent.com/16435896/54486197-d0c47480-48bf-11e9-980b-83c2fa41becb.png)
第一笔转账交易后的余额情况:
![image](https://user-images.githubusercontent.com/16435896/54486202-f8b3d800-48bf-11e9-8a28-a89be9178448.png)
`./coin send -from Helen -to Pedro -amount 3`
![image](https://user-images.githubusercontent.com/16435896/54486212-31ec4800-48c0-11e9-9883-f2f4caa97b80.png)
`./coin send -from Ivan -to Pedro -amount 10`
![image](https://user-images.githubusercontent.com/16435896/54486221-4cbebc80-48c0-11e9-9a14-7c73d5297af3.png)
可以发现, 随便乱发钱是发不了的
再做一笔转账交易:
![image](https://user-images.githubusercontent.com/16435896/54486247-93141b80-48c0-11e9-88e2-1db209e78594.png)
再做一笔转账交易:
![image](https://user-images.githubusercontent.com/16435896/54486258-cce52200-48c0-11e9-9872-77f2c40a0a9c.png)

### [6-1 数字货币地址及身份标识](https://www.imooc.com/video/17562)
**数字货币的地址及身份标识**
在互联网上身份是用地址标识的
- 某种途径可以识别出你是交易输出的所有者,称之为身份.
- 在比特币中,你的**身份(identity)/钱包**就是一对(或者多对)公钥(public key)和私钥(private key).
  - 使用公钥加密待传输的信息, 只能使用配对的私钥解密加密的信息.
- 所谓的地址,只不过是将公钥表示成人类可读的形式而已.
  - 这里借用了公钥的唯一性复用为地址,当然可以用额外的地址,不过如果你用多对公钥,其实也差不多是这个意思了.
  - 公钥就是锁嘛,锁本来就是给别人看的,私钥(钥匙)才是关键.

**公钥算法和数字签名**
- 公钥加密(public-key cryptography)算法使用的是成对的密钥: 公钥和私钥.
- 公钥不是敏感信息,可以告诉其他人. 但是, 私钥绝对不能告诉其他人: 只有所有者(owner)才能知道私钥,能够识别、鉴定和证明所有者身份的就是私钥.
  - 私钥也可以理解为**数字签名**, 因为二者的作用是一致的.
- 在加密货币的世界中,你的私钥代表的就是你,私钥就是一切,可以理解为银行卡密码.
  - 那这个私钥岂不是太不安全了,银行卡密码取钱还可以挂失呢,私钥没了不是很麻烦吗?

bitcoin_part5初始化 [52ee686](https://github.com/chaozding/go-ethereum-toy/commit/52ee686ec47d775e678e96950a595762e9508b5a)
原来的代码文件拆分成了好多份 [1412dda](https://github.com/chaozding/go-ethereum-toy/commit/1412dda2fe7f89c64df8e58fca1fa3f08a01c02a) 
补全printUsage, 增加了两条命令createwallet、listaddresses [ff81777](https://github.com/chaozding/go-ethereum-toy/commit/ff81777f1ec68fedb54a6ea888f37ab4c179215a) 

### [6-2 数字钱包创建交易过程及代码实现](https://www.imooc.com/video/17563)
看看钱包如何创建?
钱包在交易当中有什么作用?
公钥地址背后的实际所有人是匿名的
`./coin createwallet`
得到一个地址: 一个34位的字符串
生成存储钱包的文件 wallet.dat
`./coin listaddresses`
得到公钥地址, 公钥地址对应的就是各个账户, 可以理解为一个公钥地址对应一个人, 当然一个人其实可以开多个账户的.
不同的地址可以理解为不同的人.
`./coin createblockchain -address 钱包公钥地址`
生成blockchain.db, 这个就是挖矿的过程
`./coin getbalance -address 钱包公钥地址`
查询余额, 账户都要填地址了现在
`./coin send -from 钱包公钥地址 -to 钱包公钥地址 -amount 3`

完成命令行参数切分, 主要是完成Run函数 [55a88b9](https://github.com/chaozding/go-ethereum-toy/commit/55a88b90d2cf49e965ad949b449b12b3b64b2fcd)

添加命令行函数 func (cli *CLI) createWallet() [714baf3](https://github.com/chaozding/go-ethereum-toy/commit/714baf3953edc2aa99ea39ea7f54a730724eea30)

添加函数 func NewWallets() (*Wallets, error) 用来创建钱包集合 [e5685c9](https://github.com/chaozding/go-ethereum-toy/commit/e5685c903ba04105268bdbefb4532f3429955882) 

创建单个钱包, 会创建钱包对象 func NewWallet() *Wallet [cae393c](https://github.com/chaozding/go-ethereum-toy/commit/cae393cc78596cfcc7a77b842979a8f32854d9f1)  

钱包里存的是一对两个变量,一个是私钥一个是公钥,使用函数 func newKeyPair() (ecdsa.PrivateKey, []byte) [d925ad3](https://github.com/chaozding/go-ethereum-toy/commit/d925ad3859e1ffaf528d2e9270cac3692546fa60) 

Wallets.go 里面的创建钱包函数 func (ws *Wallets) CreateWallet() string [29e6c2d](https://github.com/chaozding/go-ethereum-toy/commit/29e6c2d6ed590542505f528172b34084fc6c906d)

创建晚了钱包(公钥私钥对), 还要生成地址, 地址是字节数组类型的 [c27e722](https://github.com/chaozding/go-ethereum-toy/commit/c27e722c1b88ee08f65b4f308a900264f0d51f91) 

把公钥哈希化 func HashPubKey(pubKey []byte) []byte [c7c8560](https://github.com/chaozding/go-ethereum-toy/commit/c7c85604baad1c95b46d5f5407e878b600497f4e)

创建钱包后,需要把钱包的地址取得作为钱包集合的索引用 func (ws Wallets) SaveToFile() [a15010f](https://github.com/chaozding/go-ethereum-toy/commit/a15010fc5da06e2457d35bcf9c7c006fcb809635)


### 可执行调试

### 《私有区块链,我们一起GO-慕课网》课程总结
区块链的Go语言开源项目
- Go Ethereum 是官方使用Go语言实现的以太坊协议
- 金融领域的区块链项目Chain
- Hyperledger Fabric: 基于区块链的开源分布式账本

## Go语法
- 跨文件调用函数需要大写字母开头, 文件内调用小写就可以了;调用和函数的顺序无关;
```
if _, err := os.Stat(dbFile); os.IsNotExist(err) { //如果dbFile存在的话, 应该不执行
    return false
}
```
- txHash = sha256.Sum256(bytes.Join(txHashes, []byte{})) //为什么需要添加一个空的[]byte{}
- txID := hex.EncodeToString(tx.ID) //[]byte转换为string格式
- []byte 和 string类型是什么区别?

## 疑问
- 挖矿是按照地址分的是什么意思?
- 公钥做了一个哈希, 然后用base58编码出来了, 为什么公钥要做一个哈希呢?
- 钱包里存的是公钥私钥对, 数据是需要经过公钥加密的?
  - 钱包看上去就是一个34位字符串的地址
- 钱包是如何创建的?
  - 钱包是公钥私钥的集合.
- 公钥和私钥有什么区别吗?
  - 私钥被藏起来了?
  - 私钥使用的是返回的原始格式, 公钥使用的是字节数组格式[]byte, 然后还要编码为可读的字符数组形式.
    - 编码方式是先哈希再使用base58编码, 为什么不直接用哈希的结果, 而是绕了一圈?
- 为什么私钥也存储在钱包里, 这样不是泄漏了吗?
- 钱包、钱包地址、地址、公钥私钥、加密之间是什么关系?
- 公钥和地址是什么关系?
  - 钱包里面存放的是公钥和私钥对, 然后地址是由公钥和私钥对生成来的, 一个钱包有一个地址, 里面存的是公钥私钥对.
- 钱包在交易中有什么作用?
- 哈希值有什么作用?
- 校验是什么?
- Base58
  - Base58是用于Bitcoin中使用的一种独特的编码方式, 主要用于产生Bitcoin的钱包地址.相比Base64, Base58不使用互相混淆的字符.
  至少比01010101好多了.
## 小结
- 

## 参考资料
- [blockchain-tutorial/content/part-4 at master · liuchengxu/blockchain-tutorial](https://github.com/liuchengxu/blockchain-tutorial/tree/master/content/part-4) bitcoin_part4的示例代码
- [go - Does golangs os.Stat open the file? - Stack Overflow](https://stackoverflow.com/questions/50311374/does-golangs-os-stat-open-the-file)

## ChangeLog
- 190324
- 190323
- 190310 
- 190309 init  