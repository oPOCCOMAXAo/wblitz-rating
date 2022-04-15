package main

import (
	"fmt"
	"os"
	"time"
	"wblitz-rating/api"
	"wblitz-rating/crawler"
	"wblitz-rating/data"
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
	case "a":
		analyze()
	default:
		help()
	}
	fmt.Println(time.Since(start))
}

func crawl() {
	apiClient := api.New(api.Config{
		Parallel:      10,
		Timeout:       time.Second * 30,
		ApplicationID: getArg(2),
	})

	storage, err := data.NewSqliteStorage("./data.db")
	if err != nil {
		panic(err)
	}

	/*crawlR, err := crawler.NewRatingCrawler(crawler.Config{
		Storage:   storage,
		API:       apiClient,
		BatchSize: 2000,
	})
	if err != nil {
		panic(err)
	}

	crawlR.Crawl()*/

	crawlP, err := crawler.NewStatCrawler(crawler.Config{
		Storage: storage,
		API:     apiClient,
	})
	if err != nil {
		panic(err)
	}

	crawlP.Crawl()
}

func analyze() {
	analytics := utils.NewAnalytics(
		utils.LoadRating("./rating.csv"),
		utils.LoadStats("./stats.csv"),
	)
	utils.DrawCountEntry("./count", analytics.TotalCount(4, 4))
	utils.DrawCount2dEntry("./winrate", analytics.WinRate(5, 800))
	utils.DrawCount2dEntry("./damage", analytics.Damage(5, 800))
}

func help() {
	fmt.Println(`Wblitz Rating Utility. Usage:
wblitz-rating [operation] [WarGaming application_id]
Operations:
c - crawl all needed data
a - analyze data`)
}
