package lab5

import (
	"Computer_Graphics/cmd/labs/utils"
	"gocv.io/x/gocv"
)

// getNormal calculates the normal to a plane given three points
func getNormal(p1, p2, p3 utils.Point3DFloat) utils.Point3DFloat {
	return utils.Point3DFloat{
		X: (p2.Y-p1.Y)*(p3.Z-p1.Z) - (p2.Z-p1.Z)*(p3.Y-p1.Y),
		Y: (p2.Z-p1.Z)*(p3.X-p1.X) - (p2.X-p1.X)*(p3.Z-p1.Z),
		Z: (p2.X-p1.X)*(p3.Y-p1.Y) - (p2.Y-p1.Y)*(p3.X-p1.X),
	}
}

// isFrontFace determines whether a face is visible
func isFrontFace(normal, viewDirection utils.Point3DFloat) bool {
	dotProduct := normal.X*viewDirection.X + normal.Y*viewDirection.Y + normal.Z*viewDirection.Z
	return dotProduct > 0
}

// drawEdge draws an edge between two points after perspective projection
func drawEdge(img *gocv.Mat, p1, p2 utils.Point3DFloat, projectionScale float64) {
	projP1 := PerspectiveProjectionOZ(p1, projectionScale)
	projP2 := PerspectiveProjectionOZ(p2, projectionScale)

	scale := 100.0
	offset := float64(img.Rows() / 2)

	point1 := utils.PointFloat{X: projP1.X*scale + offset, Y: projP1.Y*scale + offset}
	point2 := utils.PointFloat{X: projP2.X*scale + offset, Y: projP2.Y*scale + offset}
	utils.BresenhamLineAlgorithmFloat(img, point1, point2)
}
