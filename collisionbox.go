package physics

// CollisionBox is an object for collision handling
type CollisionBox struct {
	OnCollision func(Geometry, ShapeComparison) // What do I do when a collision occures?
	Shape                                       // what shape do I use to check collisions with?
}

// Process() satisfies the Process() method on interface Sprite.
// It processes a given shape and invokes any OnCollision callback.
func (collisionbox CollisionBox) Process(shape Geometry, comparison ShapeComparison) {
	// if `collisionbox`'s shape is != nil AND a collision happened AND OnCollision is SET, invoke it
	if collisionbox.Shape.Geometry != nil && comparison.Overlaps && collisionbox.OnCollision != nil {
		collisionbox.OnCollision(shape, comparison)
	}
}

// CollisionBoxConfigOption is an alias to uint8 and is equivilant to uint8 in all ways.
// It is used to distinguish between integer values and configuration options for a CollisionBoxConfig.
type CollisionBoxConfigOption uint8

// CollisionBoxConfig is a helper struct for instantiating CollisionBox objects.
type CollisionBoxConfig struct {
	Config CollisionBoxConfigOption
}

// Init() creates, and returns, a CollisionBox based off info stored in the CollisionBoxConfig.
func (collisionBoxConfig CollisionBoxConfig) Init() CollisionBox {
	// create a CollisionBox
	box := CollisionBox{}

	// decide what to do based off of collisionBoxConfig.Config
	switch collisionBoxConfig.Config {
	case RIGID_BODY:
		box.OnCollision = func(shape Geometry, comparison ShapeComparison) {
			println(comparison)
		}
	}

	// return CollisionBox created earlier
	return box
}
