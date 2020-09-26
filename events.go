package physics

// EventStream is an object for storing input from many sources
type EventStream struct {
	Input chan byte // keyboard input
}

// CreateEventStream() is a helper function to instantiate EventStream objects
func CreateEventStream() EventStream {
	// create a new EventStream and return it
	return EventStream{
		Input: make(chan byte), // with keyboard input as a new byte channel
	}
}
