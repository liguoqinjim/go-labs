package main

import (
	"github.com/go-gota/gota/dataframe"
	"github.com/go-gota/gota/series"
	"log"
	"os"
)

func main() {
	//创建dataframe
	df := dataframe.New(
		series.New([]string{"b", "a"}, series.String, "COL.1"),
		series.New([]int{1, 2}, series.Int, "COL.2"),
		series.New([]float64{3.0, 4.0}, series.Float, "COL.3"),
	)

	//打印
	log.Println(df)

	//创建dataframe
	df = dataframe.LoadRecords(
		[][]string{
			[]string{"A", "B", "C", "D"},
			[]string{"a", "4", "5.1", "true"},
			[]string{"k", "5", "7.0", "true"},
			[]string{"k", "4", "6.0", "true"},
			[]string{"a", "2", "7.1", "false"},
		},
	)
	log.Println(df)

	//通过struct创建dataframe
	type User struct {
		Name     string
		Age      int
		Accuracy float64
		ignored  bool // ignored since unexported
	}
	users := []User{
		{"Aram", 17, 0.2, true},
		{"Juan", 18, 0.8, true},
		{"Ana", 22, 0.5, true},
	}
	df = dataframe.LoadStructs(users)
	log.Println(df)

	//通过map创建dataframe
	df = dataframe.LoadMaps(
		[]map[string]interface{}{
			map[string]interface{}{
				"A": "a",
				"B": 1,
				"C": true,
				"D": 0,
			},
			map[string]interface{}{
				"A": "b",
				"B": 2,
				"C": true,
				"D": 0.5,
			},
		},
	)
	log.Println(df)

	//dataframe保存到csv中
	f, err := os.Create("data.csv")
	if err != nil {
		log.Fatalf("create file error:%v", err)
	}
	defer f.Close()

	if err := df.WriteCSV(f); err != nil {
		log.Fatalf("write to csv error:%v", err)
	}
}
