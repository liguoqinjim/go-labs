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
	td := Todo{Name: "Test templates", Description: "Let's test a template to see the magic."}

	t, err := template.New("todos").Parse("You have a task named \"{{ .Name}}\" with description: \"{{ .Description}}\"\n")
	if err != nil {
		log.Printf("template.New error:%v", err)
	}
	err = t.Execute(os.Stdout, td)
	if err != nil {
		log.Printf("execute error:%v", err)
	}

	tdNew := Todo{Name: "Go", Description: "Contribute to any Go project."}
	err = t.Execute(os.Stdout, tdNew)
	if err != nil {
		log.Printf("execute error:%v", err)
	}
}
