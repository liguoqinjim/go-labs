### 模仿示例2
模仿lab003中的自定义通讯协议，自己写了一份。我自己写的这个比lab003中的要简单一些。
包的前四位就写一个长度。第5位开始都是实际的数据。

#### 注意点
```
func IntToBytes(i int32) []byte { //这里用int32类型，是因为binary.Write里面是不对int类型操作的
	byteBuffer := bytes.NewBuffer([]byte{})
	binary.Write(byteBuffer, binary.BigEndian, i)

	return byteBuffer.Bytes()
}
```

#### 运行结果
##### client
![Imgur](http://i.imgur.com/0g6lrXq.png)

##### server
![Imgur](http://i.imgur.com/WyYhrfd.png)

#### 参考资料
1. http://blog.csdn.net/ahlxt123/article/details/47396509
2. [CPU的大端模式(big endian)和小端(little endian)模式](http://blog.csdn.net/xiajun07061225/article/details/7295421)