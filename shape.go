package physics

import (
	"image"
	"image/color"
	"image/draw"
	"os"
)

// Shape is an interface for geometry and rendering
type Shape interface {
	Draw(draw.Image)     // draw a shape on a given image
	Overlaps(Shape) bool // report if a shape overlaps with another given one
}

// Rectangle Shape = Color +
type Rectangle struct {
	Bounds image.Rectangle // rectangle bounds
	Color  color.RGBA      // color
}

// CreateRectangle() is a helper function for creating Rectangle objects
func CreateRectangle(minX int, minY int, maxX int, maxY int, c color.RGBA) *Rectangle {
	// create a new Rectangle and return it
	return &Rectangle{
		// rectangle boundary
		Bounds: image.Rectangle{
			// minimum
			image.Point{minX, minY},
			// maximum
			image.Point{maxX, maxY},
		},
		// rectangle color
		Color: c,
	}
}

// Draw() satisfies the Draw method on interface Shape
// it draws, pixel by pixel, a rectangle
func (rectangle Rectangle) Draw(frame draw.Image) {
	// loop while y in rectangle.bounds
	for y := rectangle.Bounds.Min.Y; y < rectangle.Bounds.Max.Y; y++ {
		// loop while x in rectangle.bounds
		for x := rectangle.Bounds.Min.X; x < rectangle.Bounds.Max.X; x++ {
			// set current position on frame
			frame.Set(x, y, rectangle.Color)
		}
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

// Point Shape = image.Point + color.RGBA
type Point struct {
	Coordinates image.Point // coordinates of point
	Color       color.RGBA  // color of point
}

// Draw() satisifes the Draw method on interface Shape.
// It sets the point's x/y coordinates a draw.Image
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
