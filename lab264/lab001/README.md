# 打包

## 打包命令
 - `go build`
 - `go build main.go`
 - `go build -o lab001 main.go `，指定生成文件的名称
 - `GOOS=linux GOARCH=amd64 go build -o lab001 main.go`，打包到linux
 - `GOOS=windows GOARCH=amd64 go build -o lab001 main.go`，打包到windows