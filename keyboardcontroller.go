package physics

import "time"

// keyboard control type
type KeyboardControl uint8

// map of controls
type KeyboardMap struct {
	KEY_UP    byte
	KEY_LEFT  byte
	KEY_DOWN  byte
	KEY_RIGHT byte
}

// keyboard controller component
type KeyboardController struct {
	Input             chan byte             // where to get input from
	KeyboardMap                             // control scheme
	OnKeyboardControl func(KeyboardControl) // invoked when keyboard control hit
}

// wasd controll scheme
var WASD = KeyboardMap{
	119, // w
	97,  // a
	115, // s
	100, // d
}

// yields the user direct control over their position
func DirectPositionControl(shape Shape, delay time.Duration) func(KeyboardControl) {
	return func(control KeyboardControl) {
		// decide what to do based on the control presed
		switch control {
		case CON_UP:
			Move(shape, 0, -1)
		case CON_LEFT:
			Move(shape, -1, 0)
		case CON_DOWN:
			Move(shape, 0, 1)
		case CON_RIGHT:
			Move(shape, 1, 0)
		}
		// delay
		time.Sleep(delay)
	}
}

/*
 * grab character from input stream
 * when a character of intrest is found, invoke callback
 */
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
		// if character in map, invoke callback
		if control, ok := charmap[char]; ok {
			keyboardController.OnKeyboardControl(control)
		}
	}
}
