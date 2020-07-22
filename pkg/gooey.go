package gooey

import (
	"git.agehadev.com/elliebelly/gooey/internal/renderer"
	"git.agehadev.com/elliebelly/gooey/pkg/window"
)

type Gooey struct {
	Window   *window.WindowManager
	Renderer *renderer.Renderer
	Stop     bool
}

func (g *Gooey) Loop() {
	defer g.cleanup()

	if len(g.Window.Windows) == 0 {
		panic("No windows have been created!")
	}

	g.Window.Windows[0].MakeCurrent()

	g.Renderer.Init()

	for g.Window.WindowCount > 0 && !g.Stop {
		for _, w := range g.Window.Windows {
			//TODO proper removal from array would be nice
			if w == nil {
				continue
			}

			if w.ShouldClose() {
				g.Window.CloseWindow(w)
				continue
			}

			if !w.Layout.Initialised {
				w.Layout.Init(w)
			}

			w.MakeCurrent()
			g.Renderer.Clear()

			w.Layout.Render()

			w.Tick()
		}
	}
}

func (g *Gooey) cleanup() {
	g.Window.Cleanup()
}

func Init() (gooey *Gooey, err error) {
	gooey = &Gooey{
		Window:   window.NewWindowManager(),
		Renderer: renderer.NewRenderer(),
	}

	err = gooey.Window.Init()

	if err != nil {
		return
	}

	return
}
