package tui

import (
	"fmt"
	"time"

	"github.com/eiannone/keyboard"
)

type Color int

const (
	Black Color = iota + 30
	Red
	Green
	Yellow
	Blue
	Magenta
	Cyan
	White
)

const (
	BrightBlack Color = iota + 90
	BrightRed
	BrightGreen
	BrightYellow
	BrightBlue
	BrightMagenta
	BrightCyan
	BrightWhite
)

type DisplayMode int

const (
	DisplayInline DisplayMode = iota
	DisplayBlock
	DisplayNewLine
)

type Label struct {
	decoration string
	Text       string
	maxLength  int
}

func (l *Label) innerText() string {
	if l.decoration == "" {
		return l.Text
	}
	return fmt.Sprintf("%s%s\033[0m", l.decoration, l.Text)
}

func NewLabel(txt string, maxLength int) *Label { return &Label{Text: txt, maxLength: maxLength} }

func (lbl *Label) Colorize(clr Color) *Label {
	lbl.decoration += fmt.Sprintf("\033[%dm", clr)
	return lbl
}

func (lbl *Label) Bold() *Label {
	lbl.decoration += "\033[1m"
	return lbl
}

func (lbl *Label) Italic() *Label {
	lbl.decoration += "\033[3m"
	return lbl
}

func (lbl *Label) Underline() *Label {
	lbl.decoration += "\033[4m"
	return lbl
}

func (lbl *Label) Reverse() *Label {
	lbl.decoration += "\033[7m"
	return lbl
}

func (lbl *Label) MaxWidth() int {
	return lbl.maxLength
}

func (lbl *Label) DisplayMode() DisplayMode {
	return DisplayInline
}

func (l *Label) setIndex(int) {}

/////////////////////////////////////////////////////////////////////////////////

type spaser struct{}

func (l *spaser) innerText() string {
	return ""
}

func (l *spaser) DisplayMode() DisplayMode {
	return DisplayBlock
}

func (l *spaser) MaxWidth() int {
	return 0
}

func (l *spaser) setIndex(int) {}

func Spaser() Component { return &spaser{} }

/////////////////////////////////////////////////////////////////////////////////

type newLine struct{}

func (l *newLine) innerText() string {
	return ""
}

func (l *newLine) DisplayMode() DisplayMode {
	return DisplayNewLine
}

func (l *newLine) MaxWidth() int {
	return 0
}

func (l *newLine) setIndex(int) {}

func NewLine() Component { return &newLine{} }

type Button struct {
	clicked Component
	base    Component
	OnClick func()
	Component
	idx int
}

func NewButton(text string, key keyboard.Key) *Button {
	btn := &Button{
		clicked: NewLabel(text, len(text)).Colorize(Black),
		base:    NewLabel(text, len(text)),
		OnClick: func() {},
	}
	btn.Component = btn.base
	currentApp.AddKeyHandler(key, func() {
		if btn.OnClick != nil {
			btn.OnClick()
		}
		btn.Component = btn.clicked
		currentApp.RedrawComponent(btn.idx)
		currentApp.LogInfo("%d %s", btn.idx, btn.innerText())
		time.Sleep(time.Millisecond * 50)
		btn.Component = btn.base
		currentApp.LogInfo("%d %s", btn.idx, btn.innerText())
		currentApp.RedrawComponent(btn.idx)
	})
	return btn
}

func (btn *Button) setIndex(idx int) {
	btn.idx = idx
	btn.base.setIndex(idx)
	btn.clicked.setIndex(idx)
}
