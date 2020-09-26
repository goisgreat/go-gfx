package physics

// type CollisionBox
type CollisionBox struct {
	OnCollision func(Shape) // What do I do when a collision occures?
	Shape
}

// Process() satisfies the Process() method on interface Sprite
// it processes a given shape and invokes
func (collisionbox CollisionBox) Process(shape Shape) bool {
	// collisionbox.Shape is provided...
	if collisionbox.Shape != nil {
		// if collisionbox.Shape overlaps given shape
		if collisionbox.Shape.Overlaps(shape) {
			// if oncollision set
			if collisionbox.OnCollision != nil {
				collisionbox.OnCollision(shape)
			}
			// return true to indicate a collision
			return true
		}
	}
	// return false to indicate no collision
	return false
}
