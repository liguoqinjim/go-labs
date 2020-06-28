package main

import (
	"github.com/kataras/iris/v12"
)

func main() {
	app := iris.New()

	app.Get("/", func(ctx iris.Context) {})

	app.SetExecutionRules(iris.ExecutionRules{
		// Begin: ...
		// Main:  ...
		Done: iris.ExecutionOptions{Force: true},
	})

	app.Use(before)
	app.Done(after)

	app.UseGlobal(before)
	app.DoneGlobal(after)
}

func before(ctx iris.Context) {
	// [...]
}

func after(ctx iris.Context) {
	// [...]
}
