package physics

import (
	"image/color"
	"image/draw"
	"os"
)

type Shape interface {
	Draw(draw.Image)
	Contains(Shape) bool
}

type Rect struct {
	Start Position   // starting point
	End   Position   // ending point
	Color color.RGBA // rectangle color
}

/*
 * draw a given rectangle
 */
func (rect Rect) Draw(img draw.Image) {
	for y := rect.Start.Y; y <= rect.End.Y; y++ {
		for x := rect.Start.X; x <= rect.End.X; x++ {
			img.Set(x, y, rect.Color)
		}
	}
}

/*
 * check if a rectangle contains a given position
 */
func (rect *Rect) Contains(shape Shape) bool {
	return false
}

/*
 * move a given shape
 */
func Move(shape Shape, x int, y int) {
	rect, ok := shape.(*Rect)
	if ok {
		rect.Start.X += x
		rect.Start.Y += x
		rect.End.X += x
		rect.End.Y += x
		return
	} else {
		os.Exit(1)
	}
}
