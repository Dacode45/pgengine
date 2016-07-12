package pgengine

import sf "github.com/manyminds/gosfml"

type resource struct{}

var Resource = resource{}

//TODO fill in
func (r *resource) FindTexture(filename string) *sf.Texture {
	return &sf.Texture{}
}

func (r *resource) FindFont(filename string) *sf.Font {
	return &sf.Font{}
}
