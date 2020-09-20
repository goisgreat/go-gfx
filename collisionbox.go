package main

type CollisionBox struct {
	OnCollision func(*Position)
}

func (collisionbox CollisionBox) Process(pos0 Position, pos1 *Position) bool {
	if pos0 == *pos1 {
		collisionbox.OnCollision(pos1)
		return true
	} else {
		return false
	}
}
