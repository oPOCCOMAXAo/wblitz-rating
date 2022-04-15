package crawler

import (
	"fmt"
	"sync"

	"wblitz-rating/data"
	"wblitz-rating/utils"
)

type StatCrawler struct {
	config   Config
	progress *utils.ProgressBar
	toSave   chan []*data.Player
}

func NewStatCrawler(config Config) (*StatCrawler, error) {
	if config.API == nil {
		return nil, fmt.Errorf("%w: API", ErrNilParam)
	}

	if config.Storage == nil {
		return nil, fmt.Errorf("%w: Storage", ErrNilParam)
	}

	return &StatCrawler{
		config: config,
		toSave: make(chan []*data.Player, 1000),
	}, nil
}

func (s *StatCrawler) Crawl() {
	const batchSize = 100

	ids, _ := s.config.Storage.GetRatingIDs()

	total := len(ids)
	batchCount := total / 100
	if total%100 > 0 {
		batchCount++
	}
	s.progress = utils.NewProgressBar("Loading players statistic", float64(batchCount))
	s.progress.Add(0)

	var wgSave sync.WaitGroup
	wgSave.Add(1)
	go s.save(&wgSave)

	var wgLoad sync.WaitGroup
	wgLoad.Add(batchCount)
	for i := 0; i < batchCount; i++ {
		start, end := i*batchSize, (i+1)*batchSize
		if end > total {
			end = total
		}
		go s.processBatch(ids[start:end], &wgLoad)
	}
	wgLoad.Wait()

	close(s.toSave)

	wgSave.Wait()
}

func (s *StatCrawler) save(wg *sync.WaitGroup) {
	for batch := range s.toSave {
		s.config.Storage.SavePlayers(batch)
		s.progress.Add(1)
	}
	wg.Done()
}

func (s *StatCrawler) processBatch(ids []int64, wg *sync.WaitGroup) {
	info := s.config.API.PlayerInfo(ids)
	s.toSave <- info
	wg.Done()
}
