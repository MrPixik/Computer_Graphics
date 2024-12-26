package bdz

import (
	"Computer_Graphics/cmd/labs/lab3"
	"Computer_Graphics/cmd/labs/utils"
	"fmt"
	"gocv.io/x/gocv"
)

func firstTask() {
	// Рисуем исходный полигон
	var outputFilenamePolygon = "..\\..\\static\\images\\bdz\\Polygon.png"

	imgPolygon := gocv.NewMatWithSize(250, 250, gocv.MatTypeCV8U)

	defer imgPolygon.Close()

	pointsPolygon := []utils.Point{
		{20, 170},
		{50, 150},
		{30, 100},
		{70, 110},
		{120, 80},
		{150, 90},
		{150, 180},
		{100, 160},
	}

	lab3.DrawPolygon(&imgPolygon, pointsPolygon)

	gocv.IMWrite(outputFilenamePolygon, imgPolygon)
	if lab3.IsSelfIntersectingPolygon(pointsPolygon) {
		fmt.Println("Polygon is self-intersected")
	} else {
		fmt.Println("Polygon is non-self-intersected")
	}
	// Рисуем выпуклую оболочку
	var outputFilenameConvexHull = "..\\..\\static\\images\\bdz\\ConvexHull.png"

	imgConvexHull := gocv.NewMatWithSize(250, 250, gocv.MatTypeCV8U)

	defer imgConvexHull.Close()

	convexHull := JarvisAlgorithm(pointsPolygon)

	lab3.DrawPolygon(&imgConvexHull, convexHull)

	gocv.IMWrite(outputFilenameConvexHull, imgConvexHull)
}

func Run() {
	firstTask()
}
