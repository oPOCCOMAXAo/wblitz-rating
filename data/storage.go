package data

import (
	"io"
)

type BatchConfig struct {
	Limit  int
	Offset int
}

type Storage interface {
	io.Closer
	SavePlayers([]Player) error
	LoadPlayers(*BatchConfig) ([]Player, error)
	SaveRating([]Rating) error
	LoadRating(*BatchConfig) ([]Rating, error)
}
