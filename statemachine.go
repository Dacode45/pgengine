package pgengine

type State interface {
  GetName() string
  Render(target sf.RenderTarget, renderStates sf.RenderStates)
  Update(dt time.Duration) bool //Returns whether bottom states should update.
  Enter(...interface{})
  Exit()
}

type StateInitializer func (...interface{}) State

type StateMachine struct {
  mStates map[string]StateInitializer
  mCurrent State
}

func NewStateMachine(states [string]StateInitializer) *StateMachine{
  sm := StateMachine{
    mStates: []StateInitializer{},
    mCurrent: nil,
  }
  for k, v := range states {
    sm.mStates[k] = v
  }
  return &sm
}

func (sm *StateMachine) Change(stateName string, enterParams ..interface{}){
  if sm.mCurrent == nil {
    sm.Exit()
    sm.mCurrent = sm.mStates[stateName]()
    sm.mCurrent(enterParams...)
  }
}

func (sm *StateMachine) Update(dt time.Duration) {
  sm.mCurrent.Update(dt)
}

func (sm *StateMachine) Render(target sf.RenderTarget, renderStates sf.RenderStates) {
  sm.mCurrent.Render(...interface{})
}
