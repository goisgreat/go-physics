package physics

// CollisionBox is an object for collision handling
type CollisionBox struct {
	OnCollision func(Shape) // What do I do when a collision occures?
	Shape                   // what shape do I use to check collisions with
}

// Process() satisfies the Process() method on interface Sprite.
// It processes a given shape and invokes any OnCollision callback
func (collisionbox CollisionBox) Process(shape Shape) bool {
	// compare shape0 to shape1
	comparison := collisionbox.Shape.Compare(shape)
	// collisionbox.Shape is provided...
	if collisionbox.Shape != nil {
		// if collisionbox.Shape overlaps given shape
		if comparison.Overlaps {
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
