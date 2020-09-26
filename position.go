package physics

import (
	"image"
	"image/color"
	"image/draw"
)

// type Position
type Position struct {
	Coordinates image.Point // coordinates of point
	Color       color.RGBA
}

// Draw() sets the position's x/y coordinates a draw.Image
func (position Position) Draw(frame draw.Image) {
	// set position on frame
	frame.Set(position.Coordinates.X, position.Coordinates.Y, position.Color)
}
