package utils

import (
	"fmt"
	"sync"
	"wblitz-rating/api"
)

type StatCrawler struct {
	client   *api.API
	progress *ProgressBar
	mu       sync.Mutex
}

func NewStatCrawler(client *api.API) *StatCrawler {
	return &StatCrawler{
		client: client,
	}
}

func (s *StatCrawler) GetAllStats(ids []uint64) []api.Player {
	const batchSize = 100
	total := len(ids)
	batchCount := total / 100
	if total%100 > 0 {
		batchCount++
	}
	res := make([]api.Player, total)
	s.progress = NewProgressBar("Loading players statistic", float64(batchCount))
	s.progress.Add(0)

	var wg sync.WaitGroup
	wg.Add(batchCount)
	for i := 0; i < batchCount; i++ {
		start, end := i*batchSize, (i+1)*batchSize
		if end > total {
			end = total
		}
		go s.processBatch(ids[start:end], res[start:end], &wg)
	}
	wg.Wait()

	fmt.Println()

	return res
}

func (s *StatCrawler) processBatch(ids []uint64, result []api.Player, wg *sync.WaitGroup) {
	info := s.client.PlayerInfo(ids)
	lookup := map[uint64]api.Player{}
	for _, player := range info {
		lookup[player.AccountId] = player
	}
	for i, id := range ids {
		result[i] = lookup[id]
	}
	s.mu.Lock()
	s.progress.Add(1)
	s.mu.Unlock()
	wg.Done()
}
