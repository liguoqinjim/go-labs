# [confluent-kafka-go](https://github.com/confluentinc/confluent-kafka-go)
安装:`go get -u -v github.com/confluentinc/confluent-kafka-go`

## 注意
 1. 先安装`librdkafka`
 ```
 git clone https://github.com/edenhill/librdkafka.git
 cd librdkafka
 ./configure --prefix /usr
 make
 sudo make install
 ```

 2. mac可能需要安装pkg-config,`brew install pkg-config`
 
 3. centos，可以先使用`yum install librdkafka-devel`，这样好了之后还是不行的话。再把1的步骤做一遍
 
 4. windows. 首先下载`https://sourceforge.net/projects/pkgconfiglite/`把bin路径加到path中。
    之后的还没有试验成功，所以暂时windows上还不能运行

|实验|简介|说明|
|---|---|---|
|lab001|example_client| |
|lab002|example_server| |