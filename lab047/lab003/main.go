package main

import (
	"bufio"
	"fmt"
	_ "image/png"

	"image"
	"os"

	"golang.org/x/image/colornames"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
)

func loadPicture(path string) (pixel.Picture, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	img, _, err := image.Decode(file)
	if err != nil {
		return nil, err
	}
	return pixel.PictureDataFromImage(img), nil
}

//初始化的参数
var Armys [8]*pixel.Sprite
var chanFrame chan []byte

func init() {
	chanFrame = make(chan []byte, 1)

	pic, err := loadPicture("gongren.png")
	if err != nil {
		panic(err)
	}

	for i := 0; i < len(Armys); i++ {
		s := pixel.NewSprite(pic, pic.Bounds())
		Armys[i] = s
	}

	//读取json
}

func run() {
	cfg := pixelgl.WindowConfig{
		Title:  "战斗模拟",
		Bounds: pixel.R(0, 0, 1200, 800),
		VSync:  true,
	}
	win, err := pixelgl.NewWindow(cfg)
	if err != nil {
		panic(err)
	}
	win.Clear(colornames.Skyblue)
	for !win.Closed() {
		select {
		case f := <-chanFrame:
			fmt.Println(f)
		default:
		}

		win.Update()
	}
}

func read() {
	reader := bufio.NewReader(os.Stdin)
	for {
		strBytes, _, err := reader.ReadLine()
		// fmt.Println(strBytes, hasMore, err)
		if err == nil {
			chanFrame <- strBytes
		}
	}
}

func main() {
	go read()
	pixelgl.Run(run)
}
