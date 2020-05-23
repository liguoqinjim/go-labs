package main

import (
	"flag"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
	"log"
	"os"
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

	client := s3.New(session.New(config))

	//列出buckets
	result, err := client.ListBuckets(nil)
	if err != nil {
		log.Fatalf("ListBuckets error:%v", err)
	}

	for _, b := range result.Buckets {
		log.Printf("* %s created on %s",
			aws.StringValue(b.Name), aws.TimeValue(b.CreationDate))
	}

	//上传图片
	file, err := os.Open("../data/1.jpg")
	if err != nil {
		log.Fatalf("os.Open error:%v", err)
	}
	defer file.Close()

	uploader := s3manager.NewUploader(session.New(config))
	_, err = uploader.Upload(&s3manager.UploadInput{
		Bucket:      aws.String(bucket),
		Key:         aws.String("20200523/3.jpg"),
		ContentType: aws.String("image/jpeg"),
		Body:        file,
	})
	if err != nil {
		log.Fatalf("NewUploader error:%v", err)
	} else {
		log.Println("success upload")
	}
}
