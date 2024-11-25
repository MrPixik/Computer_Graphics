package lab3

import (
	"Computer_Graphics/cmd/labs/utils"
	"gocv.io/x/gocv"
)

// CLPointType Types of points
type CLPointType int

const (
	Left CLPointType = iota
	Right
	Behind
	Origin
	Destination
	Between
)

// IntersectType Types of intersections
type IntersectType int

const (
	Parallel IntersectType = iota
	Same
	Skew
	SkewNoCross
	SkewCross
)

// classify classifies a point relative to a segment
func classify(x1, y1, x2, y2, x, y float64) CLPointType {
	ax := x2 - x1
	ay := y2 - y1
	bx := x - x1
	by := y - y1
	s := ax*by - bx*ay
	if s > 0 {
		return Left
	}
	if s < 0 {
		return Right
	}
	if (ax*bx < 0) || (ay*by < 0) {
		return Behind
	}
	if (ax*ax + ay*ay) < (bx*bx + by*by) {
		return Behind
	}
	if x1 == x && y1 == y {
		return Origin
	}
	if x2 == x && y2 == y {
		return Destination
	}
	return Between
}

// intersect checks whether two segments intersects
func intersect(ax, ay, bx, by, cx, cy, dx, dy float64, t *float64) IntersectType {
	nx := dy - cy
	ny := cx - dx
	var typeCL CLPointType
	denom := nx*(bx-ax) + ny*(by-ay)
	if denom == 0 {
		typeCL = classify(cx, cy, dx, dy, ax, ay)
		if typeCL == Left || typeCL == Right {
			return Parallel
		} else {
			return Same
		}
	}
	num := nx*(ax-cx) + ny*(ay-cy)
	*t = -num / denom
	return Skew
}

// cross checks for intersection of two segments and returns parameters of intersection
func cross(ax, ay, bx, by, cx, cy, dx, dy float64, tab, tcd *float64) IntersectType {
	typeInt := intersect(ax, ay, bx, by, cx, cy, dx, dy, tab)
	if typeInt == Same || typeInt == Parallel {
		return typeInt
	}
	if *tab < 0 || *tab > 1 {
		return SkewNoCross
	}
	intersect(cx, cy, dx, dy, ax, ay, bx, by, tcd)
	if *tcd < 0 || *tcd > 1 {
		return SkewNoCross
	}
	return SkewCross
}

// IsPointInPolygonEvenOdd determines whether a point is inside a polygon using the even-odd rule
func IsPointInPolygonEvenOdd(polygon []utils.Point, pt utils.Point) bool {
	intersections := 0
	var t1, t2 float64
	outsidePoint := utils.Point{X: pt.X + 10000, Y: pt.Y + 1}

	for i := 0; i < len(polygon); i++ {
		v1 := polygon[i]
		v2 := polygon[(i+1)%len(polygon)]

		intersectType := cross(float64(v1.X), float64(v1.Y), float64(v2.X), float64(v2.Y), float64(pt.X), float64(pt.Y), float64(outsidePoint.X), float64(outsidePoint.Y), &t1, &t2)
		if intersectType == SkewCross {
			intersections++
		}
	}
	return (intersections % 2) == 1
}

// IsPointInPolygonEvenOdd determines whether a point is inside a polygon using the even-odd rule
func FillPolygonEvenOdd(img *gocv.Mat, polygon []utils.Point) {
	minX, minY, maxX, maxY := polygon[0].X, polygon[0].Y, polygon[0].X, polygon[0].Y
	for _, pt := range polygon {
		minX = utils.MinInt(minX, pt.X)
		minY = utils.MinInt(minY, pt.Y)
		maxX = utils.MaxInt(maxX, pt.X)
		maxY = utils.MaxInt(maxY, pt.Y)
	}

	for y := minY; y <= maxY; y++ {
		for x := minX; x <= maxX; x++ {
			pt := utils.Point{X: x, Y: y}
			if IsPointInPolygonEvenOdd(polygon, pt) {
				img.SetUCharAt(y, x, 255)
			}
		}
	}
}

// isPointInPolygonNonZeroWinding determines whether a point is inside a polygon using the non-zero winding rule
func isPointInPolygonNonZeroWinding(polygon []utils.Point, pt utils.Point) bool {
	windingNumber := 0
	for i := 0; i < len(polygon); i++ {
		v1 := polygon[i]
		v2 := polygon[(i+1)%len(polygon)]
		if v1.Y <= pt.Y {
			if v2.Y > pt.Y {
				if vectorProduct(v1, v2, pt) > 0 {
					windingNumber++
				}
			}
		} else {
			if v2.Y <= pt.Y {
				if vectorProduct(v1, v2, pt) < 0 {
					windingNumber--
				}
			}
		}
	}
	return windingNumber != 0
}

// FillPolygonNonZeroWinding fills a polygon using the non-zero winding rule
func FillPolygonNonZeroWinding(img *gocv.Mat, polygon []utils.Point) {
	minX, minY, maxX, maxY := polygon[0].X, polygon[0].Y, polygon[0].X, polygon[0].Y
	for _, pt := range polygon {
		if pt.X < minX {
			minX = pt.X
		}
		if pt.Y < minY {
			minY = pt.Y
		}
		if pt.X > maxX {
			maxX = pt.X
		}
		if pt.Y > maxY {
			maxY = pt.Y
		}
	}

	for y := minY; y <= maxY; y++ {
		for x := minX; x <= maxX; x++ {
			pt := utils.Point{X: x, Y: y}
			if isPointInPolygonNonZeroWinding(polygon, pt) {
				img.SetUCharAt(y, x, 255)
			}
		}
	}
}
