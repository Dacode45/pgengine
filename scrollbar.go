package pgengine

type ScrollbarDef struct {
	Texture string
	Height  Int
}

type Scrollbar struct {
	mX      Int
	mY      Int
	mWidth  Int
	mHeight Int
	mValue  Float

	mUpSprite         *sf.Sprite
	mDownSprite       *sf.Sprite
	mBackgroundSprite *sf.Sprite
	mCaretSprite      *sf.Sprite

	mCaretSize Float

	mLineHeight Int
	mTileHeight Int
	mUVs        []IntRect
}

var DefaultScrollbarHeight Int = 300

func NewScrollbar(params ScrollbarDef) *Scrollbar {
	text := Resource.FindTexture(params.Texture)
	bar := &Scrollbar{
		mHeight:  DefaultInt(params.Height, DefaultScrollbarHeight),
		mTexture: text,

		mUpSprite:         sf.NewSprite(text),
		mDownSprite:       sf.NewSprite(text),
		mBackgroundSprite: sf.NewSprite(text),
		mCaretSprite:      sf.NewSprite(text),

		mCaretSize: 1,
	}

	textWidth := text.GetSize().X
	textHeight := text.GetSize().Y
	// 4 equally sized pieses make up the scrollbar
	bar.mTileHeight = textHeight / 4
	bar.mUVs = GenerateUVs(text, textWidth, bar.mTileHeight)
	bar.mUpSprite.SetTextureRect(bar.mUVs[0])
	bar.mCaretSprite.SetTextureRect(bar.mUVs[1])
	bar.mBackgroundSprite.SetTextureRect(bar.mUVs[2])
	bar.mDownSprite.SetTextureRect(bar.mUVs[3])

	// Height to ignore the up and down arrows
	bar.mLineHeight = bar.mHeight - (bar.mTileHeight * 2)
}

func (bar *Scrollbar) SetPosition(x, y Int) {
	bar.mX = x
	bar.mY = y

	top := y - bar.mHeight/2
	bottom := y + bar.mHeight/2
	// halfTileHeight := bar.mTileHeight / 2

	bar.mUpSprite.SetPosition(sf.Vector2f{X: x, Y: top})
	bar.mDownSprite.SetPosition(sf.Vector2f{X: x, Y: bottom})

	bar.mBackgroundSprite.SetScale(sf.Vector2f{X: 1, Y: bar.mLineHeight / bar.mTileHeight})
	bar.mBackgroundSprite.SetPosition(sf.Vector2f{X: bar.mX, bar.mY})
	bar.SetNormalValue(bar.mValue)
}

// TODO Implement
func (bar *Scrollbar) SetNormalValue(v Float) {}

func (bar *Scrollbar) Render(target sf.RenderTarget, renderStates sf.RenderStates) {
	bar.mUpSprite.Draw(target, renderStates)
	bar.mBackgroundSprite.Draw(target, renderStates)
	bar.mDownSprite.Draw(target, renderStates)
	bar.mCaretSprite.Draw(target, renderStates)
}

func (bar *Scrollbar) SetScrollCaretScale(normalValue Float) {
	bar.mCaretSize = ((bar.mLineHeight) * normalValue) / bar.mTileHeight
	bar.mCaretSize = MaxFloat(1, bar.mCaretSize)
}
