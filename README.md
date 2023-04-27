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
wallet_3000: 177i4fJa9SQoBd1zKJq9n6F5LTvgtLTSB7
wallet_3001: 15XZRhs3saME2zCGU9q4ib3Dj1Sh1p2L6b
wallet_3002: 13rekSfXyP3LEGthYuWAz8bFzHd7NpY4Km

1. 下载完代码后, 执行go mod init, go mod tidy
2. go run . 或者 go run blockchain_go查看命令行

3. 设置NODE_ID export NODE_ID=3000 3000也是启动的端口号
4. 运行wallet_test.go创建钱包并获取地址
5. 运行创建区块链 go run . createblockchain -address 1uTGWBZdaEMvCCtYYJwRxsUPBywaWWNS6
6. go run . startnode启动节点

7. export NODE_ID=3001设置第二个节点
8. 运行创建区块链 go run . createblockchain -address 162aVDkUo9BFro12f6kTUWV3iNc9ngU3ak
9. go run . startnode启动节点

10. export NODE_ID=3002设置第二个节点
11. 运行创建区块链 go run . createblockchain -address 1DJH85nnnCu1ptmrqL82K4ivSBeVbQDrk
12. go run . startnode -miner 1JUFXGucye42HEHVjKT2WhHcT9cCGVgK6k 启动节点&支持挖矿
13. 重复7、8、9步骤创建新节点
```