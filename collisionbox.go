package physics

type CollisionBox struct {
	OnCollision func(Shape)
	Shape
}

/*
 * process a given shape, invoking callback and returning true if a collision occures
 */
func (collisionbox CollisionBox) Process(shape Shape) bool {
	if collisionbox.Shape != nil {
		if collisionbox.Shape.Overlaps(shape) {
			collisionbox.OnCollision(shape)
			return true
		}
		return false
	}
	return false
}
