# 压缩打包大小

## 打包命令
 - `go build -o lab002_1 main.go`
 - `go build -o lab002_2 -ldflags "-w -s" main.go`
 - `upx lab002_2 -o lab002_3`
 - `upx lab002_2 --brute -o lab002_4`

## 效果
![Imgur](https://imgur.com/VDmEjNT)