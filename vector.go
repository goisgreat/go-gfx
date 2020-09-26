package physics

import (
	"image"
	"os"
)

// type vector
type Vector image.Point

/*
 * accept vector & shape pointer
 * update shape based on x/y velocity
 * if shape's position exceeds boundary, wrap
 */
func (vector Vector) Hit(shape Shape) {
	if rect, ok := shape.(*Rectangle); ok {
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
