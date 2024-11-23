package main

import (
	"gocv.io/x/gocv"
)

func maxInt(x, y int) int {
	if x > y {
		return x
	}
	return y
}
func minInt(x, y int) int {
	if x < y {
		return x
	}
	return y
}

// CLPointType Types of points
type CLPointType int

const (
	LEFT CLPointType = iota
	RIGHT
	BEHIND
	ORIGIN
	DESTINATION
	BETWEEN
)

// IntersectType Types of intersections
type IntersectType int

const (
	PARALLEL IntersectType = iota
	SAME
	SKEW
	SKEW_NO_CROSS
	SKEW_CROSS
)

// Classify classifies a point relative to a segment
func Classify(x1, y1, x2, y2, x, y float64) CLPointType {
	ax := x2 - x1
	ay := y2 - y1
	bx := x - x1
	by := y - y1
	s := ax*by - bx*ay
	if s > 0 {
		return LEFT
	}
	if s < 0 {
		return RIGHT
	}
	if (ax*bx < 0) || (ay*by < 0) {
		return BEHIND
	}
	if (ax*ax + ay*ay) < (bx*bx + by*by) {
		return BEHIND
	}
	if x1 == x && y1 == y {
		return ORIGIN
	}
	if x2 == x && y2 == y {
		return DESTINATION
	}
	return BETWEEN
}

// Intersect checks whether two segments intersect
func Intersect(ax, ay, bx, by, cx, cy, dx, dy float64, t *float64) IntersectType {
	nx := dy - cy
	ny := cx - dx
	var typeCL CLPointType
	denom := nx*(bx-ax) + ny*(by-ay)
	if denom == 0 {
		typeCL = Classify(cx, cy, dx, dy, ax, ay)
		if typeCL == LEFT || typeCL == RIGHT {
			return PARALLEL
		} else {
			return SAME
		}
	}
	num := nx*(ax-cx) + ny*(ay-cy)
	*t = -num / denom
	return SKEW
}

// Cross checks for intersection of two segments and returns parameters of intersection
func Cross(ax, ay, bx, by, cx, cy, dx, dy float64, tab, tcd *float64) IntersectType {
	typeInt := Intersect(ax, ay, bx, by, cx, cy, dx, dy, tab)
	if typeInt == SAME || typeInt == PARALLEL {
		return typeInt
	}
	if *tab < 0 || *tab > 1 {
		return SKEW_NO_CROSS
	}
	Intersect(cx, cy, dx, dy, ax, ay, bx, by, tcd)
	if *tcd < 0 || *tcd > 1 {
		return SKEW_NO_CROSS
	}
	return SKEW_CROSS
}

// isPointInPolygonEvenOdd determines whether a point is inside a polygon using the even-odd rule
func isPointInPolygonEvenOdd(polygon []Point, pt Point) bool {
	intersections := 0
	var t1, t2 float64
	outsidePoint := Point{X: pt.X + 10000, Y: pt.Y}

	for i := 0; i < len(polygon); i++ {
		v1 := polygon[i]
		v2 := polygon[(i+1)%len(polygon)]

		intersectType := Cross(float64(v1.X), float64(v1.Y), float64(v2.X), float64(v2.Y), float64(pt.X), float64(pt.Y), float64(outsidePoint.X), float64(outsidePoint.Y), &t1, &t2)
		if intersectType == SKEW_CROSS {
			intersections++
		}
	}
	return (intersections % 2) == 1
}

// isPointInPolygonEvenOdd determines whether a point is inside a polygon using the even-odd rule
func fillPolygonEvenOdd(img *gocv.Mat, polygon []Point) {
	minX, minY, maxX, maxY := polygon[0].X, polygon[0].Y, polygon[0].X, polygon[0].Y
	for _, pt := range polygon {
		minX = minInt(minX, pt.X)
		minY = minInt(minY, pt.Y)
		maxX = maxInt(maxX, pt.X)
		maxY = maxInt(maxY, pt.Y)
	}

	for y := minY; y <= maxY; y++ {
		for x := minX; x <= maxX; x++ {
			pt := Point{X: x, Y: y}
			if isPointInPolygonEvenOdd(polygon, pt) {
				img.SetUCharAt(y, x, 255)
			}
		}
	}
}

// isPointInPolygonNonZeroWinding determines whether a point is inside a polygon using the non-zero winding rule
func isPointInPolygonNonZeroWinding(polygon []Point, pt Point) bool {
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

// fillPolygonNonZeroWinding fills a polygon using the non-zero winding rule
func fillPolygonNonZeroWinding(img *gocv.Mat, polygon []Point) {
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
			pt := Point{X: x, Y: y}
			if isPointInPolygonNonZeroWinding(polygon, pt) {
				img.SetUCharAt(y, x, 255)
			}
		}
	}
}
