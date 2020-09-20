package main

// list of sprites
type Sprites []Sprite

// game object
type Sprite interface {
	Update() (*Position, bool)
	Draw()
	Init()
	Process(*Position) bool
}

/*
 * accept set of sprites
 * initialize each sprite in set
 */
func (sprites Sprites) Init() {
	// iterate sprites in set
	for idx := range sprites {
		// initialize current sprite
		sprites[idx].Init()
	}
}

/*
 * accept set of sprites
 * update all sprites in set
 */
func (sprites Sprites) Update() {
	// create a list of positions (for collision detection)
	positions := []*Position{}

	// iterate sprites in set
	for idx := range sprites {
		// update current sprite. if it has a position, append it to the list
		if position, hasposition := sprites[idx].Update(); hasposition {
			positions = append(positions, position)
		}
	}

	// loop through positions, processing each one. if a collision is detected, return
	for idx := range positions {
		// get current position
		position := positions[idx]

		// loop through sprites in set, instructing each one to process the current position (unless it's the sprite that's position is being processed)
		for idx1 := range sprites {
			if idx != idx1 {
				if collision := sprites[idx1].Process(position); collision {
					return
				}
			}
		}
	}
}

/*
 * accept set of sprites
 * draw all sprites in set on frame
 */
func (sprites Sprites) Draw() {
	// iterate sprites in set
	for idx := range sprites {
		// draw current sprite on frame
		sprites[idx].Draw()
	}
}
