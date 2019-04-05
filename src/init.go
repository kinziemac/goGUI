package main

import (
	"github.com/veandco/go-sdl2/sdl"
)

func createRectArray(
	screenWidth int,
	screenHeight int,
	screenDimensions int) []sdl.Rect {

	rectArray := make([]sdl.Rect, screenDimensions)
	xCoor := 0
	yCoor := 0
	for i := 0; i < screenDimensions; i++ {
		rectArray[i] = sdl.Rect{int32(xCoor), int32(yCoor), 1, 1}
		xCoor = xCoor + 1

		// Basically moving left to right and then reset a row down
		if xCoor >= screenWidth {
			xCoor = 0
			yCoor = yCoor + 1
		}
	}

	return rectArray
}
