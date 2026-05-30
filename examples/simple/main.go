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
