package main

import (
	"log"
	"os"
	"text/template"
)

type Todo struct {
	Name        string
	Description string
}

func main() {
	defer func() {
		if err := recover(); err != nil {
			log.Printf("error:%v", err)
		}
	}()

	td := Todo{Name: "Test templates", Description: "Let's test a template to see the magic."}

	//多了一个{
	t := template.Must(template.New("todos").Parse("You have task named \"{{ {.Name}}\" with description: \"{{ .Description}}\""))

	err := t.Execute(os.Stdout, td)
	if err != nil {
		log.Printf("execute error:%v", err)
	}
}
