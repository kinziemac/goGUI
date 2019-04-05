package main

import (
	"github.com/veandco/go-sdl2/sdl"
)

type player struct {
	id           int
	currentBlock *block
	score        int
}

func colorRect(renderer *sdl.Renderer, rect *sdl.Rect) {
	renderer.SetDrawColor(0, 0, 255, 255)
	renderer.FillRect(rect)
}
