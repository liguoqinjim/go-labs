package main

import (
	"fmt"
	"github.com/xuri/excelize"
	"log"
)

const (
	FIRST_COLUMN = 'A'
)

func main() {
	xlsx := excelize.NewFile()

	//设置一行数据
	strs := []string{"你好1", "你好2", "你好3", "你好4", "你好5", "你好6"}
	SetRowValue(xlsx, 1, strs)

	xlsx.SaveAs("test.xlsx")
	log.Println("Save success")
}

//设置一行数据
func SetRowValue(xlsx *excelize.File, row int, content []string) {
	for n, v := range content {
		xlsx.SetCellValue("Sheet1", fmt.Sprintf("%c%d", FIRST_COLUMN+n, row), v)
	}
}
