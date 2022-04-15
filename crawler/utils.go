package crawler

import (
	"wblitz-rating/api"
	"wblitz-rating/data"
)

type Config struct {
	Storage   data.Storage
	API       *api.API
	BatchSize int64
}
