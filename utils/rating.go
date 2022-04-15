package utils

import (
	"bufio"
	"io"
	"os"
	"strings"
	"wblitz-rating/api"
	"wblitz-rating/data"
)

const ratingHead = "Number;Rating;ID;MMR"

func ratingToCSV(r data.Rating) string {
	return strings.Join([]string{
		api.I2S(r.Number),
		api.I2S(r.Score),
		api.I2S(r.PlayerID),
		api.F2S(r.MMR),
	}, ";")
}

func ratingFromCSV(line string) data.Rating {
	values := strings.Split(line, ";")
	return data.Rating{
		Number: api.S2I(values[0]),
		PlayerID:  api.S2I(values[2]),
		Score:  api.S2I(values[1]),
		MMR:    api.S2F(values[3]),
	}
}

func SaveRating(fname string, data []data.Rating) {
	f, err := os.Create(fname)
	panicOnNonNil(err)
	defer f.Close()
	w := bufio.NewWriter(f)
	_, err = w.WriteString(ratingHead)
	panicOnNonNil(err)
	for _, rating := range data {
		_, err = w.WriteString("\n" + ratingToCSV(rating))
		panicOnNonNil(err)
	}
	_ = w.Flush()
}

func LoadRating(fname string) []data.Rating {
	var res []data.Rating
	f, err := os.Open(fname)
	panicOnNonNil(err)
	defer f.Close()
	w := bufio.NewReader(f)
	t, _, err := w.ReadLine()
	panicOnNonNil(err)
	if string(t) != ratingHead {
		panic("unknown head")
	}
	for {
		line, _, err := w.ReadLine()
		if err == io.EOF {
			break
		}
		panicOnNonNil(err)
		res = append(res, ratingFromCSV(string(line)))
	}
	return res
}
