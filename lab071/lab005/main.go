package main

import (
	"html/template"
	"log"
	"os"
	"strconv"
)

func main() {
	t, err := template.ParseGlob("md.tmpl")
	if err != nil {
		log.Printf("template.New error:%v", err)
	}

	ids := make([][]string, 2)
	for n := range ids {
		ids[n] = make([]string, 6)

		for j := range ids[n] {
			ids[n][j] = strconv.Itoa((n - 1) * j)
		}
	}

	err = t.Execute(os.Stdout, ids)
	if err != nil {
		log.Printf("execute error:%v", err)
	}
}
