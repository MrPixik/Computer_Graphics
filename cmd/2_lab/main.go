package main

import (
	"fmt"
	"gocv.io/x/gocv"
	"math"
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

func ditheringToNbpp(pixVal uint8, n float64) (uint8, int) {
	numPoints := math.Pow(2, n) - 1 //Num of supported values
	step := int(255 / numPoints)
	halfStep := uint8(step / 2)
	intervalSlice := []uint8{} //Supported values for pixels

	for i := 0; i <= 255; i += step {
		intervalSlice = append(intervalSlice, uint8(i))
	}

	// Finding interval for current pixel
	for i := 0; i < len(intervalSlice)-1; i++ {
		low := intervalSlice[i]
		high := intervalSlice[i+1]

		if pixVal >= low && pixVal < high {
			quantizedVal := low
			if pixVal-low > halfStep {
				quantizedVal = high
			}
			err := int(pixVal) - int(quantizedVal)
			return quantizedVal, err
		}
	}

	// If pixVal == 255 return 255
	return 255, 0
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
	//var originFilename = "..\\..\\static\\images\\2_lab\\Marcus_Aurelius.jpg"
	//var outputFilename = "..\\..\\static\\images\\2_lab\\Marcus_Aurelius_dithered_1bpp.png"

	var originFilename = "..\\..\\static\\images\\2_lab\\elephant.jpg"
	var outputFilename = "..\\..\\static\\images\\2_lab\\elephant_dithered_2bpp.png"

	originImg := gocv.IMRead(originFilename, gocv.IMReadColor)
	defer originImg.Close()

	wigth := originImg.Cols()
	height := originImg.Rows()
	grayImg := rgb2Grayscale(originImg)
	defer grayImg.Close()

	processingData, _ := grayImg.DataPtrUint8()

	for y := 0; y < height; y++ {
		for x := 0; x < wigth; x++ {
			index := wigth*y + x
			currPixVal := processingData[index]
			currPixValResult, err := ditheringToNbpp(currPixVal, 2)
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
	//Test for uniq pixel values
	uniq := []uint8{processingData[0]}
	for _, pixval := range processingData {
		isUnique := true
		for _, value := range uniq {
			if value == pixval {
				isUnique = false
				break
			}
		}
		if isUnique {
			uniq = append(uniq, pixval)
		}
	}
	fmt.Println(uniq)
	ditheringImage := gocv.NewMatWithSize(height, wigth, gocv.MatTypeCV8U)
	defer ditheringImage.Close()

	ptr, _ := ditheringImage.DataPtrUint8()
	copy(ptr, processingData)

	gocv.IMWrite(outputFilename, ditheringImage)

}

func main() {

	ditheringFloydSteinberg()
}
