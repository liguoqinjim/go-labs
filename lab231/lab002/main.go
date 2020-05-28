package main

import (
	"bytes"
	"encoding/base64"
	"flag"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
	"log"
	"regexp"
)

var (
	ak     string
	sk     string
	bucket string
)

func init() {
	pflag.StringVarP(&ak, "ak", "a", "", "ak")
	pflag.StringVarP(&sk, "sk", "s", "", "sk")
	pflag.StringVarP(&bucket, "bucket", "b", "", "bucket")

	pflag.CommandLine.AddGoFlagSet(flag.CommandLine)
	pflag.Parse()
	viper.BindPFlags(pflag.CommandLine)

	if ak == "" || sk == "" {
		log.Fatalf("need ak and sk")
	}
}

func main() {
	creds := credentials.NewStaticCredentials(ak, sk, "")
	_, err := creds.Get()
	if err != nil {
		log.Fatalf("creds.Get error:%v", err)
	}

	config := &aws.Config{
		Region:      aws.String("cn"),
		Endpoint:    aws.String("oos-cn.ctyunapi.cn"),
		DisableSSL:  aws.Bool(false),
		Credentials: creds,
	}

	//上传base64数据
	pre := "data:image/jpg;base64,"
	b := `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`
	b = pre + b

	log.Println(b)
	re := regexp.MustCompile(`data:image\/\w+;base64,`)
	b = re.ReplaceAllString(b, "")

	data, err := base64.StdEncoding.DecodeString(b)
	if err != nil {
		log.Fatalf("base64 decode error:%v", err)
	}

	uploader := s3manager.NewUploader(session.New(config))
	_, err = uploader.Upload(&s3manager.UploadInput{
		Bucket:      aws.String(bucket),
		Key:         aws.String("test001/3.jpg"),
		ContentType: aws.String("image/jpeg"),
		//ContentType: aws.String("application/octet-stream"),
		Body: bytes.NewBuffer(data),
	})
	if err != nil {
		log.Fatalf("NewUploader error:%v", err)
	} else {
		log.Println("success upload")
	}
}
