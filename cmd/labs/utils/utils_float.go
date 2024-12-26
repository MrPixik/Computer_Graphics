package utils

import (
	"gocv.io/x/gocv"
	"math"
)

type Point3DFloat struct {
	X float64
	Y float64
	Z float64
}
type PointFloat struct {
	X float64
	Y float64
}

// SetWhitePixelFloat set white pixel cy coordinates in point
func SetWhitePixelFloat(img *gocv.Mat, point PointFloat) {
	//if point.X >= 0 && int(point.X) < img.Cols() && point.Y >= 0 && int(point.Y) < img.Rows() {
	img.SetUCharAt(int(point.Y), int(point.X), 255)
	//}
}

func SetPixelValueFloat(img *gocv.Mat, point PointFloat, value uint8) {
	x := int(point.X)
	y := int(point.Y)

	// Проверка границ матрицы
	if x >= 0 && x < img.Cols() && y >= 0 && y < img.Rows() {
		img.SetUCharAt(y, x, value)
	}
}

// BresenhamLineAlgorithmFloat realises Bresenham line algorithm taking into account diagonal, horizontal and vertical cases, using PointFloat
func BresenhamLineAlgorithmFloat(img *gocv.Mat, p1, p2 PointFloat) {
	if p1.X > p2.X {
		p1, p2 = p2, p1
	}
	dx := math.Abs(p2.X - p1.X)
	dy := math.Abs(p2.Y - p1.Y)

	// Определение направления приращений по x и y
	var dirX, dirY float64
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
		for i := 0; i <= int(dx); i++ {
			SetWhitePixelFloat(img, PointFloat{X: x, Y: y}) // Рисуем текущий пиксель
			x += dirX                                       // Изменяем x на каждом шаге
			if err >= 0 {
				y += dirY // Изменяем y, если ошибка >= 0
				err -= 2 * dx
			}
			err += 2 * dy // Увеличиваем ошибку после каждого шага
		}
	} else { // Линия ближе к вертикальной
		err := 2*dx - dy // Инициализация ошибки с поправкой на половину пикселя
		for i := 0; i <= int(dy); i++ {
			SetWhitePixelFloat(img, PointFloat{X: x, Y: y}) // Рисуем текущий пиксель
			y += dirY                                       // Изменяем y на каждом шаге
			if err >= 0 {
				x += dirX // Изменяем x, если ошибка >= 0
				err -= 2 * dy
			}
			err += 2 * dx // Увеличиваем ошибку после каждого шага
		}
	}
}
