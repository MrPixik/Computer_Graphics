package lab4

import (
	"Computer_Graphics/cmd/labs/lab3"
	"Computer_Graphics/cmd/labs/utils"
	"gocv.io/x/gocv"
)

func Run() {
	//Bezier Curve drawing
	var outputFilenameLine = "..\\..\\static\\images\\lab4\\bezier_curve.png"
	imgLine := gocv.NewMatWithSize(2500, 2500, gocv.MatTypeCV8U)
	defer imgLine.Close()

	point1 := utils.Point{X: 200, Y: 2200}
	point2 := utils.Point{X: 1250, Y: 200}
	point3 := utils.Point{X: 2300, Y: 2200}

	var step = 0.001

	bezierCurvePoints := BezierCurveThirdOrder(&imgLine, point1, point2, point3, step)

	gocv.IMWrite(outputFilenameLine, imgLine)

	//Bezier Curve clipping
	var outputFilenameLineClipped = "..\\..\\static\\images\\lab4\\straight_line_clipped.png"
	imgLineClipped := gocv.NewMatWithSize(2500, 2500, gocv.MatTypeCV8U)
	defer imgLineClipped.Close()

	pointsPolygonConvex := []utils.Point{
		{500, 2000},
		{2000, 2200},
		{2300, 1500},
		{1800, 500},
		{600, 800},
	}
	for i := 2; i < len(bezierCurvePoints); i += 2 {
		currPoint1 := bezierCurvePoints[i-2]
		currPoint2 := bezierCurvePoints[i]

		clippedLine, err := Cyrus_Beck_Algorithm(currPoint1, currPoint2, pointsPolygonConvex)

		if err != nil {
			continue
		}
		p1 := utils.Point{
			X: clippedLine.X1,
			Y: clippedLine.Y1,
		}
		p2 := utils.Point{
			X: clippedLine.X2,
			Y: clippedLine.Y2,
		}
		lab3.BresenhamLineAlgorithm(&imgLineClipped, p1, p2)
	}

	lab3.DrawPolygon(&imgLineClipped, pointsPolygonConvex)

	gocv.IMWrite(outputFilenameLineClipped, imgLineClipped)

}
