package utils

import (
	"fmt"
	"sync"
	"wblitz-rating/api"
)

type RatingCrawler struct {
	client     *api.API
	batchSize  uint64
	batchCache [][]api.Rating
	progress   *ProgressBar
	mu         sync.Mutex
}

func NewRatingCrawler(client *api.API, batchSize uint64) *RatingCrawler {
	return &RatingCrawler{
		client:    client,
		batchSize: batchSize,
	}
}

func (c *RatingCrawler) GetAllRating() api.RatingList {
	c.batchCache = nil

	info := c.client.RatingInfo()
	starts := make([]api.Rating, 6)
	for i := uint64(0); i < 5; i++ {
		starts[i] = c.client.RatingTop(i)[0]
	}
	starts[5] = api.Rating{Number: info.Count * 2}
	c.progress = NewProgressBar("Loading rating", float64(info.Count)/float64(c.batchSize))
	c.progress.Add(0)

	var wg sync.WaitGroup
	wg.Add(9)
	go c.getInterval(starts[0], (starts[0].Number+starts[1].Number)/2+1, &wg)
	for i := 1; i < 5; i++ {
		go c.getInterval(starts[i], (starts[i].Number+starts[i+1].Number)/2+1, &wg)
		go c.getInterval(starts[i], (starts[i].Number+starts[i-1].Number)/2-1, &wg)
	}
	wg.Wait()

	res := make([]api.Rating, info.Count)
	for _, batch := range c.batchCache {
		for _, rating := range batch {
			res[rating.Number-1] = rating
		}
	}

	fmt.Println()
	return res
}

func (c *RatingCrawler) getInterval(start api.Rating, endNumber uint64, wg *sync.WaitGroup) {
	var fSelect func([]api.Rating) api.Rating
	var check func() bool
	if start.Number > endNumber {
		fSelect = func(arr []api.Rating) api.Rating { return arr[0] }
		check = func() bool { return start.Number > endNumber }
	} else {
		fSelect = func(arr []api.Rating) api.Rating { return arr[len(arr)-1] }
		check = func() bool { return start.Number < endNumber }
	}

	var last api.Rating
	for start.SpaId != last.SpaId && check() {
		last = start
		temp := c.client.RatingNeighbors(start.SpaId, c.batchSize)

		c.mu.Lock()
		c.batchCache = append(c.batchCache, temp)
		c.progress.Add(1)
		c.mu.Unlock()

		start = fSelect(temp)
	}
	wg.Done()
}
