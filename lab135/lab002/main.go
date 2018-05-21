package main

import (
	"fmt"
	"net"
	"net/http"
	"reflect"

	"github.com/bouk/monkey"
	"log"
)

func main() {
	//go1.10中的Get方法已经不再调用Dial了，所以无效
	var d *net.Dialer
	monkey.PatchInstanceMethod(reflect.TypeOf(d), "Dial", func(_ *net.Dialer, _, _ string) (net.Conn, error) {
		return nil, fmt.Errorf("no dialing allowed")
	})
	_, err := http.Get("http://baidu.com")
	log.Println(err) // Get nil

	//patch DefaultClient的Get方法
	monkey.PatchInstanceMethod(reflect.TypeOf(http.DefaultClient), "Get", func(_ *http.Client, url string) (*http.Response, error) {
		return nil, fmt.Errorf("no dialing allowed patched2")
	})
	_, err = http.Get("http://baidu.com")
	log.Println(err) // Get http://baidu.com: no dialing allowed2
}
