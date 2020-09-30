package physics

// ControllableShape Sprite = StaticShape + KeyboardContoller.
type ControllableShape struct {
	StaticShape        // static shape base
	KeyboardController // extra keyboard controller component
}

// CreateControllableShape() is a helper function for instantiating ControllableShape objects.
func CreateControllableShape(shape Geometry, keyboardController KeyboardController, collisionBox CollisionBox) ControllableShape {
	// create a new ControllableShape and return it
	return ControllableShape{
		StaticShape:        CreateStaticShape(shape, collisionBox), // with a StaticShape derived from given shape
		KeyboardController: keyboardController,                     // ...and the given KeyboardController
	}
}

// Init() satisfies the Init method on interface Sprite.
func (controllableShape ControllableShape) Init() {
	// initialize static shape
	controllableShape.StaticShape.Init()
	// initialize keyboard controller in a seperate goroutine
	go controllableShape.KeyboardController.Init()
}
