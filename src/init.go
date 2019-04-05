package main

func createBlockArray(
	screenDim int,
	totalPixels int,
	blockDim int) []block {

	numberOfBlocks := totalPixels / (blockDim * blockDim)
	blockArray := make([]block, numberOfBlocks)

	xOffset := 0
	yOffset := 0

	for i := 0; i < numberOfBlocks; i++ {
		blockArray[i] = initBlock(i, xOffset, yOffset, screenDim, blockDim)

		xOffset = xOffset + blockDim

		if xOffset >= screenDim {
			xOffset = 0
			yOffset = yOffset + blockDim
		}
	}

	return blockArray
}
