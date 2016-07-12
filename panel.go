package goldengine

import sf "github.com/manyminds/gosfml"

type PanelDef struct {
	Texture *sf.Texture
	Size    //Size of a Single Tile in Pixels
}

type Panel struct {
	mTexture     *sf.Texture
	mUVs         []IntRect
	mTileSize    Int
	mTiles       []*sf.Sprite
	mCenterScale Float
}

type PanelChild struct {
	X   Int
	Y   Int
	Obj interface{}
}

func NewPanel(params PanelDef) *Panel {
	panel := Panel{
		mTexture:  params.Texture,
		mUVs:      GenerateUVs(params.Texture, params.Size, params.Size),
		mTileSize: params.Size,
		mTiles:    make([]*sf.Sprite, 9),
	}

	//TODO Fix up the center U, Vs by moving them 0.5 textels in.

	//The center sprite is going to be 1 pixel smaller all around.
	panel.mCenterScale = panel.mTileSize.AsFloat32() / (panel.mTileSize - 1).AsFloat32()

	// Create a sprite for each tile of the panel
	// 0. top left     1. top     2. top right
	// 3. left         4. center  5. right
	// 6. bottom left  7. bottom  8. bottom right

	for k, v := range panel.mUVs {
		sprite := sf.NewSprite(panel.mTexture)
		sprite.SetTextureRect(v)
		panel.mTiles[k] = sprite
	}

	return panel
}

func (panel *Panel) SetColor(color sf.Color) {
	for _, v := range panel.mTiles {
		v.SetColor(color)
	}
}

func (panel *Panel) SetPosition(left, top, right, bottom Int) {
	// Reset scales
	for _, v := range panel.mTiles {
		v.SetScale(sf.Vector2f{1, 1})
	}

	hSize := panel.mTileSize / 2
	// Align the corner tiles
	panel.mTiles[0].SetPosition(sf.Vector2f{
		X: left.AsFloat32(),
		Y: top.AsFloat32(),
	})
	panel.mTiles[2].SetPosition(sf.Vector2f{
		X: float32(right - hSize),
		Y: top.AsFloat32(),
	})
	panel.mTiles[6].SetPosition(sf.Vector2f{
		X: left.AsFloat32(),
		Y: float32(bottom - hSize),
	})
	panel.mTiles[8].SetPosition(sf.Vector2f{
		X: float32(right - hSize),
		Y: float32(bottom - hSize),
	})

	// Calculate how much to scale the side tiles
	widthScale := ((right - left).Abs() - panel.mTileSize) / panel.mTileSize
	centerX := left + hSize

	panel.mTiles[1].SetPosition(sf.Vector2f{
		X: centerX.AsFloat32(),
		Y: top.AsFloat32(),
	})
	panel.mTiles[1].SetScale(sf.Vector2f{X: widthScale, Y: 1})

	panel.mTiles[7].SetPosition(sf.Vector2f{
		X: centerX.AsFloat32(),
		Y: float32(bottom - hSize),
	})
	panel.mTiles[7].SetScale(sf.Vector2f{X: widthScale, Y: 1})

	heightScale := ((bottom - top).Abs() - panel.mTileSize) / panel.mTileSize
	centerY := top + panel.mTileSize

	panel.mTiles[3].SetScale(sf.Vector2f{
		X: 1,
		Y: heightScale.AsFloat32(),
	})
	panel.mTiles[3].SetPosition(sf.Vector2f{X: left, Y: centerY})

	panel.mTiles[5].SetScale(sf.Vector2f{X: 1, Y: heightScale.AsFloat32()})
	panel.mTiles[5].SetPosition(sf.Vector2f{X: right - hSize, Y: centerY.AsFloat32()})

	// Scale the middle backing panel
	panel.mTiles[4].SetScale(sf.Vector2f{
		X: (widthScale * panel.mCenterScale),
		Y: (heightScale * panel.mCenterScale),
	})
	panel.mTiles[4].SetPosition(sf.Vector2f{X: centerX, Y: centerY})

	// Hide corner tiles when scale is equal to zero
	if left-right == 0 || top-bottom == 0 {
		for _, v := range panel.mTiles {
			v.SetScale(sf.Vector2f{})
		}
	}
}

func (panel *Panel) CenterPosition(x, y, width, height Int) {
	hWidth := width / 2
	hHeight := height / 2
	panel.SetPosition(x-hWidth, y-hHeight, x+hWidth, y+hHeight)
}

func (panel *Panel) Render(target sf.RenderTarget, renderStates sf.RenderStates) {
	for _, v := range panel.mTiles {
		v.Draw(target, renderStates)
	}
}
