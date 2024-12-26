package bdz

import (
	"Computer_Graphics/cmd/labs/utils"
	"gocv.io/x/gocv"
)

func hermitBasisFunctions(t float64) (float64, float64, float64, float64) {
	return 2*t*t*t - 3*t*t + 1, t*t*t - 2*t*t + t, -2*t*t*t + 3*t*t, t*t*t - t*t
}
func calculateCoordinate(p1, p2, v1, v2 utils.Point, t float64) utils.Point {
	h0, h1, h2, h3 := hermitBasisFunctions(t)
	x := h0*float64(p1.X) + h1*float64(v1.X) + h2*float64(p2.X) + h3*float64(v2.X)
	y := h0*float64(p1.Y) + h1*float64(v1.Y) + h2*float64(p2.Y) + h3*float64(v2.Y)
	return utils.Point{X: int(x), Y: int(y)}
}
func drawHermiteCurve(img *gocv.Mat, p1, p2, v1, v2 utils.Point) {
	t := 0.0

	for {
		if t > 1.0 {
			break
		}
		currPoint := calculateCoordinate(p1, p2, v1, v2, t)
		utils.DrawPixel(img, currPoint)
		t += 0.01
	}
}

func drawCompositeHermiteCurve(img *gocv.Mat, points []utils.Point, vectors []utils.Point) {
	for i := 1; i < len(points); i++ {
		drawHermiteCurve(img, points[i-1], points[i], vectors[i-1], vectors[i])
	}
}
