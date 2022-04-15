package utils

import "image/color"

var colorsLeague = []color.RGBA{
	{R: 205, G: 127, B: 50, A: 255},
	{R: 191, G: 191, B: 191, A: 255},
	{R: 212, G: 175, B: 55, A: 255},
	{R: 223, G: 220, B: 217, A: 255},
	{R: 163, G: 67, B: 230, A: 255},
}

func getLeagueColor(rating int64) color.RGBA {
	if rating < 2000 {
		return colorsLeague[0]
	}
	if rating < 3000 {
		return colorsLeague[1]
	}
	if rating < 4000 {
		return colorsLeague[2]
	}
	if rating < 5000 {
		return colorsLeague[3]
	}
	return colorsLeague[4]
}
