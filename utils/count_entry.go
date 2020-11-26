package utils

import (
	"bufio"
	"image"
	"image/png"
	"os"
	"strings"
	"wblitz-rating/api"
)

const countEntryHead = "Rating;Count"

type CountEntry struct {
	Count  uint64
	Rating uint64
}

func countEntryToCSV(entry CountEntry) string {
	return strings.Join([]string{
		api.U2S(entry.Rating),
		api.U2S(entry.Count),
	}, ";")
}

func SaveCountEntry(fname string, data []CountEntry) {
	f, err := os.Create(fname)
	panicOnNonNil(err)
	defer f.Close()
	w := bufio.NewWriter(f)
	_, err = w.WriteString(countEntryHead)
	panicOnNonNil(err)
	for _, entry := range data {
		_, err = w.WriteString("\n" + countEntryToCSV(entry))
		panicOnNonNil(err)
	}
	_ = w.Flush()
}

func DrawCountEntry(fname string, data []CountEntry) {
	width := len(data)
	height := width / 2
	img := image.NewRGBA(image.Rect(0, 0, width, height))
	maxCount := uint64(0)
	for _, entry := range data {
		if maxCount < entry.Count {
			maxCount = entry.Count
		}
	}

	for x := 0; x < width; x++ {
		color := getLeagueColor(data[x].Rating)
		for y := height - height*int(data[x].Count)/int(maxCount); y < height; y++ {
			img.Set(x, y, color)
		}
	}

	f, err := os.Create(fname + ".png")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	_ = png.Encode(f, img)
}
