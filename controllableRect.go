package physics

type ControllableRect struct {
	KeyboardController
	CollisionBox
	Rect
	Callbacks
}

/*
 * update given ControllableRect
 */
func (controllableRect ControllableRect) Update() (Shape, bool) {
	// update using "callbacks" object
	controllableRect.Callbacks.Update()
	// return shape
	return &controllableRect.Rect, true
}
