package api

import (
	"context"
	"encoding/json"
	"strings"
	"time"

	"wblitz-rating/data"

	"github.com/opoccomaxao-go/generic-collection/slice"
	"github.com/pkg/errors"
	"github.com/valyala/fasthttp"
	"golang.org/x/sync/semaphore"
)

type API struct {
	client *fasthttp.Client
	sem    *semaphore.Weighted
	config Config
}

type Config struct {
	Parallel      int64
	Timeout       time.Duration
	ApplicationID string
}

func New(config Config) *API {
	if config.Parallel < 1 {
		config.Parallel = 1
	}

	if config.Timeout < 0 {
		config.Timeout = 0
	}

	return &API{
		client: &fasthttp.Client{
			NoDefaultUserAgentHeader: true,
			ReadTimeout:              config.Timeout,
			WriteTimeout:             config.Timeout,
		},
		sem:    semaphore.NewWeighted(config.Parallel),
		config: config,
	}
}

func (a *API) request(url string, resPtr interface{}) error {
	err := a.sem.Acquire(context.Background(), 1)
	if err != nil {
		return err
	}
	defer a.sem.Release(1)

	req := fasthttp.AcquireRequest()
	defer fasthttp.ReleaseRequest(req)

	res := fasthttp.AcquireResponse()
	defer fasthttp.ReleaseResponse(res)

	req.SetRequestURI(url)

	err = a.client.Do(req, res)
	if err != nil {
		return err
	}

	code := res.StatusCode()
	if code != fasthttp.StatusOK {
		return errors.Errorf("status: %d", code)
	}

	return json.Unmarshal(res.Body(), resPtr)
}

func (a *API) PlayerInfo(ids []int64) []*data.Player {
	var res PlayerResponse

	err := a.request(
		"https://api.wotblitz.ru/wotb/account/info/?application_id="+a.config.ApplicationID+"&account_id="+strings.Join(slice.Map(ids, I2S), ","),
		&res,
	)
	if err != nil {
		println(err.Error())
	}

	return res.All()
}

func (a *API) RatingInfo() RatingInfo {
	var res RatingInfo

	err := a.request("https://ru.wotblitz.com/ru/api/rating-leaderboards/season/", &res)
	if err != nil {
		println(err.Error())
	}

	return res
}

func (a *API) RatingTop(leagueID int64) []*data.Rating {
	var res RatingTopResponse

	err := a.request(
		"https://ru.wotblitz.com/ru/api/rating-leaderboards/league/"+I2S(leagueID)+"/top/",
		&res,
	)
	if err != nil {
		println(err.Error())
	}

	return slice.Map(res.Result, RatingMap)
}

func (a *API) RatingNeighbors(userID int64, neighborsCount int64) []*data.Rating {
	var res RatingNeighborsResponse

	err := a.request(
		"https://ru.wotblitz.com/ru/api/rating-leaderboards/user/"+I2S(userID)+"/?neighbors="+I2S(neighborsCount),
		&res,
	)
	if err != nil {
		println(err.Error())
	}

	return slice.Map(res.Neighbors, RatingMap)
}

func (a *API) GetRatingTop1() *data.Rating {
	res := a.RatingTop(0)
	if len(res) > 0 {
		return res[0]
	}
	return &data.Rating{}
}
