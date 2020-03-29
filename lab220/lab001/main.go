package main

import (
	"github.com/360EntSecGroup-Skylar/excelize"
	"log"
)

func main() {
	filename := "test.xlsx"
	xlsx, err := excelize.OpenFile(filename)
	if err != nil {
		log.Fatalf("excelize.OpenFile error:%v", err)
	}

	cell := xlsx.GetCellValue("Sheet1", "B2")
	log.Println("B2.cell=", cell)

	rows := xlsx.GetRows("Sheet1")
	for _, row := range rows {
		for n, cell := range row {
			log.Printf("cell[%d]:%v", n, cell)
		}
	}
}
