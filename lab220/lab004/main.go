package main

import (
	"errors"
	"fmt"
	"github.com/360EntSecGroup-Skylar/excelize"
	"log"
)

func main() {
	xlsx := excelize.NewFile()

	//设置一行数据
	strs := []interface{}{"你好1", "你好2", "你好3", "你好4", "你好5", "你好6"}
	//从第三行,C列开始赋值
	if err := SetRowValue(xlsx, "Sheet1", 3, "C", strs); err != nil {
		log.Fatalf("SetRowValue error:%v", err)
	}

	if err := xlsx.SaveAs("test.xlsx"); err != nil {
		log.Fatalf("xlsx save error:%v", err)
	}
}

//设置一行数据
func SetRowValue(xlsx *excelize.File, sheet string, row int, firstColumn string, content []interface{}) error {
	if len(content) == 0 {
		return errors.New("content length == 0")
	}

	if firstColumn == "ZZ" {
		if len(content) > 1 {
			return errors.New("when set ZZ column value ,content length mush be 1")
		}
	}

	nowColumn := firstColumn
	var err error
	for n, v := range content {
		log.Println("n=", n)
		if n == 0 {
			xlsx.SetCellValue(sheet, fmt.Sprintf("%s%d", firstColumn, row), v)
		} else {
			nowColumn, err = GetNextColumn(nowColumn)
			log.Println("nowColumn=", nowColumn)
			if err != nil {
				log.Println("GetNextColumn error:", err)
				return err
			}
			xlsx.SetCellValue(sheet, fmt.Sprintf("%s%d", nowColumn, row), v)
		}
	}
	return nil
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
