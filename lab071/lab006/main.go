package main

import (
	"html/template"
	"log"
	"os"
	"fmt"
)

func main() {
	t, err := template.ParseGlob("md.tmpl")
	if err != nil {
		log.Printf("template.New error:%v", err)
	}

	vals := make(map[int]string)
	for i := 1; i <= 3; i++ {
		vals[i] = fmt.Sprintf("你好:%d", i)
	}

	err = t.Execute(os.Stdout, vals)
	if err != nil {
		log.Printf("execute error:%v", err)
	}
}
