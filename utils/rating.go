package utils

import (
	"bufio"
	"io"
	"os"
	"strings"
	"wblitz-rating/api"
)

const ratingHead = "Number;Rating;Nickname;Clan;ID;MMR"

func ratingToCSV(r api.Rating) string {
	return strings.Join([]string{
		api.U2S(r.Number),
		api.U2S(r.Score),
		r.Nickname,
		r.ClanTag,
		api.U2S(r.SpaId),
		api.F2S(r.MMR),
	}, ";")
}

func ratingFromCSV(line string) api.Rating {
	values := strings.Split(line, ";")
	return api.Rating{
		SpaId:    api.S2U(values[4]),
		MMR:      api.S2F(values[5]),
		Number:   api.S2U(values[0]),
		Score:    api.S2U(values[1]),
		Nickname: values[2],
		ClanTag:  values[3],
	}
}

func SaveRating(fname string, data api.RatingList) {
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

func LoadRating(fname string) api.RatingList {
	var res []api.Rating
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
