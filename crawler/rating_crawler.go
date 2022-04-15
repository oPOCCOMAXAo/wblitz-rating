package crawler

import (
	"fmt"
	"sync"
	"wblitz-rating/data"
	"wblitz-rating/utils"
)

type RatingCrawler struct {
	config   Config
	progress *utils.ProgressBar
	toSave   chan []*data.Rating
}

func NewRatingCrawler(config Config) (*RatingCrawler, error) {
	if config.API == nil {
		return nil, fmt.Errorf("%w: API", ErrNilParam)
	}

	if config.Storage == nil {
		return nil, fmt.Errorf("%w: Storage", ErrNilParam)
	}

	return &RatingCrawler{
		config: config,
		toSave: make(chan []*data.Rating, 1000),
	}, nil
}

func (c *RatingCrawler) Crawl() {
	info := c.config.API.RatingInfo()
	starts := make([]*data.Rating, 6)
	for i := int64(0); i < 5; i++ {
		starts[i] = c.config.API.RatingTop(i)[0]
	}
	starts[5] = &data.Rating{Number: info.Count * 2}
	c.progress = utils.NewProgressBar("Loading rating", float64(info.Count)/float64(c.config.BatchSize))
	c.progress.Add(0)

	var wgSave sync.WaitGroup
	wgSave.Add(1)
	go c.save(&wgSave)

	var wgLoad sync.WaitGroup
	wgLoad.Add(9)
	go c.getInterval(starts[0], (starts[0].Number+starts[1].Number)/2+1, &wgLoad)
	for i := 1; i < 5; i++ {
		go c.getInterval(starts[i], (starts[i].Number+starts[i+1].Number)/2+1, &wgLoad)
		go c.getInterval(starts[i], (starts[i].Number+starts[i-1].Number)/2-1, &wgLoad)
	}

	wgLoad.Wait()
	close(c.toSave)

	wgSave.Wait()
}

func (c *RatingCrawler) save(wg *sync.WaitGroup) {
	for batch := range c.toSave {
		c.config.Storage.SaveRating(batch)
		c.progress.Add(1)
	}
	wg.Done()
}

func (c *RatingCrawler) getInterval(start *data.Rating, endNumber int64, wg *sync.WaitGroup) {
	var fSelect func([]*data.Rating) *data.Rating
	var check func() bool
	if start.Number > endNumber {
		fSelect = func(arr []*data.Rating) *data.Rating { return arr[0] }
		check = func() bool { return start.Number > endNumber }
	} else {
		fSelect = func(arr []*data.Rating) *data.Rating { return arr[len(arr)-1] }
		check = func() bool { return start.Number < endNumber }
	}

	var last *data.Rating = &data.Rating{}
	for start.PlayerID != last.PlayerID && check() {
		last = start
		temp := c.config.API.RatingNeighbors(start.PlayerID, c.config.BatchSize)
		c.toSave <- temp
		start = fSelect(temp)
	}
	wg.Done()
}
