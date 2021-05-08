# 压缩打包大小

## 打包命令
 - `go build -o lab002_1 main.go`
 - `go build -o lab002_2 -ldflags "-w -s" main.go`
 - `upx lab002_2 -o lab002_3`
 - `upx lab002_2 --brute -o lab002_4`

## 效果
![Imgur](https://imgur.com/VDmEjNT)

## 参考资料
 - https://abelsu7.top/2019/10/24/go-build-compress-using-upx/
 - https://www.jianshu.com/p/cd3c766b893c