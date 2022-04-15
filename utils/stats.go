package utils

import (
	"bufio"
	"io"
	"os"
	"strings"
	"wblitz-rating/api"
	"wblitz-rating/data"
)

const statsHead = "ID;Nickname;Battles;Wins;DamageDealt"

func statToCSV(r data.Player) string {
	return strings.Join([]string{
		api.I2S(r.ID),
		r.Nickname,
		api.I2S(r.Battles),
		api.I2S(r.Wins),
		api.I2S(r.Damage),
	}, ";")
}

func statFromCSV(line string) data.Player {
	values := strings.Split(line, ";")
	return data.Player{
		ID:       api.S2I(values[0]),
		Nickname: values[1],
		Battles:  api.S2I(values[2]),
		Wins:     api.S2I(values[3]),
		Damage:   api.S2I(values[4]),
	}
}

func SaveStats(fname string, data []data.Player) {
	f, err := os.Create(fname)
	panicOnNonNil(err)
	defer f.Close()
	w := bufio.NewWriter(f)
	_, err = w.WriteString(statsHead)
	panicOnNonNil(err)
	for _, stat := range data {
		_, err = w.WriteString("\n" + statToCSV(stat))
		panicOnNonNil(err)
	}
	_ = w.Flush()
}

func LoadStats(fname string) []data.Player {
	var res []data.Player
	f, err := os.Open(fname)
	panicOnNonNil(err)
	defer f.Close()
	w := bufio.NewReader(f)
	t, _, err := w.ReadLine()
	panicOnNonNil(err)
	if string(t) != statsHead {
		panic("unknown head")
	}
	for {
		line, _, err := w.ReadLine()
		if err == io.EOF {
			break
		}
		panicOnNonNil(err)
		res = append(res, statFromCSV(string(line)))
	}
	return res
}
