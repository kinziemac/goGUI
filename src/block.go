package main

import (
	"github.com/veandco/go-sdl2/sdl"
)

type block struct {
	pixels        []pixel
	isFilled      bool
	busy          bool
	owner         int
	coloredPixels int

	minColored int
	offsetX    int
	offsetY    int
	dimension  int
	blockID    int
}

//make block 50x50 -> 2500 pixels
//window Dimension
func initBlock(blockID int, xIndex int, yIndex int, windowD int) (b block) {

	//Configure fields for Block
	b.dimension = 50
	b.minColored = 1500
	b.isFilled = false
	b.owner = -1
	b.coloredPixels = 0
	b.minColored = 1500
	b.blockID = blockID

	//these must be set layer above, similar to loop below
	b.offsetX = xIndex
	b.offsetY = yIndex

	b.pixels = createPixelArray(b.offsetX, b.offsetY, b.dimension)

	return b
}

// func createPixelArray(offsetX int, offsetY int, dimension int) []sdl.Rect {
func createPixelArray(offsetX int, offsetY int, dimension int) []pixel {
	numberOfBlockPixels := dimension * dimension
	bottomBorder := numberOfBlockPixels - (2 * dimension)
	pixelArray := make([]pixel, numberOfBlockPixels)
	xCoor := 0
	yCoor := 0

	for i := 0; i < numberOfBlockPixels; i++ {
		xWithOffset := int32(xCoor + offsetX)
		yWithOffset := int32(yCoor + offsetY)
		pixelNumber := xCoor + yCoor*dimension

		// if the pixel is regular pixel or border pixel
		// arbitrarily chose if border pixels are 2 pixels from edge
		canChange := true
		if pixelNumber < 2*dimension ||
			(pixelNumber+1)%dimension == 0 ||
			(pixelNumber+2)%dimension == 0 ||
			(pixelNumber)%dimension == 0 ||
			(pixelNumber-1)%dimension == 0 ||
			pixelNumber > bottomBorder {

			canChange = false
		}

		pixelArray[i] = newPixel(canChange, sdl.Rect{xWithOffset, yWithOffset, 1, 1})
		xCoor = xCoor + 1

		if xCoor >= dimension {
			xCoor = 0
			yCoor = yCoor + 1
		}
	}

	return pixelArray
}

//renderBlock draws boxes on screen, they are either white and in the middle, or black and a border
func (b *block) renderBlock(renderer *sdl.Renderer) {
	for i := 0; i < len(b.pixels); i++ {
		if b.pixels[i].canChange {
			renderer.SetDrawColor(255, 255, 255, 255)
			renderer.FillRect(&b.pixels[i].val)

		} else {
			renderer.SetDrawColor(0, 0, 0, 255)
			renderer.FillRect(&b.pixels[i].val)
		}
	}
}

//drawOnBlock determines if a user can draw on a block
func (b *block) drawOnBlock(renderer *sdl.Renderer, mouseX int, mouseY int, blockDim int) {
	if b.isAllowed() {
		blockIndex := (mouseX - b.offsetX) + (mouseY-b.offsetY)*blockDim

		if b.pixels[blockIndex].canChange {
			renderer.SetDrawColor(255, 0, 0, 255)
			renderer.FillRect(&b.pixels[blockIndex].val)
			b.coloredPixels++
		}
	}
}

// isAllowed checks to see if block can be coloured - need to setup w/ network though
func (b *block) isAllowed() bool {
	return true
}

// blockFilled checks if minimum number of blocks are filled
func (b *block) blockFilled() bool {
	if b.coloredPixels > b.minColored {
		return true
	}

	return false
}

func (b *block) resetBlock(renderer *sdl.Renderer) {
	b.fillBlock(0, 0, 0, renderer)
	b.coloredPixels = 0
}

// fillBlock will either fill the block or undo the changes made by the player
func (b *block) fillBlock(red uint8, green uint8, blue uint8, renderer *sdl.Renderer) {
	for i := 0; i < len(b.pixels); i++ {
		if b.pixels[i].canChange {
			renderer.SetDrawColor(red, green, blue, 255)
			renderer.FillRect(&b.pixels[i].val)
		}
	}

}
