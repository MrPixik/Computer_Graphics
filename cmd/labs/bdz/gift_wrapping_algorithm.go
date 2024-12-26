package bdz

import (
	"Computer_Graphics/cmd/labs/utils"
	"math"
)

func dotProduct(p1, p2, p3 utils.Point) float64 {
	aX := p2.X - p1.X
	aY := p2.Y - p1.Y

	bX := p3.X - p2.X
	bY := p3.Y - p2.Y

	product := aX*bX + aY*bY
	return float64(product)
}

func crossProduct(p1, p2, p3 utils.Point) float64 {
	aX := p2.X - p1.X
	aY := p2.Y - p1.Y

	bX := p3.X - p2.X
	bY := p3.Y - p2.Y

	product := aX*bY - aY*bX

	return float64(product)
}
func vectorMagnitude(p1, p2 utils.Point) float64 {
	aX := p2.X - p1.X
	aY := p2.Y - p1.Y

	return math.Sqrt(float64(aX*aX + aY*aY))
}

func calculateAngle(p1, p2, p3 utils.Point) float64 {
	dotPr := dotProduct(p1, p2, p3)

	magnitudeA := vectorMagnitude(p1, p2)
	magnitudeB := vectorMagnitude(p2, p3)

	cosTheta := dotPr / (magnitudeA * magnitudeB)
	theta := math.Acos(cosTheta)

	crossPr := crossProduct(p1, p2, p3)
	if crossPr > 0 {
		theta = 2*math.Pi - theta
	}
	return theta
}

func maxAngleIndex(prevPoint, currPoint utils.Point, points []utils.Point) int {
	maxPolarAngle := 0.0
	index := 0
	for i := 0; i < len(points); i++ {
		currAngle := calculateAngle(prevPoint, currPoint, points[i])
		if currAngle > maxPolarAngle {
			maxPolarAngle = currAngle
			index = i
		}
	}
	return index
}

func JarvisAlgorithm(originalPolygon []utils.Point) []utils.Point {
	yMax := originalPolygon[0].Y
	firstPointIndex := 0
	// Начинаем отчет с точки с минимальным значением Y
	for i := 0; i < len(originalPolygon); i++ {
		if originalPolygon[i].Y > yMax {
			yMax = originalPolygon[i].Y
			firstPointIndex = i
			// Если точек с минимальным значением Y несколько, то берем самую правую
		} else if originalPolygon[i].Y == yMax {
			if originalPolygon[i].X > originalPolygon[firstPointIndex].X {
				firstPointIndex = i
			}
		}
	}
	firstPoint := originalPolygon[firstPointIndex]
	prevPoint := utils.Point{X: firstPoint.X - 1, Y: firstPoint.Y}

	convexHull := make([]utils.Point, 1)
	convexHull[0] = firstPoint

	currIndex := firstPointIndex
	vertexNum := 1
	currPoint := originalPolygon[currIndex]
	for {
		currIndex = maxAngleIndex(prevPoint, currPoint, originalPolygon)
		prevPoint = currPoint
		currPoint = originalPolygon[currIndex]
		if (originalPolygon[currIndex].X == firstPoint.X) && (originalPolygon[currIndex].Y == firstPoint.Y) {
			return convexHull
		}
		//convexHull[vertexNum] = originalPolygon[currIndex]
		convexHull = append(convexHull, currPoint)
		vertexNum++

		originalPolygon = removePointUnordered(originalPolygon, currIndex)

	}

}
func removePointUnordered(originalPolygon []utils.Point, index int) []utils.Point {
	originalPolygon[index] = originalPolygon[len(originalPolygon)-1]
	return originalPolygon[:len(originalPolygon)-1]
}
