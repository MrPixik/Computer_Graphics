# Computer Graphics Project

This project contains various implementations and algorithms related to computer graphics, organized into multiple labs. Each lab focuses on specific topics in computer graphics, such as line drawing, polygon filling, Bézier curves, and line clipping.

## Project structure
```bash
Computer_Graphics/
│
├── cmd/
│   ├── labs/
│   │   ├── lab1/                  # Lab 1 - Basic image processing
│   │   │   ├── blending.go        # Implementation of image blending
│   │   │   ├── gray_round_image.go # Creates grayscale rounded images
│   │   │   ├── main.go            # Main entry point for Lab 1
│   │   │
│   │   ├── lab2/                  # Lab 2 - Dithering techniques
│   │   │   ├── dithering.go       # Implements image dithering algorithms
│   │   │   ├── main.go            # Main entry point for Lab 2
│   │   │
│   │   ├── lab3/                  # Lab 3 - Line drawing and polygons
│   │   │   ├── line_draw_check.go # Validates custom line drawing algorithms
│   │   │   ├── polygon.go         # Polygon representation and utility functions
│   │   │   ├── polygon_filling.go # Implements polygon filling algorithms
│   │   │   ├── main.go            # Main entry point for Lab 3
│   │   │
│   │   ├── lab4/                  # Lab 4 - Advanced algorithms
│   │   │   ├── bezier_curve.go    # Bézier curve drawing algorithm
│   │   │   ├── cyrus_beck_algorithm.go # Cyrus-Beck line clipping algorithm
│   │   │   ├── main.go            # Main entry point for Lab 4
│   │   │
│   ├── utils/
│   │   ├── utils.go               # Shared utility functions and structures
│   │
│   ├── main/
│   │   ├── line_test.txt          # Test cases for line drawing algorithms
│   │   ├── main.go                # Main entry point for running all labs
│   │
├── static/
│   ├── images/                    # Images used for testing and visualization
│       ├── lab1/                 
│       ├── lab2/                  
│       ├── lab3/                  
│       ├── lab4/                  
│
├── go.mod                         # Go module definition
```

## Labs Overview
### Lab 1 - Basic Image Processing
Blending: Combines two images using blending algorithms.
Grayscale and Rounded Images: Converts images to grayscale and rounds their corners.
### Lab 2 - Dithering Techniques
Implements dithering algorithms for reducing image colors while maintaining visual quality.
### Lab 3 - Line Drawing and Polygons
Line Drawing: Custom line-drawing algorithm validation.
Polygon Filling: Implements algorithms to fill polygons using scan-line techniques.
### Lab 4 - Advanced Algorithms
Bézier Curves: Draws smooth curves using the Bézier curve algorithm.
Cyrus-Beck Clipping: Clips a line segment to a convex polygon using the Cyrus-Beck algorithm.

## Installation

To work with this project, you must install the GoCV library. Please see the installation documentation [there](https://gocv.io/).

## How to run

To start laboratory work you need to run the script cmd/main/main.go using this command:

```bash
go run cmd/main/main.go

