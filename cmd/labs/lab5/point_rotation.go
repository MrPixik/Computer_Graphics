package lab5

import (
	"Computer_Graphics/cmd/labs/utils"
	"math"
)

// rotatePoint performs rotation of a point around the X, Y and Z axes
func rotatePoint(point utils.Point3DFloat, angleX, angleY, angleZ float64) utils.Point3DFloat {
	// Углы поворота переводятся в радианы
	angleX, angleY, angleZ = angleX*math.Pi/180, angleY*math.Pi/180, angleZ*math.Pi/180

	// Вращение вокруг X
	y1 := point.Y*math.Cos(angleX) - point.Z*math.Sin(angleX)
	z1 := point.Y*math.Sin(angleX) + point.Z*math.Cos(angleX)
	point.Y, point.Z = y1, z1

	// Вращение вокруг Y
	x1 := point.X*math.Cos(angleY) + point.Z*math.Sin(angleY)
	z1 = -point.X*math.Sin(angleY) + point.Z*math.Cos(angleY)
	point.X, point.Z = x1, z1

	// Вращение вокруг Z
	x1 = point.X*math.Cos(angleZ) - point.Y*math.Sin(angleZ)
	y1 = point.X*math.Sin(angleZ) + point.Y*math.Cos(angleZ)

	return utils.Point3DFloat{X: x1, Y: y1, Z: z1}
}

// rotatePointAroundAxis performs rotation of a point around an arbitrary axis in three-dimensional space.
func rotatePointAroundAxis(point utils.Point3DFloat, axis utils.Point3DFloat, angle float64) utils.Point3DFloat {
	// Нормализация оси вращения
	length := math.Sqrt(axis.X*axis.X + axis.Y*axis.Y + axis.Z*axis.Z)
	axis.X, axis.Y, axis.Z = axis.X/length, axis.Y/length, axis.Z/length

	// Углы для вращения
	cosTheta := math.Cos(angle)
	sinTheta := math.Sin(angle)
	dot := axis.X*point.X + axis.Y*point.Y + axis.Z*point.Z

	// Формула вращения
	x := point.X*cosTheta + (1-cosTheta)*dot*axis.X + sinTheta*(-axis.Z*point.Y+axis.Y*point.Z)
	y := point.Y*cosTheta + (1-cosTheta)*dot*axis.Y + sinTheta*(axis.Z*point.X-axis.X*point.Z)
	z := point.Z*cosTheta + (1-cosTheta)*dot*axis.Z + sinTheta*(-axis.Y*point.X+axis.X*point.Y)

	return utils.Point3DFloat{X: x, Y: y, Z: z}
}
