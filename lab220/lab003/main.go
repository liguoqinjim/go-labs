package main

import (
	"github.com/360EntSecGroup-Skylar/excelize"
	"log"
)

func main() {
	categories := map[string]string{"A2": "Small", "A3": "Normal", "A4": "Large", "B1": "Apple", "C1": "Orange", "D1": "Pear"}
	values := map[string]int{"B2": 2, "C2": 3, "D2": 3, "B3": 5, "C3": 2, "D3": 4, "B4": 6, "C4": 7, "D4": 8}

	//创建xlsx
	xlsx1 := excelize.NewFile()
	for k, v := range categories {
		xlsx1.SetCellValue("Sheet1", k, v)
	}
	for k, v := range values {
		xlsx1.SetCellValue("Sheet1", k, v)
	}
	//chart是以json的形式，按照一定格式生成
	if err := xlsx1.AddChart("Sheet1", "E1", `{"type":"col3DClustered","series":[{"name":"Sheet1!$A$2","categories":"Sheet1!$B$1:$D$1","values":"Sheet1!$B$2:$D$2"},{"name":"Sheet1!$A$3","categories":"Sheet1!$B$1:$D$1","values":"Sheet1!$B$3:$D$3"},{"name":"Sheet1!$A$4","categories":"Sheet1!$B$1:$D$1","values":"Sheet1!$B$4:$D$4"}],"title":{"name":"Fruit 3D Clustered Column Chart"}}`); err != nil {
		log.Fatalf("addChart error:%v", err)
	}

	//save
	if err := xlsx1.SaveAs("test.xlsx"); err != nil {
		log.Fatalf("xlsx save error:%v", err)
	}
}
