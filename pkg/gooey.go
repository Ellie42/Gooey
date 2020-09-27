package gooey

import (
	"git.agehadev.com/elliebelly/gooey/pkg/draw"
	"git.agehadev.com/elliebelly/gooey/pkg/widget"
	"log"
	"os"
	"runtime"
	"runtime/pprof"
)

var Renderer *draw.Renderer

type Gooey struct {
	Window   *widget.WindowManager
	Renderer *draw.Renderer
	Stop     bool
}

func (g *Gooey) Loop() {
	defer g.cleanup()

	if len(g.Window.Windows) == 0 {
		panic("No windows have been created!")
	}

	g.Window.Windows[0].MakeCurrent()
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

			if !w.Initialised {
				w.Init()
			}

			w.MakeCurrent()

			g.Renderer.Clear(widget.Context.Resolution)

			w.RecalculateRect()
			w.Render()

			w.Tick()
		}
	}

	f, err := os.Create("dumps/memprofile")
	if err != nil {
		log.Fatal("could not create memory profile: ", err)
	}
	defer f.Close() // error handling omitted for example
	runtime.GC()    // get up-to-date statistics
	if err := pprof.WriteHeapProfile(f); err != nil {
		log.Fatal("could not write memory profile: ", err)
	}
}

func (g *Gooey) cleanup() {
	g.Window.Cleanup()
}

func Init() (gooey *Gooey, err error) {
	Renderer = draw.NewRenderer()
	gooey = &Gooey{
		Window:   widget.NewWindowManager(),
		Renderer: Renderer,
	}

	err = gooey.Window.Init()

	if err != nil {
		return
	}

	return
}
