package main

import (
	"image/color"

	"github.com/veandco/go-sdl2/sdl"
)

const (
	FPS       = 33
	FrameTime = 1000 / FPS
	Delay     = true
)

const (
	WindowWidth  = 640
	WindowHeight = 560
)

const (
	PixelSize   = 2
	BoardCols   = 200
	BoardRows   = 200
	BoardWidth  = BoardCols * PixelSize
	BoardHeight = BoardRows * PixelSize
)

const LivePercent = 15

// this make sure the board was center align
const (
	startX = (WindowWidth - BoardWidth) / 2
	startY = (WindowHeight - BoardHeight) / 2
)

var FullWindow = &sdl.Rect{
	X: 0,
	Y: 0,
	W: WindowWidth,
	H: WindowHeight,
}

var White = color.Gray{Y: 255}
var Black = color.Gray{Y: 0}

func TopLeft(width, height int32) (x, y int32) {
	return (WindowWidth - width) / 2, (WindowHeight - height) / 2
}
