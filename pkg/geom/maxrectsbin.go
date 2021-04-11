package geom

// Ported to Go from the C++ implementation made by Jukka Jylänki (https://github.com/juj/RectangleBinPack/)

import (
	"github.com/maxfish/go-libs/pkg/imath"
	"math"
)

type RectNode struct {
	Rect         // Width and Height include padding
	Index   int  // This index is used to keep track of which of the input rects this is
	Rotated bool // True if the rect has been rotated 90 degrees
}

func NewRectNode(index, width, height int) RectNode {
	return RectNode{
		Index: index,
		Rect:  Rect{W: width, H: height},
	}
}

func NewRectNodeFrom(b RectNode) RectNode {
	return RectNode{
		Index: b.Index,
		Rect:  b.Rect,
	}
}

type MaxRectsBinResult struct {
	PlacedRects    []RectNode              // Rects that have been placed
	NotPlacedRects []RectNode              // Rects that didn't fit in the provided area
	Width          int                     // Total Width used by the rects
	Height         int                     // Total Height used by the rects
	Method         FreeRectChoiceHeuristic // Method used to pack the rects
}

type FreeRectChoiceHeuristic int

const (
	// BSSF: Positions the rectangle against the short side of a free rectangle into which it fits the best.
	RectBestShortSideFit FreeRectChoiceHeuristic = iota
	// BAF: Positions the rectangle into the smallest free rect into which it fits.
	RectBestAreaFit
	// BL: Does the Tetris placement.
	RectBottomLeftRule
	// BLSF: Positions the rectangle against the long side of a free rectangle into which it fits the best.
	// RectBestLongSideFit
	// CP: Chooses the placement where the rectangle touches other rects as much as possible.
	// RectContactPointRule
)

type MaxRectsBinPacker struct {
	binWidth, binHeight int
	paddingX, paddingY  int
	allowRotation       bool
	usedRectangles      []RectNode
	freeRectangles      []RectNode
}

// Create a new MaxRectsBinPacker packer
func NewMaxRectsBinPacker(width, height int, paddingX, paddingY int, allowRotation bool) *MaxRectsBinPacker {
	mr := &MaxRectsBinPacker{
		binWidth:       width,
		binHeight:      height,
		paddingX:       paddingX,
		paddingY:       paddingY,
		allowRotation:  allowRotation,
		usedRectangles: make([]RectNode, 0),
		freeRectangles: make([]RectNode, 0),
	}

	r := NewRectNode(-1, width, height)
	mr.freeRectangles = append(mr.freeRectangles, r)
	return mr
}

func (mr *MaxRectsBinPacker) Pack(inputRects []RectNode, method FreeRectChoiceHeuristic) *MaxRectsBinResult {
	rects := inputRects
	for len(rects) > 0 {
		bestRectIndex := -1
		var bestNode RectNode
		bestScore1 := math.MaxInt32
		bestScore2 := math.MaxInt32

		for i := 0; i < len(rects); i++ {
			var score1, score2 int
			newNode := mr.scoreRect(rects[i], method, &score1, &score2)
			if score1 < bestScore1 || (score1 == bestScore1 && score2 < bestScore2) {
				bestScore1 = score1
				bestScore2 = score2
				bestNode = newNode
				bestNode.Index = rects[i].Index
				bestRectIndex = i
			}
		}

		if bestRectIndex == -1 {
			break
		}

		mr.placeRect(bestNode)
		rects = append(rects[:bestRectIndex], rects[bestRectIndex+1:]...)
	}

	// Put the result together
	result := &MaxRectsBinResult{
		PlacedRects:    mr.usedRectangles,
		NotPlacedRects: rects,
	}
	// Compute the total area used by the rects
	for i := 0; i < len(mr.usedRectangles); i++ {
		rect := mr.usedRectangles[i]
		result.Width = imath.Max(result.Width, rect.Right())
		result.Height = imath.Max(result.Height, rect.Bottom())
	}
	return result
}

func (mr *MaxRectsBinPacker) placeRect(node RectNode) {
	numRectanglesToProcess := len(mr.freeRectangles)
	for i := 0; i < numRectanglesToProcess; i++ {
		if mr.splitFreeNode(mr.freeRectangles[i], node) {
			mr.freeRectangles = append(mr.freeRectangles[:i], mr.freeRectangles[i+1:]...) // remove index 'i'
			i--
			numRectanglesToProcess--
		}
	}

	mr.pruneFreeList()
	mr.usedRectangles = append(mr.usedRectangles, node)
}

func (mr *MaxRectsBinPacker) scoreRect(rect RectNode, method FreeRectChoiceHeuristic, score1, score2 *int) RectNode {
	width := rect.W + mr.paddingX
	height := rect.H + mr.paddingY
	rotatedWidth := rect.H + mr.paddingX
	rotatedHeight := rect.W + mr.paddingY
	*score1 = math.MaxInt32
	*score2 = math.MaxInt32

	var newNode RectNode
	switch method {
	case RectBestShortSideFit:
		newNode = mr.findPositionForNewNodeBestShortSideFit(width, height, rotatedWidth, rotatedHeight, mr.allowRotation, score1, score2)
		break
	case RectBottomLeftRule:
		newNode = mr.findPositionForNewNodeBottomLeft(width, height, rotatedWidth, rotatedHeight, mr.allowRotation, score1, score2)
		break
	case RectBestAreaFit:
		newNode = mr.findPositionForNewNodeBestAreaFit(width, height, rotatedWidth, rotatedHeight, mr.allowRotation, score1, score2)
		break
	default:
		panic("Unknown free-rect choice heuristic")
	}

	// Cannot fit the current rectangle.
	if newNode.H == 0 {
		*score1 = math.MaxInt32
		*score2 = math.MaxInt32
	}

	return newNode
}

// Computes the ratio of used surface area
func (mr *MaxRectsBinPacker) Occupancy() float32 {
	usedSurfaceArea := 0
	for i := 0; i < len(mr.usedRectangles); i++ {
		usedSurfaceArea += mr.usedRectangles[i].W * mr.usedRectangles[i].H
	}
	return float32(usedSurfaceArea) / float32(mr.binWidth*mr.binHeight)
}

func (mr *MaxRectsBinPacker) findPositionForNewNodeBestAreaFit(width, height, rotatedWidth, rotatedHeight int, rotate bool, bestAreaFit, bestShortSideFit *int) RectNode {
	bestNode := RectNode{}
	*bestAreaFit = math.MaxInt32
	*bestShortSideFit = math.MaxInt32

	for i := 0; i < len(mr.freeRectangles); i++ {
		areaFit := mr.freeRectangles[i].W*mr.freeRectangles[i].H - width*height

		// Try to place the rectangle in upright (non-rotated) orientation
		if mr.freeRectangles[i].W >= width && mr.freeRectangles[i].H >= height {
			leftoverH := imath.Abs(mr.freeRectangles[i].W - width)
			leftoverV := imath.Abs(mr.freeRectangles[i].H - height)
			shortSideFit := imath.Min(leftoverH, leftoverV)

			if areaFit < *bestAreaFit || (areaFit == *bestAreaFit && shortSideFit < *bestShortSideFit) {
				bestNode.X = mr.freeRectangles[i].X
				bestNode.Y = mr.freeRectangles[i].Y
				bestNode.W = width
				bestNode.H = height
				*bestShortSideFit = shortSideFit
				*bestAreaFit = areaFit
				bestNode.Rotated = false
			}
		}

		if rotate && mr.freeRectangles[i].W >= rotatedWidth && mr.freeRectangles[i].H >= rotatedHeight {
			leftoverH := imath.Abs(mr.freeRectangles[i].W - rotatedWidth)
			leftoverV := imath.Abs(mr.freeRectangles[i].H - rotatedHeight)
			shortSideFit := imath.Min(leftoverH, leftoverV)

			if areaFit < *bestAreaFit || (areaFit == *bestAreaFit && shortSideFit < *bestShortSideFit) {
				bestNode.X = mr.freeRectangles[i].X
				bestNode.Y = mr.freeRectangles[i].Y
				bestNode.W = rotatedWidth
				bestNode.H = rotatedHeight
				*bestShortSideFit = shortSideFit
				*bestAreaFit = areaFit
				bestNode.Rotated = true
			}
		}
	}
	return bestNode
}

func (mr *MaxRectsBinPacker) findPositionForNewNodeBestShortSideFit(width, height, rotatedWidth, rotatedHeight int, rotate bool, bestShortSideFit, bestLongSideFit *int) RectNode {
	bestNode := RectNode{}
	*bestShortSideFit = math.MaxInt32
	*bestLongSideFit = math.MaxInt32

	for i := 0; i < len(mr.freeRectangles); i++ {
		// Try to place the rectangle in upright (non-rotated) orientation
		if mr.freeRectangles[i].W >= width && mr.freeRectangles[i].H >= height {
			leftoverH := imath.Abs(mr.freeRectangles[i].W - width)
			leftoverV := imath.Abs(mr.freeRectangles[i].H - height)
			shortSideFit := imath.Min(leftoverH, leftoverV)
			longSideFit := imath.Max(leftoverH, leftoverV)

			if shortSideFit < *bestShortSideFit || (shortSideFit == *bestShortSideFit && longSideFit < *bestLongSideFit) {
				bestNode.X = mr.freeRectangles[i].X
				bestNode.Y = mr.freeRectangles[i].Y
				bestNode.W = width
				bestNode.H = height
				*bestShortSideFit = shortSideFit
				*bestLongSideFit = longSideFit
				bestNode.Rotated = false
			}
		}

		if rotate && mr.freeRectangles[i].W >= rotatedWidth && mr.freeRectangles[i].H >= rotatedHeight {
			flippedLeftoverHoriz := imath.Abs(mr.freeRectangles[i].W - rotatedWidth)
			flippedLeftoverVert := imath.Abs(mr.freeRectangles[i].H - rotatedHeight)
			flippedShortSideFit := imath.Min(flippedLeftoverHoriz, flippedLeftoverVert)
			flippedLongSideFit := imath.Max(flippedLeftoverHoriz, flippedLeftoverVert)

			if flippedShortSideFit < *bestShortSideFit || (flippedShortSideFit == *bestShortSideFit && flippedLongSideFit < *bestLongSideFit) {
				bestNode.X = mr.freeRectangles[i].X
				bestNode.Y = mr.freeRectangles[i].Y
				bestNode.W = rotatedWidth
				bestNode.H = rotatedHeight
				*bestShortSideFit = flippedShortSideFit
				*bestLongSideFit = flippedLongSideFit
				bestNode.Rotated = true
			}
		}
	}

	return bestNode
}

func (mr *MaxRectsBinPacker) findPositionForNewNodeBottomLeft(width, height, rotatedWidth, rotatedHeight int, rotate bool, bestY, bestX *int) RectNode {
	bestNode := RectNode{}
	*bestX = math.MaxInt32
	*bestY = math.MaxInt32

	for i := 0; i < len(mr.freeRectangles); i++ {
		// Try to place the rectangle in upright (non-rotated) orientation
		if mr.freeRectangles[i].W >= width && mr.freeRectangles[i].H >= height {
			topSideY := mr.freeRectangles[i].Y + height
			if topSideY < *bestY || (topSideY == *bestY && mr.freeRectangles[i].X < *bestX) {
				bestNode.X = mr.freeRectangles[i].X
				bestNode.Y = mr.freeRectangles[i].Y
				bestNode.W = width
				bestNode.H = height
				*bestY = topSideY
				*bestX = mr.freeRectangles[i].X
				bestNode.Rotated = false
			}
		}
		if rotate && mr.freeRectangles[i].W >= rotatedWidth && mr.freeRectangles[i].H >= rotatedHeight {
			topSideY := mr.freeRectangles[i].Y + rotatedHeight
			if topSideY < *bestY || (topSideY == *bestY && mr.freeRectangles[i].X < *bestX) {
				bestNode.X = mr.freeRectangles[i].X
				bestNode.Y = mr.freeRectangles[i].Y
				bestNode.W = rotatedWidth
				bestNode.H = rotatedHeight
				*bestY = topSideY
				*bestX = mr.freeRectangles[i].X
				bestNode.Rotated = true
			}
		}
	}
	return bestNode
}

func (mr *MaxRectsBinPacker) splitFreeNode(freeNode, usedNode RectNode) bool {
	if !usedNode.Rect.Intersect(freeNode.Rect) {
		return false
	}

	if usedNode.X < freeNode.X+freeNode.W && usedNode.X+usedNode.W > freeNode.X {
		// New node at the top side of the used node
		if usedNode.Y > freeNode.Y && usedNode.Y < freeNode.Y+freeNode.H {
			newNode := NewRectNodeFrom(freeNode)
			newNode.H = usedNode.Y - newNode.Y
			mr.freeRectangles = append(mr.freeRectangles, newNode)
		}

		// New node at the bottom side of the used node
		if usedNode.Y+usedNode.H < freeNode.Y+freeNode.H {
			newNode := NewRectNodeFrom(freeNode)
			newNode.Y = usedNode.Y + usedNode.H
			newNode.H = freeNode.Y + freeNode.H - (usedNode.Y + usedNode.H)
			mr.freeRectangles = append(mr.freeRectangles, newNode)
		}
	}

	if usedNode.Y < freeNode.Y+freeNode.H && usedNode.Y+usedNode.H > freeNode.Y {
		// New node at the left side of the used node.
		if usedNode.X > freeNode.X && usedNode.X < freeNode.X+freeNode.W {
			newNode := NewRectNodeFrom(freeNode)
			newNode.W = usedNode.X - newNode.X
			mr.freeRectangles = append(mr.freeRectangles, newNode)
		}

		// New node at the right side of the used node
		if usedNode.X+usedNode.W < freeNode.X+freeNode.W {
			newNode := NewRectNodeFrom(freeNode)
			newNode.X = usedNode.X + usedNode.W
			newNode.W = freeNode.X + freeNode.W - (usedNode.X + usedNode.W)
			mr.freeRectangles = append(mr.freeRectangles, newNode)
		}
	}

	return true
}

// pruneFreeList Find the next rectangle that packs best.
func (mr *MaxRectsBinPacker) pruneFreeList() {
	n := len(mr.freeRectangles)
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			rect1 := mr.freeRectangles[i]
			rect2 := mr.freeRectangles[j]
			if rect1.IsContainedIn(rect2.Rect) {
				mr.freeRectangles = append(mr.freeRectangles[:i], mr.freeRectangles[i+1:]...)
				i--
				n--
				break
			}
			if rect2.IsContainedIn(rect1.Rect) {
				mr.freeRectangles = append(mr.freeRectangles[:j], mr.freeRectangles[j+1:]...)
				j--
				n--
			}
		}
	}
}
