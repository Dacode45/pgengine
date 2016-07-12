package pgengine

import "time"

type Storyboard struct {
	mId            string
	mStack         *StateStack
	mEvents        []StoryboardEvent
	mStates        []State
	mSubStack      *StateStack
	mPlayingSounds [string]string
}

type StoryboardEvent interface {
	IsFunc() bool
	Gen(s *Storyboard) StoryboardEvent
	Update(time.Duration, *Storyboard)
	IsFinished() bool
	IsBlocking() bool
}

const StoryboardId = "storyboard"
const HandInId = "handin"

func NewStoryBoard(stack *StateStack, events []StoryboardEvent, handIn bool) *Storyboard {
	board := &Storyboard{
		mId:     StoryboardId,
		mStack:  stack,
		mEvents: events,
		mStates, []State{},
		mSubStack:      NewStateStack(),
		mPlayingSounds: [string]string{},
	}

	if handIn {
		state := board.mStack.Pop()
		board.PushState(HandInId, state)
	}

	return board
}

func (board *Storyboard) Enter(params ...interface{}) {}

func (board *Storyboard) Exit() {
	for _, v := range board.mPlayingSounds {
		Sound.Stop(v)
	}
}

func (board *Storyboard) AddSound(name, id string) {
	board.mPlayingSounds[name] = id
}

func (board *Storyboard) StopSound(name) {
	id := board.mPlayingSounds[name]
	delete(board.mPlayingSounds, name)
	Sound.Stop(id)
}

func (board *Storyboard) Update(dt time.Duration) {
	board.mSubStack.Update(dt)

	if len(board.mEvents) == 0 {
		board.mStack.Pop()
	}

	var shouldDelete bool
	var deleteIndex Int

	for k, v := range board.mEvents {
		if v.IsFunc() {
			board.mEvents[k] = v(v.Gen(board))
			v = board.mEvents[k]
		}

		v.Update(dt, board)
		if v.IsFinished() {
			shouldDelete = true
			deleteIndex = k
			break
		}

		if v.IsBlocking() {
			break
		}
	}

	if shouldDelete {
		board.mEvents = append(board.mEvents[:deleteIndex], board.mEvents[deleteIndex:])
	}

}

func (board *Storyboard) Render(target sf.RenderTarget, renderStates sf.RenderStates) {
	board.mSubStack.Render(target, renderstates)
}

func (board *Storyboard) HandleInput() {}

func (board *Storyboard) PushState(id string, state State) {
	board.mStates[id] = state
	board.mSubStack.Push(state)
}

func (board *Storyboard) RemoveState(id string) {
	state := board.mStates[id]
	delete(board.mStates, id)
	for i := len(board.mSubStack.mStates) - 1; i >= 0; i-- {
		v := board.mSubStack.mStates[i]
		if v == state {
			board.mSubStack.mStates = append(board.mSubStack.mStates[:i], board.mSubStack.mStates[i:])
		}
	}
}
