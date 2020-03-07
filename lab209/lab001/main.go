package main

import (
	"fmt"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/spf13/viper"
	"os"
	"strings"
)

var (
	endpoint        string
	accessKeyId     string
	accessKeySecret string
	bucketName      string
)

func init() {
	readConfig()
}

func main() {
	client, err := oss.New(endpoint, accessKeyId, accessKeySecret)
	if err != nil {
		panic(err)
	}

	//list bucket
	lsRes, err := client.ListBuckets()
	if err != nil {
		panic(err)
	}
	for _, bucket := range lsRes.Buckets {
		fmt.Println("Buckets:", bucket.Name)

		if strings.Contains(bucket.Name, "01") {
			bucketName = bucket.Name
		}
	}

	//put object
	bucket, err := client.Bucket(bucketName)
	if err != nil {
		panic(err)
	}
	//名字一样的话，会被覆盖
	if err := bucket.PutObjectFromFile("test001", "../data/1.png"); err != nil {
		panic(err)
	}
}

func readConfig() {
	v := viper.New()
	v.SetConfigName("config")
	v.AddConfigPath("..")
	v.SetConfigType("yaml")

	if err := v.ReadInConfig(); err != nil {
		fmt.Printf("err:%s\n", err)
		os.Exit(1)
	}

	endpoint = v.GetString("Endpoint")
	accessKeyId = v.GetString("AccessKeyId")
	accessKeySecret = v.GetString("AccessKeySecret")
}
