package utils

import (
	"bufio"
	"io"
	"os"
	"strings"
	"wblitz-rating/api"
)

const statsHead = "ID;Nickname;Battles;Wins;DamageDealt"

func statToCSV(r api.Player) string {
	return strings.Join([]string{
		api.U2S(r.AccountId),
		r.Nickname,
		api.U2S(r.Battles),
		api.U2S(r.Wins),
		api.U2S(r.DamageDealt),
	}, ";")
}

func statFromCSV(line string) api.Player {
	values := strings.Split(line, ";")
	return api.Player{
		AccountId:   api.S2U(values[0]),
		Nickname:    values[1],
		Battles:     api.S2U(values[2]),
		Wins:        api.S2U(values[3]),
		DamageDealt: api.S2U(values[4]),
	}
}

func SaveStats(fname string, data []api.Player) {
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

func LoadStats(fname string) []api.Player {
	var res []api.Player
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
