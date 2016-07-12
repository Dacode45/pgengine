package engine

import (
	"time"

	sf "github.com/manyminds/gosfml"
)

type PGEngine struct {
	Ticker *time.Ticker
	Update func(dt time.Duration)

	window *sf.RenderWindow
	config *PGEngineConfig
}

type PGEngineConfig struct {
	windowTitle string
	VideoMode   sf.VideoMode
	ClearColor  sf.Color
}

func NewPGEngine(config PGEngineConfig) *PGEngine {
	e := PGEngine{config: &config}
	e.LoadDefaults()
	config = *e.config
	e.Ticker = time.NewTicker(time.Second / 60)
	//Window stuff
	e.window = sf.NewRenderWindow(config.VideoMode,
		config.windowTitle,
		sf.StyleDefault,
		sf.DefaultContextSettings())
	//MapInitialization
	return &e
}

func (e *PGEngine) GetRenderer() *sf.RenderWindow {
	return e.window
}

func (e *PGEngine) Run() {
	var (
		renderWindow = e.window
		frameTicker  = e.Ticker
		clearColor   = e.config.ClearColor
	)
	now := time.Now()
	for renderWindow.IsOpen() {
		select {
		case <-frameTicker.C:
			renderWindow.Clear(clearColor)
			e.Update(time.Since(now))
			now = time.Now()
			//Drawing
			renderWindow.Display()
		}
	}
}

func (e *PGEngine) LoadDefaults() {
	config := e.config
	mode := config.VideoMode
	if !mode.IsValid() {
		mode.Width = 800
		mode.Height = 600
		mode.BitsPerPixel = 32
	} else {
		if mode.Width == 0 {
			mode.Width = 800
		}
		if mode.Height == 0 {
			mode.Height = 600
		}
		if mode.BitsPerPixel == 0 {
			mode.BitsPerPixel = 32
		}
	}
	config.VideoMode = mode
}
