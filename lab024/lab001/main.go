package main

import (
	"github.com/tealeg/xlsx/v2"
	"log"
)

func main() {
	excelFileName := "test.xlsx"
	xlFile, err := xlsx.OpenFile(excelFileName)
	if err != nil {
		log.Fatalf("xlsx.OpenFile error:%v", err)
	}

	for _, sheet := range xlFile.Sheets {
		for _, row := range sheet.Rows {
			for _, cell := range row.Cells {
				text := cell.String()
				log.Printf("%s", text)
			}
		}
	}
}
