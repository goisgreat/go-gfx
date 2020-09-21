package gogfx

import (
	"image/color"
	"image/draw"
)

type Shape interface {
	Draw()
	Contains(Position) bool
}

type Rect struct {
	Start Position   // starting point
	End   Position   // ending point
	Image draw.Image // image to draw the rectangle on
	Color color.RGBA // rectangle color
}

/*
 * draw a given rectangle
 */
func (rect Rect) Draw() {
	for y := rect.Start.Y; y <= rect.End.Y; y++ {
		for x := rect.Start.X; x <= rect.End.X; x++ {
			rect.Image.Set(x, y, rect.Color)
		}
	}
}

/*
 * check if a rectangle contains a given position
 */
func (rect *Rect) Contains(position Position) bool {
	return false
}
