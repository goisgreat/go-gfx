package gogfx

// type vector
type Vector struct {
	X int
	Y int
}

/*
 * accept frame pointer & position pointer
 * update position based on x/y velocity
 * if position exceeds boundary, wrap
 */
func (vector Vector) Hit(position *Position) {
	// make a new position
	position1 := Position{position.X, position.Y}

	// update position1 with vector data
	position1.X += vector.X
	position1.Y += vector.Y

	// assert bounds
	position1.AssertBounds()

	// save position1 into position
	position.X = position1.X
	position.Y = position1.Y
}
