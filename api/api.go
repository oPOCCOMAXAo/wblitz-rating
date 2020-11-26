package api

import (
	"encoding/json"
	"gitlab.com/opoccomaxao-go/request"
	"strings"
	"time"
)

type API struct {
	client        *request.Client
	timeout       time.Duration
	applicationId string
}

func New(timeout time.Duration, parallel int64, applicationId string) *API {
	return &API{
		client:        request.New(parallel),
		timeout:       timeout,
		applicationId: applicationId,
	}
}

func (a *API) PlayerInfo(ids []uint64) []Player {
	r := a.client.Get("https://api.wotblitz.ru/wotb/account/info/?application_id="+a.applicationId+
		"&account_id="+strings.Join(U2SArr(ids), ","),
		a.timeout, nil)
	if r.Status == 200 {
		var res PlayerResponse
		if err := json.Unmarshal(r.Response, &res); err != nil {
			println(err.Error())
			return nil
		}
		return res.All()
	}
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
	r := a.client.Get("https://ru.wotblitz.com/ru/api/rating-leaderboards/league/"+U2S(leagueId)+"/top/", a.timeout, nil)
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
	r := a.client.Get("https://ru.wotblitz.com/ru/api/rating-leaderboards/user/"+U2S(userId)+"/?neighbors="+U2S(neighborsCount), a.timeout, nil)
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
