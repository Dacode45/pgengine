package pgengine

import sf "github.com/manyminds/gosfml"

type ProgressBarDef struct {
	X       Int
	Y       Int
	Value   Number
	Maximum Number

	Background string
	Foreground string
}

type ProgressBar struct {
	mX          Int
	mY          Int
	mBackground *sf.Sprite
	mForeground *sf.Sprite
	mValue      Float
	mMaximum    Float

	mHalfWidth Int
}

func NewProgressBar(params ProgressBarDef) *ProgressBar {
	bar := &ProgressBar{
		mX:          params.X,
		mY:          params.Y,
		mBackground: sf.NewSprite(Resource.FindTexture(params.Background)),
		mForeground: sf.NewSprite(Resource.FindTexture(params.Foreground)),
		mValue:      params.Value,
		mMaximum:    DefaultFloat(params.Maximum, 1),
	}

	bar.mHalfWidth = Int(bar.mForeground.GetSize().X / 2)
	bar.SetValue(bar.mValue)
	return bar
}

func (bar *ProgressBar) SetValue(value, max Float) {
	bar.mMaximum = max
	bar.SetNormalValue(value / bar.mMaximum)
}

// TODO fill out
func (bar *ProgressBar) SetNormalValue(value Float) {
}

func (bar *ProgressBar) SetPosition(x, y Int) {
	bar.mX = x
	bar.mY = y
	position := sf.Vector2f{X: bar.mX, Y: bar.mY}
	bar.mBackground.SetPosition(position)
	bar.mForeground.SetPosition(position)
	// Make sure foreground position is set correctly
	bar.SetValue(bar.mValue, bar.mMaximum)
}

func (bar *ProgressBar) GetPosition() (Int, Int) {
	return bar.mX, bar.mY
}

func (bar *ProgressBar) Render(target sf.RenderTarget, renderStates sf.RenderStates) {
	bar.mBackground.Draw(target, renderStates)
	bar.mForeground.Draw(target, renderStates)
}
