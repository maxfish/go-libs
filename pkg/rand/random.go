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

func (g *RandomGenerator) NextInt() int {
	return g.generator.Int()
}

func (g *RandomGenerator) NextFloat() float32 {
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
