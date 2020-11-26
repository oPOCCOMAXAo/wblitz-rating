package utils

import (
	"image"
	"image/color"
	"image/png"
	"os"
)

type Count2dEntry struct {
	Count  []uint64
	Rating uint64
}

func DrawCount2dEntry(fname string, data []Count2dEntry) {
	width := len(data)
	height := len(data[0].Count)
	img := image.NewRGBA(image.Rect(0, 0, width, height))

	for x := 0; x < width; x++ {
		entry := data[x]
		c := getLeagueColor(entry.Rating)
		maxCount := uint64(1)
		for _, count := range entry.Count {
			if maxCount < count {
				maxCount = count
			}
		}
		for y := 0; y < height; y++ {
			c2 := color.RGBA{
				R: uint8(uint64(c.R) * entry.Count[height-y-1] / maxCount),
				G: uint8(uint64(c.G) * entry.Count[height-y-1] / maxCount),
				B: uint8(uint64(c.B) * entry.Count[height-y-1] / maxCount),
				A: 255,
			}
			img.SetRGBA(x, y, c2)
		}
	}

	f, err := os.Create(fname + ".png")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	_ = png.Encode(f, img)
}
