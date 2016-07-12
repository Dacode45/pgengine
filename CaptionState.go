package pgengine

type CaptionState struct {
  mStyle CaptionStyle
  mText string
}

func (c *CaptionState) Enter() {}
func (c *CaptionState) Exit() {}
func (c *CaptionState) HandleInput() {}
func (c *CaptionState) Update(dt time.Duration) {
  return true
}
func (c *CaptionState) Render(target sf.RenderTarget, renderStates sf.RenderStates) {
  c.mStyle.RenderFunc(target, renderStates, c.mStyle, c.mText)
}

func NewCaptionState(params ...interface{}) return *CaptionState {
  style := params[0].(CaptionStyle)
  text := params[1].(string)
  return &CaptionState{
    mStyle: style,
    mText: text,
  }
}
