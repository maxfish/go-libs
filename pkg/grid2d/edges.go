package grid2d

import (
	"github.com/maxfish/go-libs/pkg/geom"
)

//cellEdges is only used during the construction phase of the edges.
type cellEdges struct {
	top, bottom, left, right *geom.Segment
}

type ComputeEdgesCallback func(x int, y int) bool

//ComputeEdges creates a list of segments covering all edges of the 2d grid.
//The segments coordinates assume each cell has a dimension of 1x1 units.
//The callback has to return 'true' for each cell which is considered solid, and for which an edge has to be computed.
func ComputeEdges(gridWidth, gridHeight int, isSolid ComputeEdgesCallback) []*geom.Segment {
	segments := make([]*geom.Segment, 0)
	edgesGrid := make([][]cellEdges, gridWidth)
	for i := range edgesGrid {
		edgesGrid[i] = make([]cellEdges, gridHeight)
	}

	for y := 0; y < gridHeight; y++ {
		for x := 0; x < gridWidth; x++ {
			if !isSolid(x, y) {
				continue
			}
			// Top segment
			if y-1 < 0 || !isSolid(x, y-1) {
				if x-1 < 0 || edgesGrid[x-1][y].top == nil {
					// Create a new segment
					edgesGrid[x][y].top = &geom.Segment{
						A: geom.Point{X: x, Y: y},
						B: geom.Point{X: x + 1, Y: y},
					}
					segments = append(segments, edgesGrid[x][y].top)
				} else {
					// Reuse, and extend, the segment of the cell at the left
					edgesGrid[x-1][y].top.B.X++
					edgesGrid[x][y].top = edgesGrid[x-1][y].top
				}
			}
			// Bottom segment
			if y+1 >= gridHeight || !isSolid(x, y+1) {
				if x-1 < 0 || edgesGrid[x-1][y].bottom == nil {
					edgesGrid[x][y].bottom = &geom.Segment{
						A: geom.Point{X: x + 1, Y: y + 1},
						B: geom.Point{X: x, Y: y + 1},
					}
					segments = append(segments, edgesGrid[x][y].bottom)
				} else {
					// Reuse, and extend, the segment of the cell at the left
					edgesGrid[x-1][y].bottom.A.X++
					edgesGrid[x][y].bottom = edgesGrid[x-1][y].bottom
				}
			}
			// Left segment
			if x-1 < 0 || !isSolid(x-1, y) {
				if y-1 < 0 || edgesGrid[x][y-1].left == nil {
					edgesGrid[x][y].left = &geom.Segment{
						A: geom.Point{X: x, Y: y + 1},
						B: geom.Point{X: x, Y: y},
					}
					segments = append(segments, edgesGrid[x][y].left)
				} else {
					// Reuse, and extend, the segment of the cell above
					edgesGrid[x][y-1].left.A.Y++
					edgesGrid[x][y].left = edgesGrid[x][y-1].left
				}
			}
			// Right segment
			if x+1 >= gridWidth || !isSolid(x+1, y) {
				if y-1 < 0 || edgesGrid[x][y-1].right == nil {
					edgesGrid[x][y].right = &geom.Segment{
						A: geom.Point{X: x + 1, Y: y},
						B: geom.Point{X: x + 1, Y: y + 1},
					}
					segments = append(segments, edgesGrid[x][y].right)
				} else {
					// Reuse, and extend, the segment of the cell above
					edgesGrid[x][y-1].right.B.Y++
					edgesGrid[x][y].right = edgesGrid[x][y-1].right
				}
			}
		}
	}
	return segments
}
