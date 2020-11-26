package main

import (
	"fmt"
	"os"
	"time"
	"wblitz-rating/api"
	"wblitz-rating/utils"
)

func getArg(index int) string {
	if len(os.Args) > index {
		return os.Args[index]
	}
	return ""
}

func main() {
	start := time.Now()
	switch getArg(1) {
	case "c":
		crawl()
		break
	case "a":
		analyze()
		break
	case "t":
		test()
		break
	default:
		help()
	}
	fmt.Println(time.Since(start))
}

func crawl() {
	apiClient := api.New(time.Millisecond*100, 10, getArg(2))
	total := utils.NewRatingCrawler(apiClient, 4000).GetAllRating()
	utils.SaveRating("./rating.csv", total)
	stats := utils.NewStatCrawler(apiClient).GetAllStats(total.GetIds())
	utils.SaveStats("./stats.csv", stats)
}

func analyze() {

}

func test() {
	total := utils.LoadRating("./ratingTest.csv")
	apiClient := api.New(time.Millisecond*100, 10, getArg(2))
	stats := utils.NewStatCrawler(apiClient).GetAllStats(total.GetIds())
	utils.SaveStats("./statsTest.csv", stats)
}

func help() {
	fmt.Println(`Wblitz Rating Utility. Usage:
wblitz-rating [operation] [WarGaming application_id]
Operations:
c - crawl all needed data
a - analyze data`)
}
