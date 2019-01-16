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

|实验|简介|说明|
|---|---|---|
|lab001|example_client||
|lab002|example_server||