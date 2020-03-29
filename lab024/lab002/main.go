package main

import (
	"github.com/tealeg/xlsx/v2"
	"log"
)

func main() {
	var file *xlsx.File
	var sheet *xlsx.Sheet
	var row *xlsx.Row
	var err error

	file = xlsx.NewFile()
	sheet, err = file.AddSheet("Sheet1")
	if err != nil {
		log.Fatalf("AddSheet error:%v", err)
	}

	row = sheet.AddRow()
	cell := row.AddCell()
	cell.Value = "I am a cell"

	//保存文件
	if err := file.Save("test.xlsx"); err != nil {
		log.Fatalf("file Save error:%v", err)
	}
}
