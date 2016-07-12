package pgengine

type InGameMenuState struct {
	mMapDef    MapDef
	mStack     *StateStack
	mTitleSize Float
	mLabelSize Float
	mTextSize  Int
	mStateMachine
}

func NewInGameMenuState(params ...interface{}) *InGameMenuState {
	stack := params[0].(*StateStack)
	mapDef := params[1].(MapDef)

	state := InGameMenuState{
		mMapDef:    mapDef,
		mStack:     *StateStack,
		mTitleSize: 1.2,
		mLabelSize: 0.88,
		mTextSize:  1,
	}

	return state
}
