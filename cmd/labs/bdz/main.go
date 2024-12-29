package bdz

import (
	"Computer_Graphics/cmd/labs/lab3"
	"Computer_Graphics/cmd/labs/utils"
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

	// Рисуем выпуклую оболочку
	var outputFilenameConvexHull = "..\\..\\static\\images\\bdz\\ConvexHull.png"

	imgConvexHull := gocv.NewMatWithSize(250, 250, gocv.MatTypeCV8U)

	defer imgConvexHull.Close()

	convexHull := JarvisAlgorithm(pointsPolygon)

	lab3.DrawPolygon(&imgConvexHull, convexHull)

	gocv.IMWrite(outputFilenameConvexHull, imgConvexHull)
}

func secondTask() {
	var outputFilename = "..\\..\\static\\images\\bdz\\HermiteCurve.png"

	img1 := gocv.NewMatWithSize(1000, 1000, gocv.MatTypeCV8U)

	defer img1.Close()

	points := []utils.Point{
		{250, 250},
		{450, 550},
		{750, 350},
		{950, 650},
	}
	vectors := []utils.Point{
		{100, 100},
		{200, -100},
		{100, 200},
		{0, -100},
	}
	drawCompositeHermiteCurve(&img1, points, vectors)

	gocv.IMWrite(outputFilename, img1)
}
func thirdTask() {
	orderedDithering()
}

func Run() {
	//Напишите программу, которая находит полигон, который является внешним контуром
	//полигона с самопересечениями. Исходный и получившийся полигоны выведите на экран.
	//firstTask()

	//Напишите программу, которая строит составную кубическую кривую Эрмита. Вершины,
	//через которые проходит кривая, и направляющие вектора задаются отдельными массивами.
	//secondTask()

	//Напишите программу, которая преобразует полутоновое изображение 8 bpp в заданное
	//количество оттенков серого с использованием равномерной палитры и упорядоченного
	//псевдосмешения.
	thirdTask()

}
