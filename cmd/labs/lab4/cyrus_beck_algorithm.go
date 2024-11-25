package lab4

import (
	"Computer_Graphics/cmd/labs/lab3"
	"Computer_Graphics/cmd/labs/utils"
	"errors"
	"math"
)

func Cyrus_Beck_Algorithm(p1, p2 utils.Point, polygon []utils.Point) (utils.Segment, error) {
	segment := utils.Segment{
		X1: p1.X,
		X2: p2.X,
		Y1: p1.Y,
		Y2: p2.Y,
	}
	segmentDirectionV := segment.SegmentDirectionVector() //D

	paramEnter := 0.0
	paramExit := 1.0

	var scalarDN int
	var scalarWN int

	if !lab3.IsPointInPolygonEvenOdd(polygon, p1) && !lab3.IsPointInPolygonEvenOdd(polygon, p2) {
		return utils.Segment{}, errors.New("line is outside polygon")
	}

	for i := 1; i < len(polygon); i++ {

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

		if scalarDN == 0 {
			if scalarWN < 0 {
				return utils.Segment{}, errors.New("line is outside polygon")
			} else {
				continue
			}
		}

		param := float64(-scalarWN) / float64(scalarDN)

		if param > 0 && param < 1 {
			if scalarDN < 0 {
				paramEnter = math.Max(paramEnter, param)
			} else {
				paramExit = math.Min(paramExit, param)
			}
		}

	}
	if paramEnter > paramExit {
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
