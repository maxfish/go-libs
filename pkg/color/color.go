package color

import (
	"github.com/go-gl/mathgl/mgl32"
	"image/color"
)

// Components are normalized floats ranging from 0 to 1
type Color = mgl32.Vec4

// NewColorFromImageRGBA creates a new color from color.RGBA
func NewColorFromImageRGBA(rgba color.RGBA) Color {
	return NewColorFromBytes(rgba.R, rgba.G, rgba.B, rgba.A)
}

// NewColorFromFloats creates a new color from float components (0->1)
func NewColorFromFloats(r, g, b, a float32) Color {
	return Color{r, g, b, a}
}

// NewColorFromBytes creates a new color from 8bit components (0->255)
func NewColorFromBytes(r, g, b, a uint8) Color {
	return Color{
		float32(r) / 255.0,
		float32(g) / 255.0,
		float32(b) / 255.0,
		float32(a) / 255.0,
	}
}

// NewColorFromHex creates a new color from an hex number
func NewColorFromHex(hex uint32) Color {
	c := Color{}
	c[3] = float32(hex&0xFF) / 255
	hex >>= 8
	c[2] = float32(hex&0xFF) / 255
	hex >>= 8
	c[1] = float32(hex&0xFF) / 255
	hex >>= 8
	c[0] = float32(hex&0xFF) / 255
	return c
}

// NewColorFromGrayInt creates a gray shade from float components (0->1)
func NewColorFromGrayInt(g, a uint8) Color {
	return Color{
		float32(g) / 255.0,
		float32(g) / 255.0,
		float32(g) / 255.0,
		float32(a) / 255.0,
	}
}

func ColorScaled(color Color, scale float32) Color {
	return color.Mul(scale)
}
