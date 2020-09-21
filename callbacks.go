package physics

type Callbacks struct {
	OnUpdate func()
}

func (callbacks Callbacks) Update() {
	callbacks.OnUpdate()
}
