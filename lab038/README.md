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
|lab004|测试repeated|
|lab005|测试map|
|lab006|测试嵌套|
|lab007|测试Enum|
|lab008|测试any|
|lab009|测试singular|
|lab010|测试pb的二进制的字节大小|