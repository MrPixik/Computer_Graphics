package lab1

import (
	"gocv.io/x/gocv"
	"math"
)

func grayRoundImage() {
	var originFilename = "..\\..\\static\\images\\lab1\\apple.jpeg"
	var outputFilename = "..\\..\\static\\images\\lab1\\gray_apple.jpeg"

	//Reading origin image
	originImg := gocv.IMRead(originFilename, gocv.IMReadColor)
	defer originImg.Close()

	//Image info
	wigth := originImg.Cols()
	height := originImg.Rows()
	channels := originImg.Channels() //num of colour channels

	//Circle data
	radSquare := math.Pow(float64(min(wigth, height))/2, 2)

	centerX := wigth / 2
	centerY := height / 2

	//Creating pointer on origin image's data slice
	processingImage, _ := originImg.DataPtrUint8()

	//Creating data slice for gray image
	grayData := make([]uint8, height*wigth)

	for y := 0; y < height; y++ {
		for x := 0; x < wigth; x++ {
			index := (y*wigth + x) * channels
			grayIndex := y*wigth + x

			//test := math.Pow(float64(centerX-x), 2) + math.Pow(float64(centerY-y), 2)
			//fmt.Println(test)
			if math.Pow(float64(centerX-x), 2)+math.Pow(float64(centerY-y), 2) <= radSquare {
				blue := processingImage[index]
				red := processingImage[index+1]
				green := processingImage[index+2]

				gray := uint8(0.299*float32(red) + 0.587*float32(green) + 0.114*float32(blue))

				grayData[grayIndex] = gray
			} else {
				grayData[grayIndex] = 0
			}
		}
	}
	//Creating new gocv.Mat for writing grayscale image
	grayImage := gocv.NewMatWithSize(height, wigth, gocv.MatTypeCV8U)
	defer grayImage.Close()

	ptr, _ := grayImage.DataPtrUint8()
	copy(ptr, grayData)

	gocv.IMWrite(outputFilename, grayImage)
}
