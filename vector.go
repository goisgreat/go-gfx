package physics

import (
	"image"
	"os"
)

// type vector
type Vector image.Point

// Hit() moves [shape] using the X/Y coordinates from a vector
func (vector Vector) Hit(shape Shape) {
	if rect, ok := shape.(*Rectangle); ok { // we are moving a rectangle.
		rect.Bounds.Min.X += vector.X
		rect.Bounds.Min.Y += vector.Y
		rect.Bounds.Max.X += vector.X
		rect.Bounds.Max.Y += vector.Y
		return
	} else { // we are moving an unknown shape and we freak out.
		println("Moving unknown shape")
		os.Exit(1)
	}
}
