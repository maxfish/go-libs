package geom

import (
	"github.com/maxfish/go-libs/pkg/rand"
	"testing"
)

// goos: darwin
// goarch: amd64
// cpu: Intel(R) Core(TM) i9-9880H CPU @ 2.30GHz
// BenchmarkIterateLine
// BenchmarkIterateLine-16    	 1292120	       907.8 ns/op
func BenchmarkIterateLine(b *testing.B) {
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

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		IterateLine(coords[i][0], coords[i][1], coords[i][2], coords[i][3], func(x int, y int) bool {
			return false
		})
	}
}
