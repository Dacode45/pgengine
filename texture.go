package engine

import sf "github.com/manyminds/gosfml"

type UV struct {
	Left   float32
	Top    float32
	Right  float32
	Bottom float32
}

func (uv *UV) GetRectangle(text *sf.Texture) sf.IntRect {
	width := float32(text.GetSize().X)
	height := float32(text.GetSize().Y)
	r := sf.IntRect{
		Left:   int(uv.Left * width),
		Top:    int(uv.Top * height),
		Width:  int((uv.Right*width - uv.Left*width)),
		Height: int((uv.Bottom*height - uv.Top*height)),
	}
	// fmt.Println(uv.Right, uv.Left, (uv.Right*width - uv.Left), uv.Left*width)
	// fmt.Printf("%+v\n", r)
	return r
}

func GenerateUVs(text *sf.Texture, tileWidth, tileHeight Int) []sf.IntRect {
	size := text.GetSize()
	textWidth := Int(size.X)
	textHeight := Int(size.Y)
	cap := textWidth / tileWidth * textHeight / tileHeight
	uvs := make([]sf.IntRect, 0, cap)
	for x := Int(0); x < textWidth; x += tileWidth {
		for y := Int(0); y < textHeight; y += tileHeight {
			uvs = append(uvs, sf.IntRect{
				Left:   x,
				Top:    y,
				Width:  tileWidth,
				Height: tileHeight,
			})
		}
	}
	return uvs
}

const CollisionTileSetName = "collision_graphic"

type TileSet struct {
	Image    string
	Name     string
	FirstGid Int
}
