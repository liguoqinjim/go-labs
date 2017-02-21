package main

import (
	"fmt"
	"github.com/henrylee2cn/pholcus/common/xlsx"
)

func main() {
	excelFileName := "test.xlsx"
	xlsxFile, err := xlsx.OpenFile(excelFileName)
	if err != nil {
		fmt.Println(err)
	}

	for sheetName, sheet := range xlsxFile.Sheet {
		fmt.Println("sheetName", sheetName)

		for rowNum, row := range sheet.Rows {
			fmt.Println("rowNum", rowNum)

			for cellNum, cell := range row.Cells {
				fmt.Println("cellNum", cellNum, "content", cell.Value)
			}

			if rowNum > 1 {
				break
			}
		}
	}
}
