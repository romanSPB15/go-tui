//go:build windows

package tui

import "time"

func (w *window) startScreenResizeChecker() {
	prevW, prevH := w.Width(), w.Height()
	ticker := time.NewTicker(100 * time.Millisecond)
	go func() {
		defer ticker.Stop()
		for {
			select {
			case <-ticker.C:
				newW, newH := w.Width(), w.Height()
				if newW != prevW || newH != prevH {
					prevW, prevH = newW, newH
					w.doWithMessageAndWait(func() {
						w.currentPos = pos{0, 0}
						w.index()
						w.Redraw()
					}, "window resize (Windows)")
				}
			case <-w.stopCh:
				return
			}
		}
	}()
}
