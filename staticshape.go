package physics

type StaticShape struct {
	Vector
	Shape
	CollisionBox
}

func (staticShape *StaticShape) Update() (Shape, bool) {
	return staticShape.Shape, true
}

func (staticShape *StaticShape) Init() {

}
