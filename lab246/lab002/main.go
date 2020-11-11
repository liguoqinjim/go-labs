package main

import (
	"github.com/shopspring/decimal"
	"log"
	"strconv"
	"strings"
)

func main() {
	sum := 100 //单位:分 (100是1元)

	shareStrs := strings.Split("5000,3000,2000", ",")
	shares := make([]int, len(shareStrs))
	for n, v := range shareStrs {
		share, err := strconv.Atoi(v)
		if err != nil {
			log.Fatalf("strconv.Atoi error:%v", err)
		}
		shares[n] = share
	}

	for n, share := range shares {
		log.Println("n=", n)

		num := decimal.NewFromInt(int64(sum)).Div(decimal.NewFromInt(10000)).Mul(decimal.NewFromInt(int64(share)))
		log.Println("num=", num)
	}
}
