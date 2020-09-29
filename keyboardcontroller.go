package physics

import "time"

// KeyboardControl is an alias for uint8 and is equivilant to uint8 in all ways
// It is used to distinguish keyboard controls from integer values
type KeyboardControl uint8

// map of controls
type KeyboardMap struct {
	KEY_UP    byte // ascii character for key up
	KEY_LEFT  byte // ascii character for key down
	KEY_DOWN  byte // ascii character for key left
	KEY_RIGHT byte // ascii character for key right
}

// wasd controll scheme
var WASD = KeyboardMap{
	119, // w
	97,  // a
	115, // s
	100, // d
}

// keyboard controller component
type KeyboardController struct {
	Input             chan byte             // where to get input from
	KeyboardMap                             // control scheme
	OnKeyboardControl func(KeyboardControl) // invoked when keyboard control hit
}

// DirectPositionControl() yields the user direct control over their position
func DirectPositionControl(shape Geometry, delay time.Duration) func(KeyboardControl) {
	return func(control KeyboardControl) {
		// decide what to do based on the control presed
		switch control {
		case CON_UP:
			// move shape up
			Vector{0, -1}.Hit(shape)
		case CON_LEFT:
			// move shape left
			Vector{-1, 0}.Hit(shape)
		case CON_DOWN:
			// move shape down
			Vector{0, 1}.Hit(shape)
		case CON_RIGHT:
			// move shape right
			Vector{1, 0}.Hit(shape)
		}
		// delay
		time.Sleep(delay)
	}
}

// Init() initializes a KeyboardController
// please call Init() upon creating a KeyboardController
func (keyboardController KeyboardController) Init() {
	// maps character to CON_UP, CON_DOWN, etc.
	charmap := map[byte]KeyboardControl{
		keyboardController.KeyboardMap.KEY_UP:    CON_UP,
		keyboardController.KeyboardMap.KEY_DOWN:  CON_DOWN,
		keyboardController.KeyboardMap.KEY_LEFT:  CON_LEFT,
		keyboardController.KeyboardMap.KEY_RIGHT: CON_RIGHT,
	}
	// loop
	for {
		// get character
		char := <-keyboardController.Input
		// if character in map AND OnKeyboardControl != nil, invoke it
		if control, ok := charmap[char]; ok && keyboardController.OnKeyboardControl != nil {
			keyboardController.OnKeyboardControl(control)
		}
	}
}

// KeyboardControllerConfigOption is an alias for uint8 and is equivilant to uint8 in all ways.
// It is used to distinguish integer values from configuration options for a KeyboardControllerConfig.
type KeyboardControllerConfigOption uint8

// KeyboardControllerConfig provides a shorthand syntax for instantiating KeyboardControllers.
type KeyboardControllerConfig struct {
	Input    chan byte                      // channel to get keyboard input from
	Geometry                                // shape to control
	Config   KeyboardControllerConfigOption // configuration option
}

// Init() Initializes a KeyboardControllerConfig and returns a KeyboardController
func (keyboardControllerConfig KeyboardControllerConfig) Init() KeyboardController {
	// declare parameters to put in result
	var keyboardMap KeyboardMap
	var onKeyboardControl func(KeyboardControl)

	// decide what to do based on config option
	switch keyboardControllerConfig.Config {
	case DIRECT_WASD:
		// direct WASD setup
		keyboardMap = WASD
		onKeyboardControl = DirectPositionControl(keyboardControllerConfig.Geometry, time.Millisecond*5)
	}

	// create a KeyboardController object and return it
	return KeyboardController{
		Input:             keyboardControllerConfig.Input,
		KeyboardMap:       keyboardMap,
		OnKeyboardControl: onKeyboardControl,
	}
}
