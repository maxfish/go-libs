package rand

import "math/rand"

type RandomGenerator struct {
	generator *rand.Rand
}

func NewRandomGenerator(seed int64) *RandomGenerator {
	r := &RandomGenerator{
		generator: rand.New(rand.NewSource(seed)),
	}
	return r
}

func (g *RandomGenerator) NextUint32() uint32 {
	return g.generator.Uint32()
}

func (g *RandomGenerator) NextUint32LessThan(n int) (result uint32) {
	result = g.NextUint32() % uint32(n)
	return
}

func (g *RandomGenerator) NextUint32InRange(min, max int) (result uint32) {
	result = g.NextUint32LessThan((max - min) + 1) // +1 to include max
	result += uint32(min)
	return
}

func (g *RandomGenerator) NextFloat32() float32 {
	return g.generator.Float32()
}

func (g *RandomGenerator) Permutation(max int) []int {
	return g.generator.Perm(max)
}

func (g *RandomGenerator) Event(chance int) bool {
	return g.generator.Uint64()%100 < uint64(chance)
}

func (g *RandomGenerator) Maybe() bool {
	return g.generator.Float32() < 0.5
}
