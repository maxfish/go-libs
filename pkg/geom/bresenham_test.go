package geom

import (
	"fmt"
	"github.com/maxfish/go-libs/pkg/rand"
	"github.com/maxfish/go-libs/pkg/testx"
	"testing"
)

func TestIterateLineOrderedAA(t *testing.T) {
	var tests = []struct {
		coords [4]int
		points [][2]int
	}{
		{[4]int{0, 0, 4, 0}, [][2]int{{0, 0}, {1, 0}, {2, 0}, {3, 0}, {4, 0}}},
		{[4]int{0, 3, -5, 3}, [][2]int{{0, 3}, {-1, 3}, {-2, 3}, {-3, 3}, {-4, 3}, {-5, 3}}},
		{[4]int{0, 0, 0, 4}, [][2]int{{0, 0}, {0, 1}, {0, 2}, {0, 3}, {0, 4}}},
		{[4]int{-1, 10, -1, 5}, [][2]int{{-1, 10}, {-1, 9}, {-1, 8}, {-1, 7}, {-1, 6}, {-1, 5}}},
	}

	for i, test := range tests {
		result := make([][2]int, 0)
		IterateLineOrdered(test.coords[0], test.coords[1], test.coords[2], test.coords[3],
			func(x int, y int) bool {
				result = append(result, [2]int{x, y})
				return false
			},
		)
		errorText := fmt.Sprintf("TestIterateLineOrderedAA #%d", i)
		testx.AssertEqual(t, errorText, test.points, result)
	}
}

func TestIterateLineOrdered(t *testing.T) {
	var tests = []struct {
		coords [4]int
		points [][2]int
	}{
		{[4]int{0, 0, 4, 7}, [][2]int{{0, 0}, {1, 1}, {1, 2}, {2, 3}, {2, 4}, {3, 5}, {3, 6}, {4, 7}}},
		{[4]int{-5, 3, 4, 6}, [][2]int{{-5, 3}, {-4, 3}, {-3, 4}, {-2, 4}, {-1, 4}, {0, 5}, {1, 5}, {2, 5}, {3, 6}, {4, 6}}},
		{[4]int{10, 10, 4, 4}, [][2]int{{10, 10}, {9, 9}, {8, 8}, {7, 7}, {6, 6}, {5, 5}, {4, 4}}},
		{[4]int{6, 6, -1, 0}, [][2]int{{6, 6}, {5, 5}, {4, 4}, {3, 3}, {2, 3}, {1, 2}, {0, 1}, {-1, 0}}},
	}

	for i, test := range tests {
		result := make([][2]int, 0)
		IterateLineOrdered(test.coords[0], test.coords[1], test.coords[2], test.coords[3],
			func(x int, y int) bool {
				result = append(result, [2]int{x, y})
				return false
			},
		)
		errorText := fmt.Sprintf("TestIterateLine #%d", i)
		testx.AssertEqual(t, errorText, test.points, result)
	}
}

func TestIterateLineAA(t *testing.T) {
	var tests = []struct {
		coords [4]int
		points [][2]int
	}{
		{[4]int{0, 0, 4, 0}, [][2]int{{0, 0}, {1, 0}, {2, 0}, {3, 0}, {4, 0}}},
		{[4]int{0, 3, -5, 3}, [][2]int{{-5, 3}, {-4, 3}, {-3, 3}, {-2, 3}, {-1, 3}, {0, 3}}},
		{[4]int{0, 0, 0, 4}, [][2]int{{0, 0}, {0, 1}, {0, 2}, {0, 3}, {0, 4}}},
		{[4]int{-1, 10, -1, 5}, [][2]int{{-1, 5}, {-1, 6}, {-1, 7}, {-1, 8}, {-1, 9}, {-1, 10}}},
	}

	for i, test := range tests {
		result := make([][2]int, 0)
		IterateLine(test.coords[0], test.coords[1], test.coords[2], test.coords[3],
			func(x int, y int) bool {
				result = append(result, [2]int{x, y})
				return false
			},
		)
		errorText := fmt.Sprintf("TestIterateLineAA #%d", i)
		testx.AssertEqual(t, errorText, test.points, result)
	}
}

func TestIterateLine(t *testing.T) {
	var tests = []struct {
		coords [4]int
		points [][2]int
	}{
		{[4]int{0, 0, 4, 7}, [][2]int{{0, 0}, {1, 1}, {1, 2}, {2, 3}, {2, 4}, {3, 5}, {3, 6}, {4, 7}}},
		{[4]int{-5, 3, 4, 6}, [][2]int{{-5, 3}, {-4, 3}, {-3, 4}, {-2, 4}, {-1, 4}, {0, 5}, {1, 5}, {2, 5}, {3, 6}, {4, 6}}},
		{[4]int{10, 10, 4, 4}, [][2]int{{4, 4}, {5, 5}, {6, 6}, {7, 7}, {8, 8}, {9, 9}, {10, 10}}},
		{[4]int{6, 6, -1, 0}, [][2]int{{-1, 0}, {0, 1}, {1, 2}, {2, 3}, {3, 3}, {4, 4}, {5, 5}, {6, 6}}},
	}

	for i, test := range tests {
		result := make([][2]int, 0)
		IterateLine(test.coords[0], test.coords[1], test.coords[2], test.coords[3],
			func(x int, y int) bool {
				result = append(result, [2]int{x, y})
				return false
			},
		)
		errorText := fmt.Sprintf("TestIterateLine #%d", i)
		testx.AssertEqual(t, errorText, test.points, result)
	}
}

// === Benchmarks

func setupPoints(b *testing.B) [][]int {
	// Prepare a set, random but predictable, of coords
	rng := rand.NewHashRngWithSeed(0)
	coords := make([][]int, b.N)
	for i := 0; i < b.N; i++ {
		coords[i] = make([]int, 4)
		coords[i][0] = int(rng.NextUint32() % 1000)
		coords[i][1] = int(rng.NextUint32() % 1000)
		coords[i][2] = int(rng.NextUint32() % 1000)
		coords[i][3] = int(rng.NextUint32() % 1000)
	}
	return coords
}

// goos: darwin
// goarch: amd64
// pkg: github.com/maxfish/go-libs/pkg/geom
// cpu: Intel(R) Core(TM) i9-9880H CPU @ 2.30GHz
// BenchmarkIterateLine
// BenchmarkIterateLine-16    	 1245522	       921.0 ns/op
func BenchmarkIterateLine(b *testing.B) {
	coords := setupPoints(b)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		IterateLine(coords[i][0], coords[i][1], coords[i][2], coords[i][3], func(x int, y int) bool {
			return false
		})
	}
}

// goos: darwin
// goarch: amd64
// pkg: github.com/maxfish/go-libs/pkg/geom
// cpu: Intel(R) Core(TM) i9-9880H CPU @ 2.30GHz
// BenchmarkIterateLineOrdered
// BenchmarkIterateLineOrdered-16    	  799668	      1344 ns/op
func BenchmarkIterateLineOrdered(b *testing.B) {
	coords := setupPoints(b)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		IterateLineOrdered(coords[i][0], coords[i][1], coords[i][2], coords[i][3], func(x int, y int) bool {
			return false
		})
	}
}
