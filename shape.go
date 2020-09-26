package physics

import (
	"image"
	"image/color"
	"image/draw"
	"os"
)

// type Shape
type Shape interface {
	Draw(draw.Image)
	Overlaps(Shape) bool
}

// type Rectangle
type Rectangle struct {
	Bounds image.Rectangle // rectangle bounds
	Color  color.RGBA      // color
}

// Draw() satisfies the Draw method on interface Shape
// it draws, pixel by pixel, a rectangle
func (rectangle Rectangle) Draw(frame draw.Image) {
	// variables
	var x int = rectangle.Bounds.Min.X // current x coordinate
	var y int = rectangle.Bounds.Min.Y // current y coordinate

	// loop while y in rectangle.bounds
	for y < rectangle.Bounds.Max.Y {
		// loop while x in rectangle.bounds
		for x < rectangle.Bounds.Max.X {
			// set position
			frame.Set(x, y, rectangle.Color)
			// increment
			x++
		}
		// increment
		y++
	}
}

// Overlaps() satisfies the Overlaps method on interface Shape
// it reports whether a rectangle overlaps with a given shape
func (rectangle Rectangle) Overlaps(shape Shape) bool {
	if rectangle1, ok := shape.(*Rectangle); ok { // we are processing a rectangle.
		return rectangle.Bounds.Overlaps(rectangle1.Bounds)
	}
	// we are processing an unknown shape and we freak out.
	println("Moving unknown shape")
	os.Exit(1)
	// unreachable code but we still need to return a boolean
	return false
}

// type Point
type Point struct {
	Coordinates image.Point // coordinates of point
	Color       color.RGBA  // color of point
}

// Draw() satisifes the Draw method on interface Shape
// it sets the point's x/y coordinates a draw.Image
func (point Point) Draw(frame draw.Image) {
	// set Point on frame
	frame.Set(point.Coordinates.X, point.Coordinates.Y, point.Color)
}

// Overlaps() reports if point overlaps a given shape
func (point Point) Overlaps(shape Shape) bool {
	if point1, ok := shape.(*Point); ok { // we are processing a point
		return point1.Coordinates == point.Coordinates
	}
	// we are processing an unknown shape and we panic
	println("Processing unknown shape.")
	os.Exit(1)
	// unreachable code but we still need to return a boolean
	return false
}
