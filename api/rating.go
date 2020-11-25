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
