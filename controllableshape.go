package physics

type ControllableShape struct {
	StaticShape
	KeyboardController
}

func (controllableShape *ControllableShape) Update() Shape {
	return controllableShape.Shape
}
