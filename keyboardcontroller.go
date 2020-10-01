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

// WASD controll scheme
var WASD = KeyboardMap{
	119, // w
	97,  // a
	115, // s
	100, // d
}

// KeyboardController is a component for handling keyboard input
type KeyboardController struct {
	Input             chan byte             // where to get input from
	KeyboardMap                             // control scheme
	OnKeyboardControl func(KeyboardControl) // invoked when keyboard control hit
}

// Init() initializes a KeyboardController.
// Please call Init() upon creating a KeyboardController.
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

// DirectWASDKeyboardController() provides a default handler for keyboard events.
func DirectWASDKeyboardController(input chan byte, shape Geometry, delay time.Duration) KeyboardController {
	return KeyboardController{
		KeyboardMap: WASD,
		OnKeyboardControl: func(control KeyboardControl) {
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
		},
	}
}
