package main

import (
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
var sprite *pixel.Sprite

func init() {
	chanFrame = make(chan []byte, 1)

	pic, err := loadPicture("gongren.png")
	if err != nil {
		panic(err)
	}

	for i := 0; i < len(Armys); i++ {
		s := pixel.NewSprite(pic, pic.Bounds())
		s.SetMatrix(pixel.IM.Moved(pixel.V(0, 0)))
		Armys[i] = s
	}

	sprite = pixel.NewSprite(pic, pic.Bounds())
}

func run() {
	cfg := pixelgl.WindowConfig{
		Title:  "战斗模拟",
		Bounds: pixel.R(0, 0, 1024, 768),
		VSync:  true,
	}
	win, err := pixelgl.NewWindow(cfg)
	if err != nil {
		panic(err)
	}
	win.Clear(colornames.Skyblue)

	//先画一个
	sprite.SetMatrix(pixel.IM.Moved(pixel.V(100, 100)))
	sprite.Draw(win)

	for !win.Closed() {

		win.Update()
	}
}

func drawArmy(win *pixelgl.Window) {
	for _, v := range Armys {
		v.Draw(win)
		break
	}
}

func convertPos(posx, posy int) (float64, float64) {
	return float64(posx / 8), float64(posy / 8)
}

func main() {
	pixelgl.Run(run)
}
