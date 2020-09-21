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
	Start Position
	End   Position
	Image draw.Image
	Color color.RGBA
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

func (rect *Rect) Contains(position Position) bool {
	return false
}

var _ Shape = &Rect{}
