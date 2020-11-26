package api

type PlayerResponse struct {
	Status string `json:"status"`
	Data   map[string]struct {
		AccountId  uint64 `json:"account_id"`
		Nickname   string `json:"nickname"`
		Statistics struct {
			All struct {
				Battles     uint64 `json:"battles"`
				DamageDealt uint64 `json:"damage_dealt"`
				Wins        uint64 `json:"wins"`
			} `json:"all"`
		} `json:"statistics"`
	} `json:"data"`
}

type Player struct {
	AccountId   uint64 `json:"account_id"`
	Nickname    string `json:"nickname"`
	Battles     uint64 `json:"battles"`
	DamageDealt uint64 `json:"damage_dealt"`
	Wins        uint64 `json:"wins"`
}

func (p *PlayerResponse) All() []Player {
	res := make([]Player, 0, len(p.Data))
	for _, player := range p.Data {
		res = append(res, Player{
			AccountId:   player.AccountId,
			Nickname:    player.Nickname,
			Battles:     player.Statistics.All.Battles,
			DamageDealt: player.Statistics.All.DamageDealt,
			Wins:        player.Statistics.All.Wins,
		})
	}
	return res
}
