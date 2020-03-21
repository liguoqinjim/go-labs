package main

import (
	"github.com/willf/bloom"
	"log"
)

const (
	n = 2000000 //多少个元素
	p = 0.001   //false positive 可以理解为存储这么多个元素的情况下的错误率
)

func main() {
	//预估参数 需要多少个bit位，和多少个hash函数
	m, k := bloom.EstimateParameters(n, p)
	log.Printf("m=%d,k=%d", m, k)

	//使用m和k创建过滤器
	filter := bloom.New(m, k)

	//预估过滤器在n个元素下的的false positive
	log.Println(filter.EstimateFalsePositiveRate(n))

	filter.Add([]byte("Love"))
	if filter.Test([]byte("Love")) {
		log.Println("Love already has")
	}

	if filter.TestString("Hello") {
		log.Println("Hello already has")
	} else {
		log.Println("Hello not has")
	}
}
