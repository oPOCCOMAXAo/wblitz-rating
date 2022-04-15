package api

import (
	"strconv"
	"wblitz-rating/data"
)

func I2S(u int64) string {
	return strconv.FormatInt(u, 10)
}

func F2S(f float64) string {
	return strconv.FormatFloat(f, 'f', -1, 64)
}

func S2I(s string) int64 {
	res, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		return 0
	}
	return res
}

func S2F(s string) float64 {
	res, err := strconv.ParseFloat(s, 64)
	if err != nil {
		return 0
	}
	return res
}

func RatingMap(r Rating) *data.Rating {
	return &data.Rating{
		Number:   r.Number,
		PlayerID: r.SpaId,
		Score:    r.Score,
		MMR:      r.MMR,
	}
}
