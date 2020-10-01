package physics

import "os"

// CollisionBox is the component responsible for collision handling.
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

// CreateRigidBody() provides a default functionality for basic 2d collisoin physics.
func CreateRigidBody(shape Geometry) CollisionBox {
	return CollisionBox{
		// oncollision callback
		OnCollision: func(shape Geometry, comparison ShapeComparison) {
			println("collision!")
			os.Exit(0)
		},
		// derive shape component from given shape
		Shape: Shape{shape},
	}
}
