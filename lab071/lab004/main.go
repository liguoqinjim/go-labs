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

	t, err := template.ParseGlob("*.tmpl")
	if err != nil {
		log.Printf("template.New error:%v", err)
	}

	//注：我们在上面是直接Parse方法的，之前没有调用过New方法
	err = t.Execute(os.Stdout, tds)
	//err = t.ExecuteTemplate(os.Stdout, "todo.tmpl", tds)
	if err != nil {
		log.Printf("execute error:%v", err)
	}
}
