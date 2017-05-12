package main

import (
	"image"
	"os"
	"time"

	"golang.org/x/image/colornames"

	_ "image/png"

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

func run() {
	cfg := pixelgl.WindowConfig{
		Title:  "核心战斗",
		Bounds: pixel.R(0, 0, 1024, 768),
		VSync:  true,
	}
	win, err := pixelgl.NewWindow(cfg)
	if err != nil {
		panic(err)
	}

	pic, err := loadPicture("gongren.png")
	if err != nil {
		panic(err)
	}
	win.Clear(colornames.Skyblue)
	sprite := pixel.NewSprite(pic, pic.Bounds())

	sprite.SetMatrix(pixel.IM.Moved(win.Bounds().Center()))
	sprite.Draw(win)

	angle := 0.0
	last := time.Now()
	for !win.Closed() {
		dt := time.Since(last).Seconds()
		last = time.Now()

		angle += 3 * dt

		win.Clear(colornames.Firebrick)

		mat := pixel.IM
		mat = mat.Rotated(0, angle)
		mat = mat.Moved(win.Bounds().Center())
		sprite.SetMatrix(mat)
		sprite.Draw(win)

		win.Update()
	}
}

func main() {
	pixelgl.Run(run)
}
