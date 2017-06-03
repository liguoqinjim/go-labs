package main

import (
	_ "github.com/murlokswarm/mac"
)

type Hello struct {
	Greeting string
}

func (h *Hello) Render() string {
	return `
<div class="WindowLayout">
    <div class="HelloBox">
        <h1>
            Hello,
            <span>{{if .Greeting}}{{html .Greeting}}{{else}}World{{end}}</span>
        </h1>
        <input type="text" placeholder="What is your name?" onchange="OnInputChange" />
    </div>
</div>
    `
}

func (h *Hello) OnInputChange(arg app.ChangeArg) {
	h.Greeting = arg.Value
	app.Render(h)
}

func init() {
	app.RegisterComponent(&Hello{})
}

func main() {
	app.OnLaunch = func() {
		win := app.NewWindow(app.Window{
			Title:          "Hello World",
			Width:          1280,
			Height:         720,
			TitlebarHidden: true,
		})

		hello := &Hello{}
		win.Mount(hello)
	}

	app.Run()
}
