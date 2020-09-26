package physics

// EventStream is an object for storing input from many sources
type EventStream struct {
	Input chan byte // keyboard input
}
