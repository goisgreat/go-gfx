package physics

// StaticShape Sprite = Vector + Shape + CollisionBox
type StaticShape struct {
	Vector
	Shape
	CollisionBox
}

// Update() satisfies the Update method on Sprite interface.
func (staticShape *StaticShape) Update() {}

// Init() satisfies a the Init method on Sprite interface.
// Please invoke Init() upon creating a StaticShape.
func (staticShape *StaticShape) Init() {
	staticShape.CollisionBox.Shape = staticShape.Shape
}

// CreateStaticShape() is a helper function for instantiating `StaticShape`s.
func CreateStaticShape(shape Geometry, collisionBox CollisionBox) StaticShape {
	// store a new StaticShape
	sprite := StaticShape{
		Vector{0, 0}, // use a null vector...
		Shape{shape}, // ...a shape component derived from the given shape
		collisionBox, // ...and the given collision box
	}
	// initialize it
	sprite.Init()
	// return
	return sprite
}
