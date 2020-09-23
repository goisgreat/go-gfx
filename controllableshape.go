package physics

type ControllableShape struct {
	Vector
	Shape
	KeyboardController
	CollisionBox
}

func (controllableShape *ControllableShape) Update() Shape {
	return controllableShape.Shape
}
