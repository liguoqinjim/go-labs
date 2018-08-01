package main

import (
	"html/template"
	"log"
	"os"
)

type entry struct {
	Name string
	Done bool
}

type Todo struct {
	User string
	List []entry
}

func main() {
	tds := Todo{User: "gopher", List: []entry{{"learn text/template", true}, {"learn html/template", false}}}

	paths := []string{
		"todo.tmpl",
	}

	t, err := template.New("html-tmpl").ParseFiles(paths...)
	if err != nil {
		log.Printf("template.New error:%v", err)
	}

	//注：第一种方法调用会报错
	//err = t.Execute(os.Stdout, tds)
	err = t.ExecuteTemplate(os.Stdout, "todo.tmpl", tds)
	if err != nil {
		log.Printf("execute error:%v", err)
	}
}
