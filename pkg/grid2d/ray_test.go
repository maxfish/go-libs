package grid2d

import (
	"fmt"
	"github.com/maxfish/go-libs/pkg/rand"
	"github.com/maxfish/go-libs/pkg/testx"
	"testing"
)

type resultPoint struct {
	tile  [2]int
	pixel [2]float32
}

type testCase struct {
	cellSize, gridWidth, gridHeight int
	ray                             [4]float32
	result                          []resultPoint
}

type testCaseEnd struct {
	cellSize, gridWidth, gridHeight int
	ray                             [4]float32
	result                          resultPoint
}

func TestIterateRayEnd(t *testing.T) {
	var tests = []testCaseEnd{
		{32, 50, 50, [4]float32{0, 0, 2, 10}, resultPoint{[2]int{9, 49}, [2]float32{313.59985, 1567.9993}}},
		{32, 10, 10, [4]float32{0, 0, 1, 1}, resultPoint{[2]int{9, 9}, [2]float32{288, 288}}},
		{32, 100, 100, [4]float32{0, 0, 1, 5}, resultPoint{[2]int{19, 99}, [2]float32{633.6001, 3168.0005}}},
		{32, 55, 55, [4]float32{0, 0, -1, -1}, resultPoint{[2]int{0, 0}, [2]float32{0, 0}}},
		{32, 55, 55, [4]float32{70, 16, 0, 1}, resultPoint{[2]int{2, 54}, [2]float32{70, 1728}}},
		{32, 55, 55, [4]float32{70, 16, 0, 1000}, resultPoint{[2]int{2, 54}, [2]float32{70, 1727.9989}}},
		{32, 55, 55, [4]float32{70, 16, -5, -5}, resultPoint{[2]int{1, 0}, [2]float32{64, 10}}},
	}

	for index, test := range tests {
		var point resultPoint
		IterateRay(test.cellSize, test.gridWidth, test.gridHeight, test.ray[0], test.ray[1], test.ray[2], test.ray[3],
			func(tileX, tileY int, x, y float32) bool {
				point = resultPoint{tile: [2]int{tileX, tileY}, pixel: [2]float32{x, y}}
				return true
			},
		)
		testIndex := fmt.Sprintf("%d", index)
		testx.AssertEqual(t, "IterateRayEnd() #"+testIndex, test.result, point)
	}
}

func TestIterateRayPoints(t *testing.T) {
	var tests = []testCase{
		{32, 5, 5, [4]float32{10, 10, 2, 10}, []resultPoint{
			{[2]int{0, 0}, [2]float32{10, 10}},
			{[2]int{0, 1}, [2]float32{14.4, 32}},
			{[2]int{0, 2}, [2]float32{20.8, 64}},
			{[2]int{0, 3}, [2]float32{27.2, 96}},
			{[2]int{1, 3}, [2]float32{32, 120.00001}},
			{[2]int{1, 4}, [2]float32{33.600002, 128.00002}},
		}},
		// Same as the one above but with a different ray length
		{32, 5, 5, [4]float32{10, 10, 1, 5}, []resultPoint{
			{[2]int{0, 0}, [2]float32{10, 10}},
			{[2]int{0, 1}, [2]float32{14.4, 32}},
			{[2]int{0, 2}, [2]float32{20.8, 64}},
			{[2]int{0, 3}, [2]float32{27.2, 96}},
			{[2]int{1, 3}, [2]float32{32, 120.00001}},
			{[2]int{1, 4}, [2]float32{33.600002, 128.00002}},
		}},
		{32, 5, 5, [4]float32{150, 69, -2, -1}, []resultPoint{
			{[2]int{4, 2}, [2]float32{150, 69}},
			{[2]int{4, 1}, [2]float32{140, 64}},
			{[2]int{3, 1}, [2]float32{128, 58}},
			{[2]int{2, 1}, [2]float32{96, 42}},
			{[2]int{2, 0}, [2]float32{76, 32}},
			{[2]int{1, 0}, [2]float32{64, 26}},
			{[2]int{0, 0}, [2]float32{32, 10}},
		}},
	}

	for index, test := range tests {
		var points = make([]resultPoint, 0)
		IterateRay(test.cellSize, test.gridWidth, test.gridHeight, test.ray[0], test.ray[1], test.ray[2], test.ray[3],
			func(tileX, tileY int, x, y float32) bool {
				points = append(points, resultPoint{tile: [2]int{tileX, tileY}, pixel: [2]float32{x, y}})
				return true
			},
		)
		testIndex := fmt.Sprintf("%d", index)
		testx.AssertEqual(t, "IterateRayPoints() #"+testIndex, test.result, points)
	}
}

// === Benchmarks

func setupRaysData(n int, gridSize int) [][]float32 {
	// Prepare random, but predictable, coords
	rng := rand.NewHashRngWithSeed(0)
	coords := make([][]float32, n)
	for i := 0; i < n; i++ {
		coords[i] = make([]float32, 4)
		coords[i][0] = float32(rng.NextUint32() % uint32(gridSize))
		coords[i][1] = float32(rng.NextUint32() % uint32(gridSize))
		coords[i][2] = float32(rng.NextUint32()%uint32(10)) - 5
		coords[i][3] = float32(rng.NextUint32()%uint32(10)) - 5
	}
	return coords
}

// goos: darwin
// goarch: amd64
// cpu: Intel(R) Core(TM) i9-9880H CPU @ 2.30GHz
// BenchmarkIterateRay
// BenchmarkIterateRay-16    	 6951666	       165.9 ns/op
func BenchmarkIterateRay(b *testing.B) {
	cellSize := 32
	gridSize := 100
	coords := setupRaysData(b.N, gridSize)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		IterateRay(cellSize, gridSize, gridSize, coords[i][0], coords[i][1], coords[i][2], coords[i][3],
			func(tileX, tileY int, x, y float32) bool {
				return true
			},
		)
	}
}
