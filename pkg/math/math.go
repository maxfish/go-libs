package math

import (
	"math"
)

const (
	Float32Size      = 4 // bytes
	RadiansToDegrees = 180.0 / float32(math.Pi)
	DegreesToRadians = float32(math.Pi) / 180.0
)

func Abs(value float32) float32 {
	if value == 0 {
		return 0
	}
	if value < 0 {
		return -value
	}
	return value
}

func Sign(value float32) float32 {
	if value == 0 {
		return 0
	}
	if value < 0 {
		return -1
	}
	return 1
}

func MinI(a, b int) int {
	if a <= b {
		return a
	} else {
		return b
	}
}

func MaxI(a, b int) int {
	if a >= b {
		return a
	} else {
		return b
	}
}

func Min(x, y float32) float32 {
	if x < y {
		return x
	}
	return y
}

func Max(x, y float32) float32 {
	if x > y {
		return x
	}
	return y
}

func ClampI(value, a, b int) int {
	if value < a {
		return a
	} else if value > b {
		return b
	}

	return value
}

func Clamp(value, a, b float32) float32 {
	if value < a {
		return a
	} else if value > b {
		return b
	}

	return value
}

func ClampApproach(value, target, amount float32) float32 {
	if value > target {
		return Max(value-amount, target)
	}
	return Min(value+amount, target)
}

func Lerp(a, b, factor float32) float32 {
	return (1-factor)*a + factor*b
}
