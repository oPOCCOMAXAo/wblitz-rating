package utils

import (
	"bufio"
	"fmt"
	"os"
	"wblitz-rating/api"
)

const ratingHead = "Number;Rating;Nickname;Clan;ID;MMR\n"

func ratingToCSV(r api.Rating) string {
	return fmt.Sprintf("%d;%d;%s;%s;%d;%f\n", r.Number, r.Score, r.Nickname, r.ClanTag, r.SpaId, r.MMR)
}

func SaveTable(fname string, data []api.Rating) {
	f, err := os.Create(fname)
	panicOnNonNil(err)
	defer f.Close()
	w := bufio.NewWriter(f)
	_, err = w.WriteString(ratingHead)
	panicOnNonNil(err)
	for _, rating := range data {
		_, err = w.WriteString(ratingToCSV(rating))
		panicOnNonNil(err)
	}
	_ = w.Flush()
}
