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
}

func (i *Input) Init(w *glfw.Window) {
	i.glfwWindow = w
	w.SetMouseButtonCallback(i.onMouseButtonCallback)
}

func (i *Input) OnClick(handler func()) {
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
