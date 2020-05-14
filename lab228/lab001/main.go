package main

import (
	"flag"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/dysmsapi"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
	"log"
)

var (
	accessKeyId  string
	accessSecret string
	phone        string
)

func init() {
	pflag.StringVarP(&accessKeyId, "key", "k", "", "accessKeyId")
	pflag.StringVarP(&accessSecret, "secret", "s", "", "accessSecret")
	pflag.StringVarP(&phone, "phone", "p", "", "phone")

	pflag.CommandLine.AddGoFlagSet(flag.CommandLine)
	pflag.Parse()
	viper.BindPFlags(pflag.CommandLine)
}

func main() {
	client, err := dysmsapi.NewClientWithAccessKey("cn-hangzhou", accessKeyId, accessSecret)
	if err != nil {
		log.Fatalf("new client error:%v", err)
	}

	request := dysmsapi.CreateSendSmsRequest()
	request.Scheme = "https"

	request.PhoneNumbers = phone
	request.SignName = "微团队"
	request.TemplateCode = "SMS_190270015"
	request.TemplateParam = "{\"code\":\"134920\"}"

	response, err := client.SendSms(request)
	if err != nil {
		log.Fatalf("client.SendSMS error:%v", err)
	}
	log.Printf("response is %#v\n", response)
}
