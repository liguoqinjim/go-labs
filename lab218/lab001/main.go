package main

import (
	"flag"
	"github.com/silenceper/wechat"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
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
	config := &wechat.Config{
		AppID:     appId,
		AppSecret: appSecret,
	}
	wc := wechat.NewWechat(config)

	wxa := wc.GetMiniProgram()

	_ = wxa

	//if resp, err := wxa.GetWXACodeUnlimit(miniprogram.QRCoder{
	//	Page:      "",
	//	Path:      "",
	//	Scene:     "",
	//}); err != nil {
	//	log.Fatalf("wxa.CreateWXAQRCode error:%v", err)
	//} else {
	//	filename := "../data/1.png"
	//
	//}

	//// CreateWXAQRCode 获取小程序二维码，适用于需要的码数量较少的业务场景
	//// 文档地址： https://developers.weixin.qq.com/miniprogram/dev/api/createWXAQRCode.html
	//func (wxa *MiniProgram) CreateWXAQRCode(coderParams QRCoder) (response []byte, err error) {
	//	return wxa.fetchCode(createWXAQRCodeURL, coderParams)
	//}
	//
	//// GetWXACode 获取小程序码，适用于需要的码数量较少的业务场景
	//// 文档地址： https://developers.weixin.qq.com/miniprogram/dev/api/getWXACode.html
	//func (wxa *MiniProgram) GetWXACode(coderParams QRCoder) (response []byte, err error) {
	//	return wxa.fetchCode(getWXACodeURL, coderParams)
	//}
	//
	//// GetWXACodeUnlimit 获取小程序码，适用于需要的码数量极多的业务场景
	//// 文档地址： https://developers.weixin.qq.com/miniprogram/dev/api/getWXACodeUnlimit.html
	//func (wxa *MiniProgram) GetWXACodeUnlimit(coderParams QRCoder) (response []byte, err error) {
	//	return wxa.fetchCode(getWXACodeUnlimitURL, coderParams)
	//}

	//
	//wxa := wc.GetMiniProgram()
	//
	//wxa.
}
