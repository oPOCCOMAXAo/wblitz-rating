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
		case "c":
			f = CrawlData
			break
		case "a":
			f = AnalyzeData
			break
		}
	}
	if f != nil {
		f()
	} else {
		Help()
	}
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

func Help() {
	fmt.Println(`Wblitz Rating Utility. Usage:
wblitz-rating [argument]
Arguments:
c - crawl all needed data
a - analyze data`)
}
