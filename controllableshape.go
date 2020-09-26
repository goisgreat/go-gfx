package physics

// ControllableShape = StaticShape + KeyboardContoller
type ControllableShape struct {
	StaticShape        // static shape base
	KeyboardController // extra keyboard controller component
}

// CreateControllableShape() is a helper function for instantiating ControllableShape objects
func CreateControllableShape(shape Shape, keyboardController KeyboardController) ControllableShape {
	// create a new ControllableShape and return it
	return ControllableShape{
		StaticShape:        CreateStaticShape(shape), // with a StaticShape derived from given shape
		KeyboardController: keyboardController,       // ...and the given KeyboardController
	}
}

// Update() satisfies the Update method on interface Sprite
func (controllableShape *ControllableShape) Update() Shape {
	return controllableShape.Shape
}

// Init() satisfies the Init method on interface Sprite
// please call Init() upon creating a ControllableShape
func (controllableShape ControllableShape) Init() {
	// initialize static shape
	controllableShape.StaticShape.Init()
	// initialize keyboard controller in a seperate goroutine
	go controllableShape.KeyboardController.Init()
}
