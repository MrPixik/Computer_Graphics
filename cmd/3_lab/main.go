package main

import (
	"fmt"
	"gocv.io/x/gocv"
	"math"
)

func MyAbs(a int) int {
	return int(math.Abs(float64(a)))
}

type Point struct {
	X, Y int
}

func drawPixel(img *gocv.Mat, p Point) {
	img.SetUCharAt(p.Y, p.X, 255)
	//fmt.Printf("(%d, %d)\n", p.X, p.Y)
}

// BresenhamLineAlgorithm func realises Bresenham line algorithm taking into account diagonal, horizontal and vertical cases
func BresenhamLineAlgorithm(img *gocv.Mat, p1, p2 Point) {
	if p1.X > p2.X {
		p1, p2 = p2, p1
	}
	dx := MyAbs(p2.X - p1.X)
	dy := MyAbs(p2.Y - p1.Y)

	// Определение направления приращений по x и y
	var dirx, diry int
	if p2.X > p1.X {
		dirx = 1
	} else if p2.X < p1.X {
		dirx = -1
	}

	if p2.Y > p1.Y {
		diry = 1
	} else if p2.Y < p1.Y {
		diry = -1
	}

	x, y := p1.X, p1.Y // Начальная точка

	// Основная логика в зависимости от соотношения dx и dy
	if dx >= dy { // Линия ближе к горизонтальной
		err := 2*dy - dx // Инициализация ошибки с поправкой на половину пикселя
		for i := 0; i <= dx; i++ {
			drawPixel(img, Point{x, y}) // Рисуем текущий пиксель
			x += dirx                   // Изменяем x на каждом шаге
			if err >= 0 {
				y += diry // Изменяем y, если ошибка >= 0
				err -= 2 * dx
			}
			err += 2 * dy // Увеличиваем ошибку после каждого шага
		}
	} else { // Линия ближе к вертикальной
		err := 2*dx - dy // Инициализация ошибки с поправкой на половину пикселя
		for i := 0; i <= dy; i++ {
			drawPixel(img, Point{x, y}) // Рисуем текущий пиксель
			y += diry                   // Изменяем y на каждом шаге
			if err >= 0 {
				x += dirx // Изменяем x, если ошибка >= 0
				err -= 2 * dy
			}
			err += 2 * dx // Увеличиваем ошибку после каждого шага
		}
	}
}

// DrawPolygon func drawing polygon by points from slice of type Point
func DrawPolygon(img *gocv.Mat, points []Point) {
	for i := 1; i < len(points); i++ {
		BresenhamLineAlgorithm(img, points[i-1], points[i])
	}
	BresenhamLineAlgorithm(img, points[0], points[len(points)-1]) //connection of first and last points
}

// vectorProduct func calculate vector product of vectors (p2,p1) and (p2,p3)
func vectorProduct(p1, p2, p3 Point) int {
	return (p2.X-p1.X)*(p3.Y-p2.Y) - (p2.Y-p1.Y)*(p3.X-p2.X)
}

// isConvexPolygon checks if polygon, created by slice of type Points, convex. True if convex. False if non-convex.
func isConvexPolygon(points []Point) bool {
	n := len(points)
	if n < 3 {
		return false // Should be at least 3 points
	}

	sign := 0
	for i := 0; i < n; i++ {
		p1 := points[i]
		p2 := points[(i+1)%n]
		p3 := points[(i+2)%n]

		// Vector product for current sides
		cp := vectorProduct(p1, p2, p3)

		// Checking if sign of vector product changed
		if cp != 0 {
			if sign == 0 {
				sign = cp
			} else if (cp > 0) != (sign > 0) {
				// If changed => polygon is non-convex
				return false
			}
		}
	}

	return true // If don't changed => polygon is convex
}

// ifIntersect func calculate direction of rotation of two. True if intersecting. False if non-intersecting.
func ifIntersect(p1, p2, p3, p4 Point) bool {
	// Check direction of rotation
	orient := func(p, q, r Point) int {
		val := (q.Y-p.Y)*(r.X-q.X) - (q.X-p.X)*(r.Y-q.Y)
		if val == 0 {
			return 0 // collinear
		} else if val > 0 {
			return 1 // clockwise
		}
		return 2 // counterclockwise
	}

	o1 := orient(p1, p2, p3)
	o2 := orient(p1, p2, p4)
	o3 := orient(p3, p4, p1)
	o4 := orient(p3, p4, p2)

	// Conditions of intersecting non-collinear segments
	if o1 != o2 && o3 != o4 {
		return true
	}

	// func for checking if point C located on AB segment
	onSegment := func(a, c, b Point) bool {
		return c.X <= max(a.X, b.X) && c.X >= min(a.X, b.X) &&
			c.Y <= max(a.Y, b.Y) && c.Y >= min(a.Y, b.Y)
	}

	// Check for collinear cases
	if o1 == 0 && onSegment(p1, p3, p2) {
		return true
	}
	if o2 == 0 && onSegment(p1, p4, p2) {
		return true
	}
	if o3 == 0 && onSegment(p3, p1, p4) {
		return true
	}
	if o4 == 0 && onSegment(p3, p2, p4) {
		return true
	}

	return false
}

// isSelfIntersectingPolygon func checks if polygon is self intersected
func isSelfIntersectingPolygon(points []Point) bool {
	n := len(points)
	for i := 0; i < n; i++ {
		for j := i + 2; j < n; j++ {
			//Skip last polygon's side
			if i == 0 && j == n-1 {
				continue
			}
			if ifIntersect(points[i], points[(i+1)%n], points[j], points[(j+1)%n]) {
				return true
			}
		}
	}
	return false
}
func main() {

	//Line drawing
	var outputFilenameLine = "..\\..\\static\\images\\3_lab\\straight_line.png"
	imgLine := gocv.NewMatWithSize(250, 250, gocv.MatTypeCV8U)
	defer imgLine.Close()

	pointsLine := []Point{
		Point{50, 100},
		Point{150, 200},
	}
	BresenhamLineAlgorithm(&imgLine, pointsLine[0], pointsLine[1])
	gocv.IMWrite(outputFilenameLine, imgLine)

	// Test for segment (0, 0) (8, 3)
	DrawlineTest()
	//Polygon drawing
	var outputFilenamePolygonConvex = "..\\..\\static\\images\\3_lab\\PolygonConvex.png"
	var outputFilenamePolygonNonConvex = "..\\..\\static\\images\\3_lab\\PolygonNonConvex.png"
	var outputFilenamePolygonSelfIntersecting = "..\\..\\static\\images\\3_lab\\PolygonSelfIntersecting.png"
	var outputFilenamePolygonStar = "..\\..\\static\\images\\3_lab\\PolygonStar.png"
	var outputFilenamePolygonExample = "..\\..\\static\\images\\3_lab\\PolygonExample.png"

	imgPolygonConvex := gocv.NewMatWithSize(250, 250, gocv.MatTypeCV8U)
	imgPolygonNonConvex := gocv.NewMatWithSize(250, 250, gocv.MatTypeCV8U)
	imgPolygonSelfIntersecting := gocv.NewMatWithSize(250, 250, gocv.MatTypeCV8U)
	imgPolygonStar := gocv.NewMatWithSize(250, 250, gocv.MatTypeCV8U)
	imgPolygonExample := gocv.NewMatWithSize(250, 250, gocv.MatTypeCV8U)

	defer imgPolygonConvex.Close()
	defer imgPolygonNonConvex.Close()
	defer imgPolygonSelfIntersecting.Close()
	defer imgPolygonStar.Close()
	defer imgPolygonExample.Close()

	pointsPolygonConvex := []Point{
		{200, 230},
		{50, 220},
		{50, 100},
		{100, 50},
		{200, 50},
	}
	pointsPolygonNonConvex := []Point{
		{50, 50},
		{200, 50},
		{200, 150},
		{150, 100},
		{50, 150},
		{50, 100},
	}
	pointsPolygonSelfIntersecting := []Point{
		{50, 50},
		{200, 50},
		{200, 150},
		{150, 100},
		{50, 150},
		{100, 200},
	}
	pointsPolygonStar := []Point{
		{60, 220},
		{125, 20},
		{180, 220},
		{20, 100},
		{220, 100},
	}
	pointsPolygonExample := []Point{
		{150, 20},  // A
		{100, 220}, // B
		{50, 110},  // C
		{200, 60},  // D
		{240, 180}, // E
		{20, 180},  // F
		{160, 110}, // G
	}

	DrawPolygon(&imgPolygonConvex, pointsPolygonConvex)
	DrawPolygon(&imgPolygonNonConvex, pointsPolygonNonConvex)
	DrawPolygon(&imgPolygonSelfIntersecting, pointsPolygonSelfIntersecting)
	DrawPolygon(&imgPolygonStar, pointsPolygonStar)
	DrawPolygon(&imgPolygonExample, pointsPolygonExample)

	gocv.IMWrite(outputFilenamePolygonConvex, imgPolygonConvex)
	gocv.IMWrite(outputFilenamePolygonNonConvex, imgPolygonNonConvex)
	gocv.IMWrite(outputFilenamePolygonSelfIntersecting, imgPolygonSelfIntersecting)
	gocv.IMWrite(outputFilenamePolygonStar, imgPolygonStar)
	gocv.IMWrite(outputFilenamePolygonExample, imgPolygonExample)

	//Checking polygon if polygon convex and self-intersected
	//Convex
	fmt.Println("Convex polygon")
	if isConvexPolygon(pointsPolygonConvex) {
		fmt.Println("Polygon is convex")
	} else {
		fmt.Println("Polygon is non-convex")
	}
	if isSelfIntersectingPolygon(pointsPolygonConvex) {
		fmt.Println("Polygon is self-intersected")
	} else {
		fmt.Println("Polygon is non-self-intersected")
	}
	fmt.Println()
	//Non-convex
	fmt.Println("Non-convex polygon")
	if isConvexPolygon(pointsPolygonNonConvex) {
		fmt.Println("Polygon is convex")
	} else {
		fmt.Println("Polygon is non-convex")
	}
	if isSelfIntersectingPolygon(pointsPolygonNonConvex) {
		fmt.Println("Polygon is self-intersected")
	} else {
		fmt.Println("Polygon is non-self-intersected")
	}
	fmt.Println()
	//Self-intersected
	fmt.Println("Self-intersected polygon")
	if isConvexPolygon(pointsPolygonSelfIntersecting) {
		fmt.Println("Polygon is convex")
	} else {
		fmt.Println("Polygon is non-convex")
	}

	if isSelfIntersectingPolygon(pointsPolygonSelfIntersecting) {
		fmt.Println("Polygon is self-intersected")
	} else {
		fmt.Println("Polygon is non-self-intersected")
	}

	//Even-Odd method
	var outputFilenamePolygonStarPainted = "..\\..\\static\\images\\3_lab\\PolygonStarEvenOdd.png"
	var outputFilenamePolygonExamplePainted = "..\\..\\static\\images\\3_lab\\PolygonExampleEvenOdd.png"

	imgPolygonStarEvenOdd := gocv.NewMatWithSize(250, 250, gocv.MatTypeCV8U)
	imgPolygonExampleEvenOdd := gocv.NewMatWithSize(250, 250, gocv.MatTypeCV8U)

	defer imgPolygonStarEvenOdd.Close()
	defer imgPolygonExampleEvenOdd.Close()

	fillPolygonEvenOdd(&imgPolygonStarEvenOdd, pointsPolygonStar)
	fillPolygonEvenOdd(&imgPolygonExampleEvenOdd, pointsPolygonExample)

	gocv.IMWrite(outputFilenamePolygonStarPainted, imgPolygonStarEvenOdd)
	gocv.IMWrite(outputFilenamePolygonExamplePainted, imgPolygonExampleEvenOdd)

	//Non-Zero Winding method
	var outputFilenamePolygonStarNonZeroWinding = "..\\..\\static\\images\\3_lab\\PolygonStarNonZeroWinding.png"
	var outputFilenamePolygonExampleNonZeroWinding = "..\\..\\static\\images\\3_lab\\PolygonExampleNonZeroWinding.png"

	imgPolygonStarNonZeroWinding := gocv.NewMatWithSize(250, 250, gocv.MatTypeCV8U)
	imgPolygonExampleNonZeroWinding := gocv.NewMatWithSize(250, 250, gocv.MatTypeCV8U)

	defer imgPolygonStarNonZeroWinding.Close()
	defer imgPolygonExampleNonZeroWinding.Close()

	fillPolygonNonZeroWinding(&imgPolygonStarNonZeroWinding, pointsPolygonStar)
	fillPolygonNonZeroWinding(&imgPolygonExampleNonZeroWinding, pointsPolygonExample)

	gocv.IMWrite(outputFilenamePolygonStarNonZeroWinding, imgPolygonStarNonZeroWinding)
	gocv.IMWrite(outputFilenamePolygonExampleNonZeroWinding, imgPolygonExampleNonZeroWinding)
}
