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

	//Straight line clipping
	var outputFilenameStraightLine1 = "..\\..\\static\\images\\lab4\\straight_line1.png"
	var outputFilenameStraightLine2 = "..\\..\\static\\images\\lab4\\straight_line2.png"
	var outputFilenameStraightLine3 = "..\\..\\static\\images\\lab4\\straight_line3.png"

	imgStraightLine1 := gocv.NewMatWithSize(2500, 2500, gocv.MatTypeCV8U)
	imgStraightLine2 := gocv.NewMatWithSize(2500, 2500, gocv.MatTypeCV8U)
	imgStraightLine3 := gocv.NewMatWithSize(2500, 2500, gocv.MatTypeCV8U)

	defer imgStraightLine1.Close()
	defer imgStraightLine2.Close()
	defer imgStraightLine3.Close()

	lineFirstPoint1 := utils.Point{X: 500, Y: 1000}
	lineFirstPoint2 := utils.Point{X: 1500, Y: 1500}

	lineSecondPoint1 := utils.Point{X: 1500, Y: 2200}
	lineSecondPoint2 := utils.Point{X: 500, Y: 2200}

	lineThirdPoint1 := utils.Point{X: 1500, Y: 700}
	lineThirdPoint2 := utils.Point{X: 800, Y: 1500}

	lab3.BresenhamLineAlgorithm(&imgStraightLine1, lineFirstPoint1, lineFirstPoint2)
	lab3.BresenhamLineAlgorithm(&imgStraightLine2, lineSecondPoint1, lineSecondPoint2)
	lab3.BresenhamLineAlgorithm(&imgStraightLine3, lineThirdPoint1, lineThirdPoint2)

	gocv.IMWrite(outputFilenameStraightLine1, imgStraightLine1)
	gocv.IMWrite(outputFilenameStraightLine2, imgStraightLine2)
	gocv.IMWrite(outputFilenameStraightLine3, imgStraightLine3)

	var outputFilenameStraightLineClipped1 = "..\\..\\static\\images\\lab4\\straight_line1_clipped.png"
	var outputFilenameStraightLineClipped2 = "..\\..\\static\\images\\lab4\\straight_line2_clipped.png"
	var outputFilenameStraightLineClipped3 = "..\\..\\static\\images\\lab4\\straight_line3_clipped.png"

	imgStraightLineClipped1 := gocv.NewMatWithSize(2500, 2500, gocv.MatTypeCV8U)
	imgStraightLineClipped2 := gocv.NewMatWithSize(2500, 2500, gocv.MatTypeCV8U)
	imgStraightLineClipped3 := gocv.NewMatWithSize(2500, 2500, gocv.MatTypeCV8U)

	defer imgStraightLineClipped1.Close()
	defer imgStraightLineClipped2.Close()
	defer imgStraightLineClipped3.Close()

	pointsPolygonConvex := []utils.Point{
		{500, 2000},
		{2000, 2200},
		{2300, 1500},
		{1800, 500},
		{600, 800},
		{500, 2000},
	}

	clippedLine1, err := Cyrus_Beck_Algorithm(lineFirstPoint1, lineFirstPoint2, pointsPolygonConvex)

	if err == nil {
		p1 := utils.Point{
			X: clippedLine1.X1,
			Y: clippedLine1.Y1,
		}
		p2 := utils.Point{
			X: clippedLine1.X2,
			Y: clippedLine1.Y2,
		}
		lab3.BresenhamLineAlgorithm(&imgStraightLineClipped1, p1, p2)
	}
	lab3.DrawPolygon(&imgStraightLineClipped1, pointsPolygonConvex)

	gocv.IMWrite(outputFilenameStraightLineClipped1, imgStraightLineClipped1)

	clippedLine2, err := Cyrus_Beck_Algorithm(lineSecondPoint1, lineSecondPoint2, pointsPolygonConvex)

	if err == nil {
		p1 := utils.Point{
			X: clippedLine2.X1,
			Y: clippedLine2.Y1,
		}
		p2 := utils.Point{
			X: clippedLine2.X2,
			Y: clippedLine2.Y2,
		}
		lab3.BresenhamLineAlgorithm(&imgStraightLineClipped2, p1, p2)
	}
	lab3.DrawPolygon(&imgStraightLineClipped2, pointsPolygonConvex)

	gocv.IMWrite(outputFilenameStraightLineClipped2, imgStraightLineClipped2)

	clippedLine3, err := Cyrus_Beck_Algorithm(lineThirdPoint1, lineThirdPoint2, pointsPolygonConvex)

	if err == nil {
		p1 := utils.Point{
			X: clippedLine3.X1,
			Y: clippedLine3.Y1,
		}
		p2 := utils.Point{
			X: clippedLine3.X2,
			Y: clippedLine3.Y2,
		}
		lab3.BresenhamLineAlgorithm(&imgStraightLineClipped3, p1, p2)
	}
	lab3.DrawPolygon(&imgStraightLineClipped3, pointsPolygonConvex)

	gocv.IMWrite(outputFilenameStraightLineClipped3, imgStraightLineClipped3)

	//Bezier Curve clipping
	var outputFilenameLineClipped = "..\\..\\static\\images\\lab4\\bezier_curve_clipped.png"
	imgLineClipped := gocv.NewMatWithSize(2500, 2500, gocv.MatTypeCV8U)
	defer imgLineClipped.Close()

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
