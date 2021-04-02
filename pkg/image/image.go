package image

import (
	"bytes"
	"encoding/base64"
	"github.com/maxfish/go-libs/pkg/geom"
	"image"
	"image/draw"
	_ "image/jpeg"
	"image/png"
)

func NewRGBAImagesFromAreas(img image.Image, areas []geom.Rect, skipEmptyAreas bool) []*image.RGBA {
	images := make([]*image.RGBA, 0, len(areas))

	for _, area := range areas {
		if IsImageAreaEmpty(img, area) {
			continue
		}
		images = append(images, NewRGBAImageFromArea(img, area))
	}

	return images
}

func NewRGBAImageFromArea(img image.Image, area geom.Rect) *image.RGBA {
	sourceRectangle := area.ToRectangle()
	destRectangle := area.MoveTo(0, 0).ToRectangle()
	newImage := image.NewRGBA(destRectangle)
	draw.Draw(newImage, destRectangle, img, sourceRectangle.Min, draw.Src)
	return newImage
}

// Check if an area is full of transparent pixels
func IsImageAreaEmpty(img image.Image, area geom.Rect) bool {
	for y := area.Top(); y < area.Bottom(); y++ {
		for x := area.Left(); x < area.Right(); x++ {
			_, _, _, a := img.At(x, y).RGBA()
			if a != 0 {
				return false
			}
		}
	}
	return true
}

func IsImageEmpty(img image.Image) bool {
	return IsImageAreaEmpty(img, geom.RectFromRectangle(img.Bounds()))
}

func ToRGBA(img image.Image) *image.RGBA {
	switch img.(type) {
	case *image.RGBA:
		return img.(*image.RGBA)
	default:
		rgba := image.NewRGBA(img.Bounds())
		draw.Draw(rgba, rgba.Bounds(), img, image.Point{0, 0}, draw.Src)
		return rgba
	}
}

func ToBytes(img image.Image) ([]byte, error) {
	var buffer bytes.Buffer
	err := png.Encode(&buffer, img)
	if err != nil {
		return nil, err
	}
	return buffer.Bytes(), nil
}

func ToBase64String(img image.Image) (string, error) {
	bytes, err := ToBytes(img)
	return base64.StdEncoding.EncodeToString(bytes), err
}

func FromBase64String(base64String string) (image.Image, error) {
	imgData, err := base64.StdEncoding.DecodeString(base64String)
	if err != nil {
		return nil, err
	}
	img, _, err2 := image.Decode(bytes.NewReader(imgData))

	return img, err2
}
