package lab4

import (
	"Computer_Graphics/cmd/labs/utils"
	"errors"
	"math"
)

// polygonOrientation checks if polygon is oriented counter-clockwise.  Returns true if CCW, false otherwise.
func polygonOrientation(polygon []utils.Point) bool {
	signedArea := 0
	for i := 0; i < len(polygon); i++ {
		signedArea += (polygon[i].X * polygon[(i+1)%len(polygon)].Y) - (polygon[(i+1)%len(polygon)].X * polygon[i].Y)
	}
	return signedArea < 0
}

// reversePolygon reverses the order of points in the polygon.
func reversePolygon(polygon []utils.Point) []utils.Point {
	for i, j := 0, len(polygon)-1; i < j; i, j = i+1, j-1 {
		polygon[i], polygon[j] = polygon[j], polygon[i]
	}
	return polygon
}

// Cyrus_Beck_Algorithm performs line clipping using the Cyrus-Beck algorithm.
func Cyrus_Beck_Algorithm(p1, p2 utils.Point, polygon []utils.Point) (utils.Segment, error) {
	// Ensure counter-clockwise orientation
	if !polygonOrientation(polygon) {
		polygon = reversePolygon(polygon)
	}
	segment := utils.Segment{
		X1: p1.X,
		X2: p2.X,
		Y1: p1.Y,
		Y2: p2.Y,
	}

	segmentDirectionV := segment.SegmentDirectionVector() //D

	paramEnter := 0.0
	paramExit := 1.0

	var scalarDN int // Scalar product of segment direction and polygon edge normal
	var scalarWN int // Scalar product of vector from segment start to polygon edge and its normal

	for i := 1; i < len(polygon); i++ {
		// Polygon's vertices
		v1 := polygon[i-1]
		v2 := polygon[i]

		normalVector := utils.Vector{ //N
			X: v2.Y - v1.Y,
			Y: v1.X - v2.X,
		}

		startToVertexVector := utils.Vector{ //W
			X: segment.X1 - v1.X,
			Y: segment.Y1 - v1.Y,
		}

		scalarDN = utils.ScalarProduct(segmentDirectionV, normalVector)
		scalarWN = utils.ScalarProduct(startToVertexVector, normalVector)

		if scalarDN == 0 { //the segment is parallel to the edge
			if scalarWN < 0 { //the segment is outside the polygon
				return utils.Segment{}, errors.New("line is outside polygon")
			} else {
				continue
			}
		}

		param := float64(-scalarWN) / float64(scalarDN)

		if scalarDN > 0 {
			paramEnter = math.Max(paramEnter, param)
		} else {
			paramExit = math.Min(paramExit, param)
		}
	}

	if paramEnter > paramExit { // the segment lies outside the polygon
		return utils.Segment{}, errors.New("line is outside polygon")
	}

	p1Clipped := utils.Point{
		X: segment.X1 + int(paramEnter*float64(segmentDirectionV.X)),
		Y: segment.Y1 + int(paramEnter*float64(segmentDirectionV.Y)),
	}
	p2Clipped := utils.Point{
		X: segment.X1 + int(paramExit*float64(segmentDirectionV.X)),
		Y: segment.Y1 + int(paramExit*float64(segmentDirectionV.Y)),
	}

	return utils.Segment{X1: p1Clipped.X, Y1: p1Clipped.Y, X2: p2Clipped.X, Y2: p2Clipped.Y}, nil
}
