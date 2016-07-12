package pgengine

import (
	"fmt"

	sf "github.com/manyminds/gosfml"
)

type Entity struct {
	mLayer  Int
	mTileX  Int
	mTileY  Int
	mHeight Int
	mWidth  Int

	mX Int
	mY Int

	mStartFrame Int
	mSprite     *sf.Sprite
	mTexture    *sf.Texture
	mUVs        []sf.IntRect

	mChildren []*Entity

	Is interface // Whatever this Enity is. Could be a character.
}

type EntityDef struct {
	Height     Int
	Width      Int
	TileX      Int
	TileY      Int
	Layer      Int
	Texture    string
	StartFrame Int
	X          Int
	Y          Int
}

var EntityDefs = map[string]EntityDef //All Possible Characters

func NewEntity(def EntityDef) *Entity {
	e := Entity{
		mSprite:     sf.NewSprite(),
		mTexture:    Resource.Find(def.Texture),
		mHeight:     def.HeHeight,
		mWidth:      def.Width,
		mTileX:      def.tileX,
		mTileY:      def.tileY,
		mLayer:      def.layer,
		mStartFrame: def.StartFrame,
		mX:          def.X,
		mY:          def.Y,
		mChildren:   [Int]*Entity{},
	}
	e.mSprite.SetTexture(e.mTexture, false)
	e.mUVs = GenerateUVs(e.mTexture, e.mWidth, e.mHeight)
	e.SetFrame(e.mStartFrame)
}

func (entity *Entity) SetFrame(frame Int) {
	e.mSprite.SetTextureRect(e.mUVs[frame])
}

func (entity *Entity) SetTilePos(x, y, layer Int, m *Map) {

	// Remove from current tile
	if m.GetEntity(entity.mTileX, entity.mTileY, entity.mLayer) == entity {
		m.RemoveEntity(entity)
	}

	//Check target tile
	if m.GetEntity(x, y, layer) != nil {
		panic(fmt.Errorf("There's something in the target position!"))
	}

	entity.mTileX = x
	entity.mTileY = y
	entity.mLayer = layer

	m.AddEntity(entity)
	x, y = m.GetTileFoot(x, y, layer)
	entity.mSprite.SetPosition(sf.Vector2f{
		X: x.AsAsFloat32(),
		Y: y.AsFloat32(),
	})
	self.mX = x
	self.mY = y
}

func (entity *Entity) GetSelectPosition() (Int, Int) {
	pos := entity.mSprite.GetPosition()
	height := entity.mHeight

	x := Int(pos.X)
	y := Int(pos.Y) + height/2

	yPad := 16

	y = y + yPad

	return x, y

}

func (entity *Entity) GetTargetPosition() (Int, Int) {
	pos := entity.mSprite.GetPosition()
	width := entity.mWidth

	x := Int(pos.X) + width/2
	y := Int(pos.Y)

	return x, y
}

func (entity *Entity) AddChild(id Int, e *Entity) {
	entity.mChildren[id] = e
}

func (entity *Entity) RemoveChild(id Int) {
	delete(entity.mChildren, id)
}

func (entity *Entity) Draw(target sf.RenderTarget, renderStates sf.RenderStates) {
	if entity.sprite != nil {
		entity.sprite.Draw(target, renderStates)
	}

	for _, v := range entity.mChildren {
		sprite := v.mSprite
		pos := sprite.GetPosition()
		pos.X = pos.X + entity.X
		pos.Y = pos.Y + entity.Y
		sprite.SetPosition(pos)
		sprite.Draw(target, renderStates)
	}
}

type NPC struct {
	Entity
}

type Hero struct {
	Entity
}
