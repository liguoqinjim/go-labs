package main

import (
	"errors"
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
	//strs := []string{"你好1", "你好2", "你好3", "你好4", "你好5", "你好6"}
	//SetRowValue(xlsx, 1, 'Y', strs)

	nextCol, _ := GetNextColumn("AZ")
	log.Println("nextCol=", nextCol)

	xlsx.SaveAs("test.xlsx")
	log.Println("Save success")
}

//设置一行数据
func SetRowValue(xlsx *excelize.File, row, firstColumn string, content []string) error {
	if len(content) == 0 {
		return errors.New("content length == 0")
	}

	if firstColumn == "ZZ" {
		if len(content) > 1 {
			return errors.New("when set ZZ column value ,content length mush be 1")
		}
	}

	for n, v := range content {
		if n == 0 {
			xlsx.SetCellValue()
		}

		xlsx.SetCellValue("Sheet1", fmt.Sprintf("%c%d", firstColumn+n, row), v)
	}
}

//这里就判断到AA1这种情况，三位的情况不考虑
func GetNextColumn(column string) (string, error) {
	if len(column) > 2 {
		return "", errors.New("column.length max is 2")
	}

	if len(column) == 1 {
		if column == "Z" {
			return "AA", nil
		}

		r := rune(column[0])

		return fmt.Sprintf("%c", r+1), nil
	} else {
		if column == "ZZ" {
			return "", errors.New("max column is ZZ")
		}

		r1, r2 := rune(column[0]), rune(column[1])
		if r2 == 'Z' {
			r1 += 1
			r2 = 'A'
		} else {
			r2 += 1
		}

		return fmt.Sprintf("%c%c", r1, r2), nil
	}
}
