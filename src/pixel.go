package main

import (
	"github.com/veandco/go-sdl2/sdl"
)

type pixel struct {
	canChange bool
	val       sdl.Rect
}

func newPixel(canChange bool, value sdl.Rect) (p pixel) {
	p.canChange = canChange
	p.val = value

	return p
}

// rectArray := createRectArray(
// 	screenWidth,
// 	screenHeight,
// 	screenDimensions,
// 	boxSize)

// func createRectArray(
// 	screenWidth int,
// 	screenHeight int,
// 	screenDimensions int,
// 	boxSize int) []sdl.Rect {

// 	rectArray := make([]sdl.Rect, screenDimensions)
// 	xCoor := 0
// 	yCoor := 0
// 	for i := 0; i < screenDimensions; i++ {
// 		rectArray[i] = sdl.Rect{int32(xCoor), int32(yCoor), int32(boxSize), int32(boxSize)}
// 		xCoor = xCoor + boxSize

// 		// Basically moving left to right and then reset a row down
// 		if xCoor >= screenWidth {
// 			xCoor = 0
// 			yCoor = yCoor + boxSize
// 		}
// 	}

// 	return rectArray
// }
