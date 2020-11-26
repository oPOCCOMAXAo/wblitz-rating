package api

type Rating struct {
	SpaId    uint64  `json:"spa_id"`
	MMR      float64 `json:"mmr"`
	Number   uint64  `json:"number"`
	Score    uint64  `json:"score"`
	Nickname string  `json:"nickname"`
	ClanTag  string  `json:"clan_tag"`
}

type RatingTopResponse struct {
	Result []Rating `json:"result"`
}

type RatingNeighborsResponse struct {
	Neighbors []Rating `json:"neighbors"`
}

type RatingInfo struct {
	Count   uint64 `json:"count"`
	Leagues []struct {
		Index      uint64  `json:"index"`
		Percentile float64 `json:"percentile"`
	} `json:"leagues"`
}

type RatingList []Rating

func (r RatingList) GetIds() []uint64 {
	l := len(r)
	res := make([]uint64, l)
	for i := 0; i < l; i++ {
		res[i] = r[i].SpaId
	}
	return res
}
