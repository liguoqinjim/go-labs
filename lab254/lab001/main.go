package main

import (
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	asr "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/asr/v20190614"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
)

func main() {
	dataWav, err := ioutil.ReadFile("/Users/li/Workspace/py3-labs/lab003/data/16k.wav")
	if err != nil {
		log.Fatalf("readFile error:%v", err)
	}
	data := base64.StdEncoding.EncodeToString(dataWav)

	credential := common.NewCredential(
		os.Getenv("secretId"),
		os.Getenv("secretKey"),
	)
	cpf := profile.NewClientProfile()
	cpf.HttpProfile.Endpoint = "asr.tencentcloudapi.com"
	client, _ := asr.NewClient(credential, "", cpf)

	request := asr.NewSentenceRecognitionRequest()

	request.ProjectId = common.Uint64Ptr(0)
	request.SubServiceType = common.Uint64Ptr(2)
	request.EngSerViceType = common.StringPtr("16k_zh")
	request.SourceType = common.Uint64Ptr(1)
	request.VoiceFormat = common.StringPtr("wav")
	request.UsrAudioKey = common.StringPtr("123")
	request.Data = common.StringPtr(data)

	response, err := client.SentenceRecognition(request)
	if _, ok := err.(*errors.TencentCloudSDKError); ok {
		fmt.Printf("An API error has returned: %s", err)
		return
	}
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s", response.ToJsonString())
}
