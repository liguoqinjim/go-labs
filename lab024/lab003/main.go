package main

import (
	"fmt"
	"github.com/xuri/excelize"
	"log"
)

func main() {
	xlsxFileName := "testFile.xlsx"
	xlsx, err := excelize.OpenFile(xlsxFileName)
	if err != nil {
		log.Println("OpenFile error:", err)
	}

	cell := xlsx.GetCellValue("Sheet1", "B2")
	fmt.Println("B2.cell=", cell)

	rows := xlsx.GetRows("Sheet1")
	for _, row := range rows {
		for _, colCell := range row {
			fmt.Print(colCell, "\t")
		}
		fmt.Println()
	}
}
