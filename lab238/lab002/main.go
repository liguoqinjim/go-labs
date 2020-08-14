package main

import (
	"flag"
	"fmt"
	"github.com/iGoogle-ink/gopay"
	"github.com/iGoogle-ink/gopay/wechat"
	"github.com/iGoogle-ink/gotil"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
	"log"
	"strconv"
	"time"
)

var (
	appId  string
	mchId  string
	apiKey string

	ip     string
	isProd = false

	client *wechat.Client
)

func init() {
	pflag.StringVarP(&appId, "appId", "a", "", "appId")
	pflag.StringVarP(&mchId, "mchId", "m", "", "mchId")
	pflag.StringVarP(&apiKey, "apiKey", "k", "", "apiKey")
	pflag.StringVarP(&ip, "ip", "i", "", "ip")

	pflag.CommandLine.AddGoFlagSet(flag.CommandLine)
	pflag.Parse()
	viper.BindPFlags(pflag.CommandLine)

	// 初始化微信客户端
	//    appId：应用ID
	//    mchId：商户ID
	//    apiKey：API秘钥值
	//    isProd：是否是正式环境
	client = wechat.NewClient(appId, mchId, apiKey, isProd)
	// 设置国家：不设置默认 中国国内
	//    gopay.China：中国国内
	//    gopay.China2：中国国内备用
	//    gopay.SoutheastAsia：东南亚
	//    gopay.Other：其他国家
	client.SetCountry(wechat.China)
	//if err := client.AddCertFilePath("../data/wechat/apiclient_cert.pem", "../data/wechat/apiclient_key.pem", "../data/wechat/apiclient_cert.p12"); err != nil {
	//	log.Fatalf("addCertFilePath error:%v", err)
	//}

}

func main() {
	sandBox()

	//orderId := "20070108115936083831"
	//order(orderId)
	//
	//log.Println("-----------------")
	//queryOrder(orderId)

	//沙盒支付
	//orderId := "20070471715938319077"
	//sandBoxPay(orderId)

	//queryOrder(orderId)
}

//统一下单
func order(orderId string) {
	// 初始化 BodyMap
	bm := make(gopay.BodyMap)
	bm.Set("nonce_str", gotil.GetRandomString(32))
	bm.Set("body", "小程序测试支付")
	bm.Set("out_trade_no", orderId)
	bm.Set("total_fee", 1)
	bm.Set("spbill_create_ip", ip)
	bm.Set("notify_url", "http://www.weituandui.co")
	bm.Set("trade_type", wechat.TradeType_H5)
	bm.Set("device_info", "WEB")
	bm.Set("sign_type", wechat.SignType_MD5)

	// 嵌套json格式数据（例如：H5支付的 scene_info 参数）
	h5Info := make(map[string]string)
	h5Info["type"] = "Wap"
	h5Info["wap_url"] = "http://www.weituandui.co"
	h5Info["wap_name"] = "H5测试支付"

	sceneInfo := make(map[string]map[string]string)
	sceneInfo["h5_info"] = h5Info

	bm.Set("scene_info", sceneInfo)

	var sign string
	if isProd {
		// 参数 sign ，可单独生成赋值到BodyMap中；也可不传sign参数，client内部会自动获取
		// 如需单独赋值 sign 参数，需通过下面方法，最后获取sign值并在最后赋值此参数
		sign = wechat.GetParamSign(appId, mchId, apiKey, bm)
	} else {
		var err error
		sign, err = wechat.GetSanBoxParamSign(appId, mchId, apiKey, bm)
		if err != nil {
			log.Fatalf("GetSanBoxParamSign error:%v", err)
		}
	}
	bm.Set("sign", sign)

	resp, err := client.UnifiedOrder(bm)
	if err != nil {
		log.Printf("resp.err=%+v", resp)
		log.Fatalf("client.UnifiedOrder error:%v", err)
	}
	log.Printf("resp:%+v", resp)

	// ====微信内H5支付 paySign====
	timeStamp := strconv.FormatInt(time.Now().Unix(), 10)
	packages := "prepay_id=" + resp.PrepayId // 此处的 wxRsp.PrepayId ,统一下单成功后得到
	// 获取微信内H5支付 paySign
	//    appId：AppID
	//    nonceStr：随机字符串
	//    packages：统一下单成功后拼接得到的值
	//    signType：签名方式，务必与统一下单时用的签名方式一致
	//    timeStamp：时间
	//    apiKey：API秘钥值
	paySign := wechat.GetH5PaySign(appId, resp.NonceStr, packages, wechat.SignType_MD5, timeStamp, apiKey)
	log.Println("paySign=", paySign)
}

func queryOrder(orderId string) {
	bm := make(gopay.BodyMap)
	bm.Set("out_trade_no", orderId)
	bm.Set("nonce_str", gotil.GetRandomString(32))

	resp, resBM, err := client.QueryOrder(bm)
	if err != nil {
		log.Fatalf("queryOrder error:%v", err)
	} else {
		log.Printf("resp:%+v", resp)
		log.Printf("resBM:%+v", resBM)
	}
}

//沙盒-发起付款码支付请求
func sandBox() {
	bm := make(gopay.BodyMap)
	bm.Set("nonce_str", gotil.GetRandomString(32))
	bm.Set("body", "小程序测试支付")
	bm.Set("out_trade_no", fmt.Sprintf("20200701AO%d", time.Now().Unix()))
	bm.Set("total_fee", 1)
	bm.Set("spbill_create_ip", ip)
	bm.Set("notify_url", "http://www.weituandui.co")
	bm.Set("trade_type", wechat.TradeType_H5)
	bm.Set("device_info", "WEB")
	bm.Set("sign_type", wechat.SignType_MD5)
	bm.Set("auth_code", "120061098828009406")
	//bm.Set("openid", "o0Df70H2Q0fY8JXh1aFPIRyOBgu8")
	log.Println("orderId=", bm.Get("out_trade_no"))

	// 嵌套json格式数据（例如：H5支付的 scene_info 参数）
	//h5Info := make(map[string]string)
	//h5Info["type"] = "Wap"
	//h5Info["wap_url"] = "http://www.weituandui.co"
	//h5Info["wap_name"] = "H5测试支付"

	//sceneInfo := make(map[string]map[string]string)
	//sceneInfo["h5_info"] = h5Info

	//bm.Set("scene_info", sceneInfo)

	//// 参数 sign ，可单独生成赋值到BodyMap中；也可不传sign参数，client内部会自动获取
	//// 如需单独赋值 sign 参数，需通过下面方法，最后获取sign值并在最后赋值此参数
	////sign := wechat.GetParamSign(appId, mchId, apiKey, bm)
	//sign, err := wechat.GetSanBoxParamSign(appId, mchId, apiKey, bm)
	//if err != nil {
	//	log.Fatalf("GetSanBoxParamSign error:%v", err)
	//}
	//bm.Set("sign", sign)

	resp, err := client.UnifiedOrder(bm)
	if err != nil {
		log.Printf("resp.err=%+v", resp)
		log.Fatalf("client.UnifiedOrder error:%v", err)
	}
	log.Printf("resp:%+v", resp)

	// ====微信内H5支付 paySign====
	timeStamp := strconv.FormatInt(time.Now().Unix(), 10)
	packages := "prepay_id=" + resp.PrepayId // 此处的 wxRsp.PrepayId ,统一下单成功后得到
	// 获取微信内H5支付 paySign
	//    appId：AppID
	//    nonceStr：随机字符串
	//    packages：统一下单成功后拼接得到的值
	//    signType：签名方式，务必与统一下单时用的签名方式一致
	//    timeStamp：时间
	//    apiKey：API秘钥值
	paySign := wechat.GetH5PaySign(appId, resp.NonceStr, packages, wechat.SignType_MD5, timeStamp, apiKey)
	log.Println("paySign=", paySign)
}

func sandBoxPay(orderId string) {
	bm := make(gopay.BodyMap)
	bm.Set("nonce_str", gotil.GetRandomString(32))
	bm.Set("body", "沙盒测试支付")
	bm.Set("out_trade_no", orderId)
	bm.Set("total_fee", 1)
	bm.Set("spbill_create_ip", ip)
	bm.Set("auth_code", "120061098828009406")

	resp, err := client.Micropay(bm)
	if err != nil {
		log.Fatalf("client.MicroPay error:%v", err)
	}
	log.Printf("resp:%+v", resp)
}
