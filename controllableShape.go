package physics

type ControllableShape struct {
	KeyboardController
	CollisionBox
	Shape
}

/*
 * nothing to see here
 */
func (controllableShape ControllableShape) Update() (Shape, bool) {
	// return shape
	return controllableShape.Shape, true
}
