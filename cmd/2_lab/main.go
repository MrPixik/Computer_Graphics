package main

import (
	"fmt"
	"gocv.io/x/gocv"
)

func rgb2Grayscale(originImg gocv.Mat) gocv.Mat {
	//Image info
	wigth := originImg.Cols()
	height := originImg.Rows()
	channels := originImg.Channels() //num of colour channels

	//Creating pointer on origin image's data slice
	processingImage, _ := originImg.DataPtrUint8()

	//Creating data slice for gray image
	grayData := make([]uint8, height*wigth)

	for y := 0; y < height; y++ {
		for x := 0; x < wigth; x++ {
			index := (y*wigth + x) * channels
			grayIndex := y*wigth + x

			blue := processingImage[index]
			red := processingImage[index+1]
			green := processingImage[index+2]

			gray := uint8(0.299*float32(red) + 0.587*float32(green) + 0.114*float32(blue))

			grayData[grayIndex] = gray

		}
	}
	//Creating new gocv.Mat for writing grayscale image
	grayImage := gocv.NewMatWithSize(height, wigth, gocv.MatTypeCV8U)

	ptr, _ := grayImage.DataPtrUint8()
	copy(ptr, grayData)
	//var outputFilename = "..\\..\\static\\images\\2_lab\\Marcus_Aurelius_grayscale.jpg"
	//gocv.IMWrite(outputFilename, grayImage)
	return grayImage
}

func ditheringTo1bpp(pixVal uint8) (uint8, int) {
	if pixVal < 0 {
		return 0, 0
	} else if pixVal < 127 {
		return 0, int(pixVal)
	} else if pixVal < 255 {
		return 255, int(pixVal) - 255
	} else {
		return 255, 0
	}
}
func ditheringTo2bpp(pixVal uint8) (uint8, int) {
	if pixVal <= 0 {
		return 0, 0
	} else if pixVal <= 31 {
		return 0, int(pixVal)
	} else if pixVal <= 95 {
		return 63, int(pixVal) - 63
	} else if pixVal <= 159 {
		return 127, int(pixVal) - 127
	} else if pixVal <= 223 {
		return 191, int(pixVal) - 191
	} else if pixVal < 255 {
		return 255, int(pixVal) - 255
	} else {
		return 255, 0
	}
}

// clamp limiting value to interval [0,255]
func clamp(value int) uint8 {
	if value < 0 {
		return 0
	} else if value > 255 {
		return 255
	}
	return uint8(value)
}
func ditheringFloydSteinberg() {
	var originFilename = "..\\..\\static\\images\\2_lab\\Marcus_Aurelius.jpg"
	var outputFilename = "..\\..\\static\\images\\2_lab\\Marcus_Aurelius_dithered_2bpp.jpg"

	//var originFilename = "..\\..\\static\\images\\2_lab\\elephant.jpg"
	//var outputFilename = "..\\..\\static\\images\\2_lab\\elephant_dithered_1bpp.jpg"

	originImg := gocv.IMRead(originFilename, gocv.IMReadColor)
	defer originImg.Close()

	wigth := originImg.Cols()
	height := originImg.Rows()
	fmt.Println(wigth, height)
	grayImg := rgb2Grayscale(originImg)
	defer grayImg.Close()

	processingData, _ := grayImg.DataPtrUint8()

	for y := 0; y < height; y++ {
		for x := 0; x < wigth; x++ {
			index := wigth*y + x
			currPixVal := processingData[index]
			currPixValResult, err := ditheringTo2bpp(currPixVal)
			if (x < wigth-1) && (y < height-1) && (x > 0) {
				processingData[index] = currPixValResult
				processingData[index+1] += clamp((7 * err) / 16)
				processingData[index+wigth-1] += clamp((3 * err) / 16)
				processingData[index+wigth] += clamp((5 * err) / 16)
				processingData[index+wigth+1] += clamp(err / 16)
			} else if (x == 0) && (y == height-1) { //Lower left corner
				processingData[index] = currPixValResult
				processingData[index+1] += clamp(7 * err / 16)
			} else if (x == wigth-1) && (y == height-1) { //Lower right corner
				processingData[index] = currPixValResult
			} else if x == 0 { //Left border
				processingData[index] = currPixValResult
				processingData[index+1] += clamp(7 * err / 16)
				processingData[index+wigth] += clamp(5 * err / 16)
				processingData[index+wigth+1] += clamp(err / 16)
			} else if x == wigth-1 { //Right border
				processingData[index] = currPixValResult
				processingData[index+wigth-1] += clamp(3 * err / 16)
				processingData[index+wigth] += clamp(5 * err / 16)
			} else if y == height-1 { //Low border
				processingData[index] = currPixValResult
				processingData[index+1] += clamp(7 * err / 16)
			}
		}
	}
	ditheringImage := gocv.NewMatWithSize(height, wigth, gocv.MatTypeCV8U)
	defer ditheringImage.Close()

	ptr, _ := ditheringImage.DataPtrUint8()
	copy(ptr, processingData)

	gocv.IMWrite(outputFilename, ditheringImage)

}

func main() {

	ditheringFloydSteinberg()
}
