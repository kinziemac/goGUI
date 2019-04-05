package main

import (
	"github.com/veandco/go-sdl2/sdl"
)

func createRectArray(
	screenDim int,
	totalPixels int) []sdl.Rect {

	rectArray := make([]sdl.Rect, totalPixels)
	xCoor := 0
	yCoor := 0
	for i := 0; i < totalPixels; i++ {
		rectArray[i] = sdl.Rect{int32(xCoor), int32(yCoor), 1, 1}
		xCoor = xCoor + 1

		// Basically moving left to right and then reset a row down
		if xCoor >= screenDim {
			xCoor = 0
			yCoor = yCoor + 1
		}
	}

	return rectArray
}

func createBlockArray(
	screenDim int,
	totalPixels int,
	blockDim int) []block {

	numberOfBlocks := totalPixels / (blockDim * blockDim)
	blockArray := make([]block, numberOfBlocks)

	xOffset := 0
	yOffset := 0

	for i := 0; i < numberOfBlocks; i++ {
		blockArray[i] = initBlock(i, xOffset, yOffset, screenDim)

		xOffset = xOffset + blockDim

		if xOffset >= screenDim {
			xOffset = 0
			yOffset = yOffset + blockDim
		}
	}

	return blockArray
}
