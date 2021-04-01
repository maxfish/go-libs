package image

import (
	"github.com/maxfish/go-libs/pkg/geom"
	"image"
	"image/draw"
)

func NewImagesFromAreas(img image.Image, areas []geom.Rect) []*image.RGBA {
	images := make([]*image.RGBA, 0, len(areas))

	for _, area := range areas {
		images = append(images, NewImageFromArea(img, area))
	}

	return images
}

func NewImageFromArea(img image.Image, area geom.Rect) *image.RGBA {
	sourceRectangle := area.ToRectangle()
	destRectangle := area.MoveTo(0, 0).ToRectangle()
	newImage := image.NewRGBA(destRectangle)
	draw.Draw(newImage, destRectangle, img, sourceRectangle.Min, draw.Src)
	return newImage
}
