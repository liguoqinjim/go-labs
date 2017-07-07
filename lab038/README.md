### protobuf

#### 库
`go get github.com/golang/protobuf/`

`go get github.com/gogo/protobuf`

#### protoc命令
`protoc -I=$SRC_DIR --go_out=$DST_DIR $SRC_DIR/addressbook.proto`

|实验|简介|
|---|---|
|lab001|运行demo|
|lab002|测试package的自动生成|
|lab003|测试message中的field名称大小写、下划线的影响|
|lab003|实验repeated|
|lab004|实验map|
|lab005|实验嵌套|