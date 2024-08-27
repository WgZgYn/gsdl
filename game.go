package main

import (
	"fmt"
	"image/color"
	"math/rand"
	"time"

	"github.com/veandco/go-sdl2/sdl"
)

func rgba(c color.Gray) (r, g, b, a uint8) {
	return c.Y, c.Y, c.Y, 0xff
}

type Game struct {
	data Model
	rend *sdl.Renderer
	ra   *rand.Rand
	last time.Time
	cost time.Duration

	frame int64
}

func (g *Game) Init(render *sdl.Renderer, seed *int64) {
	g.data = NewModel(BoardRows, BoardCols)
	if seed != nil {
		g.ra = rand.New(rand.NewSource(*seed))
	} else {
		g.ra = rand.New(rand.NewSource(time.Now().UnixNano()))
	}
	g.rend = render
	g.data.random()
	g.last = time.Now()
}

func (g *Game) Run() {
	for {
		g.Clear()
		g.Draw()
		g.data.update()
		//fmt.Printf("cost: %v; FPS: %v/%v\n", g.cost, time.Second.Milliseconds(), g.cost.Milliseconds())

		if _, ok := sdl.PollEvent().(*sdl.QuitEvent); ok {
			break
		}
		g.rend.Present()
		if Delay {
			sdl.Delay(FrameTime)
		}
	}
}

func (g *Game) Draw() {
	_ = g.rend.SetDrawColor(rgba(White))
	_ = g.rend.DrawRect(&sdl.Rect{
		X: startX,
		Y: startY,
		W: BoardWidth,
		H: BoardHeight,
	})
	g.data.drawPixels(g.rend)
	now := time.Now()
	g.cost = now.Sub(g.last)
	g.last = now
	g.frame++
}

func (g *Game) Clear() {
	_ = g.rend.SetDrawColor(rgba(Black))
	_ = g.rend.Clear()
}

func Position(start sdl.Point, r, c int, size int32) (x, y int32) {
	return start.X + int32(c)*size, start.Y + int32(r)*size
}

func DrawPixel(x, y, size int32, render *sdl.Renderer) {
	if size > 1 {
		_ = render.FillRect(&sdl.Rect{X: x, Y: y, W: size, H: size})
	} else {
		_ = render.DrawPoint(x, y)
	}
}

func DrawPixelSurface(surface *sdl.Surface, x, y, size int32) {
	_ = surface.FillRect(&sdl.Rect{
		X: x,
		Y: y,
		W: size,
		H: size,
	}, 0x0000ff)
}

var ra *rand.Rand

func init() {
	seed := time.Now().UnixNano()
	fmt.Println("The World Seed is ", seed)
	ra = rand.New(rand.NewSource(seed))
}

func ClearWindow(ren *sdl.Renderer) {
	_ = ren.SetDrawColor(rgba(Black))
	_ = ren.FillRect(FullWindow)
}
