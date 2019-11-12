package main

import (
	"fmt"
	"image/jpeg"
	"log"
	"os"
)

var imagePath = "./example.jpeg"

func main() {
	_, err := os.Stat(imagePath)
	if err != nil && os.IsNotExist(err) {
		log.Println("image not exists")
		os.Exit(1)
	}

	fileReader, err := os.Open(imagePath)
	if err != nil {
		panic(err)
	}

	image, err := jpeg.Decode(fileReader)
	if err != nil {
		panic(err)
	}

	b := image.Bounds()
	log.Println(b.Min.X, b.Min.Y)
	log.Println(b.Max.X, b.Max.Y)

	step := 8
	for y := 0; y < b.Max.Y; y += step {
		for i := 0; i < b.Max.X; i += step {
			r, g, b, _ := image.At(i, y).RGBA()
			y := 0.299*float64(r) + 0.587*float64(g) + 0.114*float64(b)

			if y >= 65535/2 {
				fmt.Print("x")
			}
			fmt.Print(" ")

		}
		fmt.Println()
	}

}
