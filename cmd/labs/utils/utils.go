package utils

import (
	"gocv.io/x/gocv"
	"math"
)

// Point type is interpretation of (x,y) coordinates of pixel
type Point struct {
	X, Y int
}

type Segment struct {
	X1 int
	X2 int
	Y1 int
	Y2 int
}

func (ps Segment) GetCoordinatesByParam(param float64) (p Point) {
	point := Point{
		X: int(float64(ps.X1) + param*float64(ps.X2-ps.X1)),
		Y: int(float64(ps.Y1) + param*float64(ps.Y2-ps.Y1)),
	}
	return point
}
func (ps Segment) SegmentDirectionVector() Vector {
	return Vector{ps.X2 - ps.X1, ps.Y2 - ps.Y1}
}

type Vector struct {
	X int
	Y int
}

func ScalarProduct(v1, v2 Vector) int {
	return v1.X*v2.X + v1.Y*v2.Y
}

// DrawPixel func sets value 255 for pixel by coordinates in Point
func DrawPixel(img *gocv.Mat, p Point) {
	img.SetUCharAt(p.Y, p.X, 255)
	//fmt.Printf("(%d, %d)\n", p.X, p.Y)
}

func MaxInt(x, y int) int {
	if x > y {
		return x
	}
	return y
}
func MinInt(x, y int) int {
	if x < y {
		return x
	}
	return y
}
func MyAbs(a int) int {
	return int(math.Abs(float64(a)))
}
