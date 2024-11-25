package lab1

import (
	"fmt"
	"gocv.io/x/gocv"
)

func blending(alpha float64) {
	var originFilename1 = "..\\..\\static\\images\\lab1\\blending1.png"
	var originFilename2 = "..\\..\\static\\images\\lab1\\blending2.png"
	var outputFilename = "..\\..\\static\\images\\lab1\\blending_result.png"

	//Reading origin images
	originImg1 := gocv.IMRead(originFilename1, gocv.IMReadColor)
	defer originImg1.Close()

	originImg2 := gocv.IMRead(originFilename2, gocv.IMReadColor)
	defer originImg2.Close()

	//s info
	wigth1 := originImg1.Cols()
	height1 := originImg1.Rows()
	channels1 := originImg1.Channels() //num of colour channels

	wigth2 := originImg2.Cols()
	height2 := originImg2.Rows()
	channels2 := originImg2.Channels() //num of colour channels

	if !((wigth1 == wigth2) && (height1 == height2) && (channels1 == channels2)) {
		fmt.Println("Blending Failed")
		fmt.Println(wigth1, wigth2)
		fmt.Println(height1, height2)
		fmt.Println(channels1, channels2)
		return
	}
	//Creating pointer on origin image's data slice
	processingImage1, _ := originImg1.DataPtrUint8()
	processingImage2, _ := originImg2.DataPtrUint8()

	blendingData := make([]uint8, height1*wigth1*channels1)

	for y := 0; y < height1; y++ {
		for x := 0; x < wigth1; x++ {
			index := (y*wigth1 + x) * channels1

			blue1 := processingImage1[index]
			red1 := processingImage1[index+1]
			green1 := processingImage1[index+2]

			blue2 := processingImage2[index]
			red2 := processingImage2[index+1]
			green2 := processingImage2[index+2]

			resultBlue := uint8(alpha*float64(blue1) + (1-alpha)*float64(blue2))
			resultRed := uint8(alpha*float64(red1) + (1-alpha)*float64(red2))
			resultGreen := uint8(alpha*float64(green1) + (1-alpha)*float64(green2))

			blendingData[index] = resultBlue
			blendingData[index+1] = resultRed
			blendingData[index+2] = resultGreen
		}
	}
	blendingImage := gocv.NewMatWithSize(height1, wigth1, gocv.MatTypeCV8UC3)
	defer blendingImage.Close()

	ptr, _ := blendingImage.DataPtrUint8()
	copy(ptr, blendingData)

	gocv.IMWrite(outputFilename, blendingImage)

}
