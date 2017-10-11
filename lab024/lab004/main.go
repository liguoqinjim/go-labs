package main

import (
	"github.com/xuri/excelize"
	"log"
)

func main() {
	//xlsx1
	xlsx1 := excelize.NewFile()

	xlsx1.SetCellValue("Sheet1", "A1", "编号")
	xlsx1.SetCellValue("Sheet1", "A2", "姓名")
	xlsx1.SetCellValue("Sheet1", "B2", "小明")

	err := xlsx1.SaveAs("test1.xlsx")
	if err != nil {
		log.Println("SaveAs test1.xlsx error:", err)
	} else {
		log.Println("SaveAs test1.xlsx success")
	}

	//xlsx2
	xlsx2 := excelize.NewFile()
	xlsx2.NewSheet(0, "第一页")
	err = xlsx2.SaveAs("test2.xlsx")
	if err != nil {
		log.Println("SaveAs test2.xlsx error:", err)
	} else {
		log.Println("SaveAs test2.xlsx success")
	}
}
