package physics

type Callbacks struct {
	OnUpdate func()
}

/*
 * use stored callbacks to update
 */
func (callbacks Callbacks) Update() {
	//
	if callbacks.OnUpdate != nil {
		callbacks.OnUpdate()
	}
}
