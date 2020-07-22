package window

import (
	"git.agehadev.com/elliebelly/gooey/pkg/dimension"
	"git.agehadev.com/elliebelly/gooey/pkg/window/widget"
	"github.com/go-gl/glfw/v3.3/glfw"
)

type Window struct {
	glfwWindow *glfw.Window
	Layout     *widget.LinearLayout
	Context    WindowContext
}

func (w *Window) GetRectRelative() *dimension.Rect {
	return &w.Context.Rect
}

func (w *Window) create(preferences Preferences) error {
	for pref, value := range preferences.GLFWHints {
		glfw.WindowHint(pref, value)
	}

	glfwWindow, err := glfw.CreateWindow(preferences.Width, preferences.Height, "Jamboy", nil, nil)

	if err != nil {
		return err
	}

	w.glfwWindow = glfwWindow

	w.Context.Resolution.Width, w.Context.Resolution.Height = w.glfwWindow.GetSize()
	w.glfwWindow.SetSizeCallback(w.onWindowSetSize)

	return nil
}

func (w *Window) onWindowSetSize(glfwWindow *glfw.Window, width int, height int) {
	w.Context.Resolution.Width = width
	w.Context.Resolution.Height = height
}

func (w *Window) MakeCurrent() {
	Context = &w.Context
	w.glfwWindow.MakeContextCurrent()
}

func (w *Window) Tick() {
	w.glfwWindow.SwapBuffers()
	glfw.PollEvents()
}

func (w *Window) ShouldClose() bool {
	return w.glfwWindow.ShouldClose()
}

func (w *Window) close() {
	w.glfwWindow.Destroy()
}

func (w *Window) SetLayout(linearLayout *widget.LinearLayout) *widget.LinearLayout {
	linearLayout.Parent = w

	w.Layout = linearLayout

	return linearLayout
}

func newWindow() *Window {
	w := &Window{
		Layout: widget.NewLinearLayout(),
	}

	w.Context.Rect = dimension.Rect{
		X:      0,
		Y:      0,
		Width:  1,
		Height: 1,
	}

	return w
}
