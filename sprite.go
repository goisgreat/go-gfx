package physics

import "image/draw"

// Sprites is a list of sprites
type Sprites []Sprite

// Sprite is an interface for all actions relating to a game object
type Sprite interface {
	Update()                           // update sprite and return a new shape
	GetShape() Geometry                // get a sprite's shape
	Draw(draw.Image)                   // draw a sprite on a given image
	Init()                             // initialize a sprite
	Process(Geometry, ShapeComparison) // process shape and return boolean value indicating a collision
}

// Init() initializes all sprites in a set
func (sprites Sprites) Init() {
	// iterate sprites in set
	for _, sprite := range sprites {
		// initialize current sprite
		sprite.Init()
	}
}

// Update() updates ALL sprites in a set
func (sprites Sprites) Update() {
	// create a list of positions (for collision detection)
	shapes := []Geometry{}

	// iterate sprites in set
	for _, sprite := range sprites {
		// update current sprite
		sprite.Update()
		// get sprite's shape. if shape != nil, append it to stored list
		if shape := sprite.GetShape(); shape != nil {
			shapes = append(shapes, shape)
		}
	}

	// loop through positions, processing each one. if a collision is detected, return
	for idx, shape1 := range shapes {
		// loop through sprites
		for idx1, sprite := range sprites {
			// is the current sprite the one who's shape is being processed?
			if idx != idx1 {
				// get the current sprite's shape
				shape0 := sprite.GetShape()
				// compare shape0 to shape1 and get a few statistics
				comparison := shape0.Compare(shape1)
				// instruct the current sprite to process this shape and pass it `comparison`
				sprite.Process(shape1, comparison)
			}
		}
	}
}

// Draw() draws ALL sprites in set on a given frame
func (sprites Sprites) Draw(image draw.Image) {
	// iterate sprites in set
	for _, sprite := range sprites {
		// draw current sprite on frame
		sprite.Draw(image)
	}
}
