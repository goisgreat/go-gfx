package physics

type Callbacks struct {
	OnUpdate func() (Shape, bool)
}

/*
 * use stored callbacks to update
 */
func (callbacks Callbacks) Update() (Shape, bool) {
	// if onupdate is set, invoke it and return shape, bool
	if callbacks.OnUpdate != nil {
		shape, hasShape := callbacks.OnUpdate()
		return shape, hasShape
	}
	// otherwise return no shape and false
	return nil, false
}
