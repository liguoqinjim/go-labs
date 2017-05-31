package main

import (
	"fmt"
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"github.com/faiface/pixel/text"
	"golang.org/x/image/colornames"
	"golang.org/x/image/font/basicfont"
)

func run() {
	cfg := pixelgl.WindowConfig{
		Title:  "Pixel Rocks!",
		Bounds: pixel.R(0, 0, 1024, 768),
		VSync:  true,
	}
	win, err := pixelgl.NewWindow(cfg)
	if err != nil {
		panic(err)
	}

	basicAtlas := text.NewAtlas(basicfont.Face7x13, text.ASCII)
	basicTxt := text.New(pixel.V(100, 500), basicAtlas)

	//fmt.Fprintln(basicTxt, "Hello, text!")
	//fmt.Fprintln(basicTxt, "I support multiple lines!")
	//fmt.Fprintf(basicTxt, "And I'm an %s, yay!", "io.Writer")

	basicTxt.Color = colornames.Red                   //修改字体颜色
	basicTxt.LineHeight = basicAtlas.LineHeight() * 3 //行间距
	fmt.Fprintln(basicTxt, "Hello, text!")
	basicTxt.LineHeight = basicAtlas.LineHeight() * 1.5 //行间距
	basicTxt.Color = colornames.Blue
	fmt.Fprintln(basicTxt, "I support multiple lines!")
	basicTxt.Color = colornames.Gray
	fmt.Fprintf(basicTxt, "And I'm an %s, yay!", "io.Writer")

	for !win.Closed() {
		win.Clear(colornames.Black)
		//basicTxt.Draw(win, pixel.IM)
		basicTxt.Draw(win, pixel.IM.Scaled(basicTxt.Orig, 4)) //放大字体

		win.Update()
	}
}

func main() {
	pixelgl.Run(run)
}
