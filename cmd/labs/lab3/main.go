package lab3

import (
	"Computer_Graphics/cmd/labs/utils"
	"fmt"
	"gocv.io/x/gocv"
)

func Run() {

	//Line drawing
	var outputFilenameLine = "..\\..\\static\\images\\lab3\\straight_line.png"
	imgLine := gocv.NewMatWithSize(250, 250, gocv.MatTypeCV8U)
	defer imgLine.Close()

	pointsLine := []utils.Point{
		{50, 100},
		{150, 200},
	}
	BresenhamLineAlgorithm(&imgLine, pointsLine[0], pointsLine[1])
	gocv.IMWrite(outputFilenameLine, imgLine)

	// Test for segment (0, 0) (8, 3)
	DrawlineTest()
	//Polygon drawing
	var outputFilenamePolygonConvex = "..\\..\\static\\images\\lab3\\PolygonConvex.png"
	var outputFilenamePolygonNonConvex = "..\\..\\static\\images\\lab3\\PolygonNonConvex.png"
	var outputFilenamePolygonSelfIntersecting = "..\\..\\static\\images\\lab3\\PolygonSelfIntersecting.png"
	var outputFilenamePolygonStar = "..\\..\\static\\images\\lab3\\PolygonStar.png"
	var outputFilenamePolygonExample = "..\\..\\static\\images\\lab3\\PolygonExample.png"

	imgPolygonConvex := gocv.NewMatWithSize(250, 250, gocv.MatTypeCV8U)
	imgPolygonNonConvex := gocv.NewMatWithSize(250, 250, gocv.MatTypeCV8U)
	imgPolygonSelfIntersecting := gocv.NewMatWithSize(250, 250, gocv.MatTypeCV8U)
	imgPolygonStar := gocv.NewMatWithSize(250, 250, gocv.MatTypeCV8U)
	imgPolygonExample := gocv.NewMatWithSize(250, 250, gocv.MatTypeCV8U)

	defer imgPolygonConvex.Close()
	defer imgPolygonNonConvex.Close()
	defer imgPolygonSelfIntersecting.Close()
	defer imgPolygonStar.Close()
	defer imgPolygonExample.Close()

	pointsPolygonConvex := []utils.Point{
		{200, 230},
		{50, 220},
		{50, 100},
		{100, 50},
		{200, 50},
	}
	pointsPolygonNonConvex := []utils.Point{
		{50, 50},
		{200, 50},
		{200, 150},
		{150, 100},
		{50, 150},
		{50, 100},
	}
	pointsPolygonSelfIntersecting := []utils.Point{
		{50, 50},
		{200, 50},
		{200, 150},
		{150, 100},
		{50, 150},
		{100, 200},
	}
	pointsPolygonStar := []utils.Point{
		{60, 220},
		{125, 20},
		{180, 220},
		{20, 100},
		{220, 100},
	}
	pointsPolygonExample := []utils.Point{
		{150, 20},  // A
		{100, 220}, // B
		{50, 110},  // C
		{200, 60},  // D
		{240, 180}, // E
		{20, 180},  // F
		{160, 110}, // G
	}

	DrawPolygon(&imgPolygonConvex, pointsPolygonConvex)
	DrawPolygon(&imgPolygonNonConvex, pointsPolygonNonConvex)
	DrawPolygon(&imgPolygonSelfIntersecting, pointsPolygonSelfIntersecting)
	DrawPolygon(&imgPolygonStar, pointsPolygonStar)
	DrawPolygon(&imgPolygonExample, pointsPolygonExample)

	gocv.IMWrite(outputFilenamePolygonConvex, imgPolygonConvex)
	gocv.IMWrite(outputFilenamePolygonNonConvex, imgPolygonNonConvex)
	gocv.IMWrite(outputFilenamePolygonSelfIntersecting, imgPolygonSelfIntersecting)
	gocv.IMWrite(outputFilenamePolygonStar, imgPolygonStar)
	gocv.IMWrite(outputFilenamePolygonExample, imgPolygonExample)

	//Checking polygon if polygon convex and self-intersected
	//Convex
	fmt.Println("Convex polygon")
	if isConvexPolygon(pointsPolygonConvex) {
		fmt.Println("Polygon is convex")
	} else {
		fmt.Println("Polygon is non-convex")
	}
	if IsSelfIntersectingPolygon(pointsPolygonConvex) {
		fmt.Println("Polygon is self-intersected")
	} else {
		fmt.Println("Polygon is non-self-intersected")
	}
	fmt.Println()
	//Non-convex
	fmt.Println("Non-convex polygon")
	if isConvexPolygon(pointsPolygonNonConvex) {
		fmt.Println("Polygon is convex")
	} else {
		fmt.Println("Polygon is non-convex")
	}
	if IsSelfIntersectingPolygon(pointsPolygonNonConvex) {
		fmt.Println("Polygon is self-intersected")
	} else {
		fmt.Println("Polygon is non-self-intersected")
	}
	fmt.Println()
	//Self-intersected
	fmt.Println("Self-intersected polygon")
	if isConvexPolygon(pointsPolygonSelfIntersecting) {
		fmt.Println("Polygon is convex")
	} else {
		fmt.Println("Polygon is non-convex")
	}

	if IsSelfIntersectingPolygon(pointsPolygonSelfIntersecting) {
		fmt.Println("Polygon is self-intersected")
	} else {
		fmt.Println("Polygon is non-self-intersected")
	}

	//Even-Odd method
	var outputFilenamePolygonStarPainted = "..\\..\\static\\images\\lab3\\PolygonStarEvenOdd.png"
	var outputFilenamePolygonExamplePainted = "..\\..\\static\\images\\lab3\\PolygonExampleEvenOdd.png"

	imgPolygonStarEvenOdd := gocv.NewMatWithSize(250, 250, gocv.MatTypeCV8U)
	imgPolygonExampleEvenOdd := gocv.NewMatWithSize(250, 250, gocv.MatTypeCV8U)

	defer imgPolygonStarEvenOdd.Close()
	defer imgPolygonExampleEvenOdd.Close()

	FillPolygonEvenOdd(&imgPolygonStarEvenOdd, pointsPolygonStar)
	FillPolygonEvenOdd(&imgPolygonExampleEvenOdd, pointsPolygonExample)

	gocv.IMWrite(outputFilenamePolygonStarPainted, imgPolygonStarEvenOdd)
	gocv.IMWrite(outputFilenamePolygonExamplePainted, imgPolygonExampleEvenOdd)

	//Non-Zero Winding method
	var outputFilenamePolygonStarNonZeroWinding = "..\\..\\static\\images\\lab3\\PolygonStarNonZeroWinding.png"
	var outputFilenamePolygonExampleNonZeroWinding = "..\\..\\static\\images\\lab3\\PolygonExampleNonZeroWinding.png"

	imgPolygonStarNonZeroWinding := gocv.NewMatWithSize(250, 250, gocv.MatTypeCV8U)
	imgPolygonExampleNonZeroWinding := gocv.NewMatWithSize(250, 250, gocv.MatTypeCV8U)

	defer imgPolygonStarNonZeroWinding.Close()
	defer imgPolygonExampleNonZeroWinding.Close()

	FillPolygonNonZeroWinding(&imgPolygonStarNonZeroWinding, pointsPolygonStar)
	FillPolygonNonZeroWinding(&imgPolygonExampleNonZeroWinding, pointsPolygonExample)

	gocv.IMWrite(outputFilenamePolygonStarNonZeroWinding, imgPolygonStarNonZeroWinding)
	gocv.IMWrite(outputFilenamePolygonExampleNonZeroWinding, imgPolygonExampleNonZeroWinding)
}
