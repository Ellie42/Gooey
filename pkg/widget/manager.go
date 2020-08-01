package widget

import (
	"github.com/go-gl/glfw/v3.3/glfw"
)

type WindowManager struct {
	Windows     []*Window
	WindowCount int
}

func NewWindowManager() *WindowManager {
	return &WindowManager{
		Windows: make([]*Window, 0),
	}
}

func (m *WindowManager) CreateWindow(preferences WindowPreferences) (*Window, error) {
	window := newWindow()

	m.Windows = append(m.Windows, window)
	m.WindowCount++

	preferences.GLFWHints[glfw.ContextVersionMajor] = 4
	preferences.GLFWHints[glfw.ContextVersionMinor] = 6

	preferences.FillDefaults()

	err := window.create(preferences)

	if err != nil {
		return window, err
	}

	if preferences.OpenedCB != nil {
		preferences.OpenedCB()
	}

	return window, nil
}

func (m *WindowManager) Init() error {
	err := glfw.Init()

	if err != nil {
		return err
	}

	return nil
}

func (m WindowManager) Cleanup() {
	glfw.Terminate()
}

func (m *WindowManager) CloseWindow(w *Window) {
	w.close()
	m.WindowCount--
}
