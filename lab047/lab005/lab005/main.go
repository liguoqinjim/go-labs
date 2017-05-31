package main

import (
	"fmt"
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"github.com/faiface/pixel/text"
	"github.com/golang/freetype/truetype"
	"golang.org/x/image/colornames"
	"golang.org/x/image/font"
	"io/ioutil"
	"os"
	"unicode"
)

func loadTTF(path string, size float64) (font.Face, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	data, err := ioutil.ReadAll(file)
	if err != nil {
		return nil, err
	}

	font, err := truetype.Parse(data)
	if err != nil {
		return nil, err
	}

	return truetype.NewFace(font, &truetype.Options{
		Size:              size,
		GlyphCacheEntries: 1,
	}), nil
}

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

	face, err := loadTTF("WRYH.ttf", 30)
	if err != nil {
		panic(err)
	}

	basicAtlas := text.NewAtlas(face, text.ASCII, text.RangeTable(unicode.Han)) //加上汉字的unicode
	basicTxt := text.New(pixel.V(100, 500), basicAtlas)

	basicTxt.Color = colornames.Gray                  //修改字体颜色
	basicTxt.LineHeight = basicAtlas.LineHeight() * 3 //行间距
	fmt.Fprintln(basicTxt, "Hello, text!你好")

	//测试中文
	fmt.Fprintln(basicTxt, "你好world")

	for !win.Closed() {
		win.Clear(colornames.Black)
		//basicTxt.Draw(win, pixel.IM)
		basicTxt.Draw(win, pixel.IM.Scaled(basicTxt.Orig, 1)) //放大字体

		win.Update()
	}
}

func main() {
	pixelgl.Run(run)
}
