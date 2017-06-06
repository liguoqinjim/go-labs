### 连接websocket

#### 注意点
##### 连接参数
```
conn, _, err := dialer.Dial("wss://io.sosobtc.com/socket.io/?EIO=3&transport=websocket",
		map[string][]string{"Origin": []string{"https://k.sosobtc.com"},
			"Cookie":                []string{"OID=aEj%252BuelTgv0RAuNv%252FFJPBfslpUnvZ26EiWsVM7TiIgNZ%252FaJQtLXiwjyAsqbnRKA%252BpQ7UkYv1rrO92kq8%252BZ4ifZQex9e7Sbgj7BVy3DtSflfIJd4koi1JTx61ElPwSY8x%7C8dad2860013668cf3e1c4aa6c4e19154; _ga=GA1.2.241055403.1496717869; _gid=GA1.2.1622829335.1496717869; _gat=1; theme=dark"},
			"User-Agent":            []string{"Mozilla/5.0 (Windows NT 10.0; WOW64; Trident/7.0; rv:11.0) like Gecko"},
			"Sec-WebSocket-Version": []string{"13"}})
```
因为这个websocket是我们抓包之后的链接，一开始连接的时候是放了所有抓包抓到的参数。
但是发现有不对的地方。`Sec-WebSocket-Key`这个参数就是不能用抓包的。
这个参数是客户端自动生成的，服务器收到这个参数之后会重新计算一个key返回，要是这个值和客户端计算的不一样，那么就认为是握手失败了。