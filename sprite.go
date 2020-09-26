package physics

import "image/draw"

// Sprites is a list of sprites
type Sprites []Sprite

// Sprite is an interface for all actions relating to a game object
type Sprite interface {
	Update() Shape      // update sprite and return a new shape
	Draw(draw.Image)    // draw a sprite on a given image
	Init()              // initialize a sprite
	Process(Shape) bool // process shape and return boolean value indicating a collision
}

// Init() initializes all sprites in a set
func (sprites Sprites) Init() {
	// iterate sprites in set
	for idx := range sprites {
		// initialize current sprite
		sprites[idx].Init()
	}
}

// Update() updates ALL sprites in a set
func (sprites Sprites) Update() {
	// create a list of positions (for collision detection)
	shapes := []Shape{}

	// iterate sprites in set
	for idx := range sprites {
		// update current sprite. if it has a position, append it to the list
		if shape := sprites[idx].Update(); shape != nil {
			shapes = append(shapes, shape)
		}
	}

	// loop through positions, processing each one. if a collision is detected, return
	for idx := range shapes {
		// get current position
		shape := shapes[idx]
		// loop through sprites in set, instructing each one to process the current position (unless it's the sprite that's position is being processed)
		for idx1 := range sprites {
			if idx != idx1 {
				if collision := sprites[idx1].Process(shape); collision {
					return
				}
			}
		}
	}
}

// Draw() draws ALL sprites in set on a given frame
func (sprites Sprites) Draw(image draw.Image) {
	// iterate sprites in set
	for idx := range sprites {
		// draw current sprite on frame
		sprites[idx].Draw(image)
	}
}
