package main

import (
	"github.com/henrylee2cn/surfer"
	"io/ioutil"
	"log"
)

const (
	HR = "------------------------------------------------------------------"
)

func main() {
	//创建phantomJsSurfer
	surfer.SetPhantomJsFilePath("../../phantomjs.exe")
	mySurfer := surfer.NewPhantom("../../phantomjs.exe", "./tmp")

	//查看cookies
	log.Println("查看cookies" + HR)
	resp, err := mySurfer.Download(&surfer.Request{
		Url:          "http://httpbin.org/cookies",
		EnableCookie: true,
		DownloaderID: 1,
	})
	handleError(err)

	b, err := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	log.Println("body=", string(b))

	//设置cookies
	log.Println()
	log.Println("设置cookies" + HR)
	resp, err = mySurfer.Download(&surfer.Request{
		Url:          "http://httpbin.org/cookies/set?k2=v2&k1=v1",
		EnableCookie: true,
		DownloaderID: 1,
	})
	handleError(err)

	b, err = ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	log.Println("body=", string(b))
	log.Println("resp.Header=", resp.Header)

	//查看cookies是否保存
	log.Println()
	log.Println("查看cookies是否保存" + HR)
	resp, err = mySurfer.Download(&surfer.Request{
		Url:          "http://httpbin.org/cookies",
		EnableCookie: true,
		DownloaderID: 1,
	})
	handleError(err)

	b, err = ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	log.Println("body=", string(b))

	//查看用新的surfer，cookies是否还存在
	log.Println()
	log.Println("查看用的新surfer，cookies是否还存在")
	resp, err = surfer.Download(&surfer.Request{
		Url:          "http://httpbin.org/cookies",
		EnableCookie: true,
		DownloaderID: 1,
	})
	handleError(err)

	b, err = ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	log.Println("body=", string(b))
}

func handleError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
