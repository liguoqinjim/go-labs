# flag

## 运行命令
 - `go run main.go --flagName=123 -t=false -s=string`
 
## NOTICE
 - `-flagName=123`是无效的，因为单横线只会用来解析shorthand
 - `pflag.BoolP("testMode", "t", true, "是否进入测试模式")`,BoolP里面就这是了t为shorthand,所以-t=false是可以解析的
 - StringVarP，可以直接赋值