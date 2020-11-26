package utils

import (
	"fmt"
	"wblitz-rating/api"
)

type Entry struct {
	Id        uint64
	Rating    uint64
	Battles   uint64
	WinRate   float64
	AvgDamage float64
}

type Analytics struct {
	entries []Entry
}

func NewAnalytics(rating []api.Rating, stats []api.Player) *Analytics {
	total := len(rating)
	if t := len(stats); t != total {
		panic("data broken, reloading needed")
	}
	entries := make([]Entry, total)
	for i := 0; i < total; i++ {
		battles := float64(stats[i].Battles) + 1e-50
		entries[i] = Entry{
			Id:        rating[i].SpaId,
			Rating:    rating[i].Score,
			Battles:   stats[i].Battles,
			WinRate:   float64(stats[i].Wins) / battles,
			AvgDamage: float64(stats[i].DamageDealt) / battles,
		}
	}
	return &Analytics{
		entries: entries,
	}
}

func (a *Analytics) Log() {
	fmt.Println(len(a.entries))
}

func (a *Analytics) GetTopDamager() Entry {
	max := a.entries[0]
	for _, entry := range a.entries {
		if entry.AvgDamage > max.AvgDamage {
			max = entry
		}
	}
	return max
}

func (a *Analytics) GetTopWinner() Entry {
	max := a.entries[0]
	for _, entry := range a.entries {
		if entry.Battles > 100 && entry.WinRate > max.WinRate {
			max = entry
		}
	}
	return max
}

func (a *Analytics) TotalCount(step, movingAverage uint64) []CountEntry {
	var max uint64 = 0
	for _, entry := range a.entries {
		if max < entry.Rating {
			max = entry.Rating
		}
	}

	total := max/step + 1
	temp := make([]CountEntry, total)
	for _, entry := range a.entries {
		index := entry.Rating / step
		temp[index].Count++
	}

	// moving average
	res := make([]CountEntry, total+movingAverage)
	for dMA := uint64(0); dMA < movingAverage; dMA++ {
		for i := uint64(0); i < total; i++ {
			res[i+dMA].Count += temp[i].Count
		}
	}
	maStart := movingAverage / 2 * step
	absMax := max * 2
	for i, end := uint64(0), total+movingAverage; i < end; i++ {
		t := i*step - maStart
		if t > absMax {
			res[i].Rating = 0
		} else {
			res[i].Rating = t
		}
		// res[i].Count /= movingAverage
	}
	return res
}
