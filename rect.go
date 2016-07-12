package engine

import sf "github.com/manyminds/gosfml"

//By16 is a 16 x 16 Rect
var By16 = sf.IntRect{Left: 0, Top: 0, Width: 16, Height: 16}

func NewIntRect(left, top, width, height int) sf.IntRect {
	return sf.IntRect{
		Left:   left,
		Top:    top,
		Width:  width,
		Height: height,
	}
}
