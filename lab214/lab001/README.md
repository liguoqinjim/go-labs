# example

## 命令
### 1 
```shell script
docker run --rm -it --name go-http-demo \
  -v $PWD:/go/src/example.com/go-http-demo \
  -p 8000:8080 \
  golang
```

在容器中到/go/src/example.com/go-http-demo下, go run main.go

### 2使用docker-compose
```shell script
docker-compose up -d
```

## 参考资料
 - https://segmentfault.com/a/1190000021660020