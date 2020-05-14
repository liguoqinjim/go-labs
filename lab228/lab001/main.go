package main

import (
	"flag"
	"fmt"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/dysmsapi"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
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

	request := dysmsapi.CreateSendSmsRequest()
	request.Scheme = "https"

	request.PhoneNumbers = phone

	response, err := client.SendSms(request)
	if err != nil {
		fmt.Print(err.Error())
	}
	fmt.Printf("response is %#v\n", response)
}
