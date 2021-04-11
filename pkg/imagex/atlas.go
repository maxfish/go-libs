package imagex

import (
	"github.com/maxfish/go-libs/pkg/geom"
	"image"
	"image/draw"
)

// BuildAtlas generates an atlas image and returns it. It also returns info on how the rects have been packed
func BuildAtlas(images []*image.RGBA, paddingX, paddingY int, allowRotation bool) (*image.RGBA, []geom.RectNode) {
	result, _ := bruteForceAtlasPacking(images, paddingX, paddingY, allowRotation)

	atlasImage := image.NewRGBA(image.Rect(0, 0, result.Width, result.Height))
	for _, node := range result.PlacedRects {
		img := images[node.Index]
		if node.Rotated {
			img = Rotate90CW(images[node.Index])
		}
		draw.Draw(atlasImage,
			image.Rectangle{Min: image.Point{X: node.X, Y: node.Y}, Max: image.Point{X: node.X + node.W-1, Y: node.Y + node.H-1}},
			img,
			image.Point{},
			draw.Src,
		)
	}

	return atlasImage, result.PlacedRects
}

// bruteForceAtlasPacking tries multiple combinations of image sizes and heuristics and then uses the best one
func bruteForceAtlasPacking(images []*image.RGBA, paddingX, paddingY int, allowRotation bool) (*geom.MaxRectsBinResult, float32) {
	var bestOccupancy float32 = 0.0
	var bestResult *geom.MaxRectsBinResult
	//var bestMethod geom.FreeRectChoiceHeuristic
	//var bestWidth int

	// Tries multiple texture widths, and it leaves the height to the algorithm
	widths := []int{128, 256, 512, 1024}
	for _, maxWidth := range widths {
		// Tries the different methods
		methods := []geom.FreeRectChoiceHeuristic{geom.RectBestShortSideFit, geom.RectBestAreaFit, geom.RectBottomLeftRule}
		for method := range methods {
			done := false
			sizeFactor := float32(1)
			var result *geom.MaxRectsBinResult
			for !done {
				rects := make([]geom.RectNode, 0)
				minSurface := 0
				for i, img := range images {
					rect := geom.NewRectNode(i, img.Bounds().Dx(), img.Bounds().Dy())
					rects = append(rects, rect)
					minSurface += rect.W * rect.H
				}
				minSurface = int(float32(minSurface) * sizeFactor)
				width := maxWidth
				height := minSurface / width

				maxRects := geom.NewMaxRectsBinPacker(width, height, paddingX, paddingY, allowRotation)
				result = maxRects.Pack(rects, geom.FreeRectChoiceHeuristic(method))
				if len(result.NotPlacedRects) > 0 {
					// Some rects didn't fit in the provided area, increase it a bit
					sizeFactor += 0.01
				} else {
					//fmt.Printf("- Method:%d Occupancy:%f\n", method, maxRects.Occupancy())
					if maxRects.Occupancy() > bestOccupancy {
						bestResult = result
						bestOccupancy = maxRects.Occupancy()
						//bestMethod = geom.FreeRectChoiceHeuristic(method)
						//bestWidth = maxWidth
					}
					done = true
					continue
				}
			}
		}
	}
	//fmt.Printf("--> DONE Method:%d width:%d Occupancy:%f\n", bestMethod, bestWidth, bestOccupancy)
	return bestResult, bestOccupancy
}
