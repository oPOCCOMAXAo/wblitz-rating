package main

import (
	"fmt"
	"wblitz-rating/api"
	"wblitz-rating/utils"
)

func GetAll() {
	const batchSize = 2500
	var batches [][]api.Rating
	var current, last api.Rating
	client := api.New()
	current = client.GetRatingTop1()
	for current.SpaId != last.SpaId {
		last = current
		temp := client.RatingNeighbors(current.SpaId, batchSize)
		batches = append(batches, temp)
		current = temp[len(temp)-1]
		fmt.Println(current)
	}
	total := make([]api.Rating, current.Number)
	for _, batch := range batches {
		for _, rating := range batch {
			total[rating.Number-1] = rating
		}
	}
	utils.SaveTable("../rating.csv", total)
}
