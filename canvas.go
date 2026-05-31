package tui

import "fmt"

// Canvas — это многострочный виджет, на котором можно "рисовать".
type Canvas struct {
	width, height int
	pole          [][]Color
	idx           int
}

// NewCanvas() создаёт виждет Canvas.
func NewCanvas(width, height int) *Canvas {
	p := make([][]Color, height)
	for i := range height {
		p[i] = make([]Color, width)
	}
	return &Canvas{
		pole:   p,
		width:  width,
		height: height,
	}
}

// Draw() устанавливает указанный цвет в указанном месте Canvas.
func (c *Canvas) Draw(x, y int, clr Color) {
	if x < 0 || x >= c.width || y < 0 || y >= c.height {
		return
	}
	c.pole[x][y] = clr
}

// Draw() устанавливает указанный цвет в указанном месте Canvas, и перерисовывает.
func (c *Canvas) DrawAndRender(x, y int, clr Color) {
	if x < 0 || x >= c.width || y < 0 || y >= c.height {
		return
	}
	c.pole[x][y] = clr
	if currentApp.IsRunned() {
		currentApp.RedrawComponent(c.idx)
	}
}

// innerText() реализует интерфейс Component
func (c *Canvas) innerText() (res string) {
	lastClr := Color(-1)
	for y := 0; y < c.height; y++ {
		for x := 0; x < c.width; x++ {
			clr := c.pole[y][x]
			if lastClr != clr {
				res += fmt.Sprintf("\033[%dm", clr)
				lastClr = clr
			}
			res += " "
		}
		res += "\r\n"
	}
	res += "\033[0m"
	return
}

// MaxLength() реализует интерфейс Component
func (c *Canvas) MaxLength() int {
	return (c.width + 2) * c.height
}

// DisplayMode() реализует интерфейс Component
func (*Canvas) DisplayMode() DisplayMode {
	return DisplayBlock
}

// setIndex() реализует интерфейс Component
func (c *Canvas) setIndex(idx int) {
	c.idx = idx
}
