package main

import (
	"flag"
	"github.com/iGoogle-ink/gopay"
	"github.com/iGoogle-ink/gopay/wechat"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

var (
	appId  string
	mchId  string
	apiKey string
)

func init() {
	pflag.StringVarP(&appId, "appId", "i", "appId", "appId")
	pflag.StringVarP(&mchId, "mchId", "m", "mchId", "mchId")
	pflag.StringVarP(&apiKey, "apiKey", "k", "apiKey", "apiKey")

	pflag.CommandLine.AddGoFlagSet(flag.CommandLine)
	pflag.Parse()
	viper.BindPFlags(pflag.CommandLine)
}

func main() {
	// 初始化微信客户端
	//    appId：应用ID
	//    mchId：商户ID
	//    apiKey：API秘钥值
	//    isProd：是否是正式环境
	client := wechat.NewClient(appId, mchId, apiKey, false)

	// 设置国家：不设置默认 中国国内
	//    gopay.China：中国国内
	//    gopay.China2：中国国内备用
	//    gopay.SoutheastAsia：东南亚
	//    gopay.Other：其他国家
	client.SetCountry(wechat.China)

	// 初始化 BodyMap
	bm := make(gopay.BodyMap)
	bm.Set("nonce_str", gopay.GetRandomString(32))
	bm.Set("body", "小程序测试支付")
	bm.Set("out_trade_no", number)
	bm.Set("total_fee", 1)
	bm.Set("spbill_create_ip", "127.0.0.1")
	bm.Set("notify_url", "http://www.gopay.ink")
	bm.Set("trade_type", wechat.TradeType_Mini)
	bm.Set("device_info", "WEB")
	bm.Set("sign_type", wechat.SignType_MD5)
	bm.Set("openid", "o0Df70H2Q0fY8JXh1aFPIRyOBgu8")

	// 嵌套json格式数据（例如：H5支付的 scene_info 参数）
	h5Info := make(map[string]string)
	h5Info["type"] = "Wap"
	h5Info["wap_url"] = "http://www.gopay.ink"
	h5Info["wap_name"] = "H5测试支付"

	sceneInfo := make(map[string]map[string]string)
	sceneInfo["h5_info"] = h5Info

	bm.Set("scene_info", sceneInfo)

	// 参数 sign ，可单独生成赋值到BodyMap中；也可不传sign参数，client内部会自动获取
	// 如需单独赋值 sign 参数，需通过下面方法，最后获取sign值并在最后赋值此参数
	sign := wechat.GetParamSign("wxdaa2ab9ef87b5497", mchId, apiKey, body)
	// sign, _ := wechat.GetSanBoxParamSign("wxdaa2ab9ef87b5497", mchId, apiKey, body)
	bm.Set("sign", sign)

	client.UnifiedOrder(bm)
}
