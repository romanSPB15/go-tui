// Package tui provides a framework for creating Text User Interfaces (TUI) in Go.
//
// Библиотека go-tui позволяет легко создавать интерактивные TUI-приложения.
// Она включает в себя набор готовых компонентов: кнопки, надписи,
// цветной прогресс-бар и другие.
//
// Быстрый старт:
//
//	package main
// import (
// 	"github.com/eiannone/keyboard"
// 	"github.com/romanSPB15/go-tui"
// )

//	func main() {
//		a := tui.NewApp()
//		a.AddComponents(tui.NewStaticLabel("Привет, Go!"))
//		btn := tui.NewButton("Нажми ↑", keyboard.KeyArrowUp)
//		btn.OnClick = a.Quit
//		a.AddComponents(btn)
//		a.Run()
//	}
package tui
