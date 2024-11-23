package main

import (
	"fmt"
	"os"
)

func drawPixelTest(file *os.File, p Point) {
	text := fmt.Sprintf("(%d, %d)\n", p.X, p.Y)
	file.WriteString(text)
}

func bresenhamLineAlgorithmTest(file *os.File, p1, p2 Point) {
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
			drawPixelTest(file, Point{x, y}) // Рисуем текущий пиксель
			x += dirx                        // Изменяем x на каждом шаге
			if err >= 0 {
				y += diry // Изменяем y, если ошибка >= 0
				err -= 2 * dx
			}
			err += 2 * dy // Увеличиваем ошибку после каждого шага
		}
	} else { // Линия ближе к вертикальной
		err := 2*dx - dy // Инициализация ошибки с поправкой на половину пикселя
		for i := 0; i <= dy; i++ {
			drawPixelTest(file, Point{x, y}) // Рисуем текущий пиксель
			y += diry                        // Изменяем y на каждом шаге
			if err >= 0 {
				x += dirx // Изменяем x, если ошибка >= 0
				err -= 2 * dy
			}
			err += 2 * dx // Увеличиваем ошибку после каждого шага
		}
	}
}

func DrawlineTest() {
	// New file for test result
	file, err := os.Create("line_test.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	file.WriteString("Линия от (0, 0) до (8, 3):\n")
	bresenhamLineAlgorithmTest(file, Point{0, 0}, Point{8, 3})

	file.WriteString("Линия от (8, 3) до (0, 0):\n")
	bresenhamLineAlgorithmTest(file, Point{8, 3}, Point{0, 0})
}
