package bdz

import (
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
	//var outputFilename = "..\\..\\static\\images\\lab2\\Marcus_Aurelius_grayscale.jpg"
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
func clamp(pixVal uint8, err int) uint8 {
	var result int = int(pixVal) + err
	if result < 0 {
		return 0
	} else if result > 255 {
		return 255
	}
	return uint8(result)
}
func orderedDithering() {
	//var originFilename = "..\\..\\static\\images\\lab2\\Marcus_Aurelius.jpg"
	//var outputFilename = "..\\..\\static\\images\\lab2\\Marcus_Aurelius_dithered_2bpp.png"

	var originFilename = "..\\..\\static\\images\\bdz\\elephant.png"
	var outputFilename = "..\\..\\static\\images\\bdz\\elephant_dithered_2bpp.png"

	originImg := gocv.IMRead(originFilename, gocv.IMReadColor)
	defer originImg.Close()

	wigth := originImg.Cols()
	height := originImg.Rows()
	grayImg := rgb2Grayscale(originImg)
	defer grayImg.Close()

	processingData, _ := grayImg.DataPtrUint8()

	bayerMatrix := [8][8]int{
		{0, 32, 8, 40, 2, 34, 10, 42},
		{48, 16, 56, 24, 50, 18, 58, 26},
		{12, 44, 4, 36, 14, 46, 6, 38},
		{60, 28, 52, 20, 62, 30, 54, 22},
		{3, 35, 11, 43, 1, 33, 9, 41},
		{51, 19, 59, 27, 49, 17, 57, 25},
		{15, 47, 7, 39, 13, 45, 5, 37},
		{63, 31, 55, 23, 61, 29, 53, 21},
	}

	for y := 0; y < height; y++ {
		for x := 0; x < wigth; x++ {
			index := wigth*y + x
			currPixVal := processingData[index]
			currPixValResult, _ := ditheringToNbpp(currPixVal, 2)
			processingData[index] = currPixValResult

			// Применение порога
			threshold := bayerMatrix[y%8][x%8] * 255 / 63

			if int(currPixValResult) > threshold {
				processingData[index] = 255
			} else {
				processingData[index] = 0
			}
		}
	}

	ditheringImage := gocv.NewMatWithSize(height, wigth, gocv.MatTypeCV8U)
	defer ditheringImage.Close()

	ptr, _ := ditheringImage.DataPtrUint8()
	copy(ptr, processingData)

	gocv.IMWrite(outputFilename, ditheringImage)

}
