package pgengine

import (
	"math"
	"time"

	sf "github.com/manyminds/gosfml"
)

type TextBoxDef struct {
	Text       []string
	TextScale  Float
	TextBounds TextBounds

	Children []PanelChild

	Wrap             Int
	Size             Int
	PanelDef         PanelDef
	SelectionMenuDef SelectionMenuDef
}

type TextBox struct {
	mX      Int
	mY      Int
	mWidth  Int
	mHeight Int

	mTime time.Duration

	mStack *StateStack

	mChunkIndex Int
	mChunks     []string

	mChildren []PanelChild

	mContinueMark    *sf.Sprite
	mTextScale       Float
	mPanel           *Panel
	mSize            TextBounds
	mBounds          TextBounds
	mAppearTween     *Tween
	mWrap            Int
	mDoClickCallback bool
	mOnFinish        func()
}

type TextBounds struct {
	Left   Int
	Right  Int
	Top    Int
	Bottom Int
}

func NewTextBox(params TextBoxDef) *TextBox {
	textBox := TextBox{
		mChunks:          params.Text,
		mContinueMark:    sf.NewSprite(Resource.FindTexture("continue_caret.png")),
		mTextScale:       DefaultFloat(params.TextScale, 1),
		mPanel:           NewPanel(params.PanelDef),
		mSize:            params.Size,
		mBounds:          params.TextBounds,
		mAppearTween:     NewTween(0, 1, 0.4, EaseOutCirc),
		mWrap:            params.Wrap,
		mChildren:        params.Children,
		mSelectionMenu:   NewSelectionMenu(params.SelectionMenuDef),
		mStack:           params.Stack,
		mDoClickCallback: false,
		mOnFinish:        params.OnFinish,
	}

	// Calculate center point from mSize
	textBox.mX = (textBox.mSize.Right + textBox.mSize.Left) / 2
	textBox.mY = (textBox.mSize.Top + textBox.mSize.Bottom) / 2
	textBox.mWidth = textBox.mSize.Right - textBox.mSize.Left
	textBox.mHeight = textBox.mSize.Bottom - textBox.mSize.Top

	return &textBox
}

func (textBox *TextBox) Update(dt time.Duration) {
	textBox.mTime = textBox.mTime + dt
	textBox.mAppearTween.Update(Float(textBox.mTime.Seconds()))
	if textBox.IsDead() {
		textBox.mStack.Pop()
	}
}

func (textBox *TextBox) HandleInput() {
	if Keyboard.JustPressed(sf.KeySpace) {
		textBox.OnClick()
	} else if textBox.mSelectionMenu != nil {
		textBox.mSelectionMenu.HandleInput()
	}
}

func (textBox *TextBox) Enter() {}

func (textBox *TextBox) Exit() {
	if textBox.mDoClickCallback {
		if textBox.mSelectionMenu != nil {
			textBox.mSelectionMenu.OnClick()
		}
	}

	if textBox.mOnFinish != nil {
		textBox.mOnFinish()
	}
}

func (textBox *TextBox) OnClick() {
	if textBox.mSelectionMenu != nil {
		textBox.mDoClickCallback = true
	}

	if textBox.mChunkIndex >= len(textBox.mChunks) {
		// If dialog is appearing or dissappearing ignore
		if !textBox.mAppearTween.IsFinished() && textBox.mAppearTween.Value() == 1 {
			return
		}
		textBox.mAppearTween = NewTween(1, 0, .2, EaseInCirc)
	} else {
		textBox.mChunkIndex = textBox.mChunkIndex + 1
	}
}

func (textBox *TextBox) IsDead() bool {
	return textBox.mAppearTween.IsFinished() && textBox.mAppearTween.Value() <= 0
}

func (textBox *TextBox) Render(target sf.RenderTarget, renderStates sf.RenderStates) {
	scale := textBox.mAppearTween.Value()
	textScale := textBox.mTextScale * scale
	text := sf.NewText(Resource.FindFont("resources/sansation.tff"))
	text.Scale(sf.Vector2f{X: textScale, Y: textScale})

	textBox.mPanel.CenterPosition(textBox.mX, textBox.mY, textBox.mWidth*scale, textBox.mHeight*scale)
	textBox.mPanel.Render(target, renderStates)

	left := textBox.mX - (textBox.mWidth / 2 * scale)
	textLeft := left + (textBox.mBounds.Left * scale)
	top := textBox.mY - (textBox.mHeight / 2 * scale)
	textTop := top - (textBox.mBounds.Top * scale)
	bottom := textBox.mY + (textBox.mHeight / 2 * scale)

	// TODO fix
	text.SetPosition(sf.Vector2f{X: textLeft, Y: textTop})
	text.SetString(str)
	str := textBox.mChunks[textBox.mChunkIndex]
	for i := textBox.mWrap; i < len(str); i += textBox.mWrap {
		str = str[:i] + "\n" + str[i:]
	}
	text.Draw(target, renderStates)

	if textBox.mSelectionMenu != nil {
		menuX := textLeft
		menuY := bottom - textBox.mSelectionMenu.GetHeight()
		menuY = menuY - textBox.mBounds.Bottom
		textBox.mSelectionMenu.mX = menuX
		textBox.mSelectionMenu.mY = menuY
		textBox.mSelectionMenu.mScale = scale
		textBox.mSelectionMenu.Render(target, renderStates)
	}

	if textBox.mChunkIndex < len(textBox.mChunks)-1 {
		// There are more chuncks
		offset := float32(12 + math.Floor(math.Sin(float64textBox.mTime*10))*scale.AsFloat64())
		textBox.mContinueMark.SetScale(sf.Vector2f{X: scale.AsFloat32(), Y: scale.AsFloat32()})
		textBox.mContinueMark.SetPosition(sf.Vector2f{X: textBox.mX.AsFloat32(), Y: bottom.AsFloat32() + offset})
		textBox.mContinueMark.Draw(target, renderStates)
	}

	for k, v := range textBox.mChildren {
		switch v.Obj.(type) {
		case *sf.Text:
			v.Obj.SetPosition(sf.Vector2f{X: textLeft + (v.X * scale), Y: textTop + (v.Y * scale)})
		case *sf.Sprite:
			v.Obj.SetPosition(sf.Vector2f{X: left + (v.X * scale), Y: top + (v.Y * scale)})
			v.Obj.SetScale(sf.Vector2f{X: scale, Y: scale})
			v.Obj.Draw(target, renderStates)
		}
	}
}
