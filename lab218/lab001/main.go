package main

import(
	"github.com/silenceper/wechat"
)

func main() {
	memCache := cache.NewMemcache("127.0.0.1:11211")
	config := &wechat.Config{
		AppID:     "xxx",
		AppSecret: "xxx",
		Cache:     memCache = cache.NewMemcache("127.0.0.1:11211"),
	}
	wc := wechat.NewWechat(config)

	wxa := wc.GetMiniProgram()
}
