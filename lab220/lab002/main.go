package main

import (
	"github.com/360EntSecGroup-Skylar/excelize"
	"log"
)

func main() {
	//xlsx1
	xlsx1 := excelize.NewFile()
	xlsx1.SetCellValue("Sheet1", "A1", "编号")
	xlsx1.SetCellValue("Sheet1", "A2", "姓名")
	xlsx1.SetCellValue("Sheet1", "B2", "小明")

	if err := xlsx1.SaveAs("test1.xlsx"); err != nil {
		log.Fatalf("xlsx save error:%v", err)
	}

	//xlsx2
	xlsx2 := excelize.NewFile()
	xlsx2.NewSheet("第一页")
	if err := xlsx2.SaveAs("test2.xlsx"); err != nil {
		log.Fatalf("xlsx02 save error:%v", err)
	}
}
