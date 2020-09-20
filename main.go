package gogfx

import (
	"time"
)

/*
 * program {
	* clear terminal
	* create an event stream
	* load event stream with keyboard input (generated asynchronously)
	* main game loop {
		* clear terminal
		* update and draw sprites
		* delay 200 milliseconds
	}
 }
*/
func main() {
	// event stream
	stream := EventStream{
		Input: make(chan byte),
	}

	// define a point object
	point := Point{
		Position{X: 0, Y: 0},
		Vector{X: 0, Y: 0},
		KeyboardController{
			Input:       stream.Input,
			KeyboardMap: WASD,
		},
		CollisionBox{},
	}

	// what happens when a WASD control is pressed?
	point.KeyboardController.OnKeyboardControl = func(control KeyboardControl) {
		switch control {
		case CON_UP:
			point.Vector.Y--
		case CON_DOWN:
			point.Vector.Y++
		case CON_LEFT:
			point.Vector.X--
		case CON_RIGHT:
			point.Vector.X++
		}
	}

	inanimate := Inanimate{
		Position{X: 1, Y: 0},
		Vector{X: 0, Y: 0},
	}

	// what happens when `point` hits something?
	point.CollisionBox.OnCollision = func(position *Position) {
		position.X++
	}

	// list of objects
	sprites := Sprites{
		&point,
		&inanimate,
	}

	// setup
	Setup()

	// grab input asynchronously
	go GrabInput(stream)

	// initialize sprites
	sprites.Init()

	// game loop
	for {
		// clear the screen
		ClearTerminal()

		// update & draw sprites
		sprites.Update()
		sprites.Draw()

		// delay
		time.Sleep(200 * time.Millisecond)
	}
}
