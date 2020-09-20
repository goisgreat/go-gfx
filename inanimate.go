package gogfx

type Inanimate struct {
	Position
	Vector
}

/*
 * nothing here
 */
func (inanimate *Inanimate) Init() {

}

/*
 * accept point pointer
 * update
 */
func (inanimate *Inanimate) Update() (*Position, bool) {
	// update vector, specifying position
	inanimate.Vector.Hit(&inanimate.Position)

	// return the new position (and true to indicate we have one)
	return &inanimate.Position, true
}

/*
 * accept inanimate and frame pointer
 * draw inanimate on screen (using position component)
 */
func (inanimate Inanimate) Draw() {
	// draw inanimate on screen
	inanimate.Position.Draw()
}

/*
 * nothing to see here
 */
func (inanimate Inanimate) Process(position *Position) bool {
	// return false
	return false
}
