package lab5

import (
	"Computer_Graphics/cmd/labs/utils"
	"gocv.io/x/gocv"
	"math"
)

func drawPixel(img *gocv.Mat, p utils.Point, intensity float64) {
	offset := img.Rows() / 2
	img.SetUCharAt(p.Y+offset, p.X+offset, uint8(intensity))
	//fmt.Printf("(%d, %d)\n", p.X, p.Y)
}

// calculateLighting calculates the color (intensity) at a vertex using a simple Lambertian model
func calculateLighting(vertex, normal, lightPos utils.Point3DFloat, intensity float64) float64 {
	// Normalize the normal and light direction

	lightDir := normalize(utils.Point3DFloat{
		X: lightPos.X - vertex.X,
		Y: lightPos.Y - vertex.Y,
		Z: lightPos.Z - vertex.Z,
	})

	// Lambertian diffuse lighting
	dot := normal.X*lightDir.X + normal.Y*lightDir.Y + normal.Z*lightDir.Z
	if dot < 0 {
		dot = 0
	}

	lighting := dot * intensity
	return lighting
}

// normalize normalizes a vector
func normalize(v utils.Point3DFloat) utils.Point3DFloat {
	length := math.Sqrt(v.X*v.X + v.Y*v.Y + v.Z*v.Z)
	if length == 0 {
		return utils.Point3DFloat{}
	}
	return utils.Point3DFloat{X: v.X / length, Y: v.Y / length, Z: v.Z / length}
}
func contains(slice []int, element int) bool {
	for _, v := range slice {
		if v == element {
			return true
		}
	}
	return false
}
func vertexNormal(vertexIndex int, faces [][]int, transformedPoints []utils.Point3DFloat) utils.Point3DFloat {
	var normal utils.Point3DFloat
	var num float64
	for _, face := range faces {
		if contains(face, vertexIndex) {
			v1 := transformedPoints[face[0]]
			v2 := transformedPoints[face[1]]
			v3 := transformedPoints[face[2]]

			faceNormal := getNormal(v1, v2, v3)

			normal.X += faceNormal.X
			normal.Y += faceNormal.Y
			normal.Z += faceNormal.Z

			num++
		}
	}
	normal.X = normal.X / num
	normal.Y = normal.Y / num
	normal.Z = normal.Z / num

	normal = normalize(normal)
	return normal
}
func sortTriangle(TIntensity []float64, TCoordinates []utils.PointFloat) ([]float64, []utils.PointFloat) {
	for i := 0; i < len(TCoordinates); i++ {
		for j := 1; j < len(TCoordinates); j++ {
			if TCoordinates[j-1].Y > TCoordinates[j].Y {
				TCoordinates[j].Y, TCoordinates[j-1].Y = TCoordinates[j-1].Y, TCoordinates[j].Y
				TCoordinates[j].X, TCoordinates[j-1].X = TCoordinates[j-1].X, TCoordinates[j].X

				TIntensity[j], TIntensity[j-1] = TIntensity[j-1], TIntensity[j]
			}
		}
	}
	return TIntensity, TCoordinates
}

func scanLineMethod2T(img *gocv.Mat, yCurr, t1, t2, xIntersect1, xIntersect2 float64, TIntensity []float64) {
	var xLeft, xRight, intensityLeft, intensityRight float64
	//Sorting
	if xIntersect1 > xIntersect2 {
		xLeft, xRight = xIntersect2, xIntersect1
		intensityLeft = TIntensity[0] + (TIntensity[2]-TIntensity[0])*t2
		intensityRight = TIntensity[0] + (TIntensity[1]-TIntensity[0])*t1
	} else {
		xLeft, xRight = xIntersect1, xIntersect2
		intensityLeft = TIntensity[0] + (TIntensity[1]-TIntensity[0])*t1
		intensityRight = TIntensity[0] + (TIntensity[2]-TIntensity[0])*t2
	}

	for xCurr := xLeft; xCurr <= xRight; xCurr++ {
		if xRight-xLeft == 0 {
			point := utils.Point{X: int(xCurr), Y: int(yCurr)}
			drawPixel(img, point, intensityLeft)
			continue
		}
		t := (xCurr - xLeft) / (xRight - xLeft)
		currIntensity := intensityLeft + (intensityRight-intensityLeft)*t

		point := utils.Point{X: int(xCurr), Y: int(yCurr)}
		drawPixel(img, point, currIntensity)
	}
}
func scanLineMethod1T(img *gocv.Mat, yCurr, t1, xIntersect1, xIntersect2 float64, TIntensity []float64) {
	var xLeft, xRight, intensityLeft, intensityRight float64
	//Sorting
	if xIntersect1 > xIntersect2 {
		xLeft, xRight = xIntersect2, xIntersect1
		intensityLeft = TIntensity[0] + (TIntensity[2]-TIntensity[0])*t1
		intensityRight = TIntensity[0] + (TIntensity[1]-TIntensity[0])*t1
	} else {
		xLeft, xRight = xIntersect1, xIntersect2
		intensityLeft = TIntensity[0] + (TIntensity[1]-TIntensity[0])*t1
		intensityRight = TIntensity[0] + (TIntensity[2]-TIntensity[0])*t1
	}

	for xCurr := xLeft; xCurr <= xRight; xCurr++ {
		if xRight-xLeft == 0 {
			point := utils.Point{X: int(xCurr), Y: int(yCurr)}
			drawPixel(img, point, intensityLeft)
			continue
		}
		t := (xCurr - xLeft) / (xRight - xLeft)
		currIntensity := intensityLeft + (intensityRight-intensityLeft)*t

		point := utils.Point{X: int(xCurr), Y: int(yCurr)}
		drawPixel(img, point, currIntensity)
	}
}
func paintTriangle(img *gocv.Mat, TIntensity []float64, TCoordinates []utils.PointFloat) {
	//fmt.Println(TIntensity, TCoordinates[0], TCoordinates[1], TCoordinates[2])
	TIntensity, TCoordinates = sortTriangle(TIntensity, TCoordinates)
	//fmt.Println(TIntensity, TCoordinates[0], TCoordinates[1], TCoordinates[2])

	//Проверка на плоский треугольник
	if !(TCoordinates[0].Y == TCoordinates[2].Y) {
		//Проверка можно ли разделить треугольник на два
		if !(TCoordinates[0].Y == TCoordinates[1].Y || TCoordinates[1].Y == TCoordinates[2].Y) {

			for yCurr := TCoordinates[0].Y; yCurr <= TCoordinates[1].Y; yCurr++ {
				//Находим t для ребер 12 и 13
				t12 := (yCurr - TCoordinates[0].Y) / (TCoordinates[1].Y - TCoordinates[0].Y)
				t13 := (yCurr - TCoordinates[0].Y) / (TCoordinates[2].Y - TCoordinates[0].Y)

				//Находим пересечение по Х для сторон 12 и 13
				xIntersect12 := TCoordinates[0].X + t12*(TCoordinates[1].X-TCoordinates[0].X)
				xIntersect13 := TCoordinates[0].X + t13*(TCoordinates[2].X-TCoordinates[0].X)

				scanLineMethod2T(img, yCurr, t12, t13, xIntersect12, xIntersect13, TIntensity)

			}
			for yCurr := TCoordinates[1].Y; yCurr <= TCoordinates[2].Y; yCurr++ {
				//Находим t для ребер 12 и 13
				t23 := (yCurr - TCoordinates[1].Y) / (TCoordinates[2].Y - TCoordinates[1].Y)
				t13 := (yCurr - TCoordinates[0].Y) / (TCoordinates[2].Y - TCoordinates[0].Y)

				//Находим пересечение по Х для сторон 12 и 13
				xIntersect23 := TCoordinates[1].X + t23*(TCoordinates[2].X-TCoordinates[1].X)
				xIntersect13 := TCoordinates[0].X + t13*(TCoordinates[2].X-TCoordinates[0].X)

				scanLineMethod2T(img, yCurr, t23, t13, xIntersect23, xIntersect13, TIntensity)

			}
		} else {
			if TCoordinates[0].Y == TCoordinates[1].Y {
				for yCurr := TCoordinates[0].Y; yCurr <= TCoordinates[1].Y; yCurr++ {
					//Находим t для ребер 13
					t13 := (yCurr - TCoordinates[0].Y) / (TCoordinates[2].Y - TCoordinates[0].Y)

					//Находим пересечение по Х для сторон 12 и 13
					xIntersect12 := TCoordinates[0].X + t13*(TCoordinates[1].X-TCoordinates[0].X)
					xIntersect13 := TCoordinates[0].X + t13*(TCoordinates[2].X-TCoordinates[0].X)

					scanLineMethod1T(img, yCurr, t13, xIntersect12, xIntersect13, TIntensity)
				}

			} else {
				for yCurr := TCoordinates[0].Y; yCurr <= TCoordinates[1].Y; yCurr++ {
					//Находим t для ребер 13
					t13 := (yCurr - TCoordinates[0].Y) / (TCoordinates[2].Y - TCoordinates[0].Y)

					//Находим пересечение по Х для сторон 12 и 13
					xIntersect12 := TCoordinates[0].X + t13*(TCoordinates[1].X-TCoordinates[0].X)
					xIntersect13 := TCoordinates[0].X + t13*(TCoordinates[2].X-TCoordinates[0].X)

					scanLineMethod1T(img, yCurr, t13, xIntersect12, xIntersect13, TIntensity)
				}
			}
		}
	}
}
func paintFace(img *gocv.Mat, vIntensity []float64, vCoordinates2D []utils.PointFloat) {
	//Separating face on two triangles
	firstTCoordinates := vCoordinates2D[:3]
	firstTIntensity := vIntensity[:3]
	//fmt.Println(vCoordinates2D)
	paintTriangle(img, firstTIntensity, firstTCoordinates)

	secondTCoordinates := vCoordinates2D[1:]
	secondTIntensity := vIntensity[1:]

	paintTriangle(img, secondTIntensity, secondTCoordinates)
}
