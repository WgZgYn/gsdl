package main

import "github.com/veandco/go-sdl2/sdl"

type Render interface {
	Render(ren *Renderer)
}

type Renderer struct {
	ren *sdl.Renderer
}

func (r *Renderer) Done() {

}
