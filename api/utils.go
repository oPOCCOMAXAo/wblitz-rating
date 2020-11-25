package api

import "strconv"

func u2s(u uint64) string {
	return strconv.FormatUint(u, 10)
}
