package api

import (
	"wblitz-rating/data"
)

type Rating struct {
	SpaId    int64   `json:"spa_id"`
	MMR      float64 `json:"mmr"`
	Number   int64   `json:"number"`
	Score    int64   `json:"score"`
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
	Count   int64 `json:"count"`
	Leagues []struct {
		Index      int64   `json:"index"`
		Percentile float64 `json:"percentile"`
	} `json:"leagues"`
}

type PlayerResponse struct {
	Status string `json:"status"`
	Data   map[string]struct {
		AccountId  int64  `json:"account_id"`
		Nickname   string `json:"nickname"`
		Statistics struct {
			All struct {
				Battles     int64 `json:"battles"`
				DamageDealt int64 `json:"damage_dealt"`
				Wins        int64 `json:"wins"`
			} `json:"all"`
		} `json:"statistics"`
	} `json:"data"`
}

type Player struct {
	AccountId   int64  `json:"account_id"`
	Nickname    string `json:"nickname"`
	Battles     int64  `json:"battles"`
	DamageDealt int64  `json:"damage_dealt"`
	Wins        int64  `json:"wins"`
}

func (p *PlayerResponse) All() []*data.Player {
	res := make([]*data.Player, 0, len(p.Data))
	for _, player := range p.Data {
		res = append(res, &data.Player{
			ID:       player.AccountId,
			Nickname: player.Nickname,
			Battles:  player.Statistics.All.Battles,
			Damage:   player.Statistics.All.DamageDealt,
			Wins:     player.Statistics.All.Wins,
		})
	}
	return res
}
