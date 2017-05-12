package main

import (
	"bufio"
	"fmt"
	_ "image/png"
	"lab047/lab003/data"

	"image"
	"os"

	"golang.org/x/image/colornames"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"github.com/tidwall/gjson"
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

	data.LoadData()
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

	initBattleState()
	// sprite.Draw(win)

	for !win.Closed() {
		select {
		case f := <-chanFrame:
			fmt.Println(f)
		default:
		}

		drawArmy(win)

		win.Update()
	}
}

func drawArmy(win *pixelgl.Window) {
	for _, v := range Armys {
		v.Draw(win)
	}
}

func initBattleState() {
	fmt.Println("战斗初始配置")

	army1 := gjson.Get(data.BattleData, "Back.Params.ArmyGroup1Init.Armys")
	for n, v := range army1.Array() {
		posx := gjson.Get(v.String(), "PosX").Int()
		posy := gjson.Get(v.String(), "PosY").Int()

		x, y := convertPos(int(posx), int(posy))
		Armys[n].SetMatrix(pixel.IM.Moved(pixel.V(x, y)))
	}

	army2 := gjson.Get(data.BattleData, "Back.Params.ArmyGroup2Init.Armys")
	for n, v := range army2.Array() {
		posx := gjson.Get(v.String(), "PosX").Int()
		posy := gjson.Get(v.String(), "PosY").Int()

		x, y := convertPos(int(posx), int(posy))
		Armys[n+4].SetMatrix(pixel.IM.Moved(pixel.V(x, y)))
	}
}

func convertPos(posx, posy int) (float64, float64) {
	return float64(posx / 8), float64(posy / 8)
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
