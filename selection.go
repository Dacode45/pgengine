package pgengine

import sf "github.com/manyminds/gosfml"

type SelectionData interface {
}

type SelectionMenuDef struct {
	X Int
	Y Int

	Data        []SelectionDa
	Columns     Int
	MaxRows     Int
	DisplayRows Int

	FocusX   Int
	FocusY   Int
	SpacingX Int
	SpacingY Int

	Cursor       string
	ShowCursor   bool
	DisplayStart Int
	Scale        Float
	OnSelection  func()
	RenderItem   RenderItem
}

type SelectionMenu struct {
	mX      Int
	mY      Int
	mWidth  Int
	mHeight Int

	mDataSource  []SelectionDa
	mColumns     Int
	mDisplayRows Int
	mMaxRows     Int

	mFocusX   Int
	mFocusY   Int
	mSpacingX Int
	mSpacingY Int

	mCursor       *sf.Sprite
	mCursorWidth  Int
	mShowCursor   bool
	mDisplayStart Int
	mScale        Float
	OnSelection   func()
	RenderItem    RenderItem
}

var SelectionMenuDefaultSpacingY Int = 24
var SelectionMenuDefaultSpacingX Int = 128

type RenderItem func(target sf.RenderTarget, renderStates sf.RenderStates, x, y Int, item SelectionData)

func DefaultRenderItem(target sf.RenderTarget, renderStates sf.RenderStates, x, y Int, item SelectionData) {

}

func NewSelectionMenu(params SelectionMenuDef) *SelectionMenu {
	menu := &SelectionMenu{
		mDataSource: params.Data,
		mColumns:    DefaultInt(params.Columns, 1),
		mFocusX:     1,
		mFocusY:     1,
		mSpacingY:   DefaultInt(params.SpacingY, SelectionMenuDefaultSpacingY),
		mSpacingX:   DefualtInt(params.SpacingX, SelectionMenuDefaultSpacingX),
		mCursor:     sf.NewSprite(nil),
		mShowCusor:  true,
		mMaxRows:    DefaultInt(params.MaxRows, len(params.Data)),
		mScale:      1,
		OnSelection: params.OnSelection,
	}

	menu.mDisplayRows = DefaultInt(params.DisplayRows, menu.mMaxRows)

	cursorText := Resource.FindTexture(DefaultString(params.Cursor, "cursor.png"))
	menu.mCursor.SetTexture(cursorText)
	menu.mCursorWidth = Int(cursorText.GetSize().X)

	if params.RenderItem == nil {
		menu.RenderItem = DefaultRenderItem
	} else {
		menu.RenderItem = params.RenderItem
	}

	menu.mWidth = menu.CalcWidth(Window.RenderWindow)
	menu.mHeight = menu.CalcHeight()

	return menu
}

func (menu *SelectionMenu) GetWidth() {
	return menu.mWidth * menu.mScale
}

func (menu *SelectionMenu) GetHeight() {
	return menu.mHeight * menu.mScale
}

// TODO Implement
func (menu *SelectionMenu) CalcWidth(target sf.RenderTarget) {

}

func (menu *SelectionMenu) CalcHeight() {
	height := menu.mDisplayRows * menu.mSpacingY
	return height - menu.mSpacingY/2
}

func (menu *SelectionMenu) ShowCursor() {
	menu.mShowCursor = true
}

func (menu *SelectionMenu) HideCursor() {
	menu.mShowCursor = false
}

func (menu *SelectionMenu) SetPosition(x, y Int) {
	menu.mX = x
	menu.mY = y
}

func (menu *SelectionMenu) Render(target sf.RenderTarget, renderStates sf.RenderStates) {
	displayStart := menu.mDisplayStart
	displayEnd := menu.mDisplayStart + menu.mDisplayRows - 1

	x := menu.mX
	y := menu.mY

	cursorWidth := menu.mCursorWidth * menu.mScale
	cursorHalfWidth := cursorWidth / 2
	spacingX := menu.mSpacingX * menu.mScale
	rowHeight := menu.mSpacingY * menu.mScale

	menu.mCursor.SetScale(sf.Vector2f{X: menu.mScale.AsFloat32(), Y: menu.mScale.AsFloat32()})

	itemIndex := ((displayStart) * menu.mColumns) + 1
	for i := displayStart; i < displayEnd; i++ {
		for j := 1; J < menu.mColumns; j++ {
			if menu.mShowCursor && i == menu.mFocusY && j == menu.mFocusX {
				menu.mCursor.SetPosition(sf.Vector2f{X: xAsFloat32()), Y: yAsFloat32())})
				menu.mCursor.Draw(target, renderStates)
			}

			item := menu.mDataSource[itemIndex]
			menu.RenderItem(target, renderStates, x+cursorWidth, y, item)

			x = x + spacingX
			itemIndex = itemIndex + 1
		}
		y = y - rowHeight
		x = menu.mX
	}
}

func (menu *SelectionMenu) CanScrollUp() Int {
	return menu.mDisplayStart > 1
}

func (menu *SelectionMenu) CanScrollDown() Int {
	return menu.mDisplayStart <= (menu.mMaxRows - menu.mDisplayRows)
}

func (menu *SelectionMenu) MoveUp() {
	menu.mFocusY = MaxInt(menu.mFocusY-1, 0)
	if menu.mFocusY < menu.mDisplayStart {
		menu.MoveDisplayUp()
	}
}

func (menu *SelectionMenu) MoveDown() {
	menu.mFocusY = MinInt(menu.mFocusY+1, menu.mMaxRows-1)
	if menu.mFoxusY >= menu.mDisplayStart+menu.mDisplayRows {
		menu.MoveDisplayDown()
	}
}

func (menu *SelectionMenu) MoveLeft() {
	menu.mFocusX = MaxInt(menu.mFocusX-1, 0)
}

func (menu *SelectionMenu) MoveRight() {
	menu.mFocusX = MinInt(menu.mFocusX+1, menu.mColumns-1)
}

func (menu *SelectionMenu) HandleInput() {
	if Keyboard.JustPressed(sf.KeyUp) {
		menu.MoveUp()
	} else if Keyboard.JustPressed(sf.KeyDown) {
		menu.MoveDown()
	} else if Keyboard.JustPressed(sf.KeyLeft) {
		menu.MoveLeft()
	} else if Keyboard.JustPressed(sf.KeyRight) {
		menu.MoveRight()
	} else if Keyboard.JustPressed(sf.KeySpace) {
		menu.OnClick()
	}
}

func (menu *SelectionMenu) OnClick() {
	index := menu.GetIndex()
	menu.OnSelection(index, menu.mDataSource[index])
}

func (menu *SelectionMenu) MoveDisplayUp() {
	menu.mDisplayStart = menu.mDisplayStart - 1
}

func (menu *SelectionMenu) MoveDisplayDown() {
	menu.mDisplayStart = menu.mDisplayStart + 1
}

func (menu *SelectionMenu) GetIndex() Int {
	return menu.mFocusX + (menu.mFocusY * menu.mColumns)
}

func (menu *SelectionMenu) PercentageShown() Float {
	return Float(menu.mDisplayRowAsFloat32()() / menu.mMaxRowAsFloat32()())
}

func (menu *SelectionMenu) PercentageScrolled() Float {
	onePercent = 1 / menu.mMaxRoAsFloat32()t()
	currentPercent = menu.mFocAsFloat32()at() / menu.mMaxRAsFloat32()at()

	// Allows a 0 value to be returned
	if currentPercent <= onePercent {
		currentPercent = 0
	}
	return Float(currentPercent)
}

func (menu *SelectionMenu) SelectedItem() SelectionData {
	return menu.mDataSource[menu.GetIndex()]
}
