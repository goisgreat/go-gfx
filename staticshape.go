package physics

type StaticShape struct {
	Vector
	Shape
	CollisionBox
}

func (staticShape *StaticShape) Update() Shape {
	return staticShape.Shape
}

func (staticShape *StaticShape) Init() {
	staticShape.CollisionBox.Shape = staticShape.Shape
}
