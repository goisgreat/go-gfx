package main

// point object
type Point struct {
	Position
	Vector
	KeyboardController
	CollisionBox
}

/*
 * listen for keyboard input asynchronously
 */
func (point *Point) Init() {
	// handle keyboard input
	go point.KeyboardController.Init()
}

/*
 * accept point pointer
 * update and return the required fields
 */
func (point *Point) Update() (*Position, bool) {
	// update vector, specifying position
	point.Vector.Hit(&point.Position)

	// return the new position (and true to indicate we have one)
	return &point.Position, true
}

/*
 * accept point and frame pointer
 * draw point on screen (using position component)
 */
func (point Point) Draw() {
	// draw point on screen
	point.Position.Draw()
}

/*
 * process a given position, returning true if a collision occured
 */
func (point Point) Process(position *Position) bool {
	// let collision box handle any collisions
	collision := point.CollisionBox.Process(point.Position, position)

	// return true if a collision happened
	return collision
}
