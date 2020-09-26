package physics

import (
	"image"
	"image/color"
	"image/draw"
)

type Shape interface {
	Draw(draw.Image)
	Overlaps(Shape) bool
}

// type Rectangle
type Rectangle struct {
	Bounds image.Rectangle // rectangle bounds
	Color  color.RGBA      // color
}

// Draw() draws, pixel by pixel, a rectangle
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
