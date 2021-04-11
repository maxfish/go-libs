package rand

import (
	"math"
	"time"
)

type HashRng struct {
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

func (hr *HashRng) NextFloat32() (result float32) {
	u32 :=  HashNoise(hr.position, hr.seed)
	return float32(u32) / math.MaxUint32
}
