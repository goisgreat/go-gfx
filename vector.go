package physics

import "os"

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
	rect, ok := shape.(*Rect)
	if ok {
		rect.Start.X += vector.X
		rect.Start.Y += vector.Y
		rect.End.X += vector.X
		rect.End.Y += vector.Y
		return
	} else {
		println("Moving unknown shape")
		os.Exit(1)
	}
}
