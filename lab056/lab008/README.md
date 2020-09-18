# 得到某个前缀有多少个key

## 命令
 - 在redis命令行中，`eval "return #redis.call('keys', 'prefix-*')" 0`
 - 使用go-redis库，`return #redis.call('keys', 'prefix-*')`，要去掉最后的0