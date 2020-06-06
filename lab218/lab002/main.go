package main

import (
	"flag"
	"github.com/silenceper/wechat"
	"github.com/silenceper/wechat/cache"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
	"log"
)

var (
	appId     string
	appSecret string
)

func init() {
	pflag.StringVarP(&appId, "appId", "i", "appId", "appId")
	pflag.StringVarP(&appSecret, "appSecret", "s", "appSecret", "appSecret")

	pflag.CommandLine.AddGoFlagSet(flag.CommandLine)
	pflag.Parse()
	viper.BindPFlags(pflag.CommandLine)
}

func main() {
	c := cache.NewRedis(&cache.RedisOpts{
		Host:     "127.0.0.1:6379",
		Database: 1,
	})

	config := &wechat.Config{
		AppID:     appId,
		AppSecret: appSecret,
		Cache:     c,
	}
	wc := wechat.NewWechat(config)

	wxa := wc.GetMiniProgram()
	userInfo, err := wxa.Decrypt("vwawOozN/k+qiTJI9wsExQ==", "SpZDbwd/OUWzGnG27JtgxAc0LMRrrSUiL309yOa71Q1uSD48e74PhxSegGbMDWSBPfqVEg+2XUlBtWx2qxGu8VkLdHpI3ehBnF14yuv2CgSu23QhbwU2Al13PE9JCq/Rs2xAY6HD3OV/5k1bp9D4m0HMxY/Yj6KRz2j5a3rKMleQ2wD+kiUaY9D4+RQEk8Qvo0EC3U6cYeIdRghZFtNiso6YruoJq9zIwZhsnW/J7x2733wvCwFYV/xQd8+zenIavsZP45kdfwQBNmC3j6Kp+rQw+yWTEYACB5j5/LO/J0k1Z1Gdmo7xyS2hCFdXGC/JvQH3cKLxveL32KdGDLTaMoIQrLzqLilIqLIfA4UlW1JxpMw2Duv5gJfOyr1LlR0WIROdiyDRw69HDa8awWwGFeKnCX6pHGFRVwJUEJ5vsTBZT1SsW6eNn8UkIWOn/NJSIwiE4KI7OZfYEFiEBtpTFf86YzANaYwFyNepsBSPrXMLdLgSZUcPOlnB15byAhVKUFsnqEKNwMvEMauzLa9YBaYIechK3+kW7X1F+FRTwgIrWAtXEwRHbjdfFzcG9n39", "JuF/y34Y8S2R8DzhDWtD0A==")
	if err != nil {
		log.Fatalf("decrypt error:%v", err)
	} else {
		log.Println(userInfo)
	}

}
