### 简单实验1

#### 注意点
##### 客户端连接生成
客户端连接生成的时候就会自动发送DH64加密的信息给服务器。

##### io.Copy
```
var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatalf("accept failed: %s", err.Error())
			return
		}

		io.Copy(conn, conn)
		conn.Close()
		log.Println("copy exit")
		wg.Done()
	}()
```
这里的`io.Copy`要注意一下

#### 运行截图
![Imgur](http://i.imgur.com/BvR0xVa.png)

#### 参考资料
https://github.com/funny/snet