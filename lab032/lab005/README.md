# 查看代码覆盖率

## 命令
### 生成测试数据
`go test -covermode=count -coverprofile=cover.out`

### 查看数据
`go tool cover -func=./cover.out`

`go tool cover -html=./cover.out`

## covermode含义
- set: 每个语句是否执行？
- count: 每个语句执行了几次？
- atomic: 类似于 count, 但表示的是并行程序中的精确计数

## 参考资料
 - https://brantou.github.io/2017/05/24/go-cover-story/