package main

import (
	"encoding/base64"
	"flag"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
	"log"
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

	// Initialize a session in us-west-2 that the SDK will use to load
	// credentials from the shared credentials file ~/.aws/credentials.
	sess, _ := session.NewSession(&aws.Config{
		Region:      aws.String("cn"),
		Endpoint:    aws.String("oos-cn.ctyunapi.cn"),
		DisableSSL:  aws.Bool(false),
		Credentials: creds,
	})

	downloader := s3manager.NewDownloader(sess)

	buf := aws.NewWriteAtBuffer([]byte{})
	numBytes, err := downloader.Download(buf,
		&s3.GetObjectInput{
			Bucket: aws.String(bucket),
			Key:    aws.String("200527/CDF14FBt.xlsx"),
		})
	if err != nil {
		log.Fatalf("download error:%v", err)
	}

	log.Println("Downloaded", numBytes, "bytes")

	//转到base64
	b64 := base64.StdEncoding.EncodeToString(buf.Bytes())
	log.Println(b64)
}
