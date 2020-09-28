package physics

// CollisionBox is an object for collision handling
type CollisionBox struct {
	OnCollision func(Geometry, ShapeComparison) // What do I do when a collision occures?
	Shape                                       // what shape do I use to check collisions with?
}

// Process() satisfies the Process() method on interface Sprite.
// It processes a given shape and invokes any OnCollision callback
func (collisionbox CollisionBox) Process(shape Geometry, comparison ShapeComparison) {
	// if `collisionbox`'s shape is != nil AND a collision happened AND OnCollision is SET, invoke it
	if collisionbox.Shape.Geometry != nil && comparison.Overlaps && collisionbox.OnCollision != nil {
		collisionbox.OnCollision(shape, comparison)
	}
}
