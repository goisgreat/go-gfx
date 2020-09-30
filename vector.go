package physics

import (
	"image"
	"os"
)

// Vector is an alias for image.Point and is equivalent to image.Point in all ways
// It is used to distinguish vectors from points on the screen
type Vector image.Point

// Hit() moves [shape] using the X/Y coordinates from a vector.
func (vector Vector) Hit(shape Geometry) {
	if rect, ok := shape.(*Rectangle); ok { // we are moving a rectangle.
		rect.Bounds.Min.X += vector.X
		rect.Bounds.Min.Y += vector.Y
		rect.Bounds.Max.X += vector.X
		rect.Bounds.Max.Y += vector.Y
		return
	}
	// we are moving an unknown shape and we freak out.
	println("Moving unknown shape")
	os.Exit(1)
}
