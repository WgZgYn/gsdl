package main

import "github.com/veandco/go-sdl2/sdl"

type Board [][]bool

func (b Board) forEach(f func(r, c int)) {
	for x := 0; x < len(b); x++ {
		for y := 0; y < len(b[x]); y++ {
			f(x, y)
		}
	}
}

type Model struct {
	last, curr    Board
	width, height int
}

func NewBoard(w, h int) Board {
	b := make([][]bool, h)
	for i := range b {
		b[i] = make([]bool, w)
	}
	return b
}

func NewModel(w, h int) Model {
	return Model{
		last:   NewBoard(w, h),
		curr:   NewBoard(w, h),
		width:  w,
		height: h,
	}
}

func (m *Model) forEach(f func(r, c int)) {
	//for r := 0; r < m.height; r++ {
	//	for c := 0; c < m.width; c++ {
	//		f(r, c)
	//	}
	//}
	m.curr.forEach(f)
}

func (m *Model) random() {
	m.forEach(func(r, c int) {
		m.curr[r][c] = ra.Intn(100) < LivePercent
	})
}

func (m *Model) update() {
	m.last, m.curr = m.curr, m.last

	m.forEach(func(r, c int) {
		live := m.count(r, c)
		m.curr[r][c] = live == 3 || live == 2 && m.last[r][c]
	})
}

func (m *Model) count(r, c int) (live int) {
	for x := max(0, r-1); x <= min(r+1, m.height-1); x++ {
		for y := max(0, c-1); y <= min(c+1, m.height-1); y++ {
			if m.last[x][y] {
				live++
			}
		}
	}
	if m.last[r][c] {
		live--
	}
	return
}

func (m *Model) drawPixel(r, c int, renderer *sdl.Renderer) {
	x, y := Position(sdl.Point{X: startX, Y: startY}, r, c, PixelSize)
	DrawPixel(x, y, PixelSize, renderer)
}

func (m *Model) drawPixels(renderer *sdl.Renderer) {
	m.forEach(func(r, c int) {
		if m.curr[r][c] {
			m.drawPixel(r, c, renderer)
		}
	})
}

func (m *Model) drawPixelsOnSurface(surface *sdl.Surface) {
	m.forEach(func(r, c int) {
		x, y := Position(sdl.Point{X: startX, Y: startY}, r, c, PixelSize)
		DrawPixelSurface(surface, x, y, PixelSize)
	})
}
