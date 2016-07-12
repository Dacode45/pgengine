package main

import (
	"fmt"
	"math"
	"time"

	pg "github.com/Dacode45/pgengine"
	sf "github.com/manyminds/gosfml"
)

func main() {
	engine := pg.NewPGEngine(pg.PGEngineConfig{})
	renderer := engine.GetRenderer()
	dim1 := sf.IntRect{Left: 0, Top: 0, Width: 32, Height: 32}
	text, _ := sf.NewTextureFromFile("grass_tile.png", &dim1)
	dim2 := sf.IntRect{Left: 0, Top: 0, Width: 512, Height: 512}
	atlas, _ := sf.NewTextureFromFile("atlas.png", &dim2)
	gTileSprite, _ := sf.NewSprite(text)

	var (
		gTileWidth  = text.GetSize().X
		gTileHeight = text.GetSize().Y

		gDisplayWidth  = renderer.GetSize().X
		gDisplayHeight = renderer.GetSize().Y

		_ = uint(math.Ceil(float64(gDisplayWidth) / float64(gTileWidth)))
		_ = uint(math.Ceil(float64(gDisplayHeight) / float64(gTileHeight)))

		gTop  float32
		gLeft float32
	)

	var gMap = []int{
		1, 1, 1, 1, 5, 6, 7, 1, // 1
		1, 1, 1, 1, 5, 6, 7, 1, // 2
		1, 1, 1, 1, 5, 6, 7, 1, // 3
		3, 3, 3, 3, 11, 6, 7, 1, // 4
		9, 9, 9, 9, 9, 9, 10, 1, // 5
		1, 1, 1, 1, 1, 1, 1, 1, // 6
		1, 1, 1, 1, 1, 1, 2, 3, // 7
	}

	var (
		gMapWidth  = uint(8)
		gMapHeight = uint(7)
	)
	view := sf.NewViewFromRect(sf.FloatRect{0, 0, 8 * 16, 7 * 16})
	renderer.SetView(view)

	var gUVs = pg.GenerateUVs(atlas, gTileWidth)

	gTileSprite.SetPosition(sf.Vector2f{
		X: float32(renderer.GetSize().X / 2),
		Y: float32(renderer.GetSize().Y / 2),
	})

	engine.Update = func(dur time.Duration) {
		for j := uint(0); j < gMapHeight; j++ {
			for i := uint(0); i < gMapWidth; i++ {
				tile := GetTile(gMap, int(gMapWidth), int(i), int(j))
				uvs := gUVs[tile]
				gTileSprite.SetTextureRect(uvs.GetRectangle(atlas))
				gTileSprite.SetPosition(pg.NewVec2f(gLeft+float32(i*gTileWidth),
					gTop+float32(j*gTileHeight)))
				renderer.Draw(gTileSprite, sf.DefaultRenderStates())
			}
		}
	}
	for x := uint(0); x < renderer.GetSize().X; x = x + 16 {
		for y := uint(0); y < renderer.GetSize().Y; y = y + 16 {
			tx, ty := PointToTile(int(x), int(y), int(gTileWidth), int(gLeft), int(gTop), int(gMapWidth), int(gMapWidth))
			fmt.Println(x, y, tx, ty)
		}
	}
	engine.Run()
}

func GetTile(Map []int, rowsize, x, y int) int {
	return Map[x+y*rowsize] - 1
}

func PointToTile(x, y, tileSize, left, top, mapWidth, mapHeight int) (int, int) {
	//Tiles rendered from center
	x = x + tileSize/2
	y = y + tileSize/2

	//Clamp point to bounds of map
	x = MaxInt(left, x)
	y = MaxInt(top, y)
	x = MinInt(left+(mapWidth*tileSize)-1, x)
	y = MinInt(top+(mapHeight*tileSize)-1, y)

	tileX := math.Floor((float64(x) - float64(left)) / float64(tileSize))
	tileY := math.Floor((float64(top) + float64(y)) / float64(tileSize))

	return int(tileX), int(tileY)
}

func MaxInt(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func MinInt(a, b int) int {
	if a < b {
		return a
	}
	return b
}
