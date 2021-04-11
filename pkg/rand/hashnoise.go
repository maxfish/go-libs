package rand

// This is "Squirrel3" from "Guildhall's Squirrel Eiserloh"'s GDC 2017 talk
func HashNoise(position, seed uint32) (value uint32) {
	const bitNoise1 uint32 = 0xB5297A4D
	const bitNoise2 uint32 = 0x68E31DA4
	const bitNoise3 uint32 = 0x1B56C4E9

	value = position
	value *= bitNoise1
	value += seed
	value ^= value >> 8
	value += bitNoise2
	value ^= value << 8
	value = bitNoise3
	value ^= value >> 8
	return
}

func HashNoise2D(x, y, seed uint32) (value uint32) {
	const primeNumber uint32 = 0xBD4BCB5
	return HashNoise(x + (primeNumber * y), seed)
}
