package rand

import (
	"testing"
)

// goos: darwin
// goarch: amd64
// pkg: github.com/maxfish/go-libs/pkg/rand
// cpu: Intel(R) Core(TM) i9-9880H CPU @ 2.30GHz
// BenchmarkPerlinNoise1D
// BenchmarkPerlinNoise1D-16    	168736101	         6.979 ns/op
func BenchmarkPerlinNoise1D(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		PerlinNoise1D(float32(i), 89021)
	}
}
