package main

import (
	"image"
	"image/jpeg"
	"log"
	"os"
)

func main() {
	imageFile, err := os.Open("images.jpg")
	if err != nil {
		log.Fatalf("os.Open error:%v", err)
	}
	defer imageFile.Close()

	img, err := jpeg.Decode(imageFile)
	if err != nil {
		log.Fatalf("jpeg.Decode error:%v", err)
	}

	//裁切
	subImg := img.(interface {
		SubImage(r image.Rectangle) image.Image
	}).SubImage(image.Rect(0, 0, 150, 150))

	log.Printf("bounds %v\n", subImg.Bounds())

	f, err := os.Create("subImages.jpg")
	if err != nil {
		log.Fatalf("os.Create error:%v", err)
	}
	defer f.Close()

	//保存
	err = jpeg.Encode(f, subImg, nil)
	if err != nil {
		log.Fatalf("jpeg.Encode error:%v", err)
	}
}
