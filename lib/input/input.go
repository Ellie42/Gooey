package input

import (
	"github.com/go-gl/glfw/v3.3/glfw"
)

type Input struct {
	MouseUpHandlers []func()

	DownThisFrame    glfw.MouseButton
	Down             glfw.MouseButton
	glfwWindow       *glfw.Window
	onClickHandler   func()
	onMouseUpHandler func()
	onKeyDownHandler func(key glfw.Key)
	onKeyUpHandler   func(key glfw.Key)
	onScrollHandler  func(xoff float64, yoff float64)
}

func (i *Input) Init(w *glfw.Window) {
	i.glfwWindow = w
	w.SetMouseButtonCallback(i.onMouseButtonCallback)
	w.SetKeyCallback(i.onKeyButtonCallback)
	w.SetScrollCallback(i.onScrollCallback)
}

func (i *Input) OnMouseDown(handler func()) {
	i.onClickHandler = handler
}

func (i *Input) Tick() {
	i.DownThisFrame = 0
}

func (i *Input) onMouseButtonCallback(w *glfw.Window, button glfw.MouseButton, action glfw.Action, mods glfw.ModifierKey) {
	if action == glfw.Press {
		i.DownThisFrame |= button
		i.Down |= button
		if i.onClickHandler != nil {
			i.onClickHandler()
		}
	} else if action == glfw.Release {
		i.Down &= ^button

		if i.onMouseUpHandler != nil {
			i.onMouseUpHandler()
		}
	}
}

func (i *Input) OnMouseUp(handler func()) {
	i.onMouseUpHandler = handler
}

func (i *Input) onKeyButtonCallback(w *glfw.Window, key glfw.Key, scancode int, action glfw.Action, mods glfw.ModifierKey) {
	if action == glfw.Press {
		if i.onKeyDownHandler != nil {
			i.onKeyDownHandler(key)
		}
	} else if action == glfw.Release {
		if i.onKeyUpHandler != nil {
			i.onKeyUpHandler(key)
		}
	}
}

func (i *Input) OnKeyDown(handler func(key glfw.Key)) {
	i.onKeyDownHandler = handler
}

func (i *Input) OnKeyUp(handler func(key glfw.Key)) {
	i.onKeyUpHandler = handler
}

func (i *Input) OnScroll(handler func(xoff float64, yoff float64)) {
	i.onScrollHandler = handler
}

func (i *Input) onScrollCallback(w *glfw.Window, xoff float64, yoff float64) {
	i.onScrollHandler(xoff, yoff)
}
