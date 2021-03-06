package rand

import "math/rand"

type StdLibGenerator struct {
	RandomNumberGenerator
	generator *rand.Rand
}

func NewStdLibGenerator(seed int64) *StdLibGenerator {
	r := &StdLibGenerator{
		generator: rand.New(rand.NewSource(seed)),
	}
	return r
}

func (g *StdLibGenerator) NextUint32() uint32 {
	return g.generator.Uint32()
}

func (g *StdLibGenerator) NextUint32LessThan(n int) (result uint32) {
	result = g.NextUint32() % uint32(n)
	return
}

func (g *StdLibGenerator) NextUint32InRange(min, max int) (result uint32) {
	result = g.NextUint32LessThan((max - min) + 1) // +1 to include max
	result += uint32(min)
	return
}

func (g *StdLibGenerator) NextFloat32() float32 {
	return g.generator.Float32()
}

func (g *StdLibGenerator) Permutation(max int) []int {
	return g.generator.Perm(max)
}

func (g *StdLibGenerator) Event(chance int) bool {
	return g.generator.Uint64()%100 < uint64(chance)
}

func (g *StdLibGenerator) Maybe() bool {
	return g.generator.Float32() < 0.5
}
