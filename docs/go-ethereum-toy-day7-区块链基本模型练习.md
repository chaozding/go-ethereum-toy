两天都没干活，因为其他的拖着，其他的因为这个拖着。

## 继续《用Go语言构建自己的区块链教程-慕课网》

### 继续创建Blockchain
Blockchain里面是Block
相当于对链式数据结构操作
如果函数名签名有变量，则函数可以理解为变量的方法
不管了，先照抄吧，不懂的先放放，出了问题再说吧
跳小括号外面，试试 ctrl shift enter，叫做 Complete Current Statement
每个代码都有无数种设计，重要的是达到需求定义，所以不要再想为什么设计这个函数而不是那个函数了，等有时间了再想这个问题
用文件夹对go文件分类
新建一个main文件调用
遇到一个错误
![image](https://user-images.githubusercontent.com/16435896/52930247-a5af3980-3382-11e9-9caa-fce428aa9904.png)
后来发现这个core本身怎么就包含区块链了，奇怪了
后来发现是import错了，golang自动关联到src目录下的其他项目了

包外引用的函数的首字母需要大写，包内可以小写
log.Fatal("invalid block")
感觉Go还挺好用的，还不会调试，智能用fmt.Printf()了

### 创建 Http Sever
1. 创建 Http Sever
2. 提供 API 访问接口
也就是远程过程调用RPC，所以这个放到rpc文件夹里面
run方法提供运行服务
感觉直接调用http方法有些无聊
记得把包改为package main

### 总结
1. 什么是区块链，从定义的角度看，然后根据定义用代码实现
2. 区块链架构模型
3. 区块链基本模型的代码实现，基本模型是很简单的，但是把最核心的概念都实现了

## 《我眼中的比特币与区块链》科普视频笔记
[我眼中的比特币与区块链 - YouTube](https://www.youtube.com/playlist?list=PLVac-xziP_7MSnL_htBOrsJZpPXOtuDZV)
[什么是钱包？（韭菜入门 ） - YouTube](https://www.youtube.com/watch?v=clKXUP1eg74&t=0s&index=13&list=PLVac-xziP_7MSnL_htBOrsJZpPXOtuDZV)
20亿人没银行账户？
50亿人无法自由参与国际投资，这我相信。
疑问评论摘录：
1. 视频中所说的交易平台是否可以理解为实际货币世界中的“银行”？如果是的话，这是否和区块链和比特币的初衷是去中心化背道而驰？
2. 如果密钥是能开钱包的唯一方法，那么丢失的秘钥的钱包最终如何处置？比如现实世界中 丢失的钱还会被捡到的人重新投入到系统中去，比特币世界中丢失的钱就被销毁/浪费掉了么？
加密货币，数字货币，虚拟货币到底有什么区别呢？
[TA们将来玩什么？我为什么要投资区块链 - YouTube](https://www.youtube.com/watch?v=USUpfYPn1nk&list=PLVac-xziP_7MSnL_htBOrsJZpPXOtuDZV&index=13)
体制是什么？->游戏规则
赌博和投资的区别是什么？
不动产+证券+技术+经济
[银行会不会消失 - YouTube](https://www.youtube.com/watch?v=CbF4tpDxU-0&index=14&list=PLVac-xziP_7MSnL_htBOrsJZpPXOtuDZV)
我一直以为银行是国有的，想了下，这背后是信任的思维在起作用。
既然信任和钱有关，那么我觉得一切营销的目的就是为了取得信任。
为什么新技术不可以逆转？
把信任转化到技术（数学）上？而区块链部分地实现了，神奇的是区块链早就被发明了，却因为比特币而闻世，为什么曾经比特币的价值会虚高不下？
账簿和区块的关系是什么？
至于银行（金融产业）是否会永远存在，银行也是一类企业，核心战略优势是信任，不确定，还是先找资料研究下再做思考。
[所有人 VS 没有人的区块链？ - YouTube](https://www.youtube.com/watch?v=vk-pk2QJgiU&list=PLVac-xziP_7MSnL_htBOrsJZpPXOtuDZV&index=15)
信息畅通无阻了吗？
方法：连接
目的：没有阻碍
东西：信息
价值传递有意思
讲的东西比《区块链之新》的还提前好多
信任->价值传递
%99的公司会牺牲，%1的公司会创造真正的价值？->尽职调查
一个新技术革命，不是所有人就是没有人。
非洲怎麼沒有銀行，越南也有銀行。問題是:他們不相信銀行。
[”智能合约“ 和你有什么关系？ - YouTube](https://www.youtube.com/watch?v=bOGCqD-5jgQ&list=PLVac-xziP_7MSnL_htBOrsJZpPXOtuDZV&index=3)
智能合约：天下人都知道的自动执行的欠条，这个有意思，应该是针对数字货币来说，用欠条比喻更好理解。
合约具有
[什么是比特币Bitcoin - YouTube](https://www.youtube.com/watch?v=wcSb1womTiI&list=PLVac-xziP_7MSnL_htBOrsJZpPXOtuDZV&index=1)
比特币：一种使用区块链技术来记录交易数据的“货币”。具有某些特殊的特性。
全世界都知道交易，但是不知道是谁交易的。
挖矿：这简直太形象了，真的就是数字“黄金”，当然得有人信。我也可以把一张纸撕成碎片，这些碎片当然是地球上独一无二的，某种程度上也可以理解为纸“黄金”，但和比特币的属性比起来就差远了。
记录越多，挖矿越难。
感觉比特币已经和股票区别不大了，给比特币价值背书的就是加密货币概念的优势。
[什么是ICO？ - YouTube](https://www.youtube.com/watch?v=2Wx1HOjqgQU&list=PLVac-xziP_7MSnL_htBOrsJZpPXOtuDZV&index=4)
讲的很形象
ICO和非法集资什么区别？
ICO=众筹+IPO 这和股票发行制度有啥本质区别？区别就是没监管了，本来股票的弊端不是出在监管上，现在最重要的监管反而没了，ICO到处是风险，感觉问题就出在众筹那一步。
[当马爸爸来敲门， 投资小白们需要注意什么？ - YouTube](https://www.youtube.com/watch?v=SCjYBZ4WTrQ&list=PLVac-xziP_7MSnL_htBOrsJZpPXOtuDZV&index=5)
余额宝？
各种基金？
投资其实是投资人，世界是人创造的，即使机器人也有某种人的成分，我们习惯看产品说明书和产品广告，却从不去了解是谁创造了这个产品，谁投资了这个产品。
[目前全球最大的骗局，每年涉入资金2000亿 - YouTube](https://www.youtube.com/watch?v=3PuqlTvR0yg&list=PLVac-xziP_7MSnL_htBOrsJZpPXOtuDZV&index=6)
传销和ICO类似，ICO和股票类似。
这个视频做的太好了，手漫画的教学效果不错。
[从“一根香蕉”到“数字货币” — 钱为何物？ - YouTube](https://www.youtube.com/watch?v=iCA8dXEn6QY&list=PLVac-xziP_7MSnL_htBOrsJZpPXOtuDZV&index=7)
[不要买比特币！DON‘T BUY BITCOINS!!! 因为... - YouTube](https://www.youtube.com/watch?v=a2YOMAhQDko&list=PLVac-xziP_7MSnL_htBOrsJZpPXOtuDZV&index=9)
玩比特币就是玩游戏，是需要本金和风险控制的，如果连比特币都买不到就算了吧。
[21+ Ways to Buy Bitcoins Online 2019 (Trusted Exchanges)](https://www.buybitcoinworldwide.com/)
[40%的比特币被1000个人所拥有，他们是谁？ - YouTube](https://www.youtube.com/watch?v=JWGUUNWbz_o&list=PLVac-xziP_7MSnL_htBOrsJZpPXOtuDZV&index=10)
这期视频有点意思。
[一只韭菜的自白 - YouTube](https://www.youtube.com/watch?v=e315PZfa1N8&list=PLVac-xziP_7MSnL_htBOrsJZpPXOtuDZV&index=11)
