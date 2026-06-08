package tui

type Component interface {
	innerText() string

	MaxWidth() int
	DisplayMode() DisplayMode
	setIndex(int)
}

type App interface {
	Components() []Component
	Redraw()
	RedrawComponent(int)
	AddComponents(...Component)
	Clear()
	Run()
	Quit()
	OnQuit() <-chan struct{}

	Window() Window

	LogInfo(message string, args ...any)
	LogFatal(message string)
}

type Window interface {
	Width() int
	Height() int
}
