package main

import (
	"bufio"
	"fmt"
	_ "image/png"
	"lab047/lab006/data"
	"strconv"

	"image"
	"os"

	"golang.org/x/image/colornames"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"github.com/faiface/pixel/text"
	"github.com/golang/freetype/truetype"
	"github.com/tidwall/gjson"
	"golang.org/x/image/font"
	"io/ioutil"
	"unicode"
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

	face, err := loadTTF("WRYH.ttf", 30)
	if err != nil {
		panic(err)
	}

	initBattleState(win)
	initBattleInfo(win, face)

	for !win.Closed() {
		select {
		case f := <-chanFrame:
			frame, err := strconv.Atoi(string(f))
			if err != nil {
				panic(err)
			}
			win.Clear(colornames.Skyblue)
			setArmyStateFrame(win, frame)
		default:
		}

		win.Update()
	}
}

func initBattleState(win *pixelgl.Window) {
	fmt.Println("战斗初始配置")

	army1 := gjson.Get(data.BattleData, "Back.Params.ArmyGroup1Init.Armys")
	for n, v := range army1.Array() {
		posx := gjson.Get(v.String(), "PosX").Int()
		posy := gjson.Get(v.String(), "PosY").Int()

		x, y := convertPos(int(posx), int(posy))
		Armys[n].Draw(win, pixel.IM.Moved(pixel.V(x, y)))
	}

	army2 := gjson.Get(data.BattleData, "Back.Params.ArmyGroup2Init.Armys")
	for n, v := range army2.Array() {
		posx := gjson.Get(v.String(), "PosX").Int()
		posy := gjson.Get(v.String(), "PosY").Int()

		x, y := convertPos(int(posx), int(posy))
		Armys[n+4].Draw(win, pixel.IM.Moved(pixel.V(x, y)))
	}
}

func initBattleInfo(win *pixelgl.Window, font font.Face) {
	basicAtlas := text.NewAtlas(font, text.ASCII, text.RangeTable(unicode.Han))
	basicTxt := text.New(pixel.V(0, 0), basicAtlas)

	basicTxt.Color = colornames.Red
	fmt.Fprintln(basicTxt, "战报")

	basicTxt.Draw(win, pixel.IM.Scaled(basicTxt.Orig, 0.5))
}

func setArmyStateFrame(win *pixelgl.Window, frame int) {
	matchPath := fmt.Sprintf("Back.Params.BattleFrameDatas.#[Frame==\"%d\"]#", frame)

	value := gjson.Get(data.BattleData, matchPath)
	for _, v := range value.Array() {
		operator := gjson.Get(v.String(), "Operator").Int()
		fid := gjson.Get(v.String(), "ArmyFieldId").Int()
		posx := gjson.Get(v.String(), "Posx").Int()
		posy := gjson.Get(v.String(), "Posy").Int()
		x, y := convertPos(int(posx), int(posy))
		if operator == 1 { //移动
			Armys[fid-1].Draw(win, pixel.IM.Moved(pixel.V(x, y)))
		}
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
