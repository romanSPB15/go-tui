package tui

import (
	"fmt"
	"strings"
)

func recoveryScreen(message string) {
	fmt.Fprint(currentApp.f, "\033[2J\033[H")
	fmt.Fprint(currentApp.f, "\033[44m")
	w := currentApp.Window().Width()
	fmt.Fprintf(currentApp.f, fmt.Sprintf("%%-%ds", w), "Ошибка")
	fmt.Fprintf(currentApp.f, message+"\r\n")
	fmt.Fprintf(currentApp.f, "Нажмите ENTER для выхода...")
	for range currentApp.Window().Height() - 3 {
		fmt.Println(strings.Repeat(" ", w))
	}
	fmt.Fprint(currentApp.f, "\033[0m")
	fmt.Scanln()
}
