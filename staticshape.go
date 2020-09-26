package physics

type StaticShape struct {
	Vector
	Shape
	CollisionBox
}

// Update() satisfies the Update method on Sprite interface
func (staticShape *StaticShape) Update() Shape {
	return staticShape.Shape
}

// Init() satisfies a the Init method on Sprite interface
// please invoke this upon creating a StaticShape
func (staticShape *StaticShape) Init() {
	staticShape.CollisionBox.Shape = staticShape.Shape
}
