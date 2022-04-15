package utils

import (
	"fmt"
	"math"
	"wblitz-rating/data"
)

type Entry struct {
	Id        int64
	Rating    int64
	Battles   int64
	WinRate   float64
	AvgDamage float64
}

type Analytics struct {
	entries []Entry
}

func NewAnalytics(rating []data.Rating, stats []data.Player) *Analytics {
	total := len(rating)
	if t := len(stats); t != total {
		panic("data broken, reloading needed")
	}
	entries := make([]Entry, total)
	for i := 0; i < total; i++ {
		battles := float64(stats[i].Battles) + 1e-50
		entries[i] = Entry{
			Id:        rating[i].PlayerID,
			Rating:    rating[i].Score,
			Battles:   stats[i].Battles,
			WinRate:   float64(stats[i].Wins) / battles,
			AvgDamage: float64(stats[i].Damage) / battles,
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

func (a *Analytics) GetMaxRating() int64 {
	var res int64 = 0
	for _, entry := range a.entries {
		if res < entry.Rating {
			res = entry.Rating
		}
	}
	return res
}

func (a *Analytics) GetMaxDamage() float64 {
	var res float64 = 0
	for _, entry := range a.entries {
		if res < entry.AvgDamage {
			res = entry.AvgDamage
		}
	}
	return res
}

func (a *Analytics) TotalCount(step, movingAverage int64) []CountEntry {
	max := a.GetMaxRating()
	total := max/step + 1
	temp := make([]CountEntry, total)
	for _, entry := range a.entries {
		index := entry.Rating / step
		temp[index].Count++
	}

	// moving average
	res := make([]CountEntry, total+movingAverage)
	for dMA := int64(0); dMA < movingAverage; dMA++ {
		for i := int64(0); i < total; i++ {
			res[i+dMA].Count += temp[i].Count
		}
	}
	maStart := movingAverage / 2 * step
	absMax := max * 2
	for i, end := int64(0), total+movingAverage; i < end; i++ {
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

func (a *Analytics) WinRate(stepRating int64, totalWinRate int64) []Count2dEntry {
	max := a.GetMaxRating()
	totalRating := max/stepRating + 1
	stepWinRate := 1.0 / float64(totalWinRate)
	res := make([]Count2dEntry, totalRating)
	for i := int64(0); i < totalRating; i++ {
		res[i].Rating = i * stepRating
		res[i].Count = make([]int64, totalWinRate+1)
	}
	for _, entry := range a.entries {
		iRating := entry.Rating / stepRating
		iWinRate := int(math.Floor(entry.WinRate / stepWinRate))
		res[iRating].Count[iWinRate]++
	}
	return res
}

func (a *Analytics) Damage(stepRating int64, totalDamage int64) []Count2dEntry {
	max := a.GetMaxRating()
	maxDamage := a.GetMaxDamage()
	fmt.Printf("Max Damage: %.0f\n", maxDamage)
	totalRating := max/stepRating + 1
	stepDamage := maxDamage / float64(totalDamage)
	res := make([]Count2dEntry, totalRating)
	for i := int64(0); i < totalRating; i++ {
		res[i].Rating = i * stepRating
		res[i].Count = make([]int64, totalDamage+1)
	}
	for _, entry := range a.entries {
		iRating := entry.Rating / stepRating
		i2 := int(math.Floor(entry.AvgDamage / stepDamage))
		res[iRating].Count[i2]++
	}
	return res
}
