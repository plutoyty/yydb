# yydb

yydb是一个准备用Go语言实现的一个KV存储数据库。待开发中.....

开发计划

- [ ] 自动过期功能(TTL)
- [ ] AOF 持久化及 AOF 重写
- [ ] 加载和导出 RDB 文件
- [ ] `Watch` 命令和 CAS 支持
- [ ] `Multi` 命令
- [ ] 加载RDB文件
- [ ] 主从模式
- [ ] 哨兵

关键功能

//TODO

架构设计

//TODO

集群模式

//TODO

待支持的命令

###  String

- Set
- SetNx
- Get
- GetSet
- Append
- StrLen
- StrExists
- StrRem
- PrefixScan
- RangeScan
- Expire
- Persist
- TTL

### List

- LPush
- RPush
- LPop
- RPop
- LIndex
- LRem
- LInsert
- LSet
- LTrim
- LRange
- LLen
- LKeyExists
- LValExists

### Hash

- HSet
- HSetNx
- HGet
- HGetAll
- HDel
- HExists
- HLen
- HKeys
- HVals

### Set

- SAdd
- SPop
- SIsMember
- SRandMember
- SRem
- SMove
- SCard
- SMembers
- SUnion
- SDiff

### Zset

- ZAdd
- ZScore
- ZCard
- ZRank
- ZRevRank
- ZIncrBy
- ZRange
- ZRangeWithScores
- ZRevRange
- ZRevRangeWithScores
- ZRem
- ZGetByRank
- ZRevGetByRank
- ZScoreRange
- ZRevScoreRange

性能测试

//TODO



