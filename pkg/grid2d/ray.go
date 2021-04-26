package grid2d

type GridRayCallback func(tileX, tileY int, x, y float32) bool

func rayAxisSetup(cellSize int, pos, dir float32) (tile, deltaTile int, deltaDistance, ddt float32) {
	cellSizeF := float32(cellSize)
	tile = int(pos / cellSizeF)

	if dir == 0 {
		return tile, 0,0,0
	}

	if dir > 0 {
		deltaTile = 1
		deltaDistance = (float32(tile+1)*cellSizeF - pos) / dir
		ddt = cellSizeF / dir
	} else {
		deltaTile = -1
		deltaDistance = ((float32(tile))*cellSizeF - pos) / dir
		ddt = -cellSizeF / dir
	}
	return
}

func IterateRay(cellSize, gridWidth, gridHeight int, rayX, rayY, rayDirX, rayDirY float32, canContinue GridRayCallback) {
	if rayDirX == 0 && rayDirY == 0 {
		return
	}

	var distance float32
	tileX, deltaTileX, deltaDistanceX, ddtX := rayAxisSetup(cellSize, rayX, rayDirX)
	tileY, deltaTileY, deltaDistanceY, ddtY := rayAxisSetup(cellSize, rayY, rayDirY)

	for tileX >= 0 && tileY >= 0 && tileX < gridWidth && tileY < gridHeight {
		if !canContinue(tileX, tileY, rayX+rayDirX*distance, rayY+rayDirY*distance) {
			return
		}
		if deltaTileY == 0 || (deltaTileX != 0 && deltaDistanceX < deltaDistanceY) {
			distance += deltaDistanceX
			tileX += deltaTileX
			deltaDistanceY -= deltaDistanceX
			deltaDistanceX = ddtX
		} else {
			distance += deltaDistanceY
			tileY += deltaTileY
			deltaDistanceX -= deltaDistanceY
			deltaDistanceY = ddtY
		}
	}
}
