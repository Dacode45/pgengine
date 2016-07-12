package pgengine

import sf "github.com/manyminds/gosfml"

type WaitState struct {
  mCharacter *Character
  mMap *Map
  mEntity *Entity
  mController *StateMachine
  mFrameResetSpeed Float
  mFrameCount Float
}

func NewWaitState(params ...interface{}) State {
  char := params[0].(*Charater)
  m := params[1].(*Map)
  state := &WaitState{
    mCharacter: char,
    mMap: m,
    mEntity: char.mEntity,
    mController: mCharacter.mController,

    mFrameRestSpeed: 0.05,
  }
  return state
}

func (state *WaitState) Enter(params ...interface{}) {
  state.mFrameCount = 0
}

func (state *WaitState) Render(target sf.RenderTarget, renderStates sf.RenderStates) {}

func (state *WaitState) Exit() {}

func (state *WaitState) Update(dt time.Duration) {
  if state.mFrameCount == -1 {
    state.mFrameCount = state.mFrameCount + dt.Seconds()
    if state.mFrameCount >= state.mFrameRese {
      state.mFrameCount = -1
      state.mEntity.SetFrame(state.mEntity.mStartFrame)
      state.mCharacter.mFacing = FacingDown
    }
  }

  if Keyboard.IsPressed(sf.KeyLeft) {
    state.mController.Change("move", Point{-1, 0})
  } else if Keyboard.IsPressed(sf.KeyRight) {
    state.mController.Change("move", Point{1, 0})
  } else if Keyboard.IsPressed(sf.KeyUp) {
    state.mController.Change("move", Point{0, -1})
  } else {
    state.mController.Change("move", Point{0, 1})
  }
}
