# [zap](https://github.com/uber-go/zap)

|实验|简介|说明|
|---|---|---|
|lab001|demo|SugarLogger和Logger |
|lab002|example/AdvancedConfiguration|https://pkg.go.dev/go.uber.org/zap 示例 |
|lab003|example/BasicConfiguration|https://pkg.go.dev/go.uber.org/zap 示例  |
|lab004|实验| |
|lab005|模板|gist |

## NOTICE
 - 不需要性能的时候使用SugarLogger，需要性能的时候就使用Logger，但是log只能
 - DPanic是特有的，是在development模式下是panic，但是production模式下不会