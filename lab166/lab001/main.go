package main

import (
	"github.com/bculberson/bloom"
	"log"
)

const (
	n  = 100000
	fp = 0.001
)

func main() {
	//计算m和k值
	m, k := bloom.EstimateParameters(n, fp)

	//新建内存
	bitSet := bloom.NewBitSet(m)

	//新建过滤器
	filter := bloom.New(m, k, bitSet)

	filter.Add([]byte("hello"))
	filter.Add([]byte("world"))

	//判断是否存在
	if ok, err := filter.Exists([]byte("hello")); err != nil {
		log.Fatalf("filter.Exists error:%v", err)
	} else {
		if ok {
			log.Println("hello is in the filter")
		} else {
			log.Println("hello not in the filter")
		}
	}
}
