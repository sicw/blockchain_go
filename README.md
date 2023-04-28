# Blockchain in Go

A blockchain implementation in Go, as described in these articles:

1. [Basic Prototype](https://jeiwan.net/posts/building-blockchain-in-go-part-1/)
2. [Proof-of-Work](https://jeiwan.net/posts/building-blockchain-in-go-part-2/)
3. [Persistence and CLI](https://jeiwan.net/posts/building-blockchain-in-go-part-3/)
4. [Transactions 1](https://jeiwan.net/posts/building-blockchain-in-go-part-4/)
5. [Addresses](https://jeiwan.net/posts/building-blockchain-in-go-part-5/)
6. [Transactions 2](https://jeiwan.net/posts/building-blockchain-in-go-part-6/)
7. [Network](https://jeiwan.net/posts/building-blockchain-in-go-part-7/)

运行步骤
```
1. 下载完代码后, 执行go mod init, go mod tidy
2. go run . 或者 go run blockchain_go查看命令行

开三个终端创建钱包
export NODE_ID=3000
go run . createwallet
wallet_3000: 1KY1ZHxTe4NiXFW3MoCirVHrbxpJNwckF5

export NODE_ID=3001
go run . createwallet
wallet_3001: 1CgemWVkwShYZ2ABocMmpX8unoAEasv4Ux

export NODE_ID=3002
go run . createwallet
wallet_3002: 12p47oV9vzSg3zKKg4uSLsRtEfSnSnwKar

创建区块链
回到3002端口, 执行 go run . createblockchain -address 12p47oV9vzSg3zKKg4uSLsRtEfSnSnwKar (3000端口为公共节点 发布地址的)
复制blockchain_3002.db 为blockchain_3000.db、blockchain_3001.db(所有节点只是拥有相同区块的副本)

启动节点
3000: go run . startnode
3001: go run . startnode -miner 1CgemWVkwShYZ2ABocMmpX8unoAEasv4Ux
3002: 不启动, 直接在terminal使用命令NODE_ID=3002
```