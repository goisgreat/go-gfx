package physics

import (
	"image"
	"image/color"
	"image/draw"
)

// ShapeComparison is the result of running Shape.Compare().
// It provides information about the relative position of two shapes.
type ShapeComparison struct {
	Higher   bool // is shape1 above shape0?
	Lower    bool // is shape1 below shape0?
	Right    bool // is shape1 to the right of shape0?
	Left     bool // is shape1 to the left of shape0?
	Overlaps bool // does shape1 overlap shape0?
}

// Shape represents the shape component and has an embedded Geometry interface
type Shape struct {
	Geometry
}

// GetShape() is used to get the geometry of a shape
func (shape Shape) GetShape() Geometry {
	return shape.Geometry
}

// Geometry is an interface for rendering and trigonometry
type Geometry interface {
	Draw(draw.Image)                  // draw a shape on a given image
	Compare(Geometry) ShapeComparison // compare with another shape
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

// Compare() satifies the Compare method on interface Shape
func (rectangle Rectangle) Compare(shape Geometry) ShapeComparison {
	// store a ShapeComparison to write results to
	var result ShapeComparison
	// are we comparing rectangles?
	if rectangle1, ok := shape.(*Rectangle); ok {
		// does rectangle1 overlap with rectangle?
		result.Overlaps = rectangle1.Bounds.Overlaps(rectangle.Bounds)
	}
	// return result
	return result
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

// Compare() satisfies the Compare method on interface Shape
func (point Point) Compare(shape Geometry) ShapeComparison {
	// store a ShapeComparison to write results to
	var result ShapeComparison
	// are we processing a point?
	if point1, ok := shape.(*Point); ok {
		// report if point overlaps point1
		result.Overlaps = point.Coordinates == point1.Coordinates
	}
	// return result
	return result
}
