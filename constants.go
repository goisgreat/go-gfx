package physics

// ENUM: keyboard controls
const (
	CON_UP KeyboardControl = iota
	CON_LEFT
	CON_DOWN
	CON_RIGHT
)

// ENUM: KeyboardControllerConfig options
const (
	DIRECT_WASD KeyboardControllerConfigOption = iota
)

// ENUM: CollisionBoxConfig options
const (
	RIGID_BODY CollisionBoxConfigOption = iota
)
