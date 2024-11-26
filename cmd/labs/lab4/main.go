package lab4

import (
	"Computer_Graphics/cmd/labs/lab3"
	"Computer_Graphics/cmd/labs/utils"
	"gocv.io/x/gocv"
	"strconv"
)

func clipSegmentsByPolygons(segments []utils.Segment, polygons [][]utils.Point) {
	for i, segment := range segments {
		originP1 := utils.Point{X: segment.X1, Y: segment.Y1}
		originP2 := utils.Point{X: segment.X2, Y: segment.Y2}
		//DrawLine
		outputFilenameLine := "..\\..\\static\\images\\lab4\\straight_line" + strconv.Itoa(i+1) + ".jpg"
		imgLine := gocv.NewMatWithSize(2500, 2500, gocv.MatTypeCV8U)
		defer imgLine.Close()

		lab3.BresenhamLineAlgorithm(&imgLine, originP1, originP2)

		gocv.IMWrite(outputFilenameLine, imgLine)
		for j, polygon := range polygons {

			//Clip line by polygon
			outputFilenameLineClipped := "..\\..\\static\\images\\lab4\\straight_line" + strconv.Itoa(i+1) + "_clipped_by_polygon" + strconv.Itoa(j+1) + ".jpg"
			imgLineCLipped := gocv.NewMatWithSize(2500, 2500, gocv.MatTypeCV8U)
			defer imgLineCLipped.Close()

			clippedLine, err := Cyrus_Beck_Algorithm(originP1, originP2, polygon)

			if err == nil {
				p1 := utils.Point{
					X: clippedLine.X1,
					Y: clippedLine.Y1,
				}
				p2 := utils.Point{
					X: clippedLine.X2,
					Y: clippedLine.Y2,
				}
				lab3.BresenhamLineAlgorithm(&imgLineCLipped, p1, p2)
			}
			lab3.DrawPolygon(&imgLineCLipped, polygon)

			gocv.IMWrite(outputFilenameLineClipped, imgLineCLipped)
		}
	}
}

func Run() {
	//Straight line clipping

	segments := []utils.Segment{
		{X1: 500, Y1: 1000, X2: 1500, Y2: 1500},
		{X1: 1500, Y1: 2200, X2: 500, Y2: 2200},
		{X1: 1500, Y1: 700, X2: 800, Y2: 1500},
	}

	polygons := [][]utils.Point{
		{
			{500, 2000},
			{2000, 2200},
			{2300, 1500},
			{1800, 500},
			{600, 800},
			{500, 2000},
		},
		{
			{X: 500, Y: 2000},
			{X: 600, Y: 800},
			{X: 1800, Y: 500},
			{X: 2300, Y: 1500},
			{X: 2000, Y: 2200},
			{X: 500, Y: 2000},
		},
	}

	clipSegmentsByPolygons(segments, polygons)

	//Bezier Curve drawing
	var outputFilenameLine = "..\\..\\static\\images\\lab4\\bezier_curve.png"
	imgLine := gocv.NewMatWithSize(2500, 2500, gocv.MatTypeCV8U)
	defer imgLine.Close()

	point1 := utils.Point{X: 200, Y: 2200}
	point2 := utils.Point{X: 1250, Y: 200}
	point3 := utils.Point{X: 2300, Y: 2200}

	var numPoints = 5000

	bezierCurvePoints := BezierCurveThirdOrder(&imgLine, point1, point2, point3, numPoints)

	gocv.IMWrite(outputFilenameLine, imgLine)

	//Bezier Curve clipping
	var outputFilenameLineClipped = "..\\..\\static\\images\\lab4\\bezier_curve_clipped.png"
	imgLineClipped := gocv.NewMatWithSize(2500, 2500, gocv.MatTypeCV8U)
	defer imgLineClipped.Close()

	for i := 2; i < len(bezierCurvePoints); i += 2 {
		currPoint1 := bezierCurvePoints[i-2]
		currPoint2 := bezierCurvePoints[i]

		clippedLine, err := Cyrus_Beck_Algorithm(currPoint1, currPoint2, polygons[0])

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

	lab3.DrawPolygon(&imgLineClipped, polygons[0])

	gocv.IMWrite(outputFilenameLineClipped, imgLineClipped)
}
