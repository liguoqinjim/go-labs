### example
先要在代码里面加上expvar，加上自己要输出的参数，然后运行程序。expvarmon另外单独运行。

#### 注意点
 - 要是不需要输出自己的参数，就只要这样导入`import _ "expvar"`

#### expvarmon运行命令
 - `expvarmon.exe -ports="1234"` (要监控的端口号是1234)
 - `expvarmon.exe -ports="1234" -vars="Goroutines,visits"` (Goroutines,visits就是在程序中自己加的要输出的参数)
