package main

import (
	"github.com/tealeg/xlsx"
	"github.com/xuri/excelize"
)

func main() {
	//tealeg/xlxs
	xlsx1 := xlsx.NewFile()
	sheet, _ := xlsx1.AddSheet("Sheet1")
	row := sheet.AddRow()
	cell1 := row.AddCell()
	cell1.SetValue("你好")
	xlsx1.Save("文件1.xlsx")

	//xuri/excelize
	xlsx2 := excelize.NewFile()
	xlsx2.SetCellValue("Sheet1", "A1", "你好2")
	xlsx2.SaveAs("文件2.xlsx")
}
