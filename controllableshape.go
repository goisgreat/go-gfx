package physics

type ControllableShape struct {
	StaticShape
	KeyboardController
}

// Update() satisfies the Update method on interface Sprite
func (controllableShape *ControllableShape) Update() Shape {
	return controllableShape.Shape
}

// Init() satisfies the Init method on interface Sprite
// please call Init() upon creating a ControllableShape
func (controllableShape ControllableShape) Init() {
	controllableShape.StaticShape.Init()
	controllableShape.KeyboardController.Init()
}
