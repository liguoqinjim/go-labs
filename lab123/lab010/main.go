package main

import (
	"github.com/sclevine/agouti"
	"image"
	"image/png"
	"log"
	"os"
	"time"
)

func main() {
	driver := agouti.ChromeDriver()
	if err := driver.Start(); err != nil {
		log.Fatalf("Failed to start driver:%v", err)
	}
	defer func() {
		time.Sleep(time.Hour)
		driver.Stop()
	}()

	page, err := driver.NewPage(agouti.Browser("chrome"))
	if err != nil {
		log.Fatalf("Failed to open page:%v", err)
	}

	//打开页面
	if err := page.Navigate("https://gocn.vip/"); err != nil {
		log.Fatalf("Failed to navigate:%v", err)
	}

	//找出img元素
	secImg := page.Find("body > div.aw-container-wrap > div.container.category > div > div > dl:nth-child(1) > dt > img")

	es, err := secImg.Elements()
	if len(es) == 0 {
		log.Fatalf("secImg.Elements error:%v", err)
	}
	eImg := es[0]
	x, y, err := eImg.GetLocation()
	if err != nil {
		log.Fatalf("v.GetLocation error:%v", err)
	}
	log.Printf("x[%d]y[%d]", x, y)
	height, width, err := eImg.GetSize()
	if err != nil {
		log.Fatalf("v.GetSize error:%v", err)
	}
	log.Printf("height[%d]width[%d]", height, width)

	//整个屏幕截图
	if err := page.Screenshot("screenshot.png"); err != nil {
		log.Fatalf("page.Screenshot error:%v", err)
	}

	//裁切图片
	fImg, err := os.Open("screenshot.png")
	if err != nil {
		log.Fatalf("os.Open error:%v", err)
	}
	defer fImg.Close()

	img, err := png.Decode(fImg)
	if err != nil {
		log.Fatalf("jpeg.Decode error:%v", err)
	}

	//裁切
	subImg := img.(interface {
		SubImage(r image.Rectangle) image.Image
	}).SubImage(image.Rect(x, y, x+height, y+width))

	log.Printf("bounds %v\n", subImg.Bounds())

	f, err := os.Create("subImages.png")
	if err != nil {
		log.Fatalf("os.Create error:%v", err)
	}
	defer f.Close()

	//保存
	err = png.Encode(f, subImg)
	if err != nil {
		log.Fatalf("jpeg.Encode error:%v", err)
	}
}
