//go:build windows

package tui

import (
	"os"

	"golang.org/x/sys/windows"
)

func enableANSI() {
	stdout := windows.Handle(os.Stdout.Fd())
	var mode uint32
	if err := windows.GetConsoleMode(stdout, &mode); err != nil {
		return
	}
	mode |= windows.ENABLE_VIRTUAL_TERMINAL_PROCESSING
	_ = windows.SetConsoleMode(stdout, mode)
}
