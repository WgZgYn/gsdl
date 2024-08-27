package main

import (
	"errors"
	"fmt"
	"math"
	"time"

	"github.com/veandco/go-sdl2/sdl"
)

type Handle struct {
	win *sdl.Window
	ren *sdl.Renderer
}

type Config struct {
	windowHeight int32
	windowWidth  int32
	renderFlag   uint32
}

type Circle struct {
	sdl.Point
	r int32
}

func (h *Handle) LoadBMP(file string) (*sdl.Texture, error) {
	bmp, err := sdl.LoadBMP(file)
	if err != nil {
		return nil, err
	}
	defer bmp.Free()
	return h.ren.CreateTextureFromSurface(bmp)
}

func (h *Handle) Destroy() error {
	fmt.Println("Exit SDL")
	defer sdl.Quit()
	return errors.Join(h.ren.Destroy(), h.win.Destroy())
}

func InitHandle(config *Config) (*Handle, error) {
	if err := sdl.Init(sdl.INIT_EVERYTHING); err != nil {
		return nil, err
	}
	window, err := sdl.CreateWindow("Test", sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_CENTERED, config.windowWidth, config.windowHeight, 0)
	if err != nil {
		return nil, err
	}
	renderer, err := sdl.CreateRenderer(window, -1, config.renderFlag)
	if err != nil {
		return nil, err
	}
	// window, renderer, err := sdl.CreateWindowAndRenderer(config.windowWidth, config.windowHeight, config.renderFlag)
	window.SetTitle("Game Of Life")
	return &Handle{
		win: window,
		ren: renderer,
	}, nil
}

func (h *Handle) DrawTexture(texture *sdl.Texture, dst *sdl.Point) {
	_ = h.ren.Clear()
	_, _, i2, i3, _ := texture.Query()
	_ = h.ren.Copy(texture, nil, &sdl.Rect{X: dst.X, Y: dst.Y, W: i2, H: i3})
	h.ren.Present()
}

func (h *Handle) DrawCircle(c *Circle, start, end float64, eps float64) {
	for degree := start; degree < end; degree += eps {
		dx, dy := float64(c.r)*math.Cos(degree), float64(c.r)*math.Sin(degree)
		_ = h.ren.DrawPointF(float32(c.X)+float32(dx), float32(c.Y)-float32(dy))
	}
	h.ren.Present()
}

func main() {
	handle, err := InitHandle(&Config{
		windowHeight: WindowHeight,
		windowWidth:  WindowWidth,
		renderFlag:   sdl.RENDERER_SOFTWARE,
	})
	if err != nil {
		panic(err)
	}
	defer exec(handle.Destroy)

	RunGame(handle)
	//RunFactory(handle)
}

func RunGame(handle *Handle) {
	var game Game
	game.Init(handle.ren, nil)
	start := time.Now()
	game.Run()
	defer func() {
		duration := time.Now().Sub(start)
		fmt.Println("frame: ", game.frame, ", duration: ", duration)
		fmt.Println(time.Duration(int64(duration)/game.frame), time.Second/time.Duration(int64(duration)/game.frame))
	}()
}

func RunFactory(handle *Handle) {
	design := NewDesign(DesignRows, DesignCols)
	design.Init(handle.ren)
	design.Run()
}

func init() {
	//go func() {
	//	expvar.Publish("GameData", &game.cost)
	//	http.Handle("/", expvar.Handler())
	//	log.Fatal(http.ListenAndServe(":8080", nil))
	//}()
}
