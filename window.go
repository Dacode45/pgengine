package pgengine

import sf "github.com/manyminds/gosfml"


type window struct {
	dt time.Duration
	Params WindowDef
	RenderWindow *sf.RenderWindow
}

type WindowDef struct {
  Width uint
  Height uint
}

var Window window

func StartWindow(params WindowDef){
  Window = window{
		Params: params,
    RenderWindow: sf.NewRenderWindow(sf.VideoMode{params.Width, params.Height, 32}, "Pong (GoSFML2)", sf.StyleDefault, sf.DefaultContextSettings())
  }
	renderWindow := Window.RenderWindow
	ticker := time.NewTicker(time.Second / 60)
	now := time.Now()
  go func () {
		for renderWindow.IsOpen() {
			select {
			case <- ticker.C:
				//Duration
				dt := time.Since(now)
				Window.dt = dt
				now = time.Now()
				// Clear Keyboard
				keyboard.Clear()
				for event := renderWindow.PollEvent(); event != nil; event = renderWindow.PollEvent() {
					switch ev := event.(type) {
					case sf.EventKeyPressed:
						Keyboard.KeyPressed(ev)
					}
				}
			}
		}
	}()
}

func GetDeltaTime() {

}

func ScreenCenterX() Int {
	return Int(Window.RenderWindow.GetSize().X)
}

func ScreenCenterY() Int {
	return Int(Window.RenderWindow.GetSize().Y)
}
