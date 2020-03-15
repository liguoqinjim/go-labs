# example

## NOTICE
 - go程序可以在容器中编译，也可以在外部交叉编译之后放入容器中运行。
 - 不在容器中编译的话，镜像使用scratch即可，这个镜像很小，编译的时候要使用  
 `CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main.go .`  
 scratch 镜像，它是空镜像，因此我们需要将生成的可执行文件静态链接所依赖的库

## 运行命令
 - `docker build -t iris_test .` 构建镜像
 - `docker run --rm -p 8000:8000 iris_test` 运行

## 参考资料
 - https://segmentfault.com/a/1190000013960558