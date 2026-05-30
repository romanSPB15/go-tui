# go-tui: Text User Interface на Go
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)
[![Go version](https://img.shields.io/github/go-mod/go-version/romanSPB15/tui)](https://github.com/romanSPB15/go-tui)

## Обзор
Лёгкая библиотека для TUI на Golang. Умеет:
* Создавать надписи
* Создавать кнопки и настраивать
* Обрабатывать нажатия клавиатуры
* Красить текст
## Установка
```
go get -u github.com/...
```
## Пример использования
```go
package main

import (
	"github.com/eiannone/keyboard"
	"github.com/romanSPB15/go-tui"
)

func main() {
	a := tui.NewApp()
	a.AddComponents(tui.NewStaticLabel("Привет, Go!"))
	btn := tui.NewButton("Нажми ↑", keyboard.KeyArrowUp)
	btn.OnClick = a.Quit
	a.AddComponents(btn)
	a.Run()
}
```