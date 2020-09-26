package physics

import (
	"image"
	"image/color"
	"image/draw"
)

// type Point
type Point struct {
	Coordinates image.Point // coordinates of point
	Color       color.RGBA  // color of point
}

// Draw() sets the point's x/y coordinates a draw.Image
func (point Point) Draw(frame draw.Image) {
	// set Point on frame
	frame.Set(point.Coordinates.X, point.Coordinates.Y, point.Color)
}
