package api

import (
	"fmt"
	"testing"
)

func TestAPI_RatingTop(t *testing.T) {
	a := New()
	res := a.RatingNeighbors(75340817, 10000)
	fmt.Println(len(res))
}
