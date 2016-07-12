package pgengine

import sf "github.com/manyminds/gosfml"

type keyboard struct {
	JustPressedKey sf.KeyCode
}

var Keyboard = keyboard{}

func (k *keyboard) KeyPressed(ev sf.EventKeyPressed) {
	k.JustPressedKey = ev.Code
}

func (k *keyboard) JustPressed(key sf.KeyCode) bool {
	return k.JustPressedKey == key
}

func (k *keyboard) IsPressed(key sf.KeyCode) bool {
	return sf.KeyboardIsKeyPressed(key)
}

func (k *keyboard) Clear(key sf.KeyCode) {
	k.JustPressedKey = -1
}
