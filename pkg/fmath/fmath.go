package fmath

import (
	"math"
)

const (
	Float32Size      = 4 // bytes
	Float32MaxValue  = math.MaxFloat32
	Float32MinNormal = float32(1.1754943508222875e-38)
	Float32Epsilon   = float32(1e-7)
	RadiansToDegrees = float32(180.0 / math.Pi)
	DegreesToRadians = float32(math.Pi / 180.0)
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

func Floor(value float32) int {
	var vI = int(value)
	var vF = float32(vI)
	if value < vF {
		return vI - 1
	} else {
		return vI
	}
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

func Round(value float32) float32 {
	return float32(math.Round(float64(value)))
}

func Sqrt(value float32) float32 {
	return float32(math.Sqrt(float64(value)))
}
