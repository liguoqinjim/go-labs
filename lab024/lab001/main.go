package main

import (
	"fmt"
	"github.com/tealeg/xlsx"
	"log"
)

func main() {
	excelFileName := "testFile.xlsx"
	xlFile, err := xlsx.OpenFile(excelFileName)
	if err != nil {
		log.Println("OpenFile error:", err)
	}

	for _, sheet := range xlFile.Sheets {
		for _, row := range sheet.Rows {
			for _, cell := range row.Cells {
				text := cell.String()
				fmt.Printf("%s\t", text)
			}
			fmt.Println()
		}
	}
}
