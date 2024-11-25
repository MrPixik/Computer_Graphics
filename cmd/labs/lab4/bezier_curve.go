package lab4

import (
	"Computer_Graphics/cmd/labs/utils"
	"gocv.io/x/gocv"
)

// BezierCurveThirdOrder plots Bezier curve by 3 points
func BezierCurveThirdOrder(img *gocv.Mat, p1, p2, p3 utils.Point, step float64) []utils.Point {
	segment1 := utils.Segment{
		X1: p1.X, X2: p2.X, Y1: p1.Y, Y2: p2.Y,
	}
	segment2 := utils.Segment{
		X1: p2.X, X2: p3.X, Y1: p2.Y, Y2: p3.Y,
	}
	var currParamValue float64
	var points []utils.Point
	for ; currParamValue < 1; currParamValue += step {
		point1 := segment1.GetCoordinatesByParam(currParamValue)
		point2 := segment2.GetCoordinatesByParam(currParamValue)

		currSegmentParam := utils.Segment{
			X1: point1.X, X2: point2.X, Y1: point1.Y, Y2: point2.Y,
		}
		currPoint := currSegmentParam.GetCoordinatesByParam(currParamValue)
		utils.DrawPixel(img, currPoint)
		points = append(points, currPoint)
	}
	return points
}
