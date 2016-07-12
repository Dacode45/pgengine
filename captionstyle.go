package pgengine

import sf "github.com/manyminds/gosfml"

type CaptionStyle struct {
	Font          string
	CharacterSize Int
	Scale         Int
	X             Int
	Y             Int
	Color         sf.Color
	Width         Int
	TextStyle     sf.TextStyle

  Duration time.Duration
	RenderFunc    func (target, renderStates, style, text)
	ApplyFunc     func()
}

func DefaultRender(target sf.RenderTarget, renderStates sf.RenderStates, style CaptionStyle, text string) {
  text := sf.NewText(Resource.FindFont(style.Font))
	text.SetCharacterSize(style.CharacterSize)
	text.SetColor(style.Color)
	text.SetStyle(style.TextStyle)
  text.SetPosition(sf.Vector2f{X: style.X.AsFloat32(), Y: style.Y.AsFloat32()})
	text.SetScale(sf.Vector2f{X: style.Scale.AsFloat32(), Y: style.Scale.AsFloat32()})
	text.Draw(target, renderStates)
}

func FadeApply(target CaptionStyle, value Int) {
	target.Color.A = vaule
}

var CaptionStyles = [string]CaptionStyle{
  "default": CaptionStyle{
    Font: "default",
    Scale: 1,
    Color: sf.ColorWhite(),
    Width: -1,
    X: ScreenCenterX(),
    Y: ScreenCenterY() + 75,
    RenderFunc: DefaultRender,
    ApplyFunc: func() {}
  },
  "title": CaptionStyle{
    Scale: 3,
    X:ScreenCenterX(),
    Y: ScreenCenterY(),
    RenderFunc: DefaultRender,
    ApplyFunc: FadeApply,
    Duration: 3 * time.Second,
  },
  "subtitle": CaptionStyle{
    Scale: 1,
    X: ScreenCenterX()
    Y: ScreenCenterY() + 5,
    Color: sf.Color{R: 0.4, G: 0.38, B:0.39, A:1},
    RenderFunc: Render,
    ApplyFunc: FadeApply,
    Duration: 1 * time.Second,
  }
}
