package lab5

import "Computer_Graphics/cmd/labs/utils"

// PerspectiveProjectionOZ performs a perspective projection of a three-dimensional point onto a plane parallel to the OZ axis
func PerspectiveProjectionOZ(point utils.Point3DFloat, k float64) utils.PointFloat {
	return utils.PointFloat{X: point.X * k / (k + point.Z), Y: point.Y * k / (k + point.Z)}
}
