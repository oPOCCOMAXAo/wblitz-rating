package main

import (
	"fmt"
	"os"
	"time"
	"wblitz-rating/api"
	"wblitz-rating/utils"
)

func main() {
	var f func()
	if len(os.Args) > 1 {
		switch os.Args[1] {
		case "a":
			f = AnalyzeData
			break
		default:
			f = CrawlData
			break
		}
	} else {
		f = CrawlData
	}
	f()
}

func CrawlData() {
	start := time.Now()
	total := utils.NewCrawler(api.New(time.Millisecond*100, 10), 4000).GetAllRating()
	utils.SaveTable("./rating.csv", total)
	fmt.Println(time.Since(start))
}

func AnalyzeData() {
	start := time.Now()

	fmt.Println(time.Since(start))
}
