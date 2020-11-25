package api

import (
	"encoding/json"
	"gitlab.com/opoccomaxao-go/request"
	"time"
)

const defaultTimeout = time.Millisecond * 100

type API struct {
	client *request.Client
}

func New() *API {
	return &API{
		client: request.New(1),
	}
}

func (a *API) get(url string) request.Response {
	return a.client.Get(url, time.Millisecond*200, nil)
}

func (a *API) PlayerInfo(ids []uint64) []Player {
	return nil
}

func (a *API) RatingTop(leagueId uint64) []Rating {
	r := a.client.Get("https://ru.wotblitz.com/ru/api/rating-leaderboards/league/"+u2s(leagueId)+"/top/", defaultTimeout, nil)
	if r.Status == 200 {
		var res RatingTopResponse
		if err := json.Unmarshal(r.Response, &res); err != nil {
			println(err)
			return nil
		}
		return res.Result
	}
	return nil
}

func (a *API) RatingNeighbors(userId uint64, neighborsCount uint64) []Rating {
	r := a.client.Get("https://ru.wotblitz.com/ru/api/rating-leaderboards/user/"+u2s(userId)+"/?neighbors="+u2s(neighborsCount), defaultTimeout, nil)
	if r.Status == 200 {
		var res RatingNeighborsResponse
		if err := json.Unmarshal(r.Response, &res); err != nil {
			println(err)
			return nil
		}
		return res.Neighbors
	}
	return nil
}

func (a *API) GetRatingTop1() Rating {
	res := a.RatingTop(0)
	if len(res) > 0 {
		return res[0]
	}
	return Rating{}
}
