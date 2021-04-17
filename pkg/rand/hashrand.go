package rand

import (
	"math"
	"time"
)

type HashRng struct {
	RandomNumberGenerator
	position uint32
	seed     uint32
}

func NewHashRng() *HashRng {
	return &HashRng{
		seed: uint32(time.Now().Nanosecond()),
	}
}

func NewHashRngWithSeed(seed uint32) *HashRng {
	return &HashRng{
		seed: seed,
	}
}

func (hr *HashRng) NextUint32() (result uint32) {
	result = HashNoise(hr.position, hr.seed)
	hr.position++
	return
}

func (hr *HashRng) NextUint32LessThan(n int) (result uint32) {
	result = HashNoise(hr.position, hr.seed) % uint32(n)
	hr.position++
	return
}

func (hr *HashRng) NextUint32InRange(min, max int) (result uint32) {
	result = hr.NextUint32LessThan((max - min) + 1) // +1 to include max
	result += uint32(min)
	hr.position++
	return
}

func (hr *HashRng) NextFloat32() (result float32) {
	result = float32(HashNoise(hr.position, hr.seed)) / math.MaxUint32
	hr.position++
	return
}

func (hr *HashRng) Event(chance int) bool {
	return hr.NextUint32()%100 < uint32(chance)
}

func (hr *HashRng) Maybe() bool {
	return hr.NextUint32() < math.MaxUint32/2
}
