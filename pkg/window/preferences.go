package window

import "github.com/go-gl/glfw/v3.3/glfw"

type Preferences struct {
	GLFWHints map[glfw.Hint]int
	Width     int
	Height    int
	Title     string
	// Called when window has been opened and initialised
	OpenedCB func()
}

func (p *Preferences) FillDefaults() {
	if p.Width == 0 {
		p.Width = 384
	}

	if p.Height == 0 {
		p.Height = 260
	}

	if p.Title == "" {
		p.Title = "Gooey Display"
	}

	if p.GLFWHints == nil {
		p.GLFWHints = make(map[glfw.Hint]int, 0)
	}
}
