package main

import (
	"github.com/tealeg/xlsx"
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
		log.Println("AddSheet error:", err)
	}
	row = sheet.AddRow()
	cell := row.AddCell()
	cell.Value = "I am a cell"
	err = file.Save("MyXlsxFile.xlsx")
	if err != nil {
		log.Println("Save error:", err)
	} else {
		log.Println("Save success")
	}
}
