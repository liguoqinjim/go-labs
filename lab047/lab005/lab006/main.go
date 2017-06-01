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
		Title:  "lab006",
		Bounds: pixel.R(0, 0, 1024, 768),
		VSync:  true,
	}

	win, err := pixelgl.NewWindow(cfg)
	if err != nil {
		panic(err)
	}

	basicAtlas := text.NewAtlas(basicfont.Face7x13, text.ASCII)
	basicTxt := text.New(pixel.V(0, 700), basicAtlas)
	basicTxt.Color = colornames.Red
	fmt.Fprintln(basicTxt, "helloworld")

	basicTxt2 := text.New(pixel.V(100, 500), basicAtlas)
	fmt.Fprintln(basicTxt2, "helloworld1!!!!")

	for !win.Closed() {
		win.Clear(colornames.Black)
		basicTxt.Draw(win, pixel.IM.Scaled(basicTxt.Orig, 1))  //放大字体
		basicTxt2.Draw(win, pixel.IM.Scaled(basicTxt.Orig, 1)) //放大字体

		win.Update()
	}
}

func main() {
	pixelgl.Run(run)
}
