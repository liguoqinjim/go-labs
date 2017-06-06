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
	"github.com/name5566/leaf/log"
	"github.com/tidwall/gjson"
	"golang.org/x/image/font"
	"golang.org/x/image/font/basicfont"
	"io/ioutil"
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
var Armys [8]*Army
var chanFrame chan []byte
var basicAtlas *text.Atlas
var battleInfoText *text.Text

type Army struct {
	id     int
	x, y   float64
	ox, oy int
	sp     *pixel.Sprite //显示图片
	txtId  *text.Text    //显示军队编号
}

func init() {
	chanFrame = make(chan []byte, 1)

	pic, err := loadPicture("gongren.png")
	if err != nil {
		panic(err)
	}
	basicAtlas = text.NewAtlas(basicfont.Face7x13, text.ASCII)
	battleInfoText = text.New(pixel.V(0, 0), basicAtlas)

	for i := 0; i < len(Armys); i++ {
		army := new(Army)
		army.sp = pixel.NewSprite(pic, pic.Bounds())
		army.id = i + 1
		army.txtId = text.New(pixel.V(army.x, army.y), basicAtlas)
		army.txtId.Color = colornames.Red
		Armys[i] = army
	}

	data.LoadData()
}

func run() {
	cfg := pixelgl.WindowConfig{
		Title:  "战斗模拟",
		Bounds: pixel.R(0, 0, 1920, 900),
		VSync:  true,
	}
	win, err := pixelgl.NewWindow(cfg)
	if err != nil {
		panic(err)
	}
	win.Clear(colornames.Skyblue)

	//face, err := loadTTF("WRYH.ttf", 30)
	//if err != nil {
	//	panic(err)
	//}

	initBattleState(win)
	//initBattleInfo(win, face)

	for !win.Closed() {
		select {
		case f := <-chanFrame:
			frame, err := strconv.Atoi(string(f))
			if err != nil {
				log.Error("输入不正确")
				break
			}
			win.Clear(colornames.Skyblue)
			setArmyStateFrame(win, frame)
			initBattleInfo(win)
		default:
		}

		win.Update()
	}
}

func initBattleState(win *pixelgl.Window) {
	fmt.Println("战斗初始配置")

	//army1 := gjson.Get(data.BattleData, "Back.Params.ArmyGroup1Init.Armys")
	army1 := gjson.Get(data.BattleData, "ArmyGroup1Init.Armys")
	for n, v := range army1.Array() {
		posx := gjson.Get(v.String(), "PosX").Int()
		posy := gjson.Get(v.String(), "PosY").Int()

		x, y := convertPos(int(posx), int(posy))
		Armys[n].x = x
		Armys[n].y = y
		Armys[n].ox = int(posx)
		Armys[n].oy = int(posy)
		Armys[n].id = n + 1
		Armys[n].sp.Draw(win, pixel.IM.Moved(pixel.V(x, y)))
		fmt.Fprintf(Armys[n].txtId, "%d", Armys[n].id)

		mat := pixel.IM
		mat = mat.Moved(pixel.V(x, y))
		mat = mat.Scaled(pixel.V(x, y), 2)

		Armys[n].txtId.Draw(win, mat)
	}

	army2 := gjson.Get(data.BattleData, "ArmyGroup2Init.Armys")
	for n, v := range army2.Array() {
		posx := gjson.Get(v.String(), "PosX").Int()
		posy := gjson.Get(v.String(), "PosY").Int()

		x, y := convertPos(int(posx), int(posy))
		Armys[n+4].x = x
		Armys[n+4].y = y
		Armys[n+4].ox = int(posx)
		Armys[n+4].oy = int(posy)
		Armys[n+4].id = n + 5
		Armys[n+4].sp.Draw(win, pixel.IM.Moved(pixel.V(x, y)))
		fmt.Fprintf(Armys[n+4].txtId, "%d", Armys[n+4].id)
		mat := pixel.IM
		mat = mat.Moved(pixel.V(x, y))
		mat = mat.Scaled(pixel.V(x, y), 2)

		Armys[n+4].txtId.Draw(win, mat)
	}

	initBattleInfo(win)
}

func initBattleInfo(win *pixelgl.Window) {
	basicTxt := text.New(pixel.V(0, 0), basicAtlas)
	basicTxt.Color = colornames.Red

	value := gjson.Get(data.BattleData, "BattleFrameDatas")
	for _, v := range value.Array() {
		operator := gjson.Get(v.String(), "Operator").Int()
		if operator == 3 { //dead
			frame := gjson.Get(v.String(), "Frame").Int()
			fid := gjson.Get(v.String(), "ArmyFieldId").Int()
			fmt.Fprintf(basicTxt, "ArmyId[%d]dead at frame[%d]\n", fid, frame)
		}
	}

	mat := pixel.IM
	mat = mat.Moved(pixel.V(1320, 300))
	mat = mat.Scaled(pixel.V(1320, 300), 2)
	basicTxt.Draw(win, mat)
}

func setArmyStateFrame(win *pixelgl.Window, frame int) {
	matchPath := fmt.Sprintf("BattleFrameDatas.#[Frame==\"%d\"]#", frame)

	value := gjson.Get(data.BattleData, matchPath)
	if len(value.Array()) == 0 {
		fmt.Println("该帧没有动作")
	} else {
		for _, v := range value.Array() {
			operator := gjson.Get(v.String(), "Operator").Int()
			fid := gjson.Get(v.String(), "ArmyFieldId").Int()
			posx := gjson.Get(v.String(), "Posx").Int()
			posy := gjson.Get(v.String(), "Posy").Int()
			x, y := convertPos(int(posx), int(posy))
			if operator == 1 { //移动
				Armys[fid-1].x = x
				Armys[fid-1].y = y
				Armys[fid-1].ox = int(posx)
				Armys[fid-1].oy = int(posy)
			}
		}
	}

	//开始画
	for _, v := range Armys {
		if v != nil {
			v.sp.Draw(win, pixel.IM.Moved(pixel.V(v.x, v.y)))
			mat := pixel.IM
			mat = mat.Moved(pixel.V(v.x, v.y))
			mat = mat.Scaled(pixel.V(v.x, v.y), 2)
			v.txtId.Draw(win, mat)
		}
	}

	//显示出位置
	setBattleInfoFrame(win)
}
func setBattleInfoFrame(win *pixelgl.Window) {
	battleInfoText.Clear()
	battleInfoText.Color = colornames.Red
	battleInfoText.Orig = pixel.V(0, 0)
	battleInfoText.Dot = pixel.V(0, 0)
	for _, v := range Armys {
		if v != nil {
			fmt.Fprintf(battleInfoText, "%d:xy[%f,%f];oxy[%d,%d]\n", v.id, v.x, v.y, v.ox, v.oy)
		}
	}

	mat := pixel.IM
	mat = mat.Moved(pixel.V(1320, 700))
	mat = mat.Scaled(pixel.V(1320, 700), 2)
	battleInfoText.Draw(win, mat)
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
