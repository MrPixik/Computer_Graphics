package lab5

import (
	"Computer_Graphics/cmd/labs/utils"
	"gocv.io/x/gocv"
	"time"
)

func firstPart() {
	var outputFilename = "..\\..\\static\\images\\lab5\\parallel_projection.png"
	img := gocv.NewMatWithSize(500, 500, gocv.MatTypeCV8U)
	defer img.Close()

	// Vertices of a parallelepiped
	parallelepiped := [8]utils.Point3DFloat{
		{-1, -1, -1}, {1, -1, -1}, {1, 1, -1}, {-1, 1, -1},
		{-1, -1, 1}, {1, -1, 1}, {1, 1, 1}, {-1, 1, 1},
	}

	// Rotation options
	angleX, angleY, angleZ := 30.0, 45.0, 60.0

	// Project the parallelepiped and scale for rendering
	scale := 100.0
	offset := float64(img.Rows() / 2)

	projVertices := make([]utils.PointFloat, len(parallelepiped))
	for i, v := range parallelepiped {
		point := rotatePoint(v, angleX, angleY, angleZ)
		projVertices[i] = utils.PointFloat{X: point.X*scale + offset, Y: point.Y*scale + offset}
	}

	// Connections between parallelepiped
	edges := [][2]int{
		{0, 1}, {1, 2}, {2, 3}, {3, 0}, // Bottom face
		{4, 5}, {5, 6}, {6, 7}, {7, 4}, // Top face
		{0, 4}, {1, 5}, {2, 6}, {3, 7}, // Vertical edges
	}

	// Draw the edges
	for _, edge := range edges {
		p1, p2 := projVertices[edge[0]], projVertices[edge[1]]

		utils.BresenhamLineAlgorithmFloat(&img, p1, p2)
	}

	gocv.IMWrite(outputFilename, img)
}

func secondPart() {
	var outputFilename = "..\\..\\static\\images\\lab5\\perspective_projection.png"
	img := gocv.NewMatWithSize(500, 500, gocv.MatTypeCV8U)
	defer img.Close()

	//Vertices of a parallelepiped
	parallelepiped := [8]utils.Point3DFloat{
		{-1, -1, -1}, {1, -1, -1}, {1, 1, -1}, {-1, 1, -1},
		{-1, -1, 1}, {1, -1, 1}, {1, 1, 1}, {-1, 1, 1},
	}

	// Project the parallelepiped and scale for rendering
	scale := 100.0
	offset := float64(img.Rows() / 2)

	projVertices := make([]utils.PointFloat, len(parallelepiped))
	for i, v := range parallelepiped {
		point := PerspectiveProjectionOZ(v, 2)
		projVertices[i] = utils.PointFloat{X: point.X*scale + offset, Y: point.Y*scale + offset}
	}

	// Connections between parallelepiped
	edges := [][2]int{
		{0, 1}, {1, 2}, {2, 3}, {3, 0}, // Bottom face
		{4, 5}, {5, 6}, {6, 7}, {7, 4}, // Top face
		{0, 4}, {1, 5}, {2, 6}, {3, 7}, // Vertical edges
	}

	// Draw the edges
	for _, edge := range edges {
		p1, p2 := projVertices[edge[0]], projVertices[edge[1]]

		utils.BresenhamLineAlgorithmFloat(&img, p1, p2)
	}

	gocv.IMWrite(outputFilename, img)
}

func thirdPart() {
	//Vertices of a parallelepiped
	parallelepiped := []utils.Point3DFloat{
		{X: -0.5, Y: -0.5, Z: 1}, {X: 0.5, Y: -0.5, Z: 1}, {X: 0.5, Y: 0.5, Z: 1}, {X: -0.5, Y: 0.5, Z: 1}, // Top face
		{X: -1, Y: -1, Z: -1}, {X: 1, Y: -1, Z: -1}, {X: 1, Y: 1, Z: -1}, {X: -1, Y: 1, Z: -1}, // Bottom face
	}

	//Faces of parallelepiped
	faces := [][]int{
		{0, 3, 2, 1}, // Bottom edge
		{4, 5, 6, 7}, // Top edge
		{0, 1, 5, 4}, // Front edge
		{1, 2, 6, 5}, // Right edge
		{2, 3, 7, 6}, // Back edge
		{3, 0, 4, 7}, // Left edge
	}

	viewDirection := utils.Point3DFloat{X: 0, Y: 0, Z: 1}
	projectionScale := 2.0

	img := gocv.NewMatWithSize(500, 500, gocv.MatTypeCV8U)
	defer img.Close()

	// Removing invisible edges
	for _, face := range faces {
		normal := getNormal(
			parallelepiped[face[0]],
			parallelepiped[face[1]],
			parallelepiped[face[2]],
		)

		if !isFrontFace(normal, viewDirection) {
			continue
		}

		for i := 0; i < len(face); i++ {
			p1 := parallelepiped[face[i]]
			p2 := parallelepiped[face[(i+1)%len(face)]]
			drawEdge(&img, p1, p2, projectionScale)
		}
	}

	var outputFilename = "..\\..\\static\\images\\lab5\\deleted_edges.png"
	gocv.IMWrite(outputFilename, img)
}

func fourthPart() {

	img := gocv.NewMatWithSize(600, 800, gocv.MatTypeCV8U)
	defer img.Close()

	//Vertices of a parallelepiped
	//parallelepiped := []utils.Point3DFloat{
	//	{X: -0.5, Y: -0.5, Z: 1}, {X: 0.5, Y: -0.5, Z: 1}, {X: 0.5, Y: 0.5, Z: 1}, {X: -0.5, Y: 0.5, Z: 1}, // Top face
	//	{X: -1, Y: -1, Z: -1}, {X: 1, Y: -1, Z: -1}, {X: 1, Y: 1, Z: -1}, {X: -1, Y: 1, Z: -1}, //Bottom face
	//}
	parallelepiped := []utils.Point3DFloat{
		{X: -0.5, Y: -0.5, Z: -0.5}, // Вершина 0
		{X: 0.5, Y: -0.5, Z: -0.5},  // Вершина 1
		{X: 0.5, Y: 0.5, Z: -0.5},   // Вершина 2
		{X: -0.5, Y: 0.5, Z: -0.5},  // Вершина 3
		{X: -0.5, Y: -0.5, Z: 0.5},  // Вершина 4
		{X: 0.5, Y: -0.5, Z: 0.5},   // Вершина 5
		{X: 0.5, Y: 0.5, Z: 0.5},    // Вершина 6
		{X: -0.5, Y: 0.5, Z: 0.5},   // Вершина 7
	}

	//Faces of parallelepiped
	faces := [][]int{
		{0, 3, 2, 1}, // Bottom edge
		{4, 5, 6, 7}, // Top edge
		{0, 1, 5, 4}, // Front edge
		{1, 2, 6, 5}, // Right edge
		{2, 3, 7, 6}, // Back edge
		{3, 0, 4, 7}, // Left edge
	}

	axis := utils.Point3DFloat{X: 1, Y: 1, Z: 0} // Axis of rotation
	angle := 0.0                                 // Angle of rotation

	// Project the parallelepiped and scale for rendering
	scale := 100.0
	offset := float64(img.Rows() / 2)

	viewDirection := utils.Point3DFloat{X: 0, Y: 0, Z: 1}

	writer, err := gocv.VideoWriterFile("..\\..\\static\\images\\lab5\\deleted_edges_cube.mp4", "mp4v", 30, img.Cols(), img.Rows(), false)
	if err != nil {
		panic(err)
	}
	defer writer.Close()

	for i := 0; i < 200; i++ {
		img.SetTo(gocv.Scalar{Val1: 0, Val2: 0, Val3: 0, Val4: 0}) // Image cleaning

		// Rotation and projection
		transformedPoints := make([]utils.Point3DFloat, len(parallelepiped))
		for i, vertex := range parallelepiped {
			rotated := rotatePointAroundAxis(vertex, axis, angle)
			transformedPoints[i] = rotated
		}
		for _, face := range faces {
			normal := getNormal(
				transformedPoints[face[0]],
				transformedPoints[face[1]],
				transformedPoints[face[2]],
			)

			if !isFrontFace(normal, viewDirection) {
				continue
			}

			for i := 0; i < len(face); i++ {
				p1 := transformedPoints[face[i]]
				p2 := transformedPoints[face[(i+1)%len(face)]]

				projectedP1 := PerspectiveProjectionOZ(p1, 500)
				projectedP2 := PerspectiveProjectionOZ(p2, 500)

				projectedP1.X = projectedP1.X*scale + offset
				projectedP1.Y = projectedP1.Y*scale + offset

				projectedP2.X = projectedP2.X*scale + offset
				projectedP2.Y = projectedP2.Y*scale + offset

				utils.BresenhamLineAlgorithmFloat(&img, projectedP1, projectedP2)
			}
		}

		// Recording a frame in video
		writer.Write(img)

		angle += 0.02 // Changing the angle for animation
		time.Sleep(10 * time.Millisecond)
	}
}

func extra() {
	img := gocv.NewMatWithSize(600, 800, gocv.MatTypeCV8U)
	defer img.Close()

	// Vertices of a parallelepiped
	parallelepiped := []utils.Point3DFloat{
		{X: -50, Y: -50, Z: -50}, // Вершина 0
		{X: 50, Y: -50, Z: -50},  // Вершина 1
		{X: 50, Y: 50, Z: -50},   // Вершина 2
		{X: -50, Y: 50, Z: -50},  // Вершина 3
		{X: -50, Y: -50, Z: 50},  // Вершина 4
		{X: 50, Y: -50, Z: 50},   // Вершина 5
		{X: 50, Y: 50, Z: 50},    // Вершина 6
		{X: -50, Y: 50, Z: 50},   // Вершина 7
	}

	//Faces of parallelepiped
	faces := [][]int{
		{0, 3, 2, 1}, // Bottom edge
		{4, 5, 6, 7}, // Top edge
		{0, 1, 5, 4}, // Front edge
		{1, 2, 6, 5}, // Right edge
		{2, 3, 7, 6}, // Back edge
		{3, 0, 4, 7}, // Left edge
	}

	axis := utils.Point3DFloat{X: 1, Y: 1, Z: 0} // Axis of rotation
	angle := 0.0                                 // Angle of rotation

	// Project the parallelepiped and scale for rendering
	//

	viewDirection := utils.Point3DFloat{X: 0, Y: 0, Z: 1}

	//Light position and light source's intensity
	lightPos := utils.Point3DFloat{X: 80, Y: 80, Z: 80}
	lightIntensity := 254.0

	writer, err := gocv.VideoWriterFile("..\\..\\static\\images\\lab5\\colored_cube.mp4", "mp4v", 30, img.Cols(), img.Rows(), false)
	if err != nil {
		panic(err)
	}
	defer writer.Close()

	for i := 0; i < 200; i++ {
		img.SetTo(gocv.Scalar{Val1: 0, Val2: 0, Val3: 0, Val4: 0}) // Image cleaning

		// Rotation and projection
		transformedPoints := make([]utils.Point3DFloat, len(parallelepiped))
		for i, vertex := range parallelepiped {
			rotated := rotatePointAroundAxis(vertex, axis, angle)
			transformedPoints[i] = rotated
		}
		for _, face := range faces {
			normal := getNormal(
				transformedPoints[face[0]],
				transformedPoints[face[1]],
				transformedPoints[face[2]],
			)

			if !isFrontFace(normal, viewDirection) {
				continue
			}

			vIntensity := make([]float64, len(face))              // Массив интенсивностей света в гранях для текущей стороны
			vCoordinates2D := make([]utils.PointFloat, len(face)) //Массив координат проекций точек на плоскость
			for i := 0; i < len(face); i++ {

				currVNormal := vertexNormal(i, faces, transformedPoints)

				vIntensity[i] = calculateLighting(transformedPoints[face[i]], currVNormal, lightPos, lightIntensity)

				p := transformedPoints[face[i]]
				projectedPoint := PerspectiveProjectionOZ(p, 500)
				vCoordinates2D[i] = utils.PointFloat{X: projectedPoint.X, Y: projectedPoint.Y}
			}
			paintFace(&img, vIntensity, vCoordinates2D)

		}

		// Recording a frame in video
		writer.Write(img)

		angle += 0.02 // Changing the angle for animation
		time.Sleep(10 * time.Millisecond)
	}
}

func Run() {

	// 1
	//firstPart()

	// 2
	//secondPart()

	// 3
	//thirdPart()

	// 4
	//fourthPart()

	// Extra
	extra()
}
