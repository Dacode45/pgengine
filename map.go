package pgengine

import (
	"fmt"

	sf "github.com/manyminds/gosfml"
)

type Position struct {
	X Int
	Y Int
	Z Int
}

type MapLayer struct {
	Width  Int
	Height Float
	Data   []Int
}

type Trigger struct {
}

type Map struct {
	mX Int
	mY Int

	mCamX Int
	mCamY Int
	mCamWidth Int
	mCamHeight Int

	mMapDef       *MapDef
	mTextureAtlas *sf.Texture
	mUVs          []sf.IntRect

	mTileSprite  *sf.Sprite
	mLayer       MapLayer
	mWidth       Int
	mHeight      Int
	mWidthPixel  Int
	mHeightPixel Int

	mTiles        []Int
	mBlockingTile Int
	mTileWidth    Int
	mTileHeight   Int

	mTriggers     []Trigger
	mTriggerTypes [string]Trigger
	mActions      []Action
	mEntities     map[Int]map[Int]*Entity
	mNPCs         []*Character
	mNPCbyID      map[string]NPC
}

type MapDef struct {
	Layers   []MapLayer
	TileSets []TileSet

	Actions     []ActionDef
	TriggerType TriggerType

	OnWake []ActionDef
}

func NewMap(mapDef MapDef) *Map {
	layer = mapDef.Layers[0]
	m := Map{
		mMapDef:       mapDef,
		mTextureAtlas: Resource.FindTexture(mapDef.TileSets[0].Image),

		mTileSprite: sf.NewSprite(nil),
		mLayer:      layer,
		mWidth:      layer.Width,
		mHeight:     layer.Height,

		mCamWidth: 800,
		mCamHeight: 600,

		mTiles:      layer.Data,
		mTileWidth:  mapDef.TileSets[0].TileWidth,
		mTileHeight: mapDef.TileSets[0].TileHeight,
		mTriggers:   []Trigger{},
		mEntities:   []Entity{},
		mNPCs:       []NPC{},
		mNPCbyId:    map[string]NPC{},
	}
	m.mTileSprite.SetTexture(m.mTextureAtlas, false)

	// Map Dimensions
	m.mWidthPixel = m.mWidth * m.mTileWidth
	m.mHeightPixel = m.mHeight * m.mTileHeight
	m.mUVs = GenerateUVs(m.mTextureAtlas, mapDef.TileSets[0].TileWidth, mapDef.TileSets[0].TileHeight)

	//Assign blocking tile id
	for _, v := range mapDef.TileSets {
		if v.Name == CollisionTileSetName {
			m.mBlockingTile = v.FirstGid
		}
	}
	// TODO add assert
	//assert(m.mBlockingTile)
	fmt.Println("blocking tile is", m.mBlockingTile)

	// Create Actions from def
	m.mActions = [string]Action{}
	for name, def := range mapDef.Actions {
		// TODO assert map has keys
		//assert(Actions[def.id])
		action := Actions[def.Id](m, def.Params...)
		m.mActions[name] = action
	}

	// Create Trigger Type from def
	for k, v := range mapDef.TriggerType {
		triggerParam := Trigger{}
		for callback, action := range v {
			_, ok := m.mActions[action]
			AssertTrue(ok)
			triggerParam[callback] = m.mActions[action]
		}
		this.mTriggerTypes[k] = NewTrigger(triggerParam)
	}

	//Place Triggers
	m.mTriggers = []Trigger{}
	for _, v := range mapDef.Triggers {
		m.AddTrigger(v)
	}

	for _, v := range mapDef.OnWake {
		action := Actions[v.Id]
		action(m, v.Params)
	}

	return m
}

func (m *Map) AddTrigger(trigger TriggerDef) {
	x := def.X
	y := def.Y
	layer := def.Layer

	if m.mTriggers[layer] == nil {
		m.mTriggers[layer] = [Int]Trigger{}
	}

	targetLayer = m.mTriggers[layer]
	trigger = m.mTriggerTypes[def.Trigger]
	//TODO add assert(trigger)
	targetLayer[m.CoordToIndex(x, y)] = trigger
	fmt.Println("Add trigger", x, y)
	fmt.Println("Trigger", m.GetTrigger(x, y, layer))
}

func (m *Map) AddFullTrigger(trigger Trigger, x, y, layer Int) {
	if m.mTriggers[layer] == nil {
		m.mTriggers[layer] = [Int]Trigger{}
	}

	targetLayer := m.mTriggers[layer]
	targetLayer[m.CoordToIndex(x, y)] = trigger
}

func (m *Map) GetEntity(x, y, layer Int) *Entity {
	if m.mEntities[layer] == nil {
		return nil
	}
	index := m.CoordToIndex(x, y)
	return m.mEntities[layer][index]
}

func (m *Map) AddEntity(entity *Entity) {
	if m.mEntities[entity.mLayer] == nil {
		m.mEntities[entity.mLayer] = map[Int]*Entity{}
	}

	layer := m.mEntities[entity.mLayer]
	index := m.CoordToIndex(entity.mTileX, entity.mTileY)

	// TODO add assert(layer[index] == entity or layer[index]==nil)
	layer[index] = entity
}

func (m *Map) RemoveEntity(entity *Entity) {
	//assert(m.mEntities[entity.mLayer])
	layer := m.mEntities[entity.mLayer]
	index := m.CoordsToIndex(entity.mTileX, entity.mTileY)
	//Entity should be at the position
	//assert(entity == layer[index])
	layer[index] = nil
}

func (m *Map) GetTile(x, y, layer Int) {
	tiles := m.mMapDef.Layers[Int].data
	return tiles[m.CoordsToIndex(x, y)]
}

type MapTile struct {
	X Int
	Y Int

	Tile   Int
	Layer  Int
	Detail Int
}

func (m *Map) WriteTile(params MapTile) {
	layer := params.Layer
	detail := params.Detail

	//Each layer has 3 parts
	layer = layer * 3

	x := params.X
	y := params.Y
	tile := params.Tile
	collision := m.mBlockingTile
	if !params.Collision {
		collision = 0
	}

	index = m.CoordToIndex(x, y)
	tiles := m.mMapDef.Layers[layer].Data
	tiles[index] = tile

	//Detail
	tiles = m.mMapDef.Layers[layer+1].Data
	tiles[index] = detail

	//Collision
	tiles = m.mMapDef.layers[layer+2].Data
	tiles[index] = collision

}

func (m *Map) GetTrigger(x, y, layer Int) Trigger {
	triggers := m.mTriggers[layer]

	if triggers == nil {
		return nil
	}

	index := m.CoordToIndex(x, y)
	triggers[index]
}

func (m *Map) RemoveTrigger(x, y, layer Int) {
	// TODO assert(m.mTriggers[layer])
	triggers := m.mTriggers[layer]
	index := m.CoordToIndex(x, y)
	// TODO assert(triggers[index])
	triggers[index] = nil
}

func (m *Map) GetNPC(x, y, layer Int) *Character{
	for _, npc := range m.mNPCs {
		if npc.mEntity.mLayer == layer && npc.mEntity.mTileX == x and npc.mEntity.mTileY == y {
			return npc
		}
	}

	return nil
}

func (m *Map) RemoveNPC(x, y, layer Int) bool {
	for i := len(m.mNPCs) - 1; i >= 0; i = i - 1 {
		npc := m.mNPCs[i]
		if npc.mEntity.mLayer == layer && npc.mEntity.mTileX == x and npc.mEntity.mTileY == y {
			m.RemoveEntity(npc.mEntity)
			npc.mEntity.Is = nil
			m.mNPCbyID[npc.mId] = nil
			return true
		}
	}

	return false
}

func (m *Map) CoordToIndex(x, y) {
	return x + y*m.mWidth
}

func (m *Map) IsBlocked(layer, tileX, tileY) {
	tile := m.GetTile(tileX, tileY, layer+2)
	entity := m.GetEntity(tileX, tileY, layer)
	return tile == m.mBlockingTile || entity != nil
}

func (m *Map) GetTileFoot(x, y Int) (Int, Int) {
	return m.mX + (x * m.mTileWidth), m.mY + (y * m.mTileHeight) + m.mTileHeight/2
}

func (m *Map) GotoTile(x, y Int) {
	fmt.Println("Goto tile:", x, y)
	m.Goto((x*m.mTileWidth)+m.mTileWidth/2, (y*m.mTileHeight)+m.mTileHeight/2)
}

func (m *Map) Goto(x, y Int) {
	m.mCamX = x - Int(Window.RenderWindow.GetSize().X/2)
	m.mCamY = y - Int(Window.RenderWindow.GetSize().Y/2)
}

func (m *Map) PointToTile(x, y Int) {
	// Tiles are rendered from the center
	x = x + m.mTileWidth/2
	y = y + m.mTileHeight/2
	//clamp the points to bound of map
	x = Max(m.mX, x)
	y = Max(m.mY, y)
	x = Min(m.mX+m.mWidthPixel-1, x)
	y = Min(m.mY+m.mHeightPixel-1, y)
	// Map from the bounded point to a tile
	tileX := Floor((x - m.mX).AsFloat32() / m.mTileWidth.AsFloat32())
	tileY := Floor((y - m.mY).AsFloat32() / m.mTileHeight.AsFloat32())

	return tileX, tileY
}

func (m *Map) LayerCount() Int {
	//TODO assert(len(m.mMapDef.Layers) %3 == 0)
	return len(m.mMapDef.Layers) / 3
}

func (m *Map) Render(target sf.RenderTarget, renderStates sf.RenderStates) {
	m.RenderLayer(target, renderStates, 1, nil)
}

func (m *Map) RenderLayer(target sf.RenderTarget, renderStates sf.RenderStates, layer Int, hero *Hero) {
	layerIndex := layer * 3

	size := Window.RenderWindow.GetSize()
	tileLeft, tileBottom := m.PointToTile(m.mCamX-size.X/2, m.mCamY+size.Y/2)
	tileRight, tileTop := m.PointToTile(m.mCamX+size.X/2, m.mCamY-size.Y/2)

	for j := tileTop; j <= tileBottom; j++ {
		for i := tileLeft; i <= tileRight; i++ {
			tile := m.GetTile(i, j, layerIndex)

			m.mTileSprite.SetPosition(sf.Vector2f{
				X: m.mX + i*m.mTileWidth,
				Y: m.mY + j*m.mTileHeight,
			})

			//there can be empty tiles
			if tile > 0 {
				uvs := m.mUVs[tile]
				m.mTileSprite.SetTextureRect(uvs)
				m.mTileSprite.Draw(target, renderStates)
			}
		}
	}

	entLayer := m.mEntities[layer]
	drawList := []*Entity{hero}

	for _, j := range entLayer {
		drawList = append(drawList, j)
	}

	for _, j := range drawList {
		j.Draw(2, renderStates)
	}
}
