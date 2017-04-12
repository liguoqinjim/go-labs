package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	_ "net/http/pprof"
	"sync"
	"time"
)

func counter() {
	list := []int{1}
	c := 1
	for i := 0; i < 10000000; i++ {
		httpGet()
		c = i + 1 + 2 + 3 + 4 - 5
		list = append(list, c)
	}
	fmt.Println(c)
	fmt.Println(list[0])
}

func work(wg *sync.WaitGroup) {
	for {
		counter()
		time.Sleep(1 * time.Second)
	}
	wg.Done()
}

func httpGet() int {
	queue := []string{"start..."}
	resp, err := http.Get("http://www.163.com")
	if err != nil {
		// handle error
	}

	//defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		// handle error
	}
	queue = append(queue, string(body))
	return len(queue)
}

func main() {
	flag.Parse()

	//这里实现了远程获取pprof数据的接口
	go func() {
		log.Println(http.ListenAndServe("localhost:7777", nil))
	}()

	var wg sync.WaitGroup
	wg.Add(10)
	for i := 0; i < 100; i++ {
		go work(&wg)
	}

	wg.Wait()
	time.Sleep(3 * time.Second)
}
