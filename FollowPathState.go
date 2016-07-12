package pgengine

type Path []PathDirection

type PathDirection Int

const (
  PathUp PathDirection Int
  PathRight
  PathDown
  PathLeft
)

const FollowPathStateName = "FollowPath"

type FollowPathState struct {
  mCharacter *Character
  mMap *Map
  mEntity *Entity
  mController *StateMachine
}

func (state *FollowPathState) GetName() string {
  return FollowPathStateName
}

func (state *FollowPathState) Enter(params ...interface{}) {
  char := state.mCharacter
  controller := state.mController

  if char.mPath == nil || char.mPathIndex >= len(char.mPath) {
    char.mDefaultState = DefaultString(char.mPrevDefaultState, char.mDefaultState)
    controller.Change(char.mDefaultState)
    return
  }

  direction := char.mPath[char.mPathIndex]
  switch direction {
  case PathUp:
    controller.Change(MoveStateString, 0, -1)
  case PathRight:
    controller.Change(MoveStateString, 1, 0)
  case PathDown:
    controller.Change(MoveStateString, 0, 1)
  case PathLeft:
    controller.Change(MoveStateString, -1, 0)
  case default:
    panic(fmt.Errorf("Bad Path Direction, %s", direction))
  }
}

func (state *FollowPathState) Exit() {
  state.mCharacter.mPathIndex = state.mCharacter.mPathIndex + 1
}

func (state *FollowPathState) Update(dt time.Duration) {}

func (state *FollowPathState) Render(target sf.RenderTarget, renderStates sf.RenderStates) {}

func NewFollowPathState(char *Character, m *Map) *FollowPathState {
  return &FollowPathState {
    mCharacter: char,
    mMap: m,
    mEntity: char.mEntity,
    mController: char.mController,
  }
}
