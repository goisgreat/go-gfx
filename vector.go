package physics

// type vector
type Vector struct {
	X int
	Y int
}

/*
 * accept vector & shape pointer
 * update shape based on x/y velocity
 * if shape's position exceeds boundary, wrap
 */
func (vector Vector) Hit(shape Shape) {
	// move given shape (based on `vector` x/y coordinates)
	Move(shape, vector.X, vector.Y)
}
