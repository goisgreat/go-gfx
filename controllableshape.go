package physics

type ControllableShape struct {
	StaticShape
	KeyboardController
}

func (controllableShape *ControllableShape) Update() Shape {
	return controllableShape.Shape
}

func (controllableShape ControllableShape) Init() {
	controllableShape.StaticShape.Init()
	controllableShape.KeyboardController.Init()
}
