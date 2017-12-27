### net/http/pprof
在程序里面import，`_ "net/http/pprof"`，然后监听端口，就可以在网页里面看到问题

1. 浏览器访问`http://localhost:7777/debug/pprof`可以看到信息
2. go tool pprof访问，`go tool pprof http://localhost:7777/debug/pprof/goroutine`

#### 参考资料
http://xiaorui.cc/2016/03/20/golang%E4%BD%BF%E7%94%A8pprof%E7%9B%91%E6%8E%A7%E6%80%A7%E8%83%BD%E5%8F%8Agc%E8%B0%83%E4%BC%98/