package pgengine

type StateStack struct {
	mStates []State
}

func NewStateStack() *StateStack {
  return &StateStack{
    mStates: []State{}
  }
}

func (ss *StateStack) Push(state State) {
  ss.mStates = append(ss.mStates, state)
  state.Enter(nil)
}

func (ss *StateStack) Pop() State {
  top, ss.mStates := ss.mStates[len(ss.mStates)-1], ss.mStates[:len(ss.mStates)-1]
  top.Exit()
  return top
}

func (ss *StateStack) Top() State {
  return ss.mStates[len(ss.mStates) - 1]
}
//Update updates each state from the top down
func (ss *StateStack) Update(dt time.Duration) {
  for k := len(ss.mStates) -1; k >= 0; k = k -1 {
    v := ss.mStates[k]
    updateNext := v.Update(dt)
    if !updateNext {
      break
    }
  }

  top := ss.Top()

  if top == nil {
    return
  }

  top.HandleInput()
}

//Render states from bottom up
func (ss *StateMachine) Render(target sf.RenderTarget, renderStates sf.RenderStates) {
  for _, v := range ss.mStates {
    v.Render(target sf.RenderTarget, renderStates sf.RenderStates)
  }
}
