package main

import (
	"fmt"
	"github.com/360EntSecGroup-Skylar/excelize"
	"log"
)

func main() {
	xlsx := excelize.NewFile()

	//得到列名
	log.Println("第1列:", excelize.ToAlphaString(0))
	log.Println("第26列:", excelize.ToAlphaString(25))
	log.Println("第27列:", excelize.ToAlphaString(26))

	//
	//第三行,第五列 批量设置数据
	strs := []interface{}{"你好1", "你好2", "你好3", "你好4", "你好5", "你好6"}
	xlsx.SetSheetRow("sheet1", fmt.Sprintf("%s%d", excelize.ToAlphaString(5-1), 3-1), &strs)

	if err := xlsx.SaveAs("test.xlsx"); err != nil {
		log.Fatalf("xlsx save error:%v", err)
	}
}
