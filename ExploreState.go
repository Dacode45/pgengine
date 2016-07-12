package pgengine

import "time"

type ExploreState struct {
	mStack *StateStack

	mMapDef MapDef
	mMap    *Map
	mHero   *Character

	mFollowCam  bool
	mFollowChar *Character
	mManualCamX Int
	mManualCamY Int
}

func NewExploreState(params ...interface{}) *ExploreState {
	stack := params[0].(*StateStack)
	mapDef := params[1].(MapDef)
	startPos := params[2].(Position)

	state := &State{
		mStack:  stack,
		mMapDef: mapDef,

		mFollowCam:  true,
		mFollowChar: nil,
		mManualCamX: 0,
		mManualCamY: 0,

		mMap: NewMap(mapDef),
	}
	state.mHero = NewCharacter(Characters["hero"], state.mMap)
	state.mHero.mEntity.SetTilePos(startPos.X, startPos.Y, startPos.Z)
	state.mHero.GotoTile(startPos.X, startPos.Y)

	state.mFollowChar = state.mHero
	return state
}

func (state *ExploreState) HideHero() {
	state.mHero.mEntity.SetTilePos(
		state.mHero.mEntity.mTileX,
		state.mHero.mEntity.mTileY,
		-1,
		state.mMap)
}

func (state *ExploreState) ShowHero(layer Int) {
  state.mHero.mEntity.SetTilePos(
    state.mHero.mEntity.mTileX,
    state.mHero.mEntity.mTileY,
    layer,
    state.mMap
  )
}

func (state *ExploreState) Enter() {}

func (state *ExploreState) Exit() {}

func (state *ExploreState) UpdateCamera(m *Map) {
	if state.mFollowCam {
		pos := state.mHero.mEntity.mSprite.GetPosition()
		m.mCamX = Int(pos.X)
		m.mCamY = Int(pos.Y)
	} else {
		m.mCamX = state.mManualCamX
		m.mCamY = state.mManualCamY
	}
}

func (state *ExploreState) SetFollowCam(flag bool, character *Character) {
	state.mFollowChar = character
	state.mFollowCam = flag
	if !state.mFollowCam {
		pos := state.mFollowChar.mEntity.mSprite.GetPosition()
		state.mManualCamX = Int(pos.X)
		state.mManualCamY = Int(pos.Y)
	}
}

func (state *ExploreState) Update(dt time.Duration) {
	m := state.mMap

	state.UpdateCamera(m)

	for _, v := range m.mNPCs {
		v.mController.Update(dt)
	}
}

func (state *ExploreState) Render(target sf.RenderTarget, renderStates sf.RenderStates) {
  hero := state.mHero
  m := state.mMap

  view := sf.NewViewFromRect(sf.FloatRect{
    Left: (mCamX - mCamWidth / 2).AsFloat32(),
    Top: (mCamY - mCamHeight / 2).AsFloat32(),
    Width: mCamWidth.AsFloat32(),
    Height: mCamHeight.AsFloat32(),
  })

  target.SetView(view)

  layerCount := m.LayerCount()

  for i := 0; i < layerCount; i++ {
    var heroEntity *Entity
    if i == hero.mEntity.mLayer {
      heroEntity = hero.mEntity
    }

    m.RenderLayer(target, renderStates, i, hero)
  }

  view := sf.NewViewFromRect(sf.FloatRect{
    Width: mCamWidth.AsFloat32(),
    Height: mCamHeight.AsFloat32(),
  })
  target.SetView(view)
}

func (state *ExploreState) HandleInput() {
  if GlobalWorld.IsInputLocked() {
    return
  }

  state.mHero.mController.Update(GetDeltaTime())

  if Keyboard.JustPressed(sf.KeySpace) {
    // Which way is the player facing?
    x, y := state.mHero.GetFacedTileCoords()
    fmt.Println("Hero facing", x, y)
    layer := state.mHero.mEntity.mLayer
    trigger := state.mMap.GetTrigger(x, y, layer)
    if trigger {
      trigger.OnUse(state.mHero, x, y, layer)
    }
  }

  if Keyboard.JustPressed(sf.KeyAult) {
    menu := NewInGameMenuState(state.mStack, state.mMapDef)
    return state.mStack.Push(menu)
  }
}
