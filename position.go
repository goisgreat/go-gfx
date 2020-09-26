package physics

import "image"

// `Position` is an alias for image.Point
type Position image.Point

/*
 * accept position
 * draw position in terminal
 */
func (position Position) Draw() {
	println("X: ", position.X, " Y: ", position.Y)
}

/*
 * assert bounds on a given position
 * if position exceeds bounds, wrap
 */
func (position *Position) AssertBounds() {
	// if position.X < 0, wrap to {WIDTH}
	if position.X < 0 {
		position.X = MAX_X
	} else if position.X > MAX_X { // if position.X > WIDTH, wrap to 0
		position.X = 0
	}

	// if position.Y < 0, wrap to {HEIGHT}
	if position.Y < 0 {
		position.Y = MAX_Y
	} else if position.Y > MAX_Y { // if position.Y > HEIGHT, wrap to 0
		position.Y = 0
	}
}
