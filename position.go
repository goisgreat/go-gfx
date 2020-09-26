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

/*
 * assert bounds on a given position
 * if position exceeds bounds, wrap
 */
func (position *Position) AssertBounds() {
	// if position.Coordinates.X < 0, wrap to {WIDTH}
	if position.Coordinates.X < 0 {
		position.Coordinates.X = MAX_X
	} else if position.Coordinates.X > MAX_X { // if position.Coordinates.X > WIDTH, wrap to 0
		position.Coordinates.X = 0
	}

	// if position.Coordinates.Y < 0, wrap to {HEIGHT}
	if position.Coordinates.Y < 0 {
		position.Coordinates.Y = MAX_Y
	} else if position.Coordinates.Y > MAX_Y { // if position.Coordinates.Y > HEIGHT, wrap to 0
		position.Coordinates.Y = 0
	}
}
