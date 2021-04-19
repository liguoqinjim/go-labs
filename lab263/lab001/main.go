package main

import (
	"fmt"
	"github.com/vitali-fedulov/images"
)

func main() {
	// Open photos.
	imgA, err := images.Open("../datas/01.jpg")
	if err != nil {
		panic(err)
	}
	imgB, err := images.Open("../datas/02.jpeg")
	if err != nil {
		panic(err)
	}
	imgC, err := images.Open("../datas/03.jpeg")
	if err != nil {
		panic(err)
	}

	// Calculate hashes and image sizes.
	hashA, imgSizeA := images.Hash(imgA)
	hashB, imgSizeB := images.Hash(imgB)
	hashC, imgSizeC := images.Hash(imgC)

	// Image comparison.
	if images.Similar(hashA, hashB, imgSizeA, imgSizeB) {
		fmt.Println("Images are similar.")
	} else {
		fmt.Println("Images are distinct.")
	}

	if images.Similar(hashA, hashC, imgSizeA, imgSizeC) {
		fmt.Println("Images are similar.")
	} else {
		fmt.Println("Images are distinct.")
	}
}
