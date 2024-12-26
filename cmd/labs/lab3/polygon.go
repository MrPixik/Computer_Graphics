package lab3

import (
	"Computer_Graphics/cmd/labs/utils"
	"gocv.io/x/gocv"
)

// BresenhamLineAlgorithm func realises Bresenham line algorithm taking into account diagonal, horizontal and vertical cases
func BresenhamLineAlgorithm(img *gocv.Mat, p1, p2 utils.Point) {
	if p1.X > p2.X {
		p1, p2 = p2, p1
	}
	dx := utils.MyAbs(p2.X - p1.X)
	dy := utils.MyAbs(p2.Y - p1.Y)

	// Определение направления приращений по x и y
	var dirX, dirY int
	if p2.X > p1.X {
		dirX = 1
	} else if p2.X < p1.X {
		dirX = -1
	}

	if p2.Y > p1.Y {
		dirY = 1
	} else if p2.Y < p1.Y {
		dirY = -1
	}

	x, y := p1.X, p1.Y // Начальная точка

	// Основная логика в зависимости от соотношения dx и dy
	if dx >= dy { // Линия ближе к горизонтальной
		err := 2*dy - dx // Инициализация ошибки с поправкой на половину пикселя
		for i := 0; i <= dx; i++ {
			utils.DrawPixel(img, utils.Point{X: x, Y: y}) // Рисуем текущий пиксель
			x += dirX                                     // Изменяем x на каждом шаге
			if err >= 0 {
				y += dirY // Изменяем y, если ошибка >= 0
				err -= 2 * dx
			}
			err += 2 * dy // Увеличиваем ошибку после каждого шага
		}
	} else { // Линия ближе к вертикальной
		err := 2*dx - dy // Инициализация ошибки с поправкой на половину пикселя
		for i := 0; i <= dy; i++ {
			utils.DrawPixel(img, utils.Point{X: x, Y: y}) // Рисуем текущий пиксель
			y += dirY                                     // Изменяем y на каждом шаге
			if err >= 0 {
				x += dirX // Изменяем x, если ошибка >= 0
				err -= 2 * dy
			}
			err += 2 * dx // Увеличиваем ошибку после каждого шага
		}
	}
}

// DrawPolygon func drawing polygon by points from slice of type Point
func DrawPolygon(img *gocv.Mat, points []utils.Point) {
	for i := 1; i < len(points); i++ {
		BresenhamLineAlgorithm(img, points[i-1], points[i])
	}
	BresenhamLineAlgorithm(img, points[0], points[len(points)-1]) //connection of first and last points
}

// vectorProduct func calculate vector product of vectors (p2,p1) and (p2,p3)
func vectorProduct(p1, p2, p3 utils.Point) int {
	return (p2.X-p1.X)*(p3.Y-p2.Y) - (p2.Y-p1.Y)*(p3.X-p2.X)
}

// isConvexPolygon checks if polygon, created by slice of type Points, convex. True if convex. False if non-convex.
func isConvexPolygon(points []utils.Point) bool {
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
func ifIntersect(p1, p2, p3, p4 utils.Point) bool {
	// Check direction of rotation
	orient := func(p, q, r utils.Point) int {
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
	onSegment := func(a, c, b utils.Point) bool {
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

// IsSelfIntersectingPolygon func checks if polygon is self intersected
func IsSelfIntersectingPolygon(points []utils.Point) bool {
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
