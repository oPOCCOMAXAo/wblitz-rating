package api

import (
	"encoding/json"
	"gitlab.com/opoccomaxao-go/request"
	"time"
)

type API struct {
	client  *request.Client
	timeout time.Duration
}

func New(timeout time.Duration, parallel int64) *API {
	return &API{
		client:  request.New(parallel),
		timeout: timeout,
	}
}

func (a *API) PlayerInfo(ids []uint64) []Player {
	return nil
}

func (a *API) RatingInfo() RatingInfo {
	r := a.client.Get("https://ru.wotblitz.com/ru/api/rating-leaderboards/season/", a.timeout, nil)
	if r.Status == 200 {
		var res RatingInfo
		if err := json.Unmarshal(r.Response, &res); err != nil {
			println(err.Error())
			return RatingInfo{}
		}
		return res
	}
	return RatingInfo{}
}

func (a *API) RatingTop(leagueId uint64) []Rating {
	r := a.client.Get("https://ru.wotblitz.com/ru/api/rating-leaderboards/league/"+u2s(leagueId)+"/top/", a.timeout, nil)
	if r.Status == 200 {
		var res RatingTopResponse
		if err := json.Unmarshal(r.Response, &res); err != nil {
			println(err.Error())
			return nil
		}
		return res.Result
	}
	return nil
}

func (a *API) RatingNeighbors(userId uint64, neighborsCount uint64) []Rating {
	r := a.client.Get("https://ru.wotblitz.com/ru/api/rating-leaderboards/user/"+u2s(userId)+"/?neighbors="+u2s(neighborsCount), a.timeout, nil)
	if r.Status == 200 {
		var res RatingNeighborsResponse
		if err := json.Unmarshal(r.Response, &res); err != nil {
			println(err.Error())
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
