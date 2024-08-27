package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/veandco/go-sdl2/sdl"
)

const (
	DesignCols   = 10
	DesignRows   = 10
	DesignSize   = 40
	DesignWidth  = DesignCols * DesignSize
	DesignHeight = DesignRows * DesignSize
)

type Design struct {
	Board
	ren *sdl.Renderer

	lastClick *uint32
	sdl.Point

	running bool
}

func NewDesign(r, c int) Design {
	return Design{Board: NewBoard(r, c)}
}

func (d *Design) Save() error {
	if info, err := os.Stat("./design"); os.IsNotExist(err) || !info.IsDir() {
		err := os.Mkdir("./design", 0777)
		if err != nil {
			return err
		}
	}

	name := "demo"
	count := 0
	for info, err := os.Stat("./design/" + name + strconv.Itoa(count) + ".obj"); os.IsExist(err) && !info.IsDir(); count++ {
		info, err = os.Stat("./design/" + name + strconv.Itoa(count) + ".obj")
	}

	f, err := os.Create("./design/" + name + strconv.Itoa(count) + ".obj")
	if err != nil {
		return err
	}
	defer exec(f.Close)
	for _, i := range d.Board {
		err = Serialize(f, i)
		if err != nil {
			return err
		}
	}

	sdl.Log("successfully save the design")
	return nil
}

func (d *Design) Init(ren *sdl.Renderer) {
	d.Board.forEach(func(r, c int) {
		d.Board[r][c] = false
	})
	d.ren = ren
	d.running = true
	d.Draw()

	// go handleInput(d)
}

func (d *Design) Run() {
	defer exec(d.Save)

Begin:
	for {
		switch event := sdl.PollEvent().(type) {
		case *sdl.MouseButtonEvent:
			if event.Button == sdl.BUTTON_LEFT {
				sdl.Log("click, x:%v, y:%v", event.X, event.Y)
				d.HandleInput(event.X, event.Y, event.GetTimestamp(), event.State)
			}
		case *sdl.QuitEvent:
			d.running = false
			break Begin
		case *sdl.MouseMotionEvent:
			if d.lastClick != nil {
				x, y := event.X, event.Y
				a := &sdl.Point{X: x, Y: y}
				x, y = TopLeft(DesignWidth, DesignHeight)
				start := &sdl.Point{X: x, Y: y}
				r, c := Cell(a, start, DesignSize)
				d.set(r, c, true)
			}
		default:
			break
		}
		// sdl.Delay(FrameTime)// There is a bug on linux sysytem, if delay it couldn't handle too many events
	}
	// for d.running {
	// 	sdl.Delay(100) // Add a delay to prevent the loop from spinning and using 100% CPU.
	// }
	// sdl.Delay(FrameTime)
}

func (d *Design) Update(r, c int) {
	if r < DesignRows && c < DesignCols && r >= 0 && c >= 0 {
		d.Board[r][c] = !d.Board[r][c]
	}
	d.Draw()
}

func (d *Design) flip(r, c int) {
	if r < DesignRows && c < DesignCols && r >= 0 && c >= 0 {
		d.Board[r][c] = !d.Board[r][c]
	}
	d.Draw()
}

func (d *Design) set(r, c int, val bool) {
	if r < DesignRows && c < DesignCols && r >= 0 && c >= 0 && d.Board[r][c] != val {
		d.Board[r][c] = val
		d.Draw()
	}
}

func (d *Design) HandleInput(X, Y int32, time uint32, stat uint8) {
	if stat == sdl.PRESSED {
		d.X, d.Y = X, Y
		d.lastClick = &time
	} else {
		click := &sdl.Point{
			X: X,
			Y: Y,
		}
		x, y := TopLeft(DesignWidth, DesignHeight)
		sdl.Log("Layout: x: %v, y: %v", x, y)
		start := &sdl.Point{X: x, Y: y}
		if OneCell(&d.Point, click, start, DesignSize) {
			d.Update(Cell(click, start, DesignSize))
		}
		d.lastClick = nil
	}
}

func handleInput(d *Design) {
	for {
		switch event := sdl.PollEvent().(type) {
		case *sdl.MouseButtonEvent:
			if event.Button == sdl.BUTTON_LEFT {
				sdl.Log("click, x:%v, y:%v", event.X, event.Y)
				d.HandleInput(event.X, event.Y, event.GetTimestamp(), event.State)
			}
		case *sdl.QuitEvent:
			d.running = false
			return
		default:
			fmt.Println(event)
		}
	}
}

func OneCell(a, b *sdl.Point, start *sdl.Point, size int32) bool {
	x1, y1 := Cell(a, start, size)
	x2, y2 := Cell(b, start, size)
	return x1 == x2 && y1 == y2
}

func Cell(a *sdl.Point, start *sdl.Point, size int32) (x, y int) {
	if a.X-start.X < 0 || a.Y-start.Y < 0 {
		return -1, -1
	}
	return int((a.Y - start.Y) / size), int((a.X - start.X) / size)
}

func (d *Design) Draw() {
	startX, startY := TopLeft(DesignWidth, DesignHeight)

	_ = d.ren.SetDrawColor(rgba(Black))
	_ = d.ren.Clear()

	_ = d.ren.SetDrawColor(rgba(White))
	_ = d.ren.DrawRect(&sdl.Rect{
		X: startX,
		Y: startY,
		W: DesignWidth,
		H: DesignHeight,
	})

	for r := 1; r < DesignRows; r++ {
		_ = d.ren.DrawLine(startX, startY+int32(r*DesignSize), startX+DesignWidth, startY+int32(r*DesignSize))
	}

	for c := 1; c < DesignCols; c++ {
		_ = d.ren.DrawLine(startX+int32(c*DesignSize), startY, startX+int32(c*DesignSize), startY+DesignHeight)
	}

	rect := &sdl.Rect{W: DesignSize, H: DesignSize}
	d.Board.forEach(func(r, c int) {
		if d.Board[r][c] {
			rect.X, rect.Y = Position(sdl.Point{X: startX, Y: startY}, r, c, DesignSize)
			err := d.ren.FillRect(rect)
			if err != nil {
				panic(err)
			}
		}
	})

	d.ren.Present()
}
