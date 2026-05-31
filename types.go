// Библитека TUI для Go
package tui

import "github.com/eiannone/keyboard"

// Объект компонента приложения.
type Component interface {
	innerText() string

	MaxWidth() int
	DisplayMode() DisplayMode
	setIndex(int)
}

// Объект приложения
type App interface {
	Components() []Component
	Redraw()
	RedrawComponent(int)
	AddComponents(...Component)
	Clear()

	Run()
	Runned() bool

	Quit()
	OnQuit() <-chan struct{}

	Window() Window

	AddKeyHandler(key keyboard.Key, h func())

	LogInfo(message string, args ...any)
	LogFatal(message string, args ...any)
}

// Объект окна приложения
type Window interface {
	Width() int  // Ширина окна в символах
	Height() int // Высота окна в символах
}
