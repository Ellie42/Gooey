package widget

import (
	"git.agehadev.com/elliebelly/gooey/lib/dimension"
	"github.com/go-gl/gl/v4.6-core/gl"
	"github.com/go-gl/glfw/v3.3/glfw"
)

type Window struct {
	glfwWindow  *glfw.Window
	Layout      Widget
	Context     WindowContext
	Initialised bool
}

func (w *Window) GetRectAbsolute() dimension.Rect {
	return w.Context.Rect
}

func (w *Window) GetChildRectAbsolute(index int) dimension.Rect {
	return w.GetRectAbsolute()
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
	w.Context.GLFWWindow = glfwWindow
	w.Context.Resolution.Width, w.Context.Resolution.Height = w.glfwWindow.GetSize()
	w.glfwWindow.SetSizeCallback(w.onWindowSetSize)

	return nil
}

func (w *Window) onWindowSetSize(glfwWindow *glfw.Window, width int, height int) {
	w.Context.Resolution.Width = width
	w.Context.Resolution.Height = height
	gl.Viewport(0, 0, int32(width), int32(height))
}

func (w *Window) MakeCurrent() {
	Context = &w.Context
	w.glfwWindow.MakeContextCurrent()
}

func (w *Window) Init() {
	w.Context.Input.Init(w.glfwWindow)
	Context.Input.OnClick(w.Context.EventManager.HandleClickCollisions)
	Context.Input.OnMouseUp(w.Context.EventManager.HandleMouseUp)
	w.Context.EventManager.Init(w, w.glfwWindow)
	w.Layout.Init()
	w.Initialised = true
}

func (w *Window) GetTotalArea() dimension.SizeFloat32 {
	return dimension.SizeFloat32{
		float32(w.Context.Resolution.Width),
		float32(w.Context.Resolution.Height),
	}
}

func (w *Window) Tick() {
	w.glfwWindow.SwapBuffers()
	w.Context.Input.Tick()
	glfw.PollEvents()
}

func (w *Window) ShouldClose() bool {
	return w.glfwWindow.ShouldClose()
}

func (w *Window) close() {
	w.glfwWindow.Destroy()
}

func (w *Window) Render() {
	w.Layout.Render()
}

func newWindow() *Window {
	layout := NewFreeLayout(nil)

	w := &Window{
		Layout: layout,
	}

	w.Context.Rect = dimension.Rect{
		X:      0,
		Y:      0,
		Width:  1,
		Height: 1,
	}

	return w
}
