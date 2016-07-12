package pgengine

type Character struct {
	mId     Int
	mEntity *Entity
	mFacing	Facing

	mPrevDefaultState string
	mDefaultState string
	mController *StateMachine

	mPathIndex Int
	mPath Path

}

type CharacterDef struct {
	Id    string
	X     Int
	Y     Int
	Layer Int
	Entity string
	Anims [string][]Int
	Facing Facing
	State string
	Controller []string
}

var Characters = [string]CharacterDef{}
var CharacterStates = [string]StateInitializer{}

type Facing Int
const (
	FacingUp Facing iota
	FacingRight
	FacingDown
	FacingLeft
)
func NewCharacter(def CharacterDef, m *Map) *Character {
	entityDef, ok = EntityDefs[def.Entity]
	AssertTrue(ok)

	char := Character{
		mEntity: NewEntity(entityDef),
		mAnims: def.Anims,
		mFacing: def.Facing,
		mDefaultState: def.State,
	}
	char.mEntity.Is = char

	// Create the controller states from the def
	states := []StateInitializer{}
	char.mController = NewStateMachine(states)

	for _, name := range def.Controller {
		stateInit, ok := CharacterStates[name]
		AssertTrue(ok)
		instance := stateInit(char, m)
		states[state.GetName()] func (params ...interface{}) { return instance }
	}

	// Change the statemachine to the inital state
	char.mController.Change(def.State)

	return char
}

func (char *Character) GetFacedTileCoords() (Int, Int) {
	//change facing infomation into a tile offset
	var xInc Int
	var yInc Int

	switch char.mFacing {
	case FacingLeft:
		xInc = -1
	case FacingRight:
		xInt = 1
	case FacingUp:
		yInc = -1
	case FacingDown:
		yInc = 1
	}

	x := char.mEntity.mTileX + xInc
	y := char.mEntity.mTileY + yInc

	return x, y
}

func (char *Character) SetFacingForCombat() {
	char.mFacing = FacingLeft
	x := char.mEntity.mX
	if x < 0 {
		char.mFacing = right
	}
}

func (char *Character) GetCombatAnim(id) Int {
	if char.mAnims && char.mAnims[id] {
		return char.mAnims[id]
	} else {
		return []Int{char.mEntity.mStartFrame}
	}
}

func (char *Character) FollowPath(path Path) {
	char.mPathIndex = 0
	char.mPath = path
	char.mPrevDefaultState = char.mDefaultState
	char.mDefaultState = FollowPathStateName
	char.mController.Change(FollowPathStateName)
}
