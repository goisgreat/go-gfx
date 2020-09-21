package physics

type ControllableShape struct {
	KeyboardController
	CollisionBox
	Shape
	Callbacks
}

/*
 * nothing to see here
 */
func (controllableShape ControllableShape) Update() (Shape, bool) {
	// update using "callbacks" object
	controllableShape.Callbacks.Update()
	// return shape
	return controllableShape.Shape, true
}
