package engine

import sf "github.com/manyminds/gosfml"

func NewVec2f(x, y float32) sf.Vector2f {
	return sf.Vector2f{
		X: x,
		Y: y,
	}
}
