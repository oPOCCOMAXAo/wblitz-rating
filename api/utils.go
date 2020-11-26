package api

import "strconv"

func U2S(u uint64) string {
	return strconv.FormatUint(u, 10)
}

func F2S(f float64) string {
	return strconv.FormatFloat(f, 'f', -1, 64)
}

func S2U(s string) uint64 {
	res, err := strconv.ParseUint(s, 10, 64)
	if err != nil {
		return 0
	}
	return res
}

func S2F(s string) float64 {
	res, err := strconv.ParseFloat(s, 64)
	if err != nil {
		return 0
	}
	return res
}

func U2SArr(u []uint64) []string {
	l := len(u)
	res := make([]string, l)
	for i := 0; i < l; i++ {
		res[i] = U2S(u[i])
	}
	return res
}
