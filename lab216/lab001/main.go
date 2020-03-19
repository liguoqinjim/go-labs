package main

import (
	"fmt"
	"github.com/pkumza/numcn"
)

func main() {
	chNum := "负十七亿零五十三万七千零一十六"
	num, _ := numcn.DecodeToInt64(chNum)
	fmt.Println(num) // -1700537016
	chNumAgain := numcn.EncodeFromInt64(num)
	fmt.Println(chNumAgain) // 负十七亿零五十三万七千零一十六

	chFloatNum := "负零点零七三零六"
	fNum, _ := numcn.DecodeToFloat64(chFloatNum)
	fmt.Printf("%f\n", fNum) // -0.073060
	chFloatNumAgain := numcn.EncodeFromFloat64(fNum)
	fmt.Println(chFloatNumAgain) // 负零点零七三零六

	fmt.Println(numcn.EncodeFromInt64(9))
	fmt.Println(numcn.EncodeFromInt64(11))
	fmt.Println(numcn.EncodeFromInt64(19))
	fmt.Println(numcn.EncodeFromInt64(20))
	fmt.Println(numcn.EncodeFromInt64(21))
	fmt.Println(numcn.EncodeFromInt64(100))
	fmt.Println(numcn.EncodeFromInt64(105))
	fmt.Println(numcn.EncodeFromInt64(3303))
	fmt.Println(numcn.EncodeFromInt64(3033))
	fmt.Println(numcn.EncodeFromInt64(3003))
	fmt.Println(numcn.EncodeFromInt64(13003))
	fmt.Println(numcn.EncodeFromInt64(10003))
	fmt.Println(numcn.EncodeFromInt64(30303))
}
