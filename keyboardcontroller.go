package physics

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

/*
 * grab character from input stream
 * when a character of intrest is found, invoke callback
 */
func (keyboard_controller KeyboardController) Init() {
	// maps character to CON_UP, CON_DOWN, etc.
	charmap := map[byte]KeyboardControl{
		keyboard_controller.KeyboardMap.KEY_UP:    CON_UP,
		keyboard_controller.KeyboardMap.KEY_DOWN:  CON_DOWN,
		keyboard_controller.KeyboardMap.KEY_LEFT:  CON_LEFT,
		keyboard_controller.KeyboardMap.KEY_RIGHT: CON_RIGHT,
	}
	// loop
	for {
		// get character
		char := <-keyboard_controller.Input
		// if character in map, invoke callback
		if control, ok := charmap[char]; ok {
			keyboard_controller.OnKeyboardControl(control)
		}
	}
}
