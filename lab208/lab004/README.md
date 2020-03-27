# flag

## 运行命令
 - `go run main.go --flagName=123 -t=false -s=string -c=1 -c=2 -c=3`
 
## NOTICE
 - `-flagName=123`是无效的，因为单横线只会用来解析shorthand
 - `pflag.BoolP("testMode", "t", true, "是否进入测试模式")`,BoolP里面就这是了t为shorthand,所以-t=false是可以解析的
 - StringVarP，可以直接赋值
 - shorthand只能单个字母
 - flag得到数组的时候命令如下：`go run main.go -c=127.0.0.1:2181 -c=127.0.0.1:2182`，这样就会得到数组了，而不是用符号切分