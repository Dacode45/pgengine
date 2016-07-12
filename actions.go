package pgengine

type ActionDef struct{
  Id string
  Params []interface{}
}

var EmptyAction = func(trigger Trigger, entity *Entity, tX, tY, tLayer Int){}
type Action func(trigger Trigger, entity *Entity, tX, tY, tLayer Int)

type ActionGenerator func (m *Map, params ...interface{}) Action

type MapScript func(m *Map, trigger Trigger, entity *Entity, tX, tY, tLayer Int)

var Actions = [string]ActionGenerator{
  "Teleport": func (map, params ...interface{}) {
    tileX := params[0].(Int)
    tileY := params[0].(Int)
    tLayer := params[0].(Int)

    return func(trigger Trigger, entity *Entity, tX, tY, tLayer Int) {
      entity.SetTilePos(x, y, layer, m)
    }
  },

  "AddNPC" : func (m *Map, params ...interface{}) {
    npc := params[0].(CharacterDef)
    return func (trigger *MapTrigger, entity *Entity, tX, tY, tLayer) {

      charDef, ok := Characters[npc.Id]
      AssertTrue(ok)
      char := NewCharacter(charDef)

      x := DefaultInt(npc.X, char.mEntity.mTileX)
      y := DefaultInt(npc.Y, char.mEntity.mTileY)
      layer := DefaultInt(npc.Layer, char.mEntity.mLayer)

      char.mEntity.SetTilePos(x, y, layer, m)

      m.mNPCs = append(m.mNPCs, char)
      AssertTrue(m.mNPCbyID[npc.Id] == nil)
      char.mId = npc.Id
      m.mNPCbyID[npc.Id] = char
    }
  },
  "RunScript" : func (m *Map, Func MapScript ) {
    return func (trigger *MapTrigger, entity *Entity, tX, tY, tLayer) {
      Func(m, trigger, entity, tX, tY, tLayer)
    }
  }
}
