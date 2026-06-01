package main

import (
	"github.com/eiannone/keyboard"
	"github.com/romanSPB15/go-tui"
)

func main() {
	a := tui.NewApp()                                    // Создаём приложение
	a.AddWidgets(tui.NewStaticLabel("Привет, TUI!"))     // Добавляем надпись
	btn := tui.NewButton("Нажми ↑", keyboard.KeyArrowUp) // Создаём кнопку
	btn.OnClick = a.Quit                                 // Назначаем обработчик
	a.AddWidgets(btn)                                    // Добавляем кнопку
	a.Run()                                              // Запускаем приложение
}
