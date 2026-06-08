//go:build !windows

package tui

import (
	"os"
	"os/signal"
	"syscall"
)

func (w *window) startScreenResizeChecker() {
	sigwinch := make(chan os.Signal, 1)
	signal.Notify(sigwinch, syscall.SIGWINCH)
	go func() {
		for range sigwinch {
			w.doWithMessageAndWait(func() {
				w.currentPos = pos{0, 0}
				w.index()
				w.Redraw()
			}, "window resize (Unix)")
		}
	}()
}
