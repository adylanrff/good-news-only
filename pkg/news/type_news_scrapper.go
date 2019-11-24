package news

import (
	"time"

	"github.com/gocolly/colly"
)

type News struct {
	title       string
	url         string
	createdDate time.Time
}

type Scrapper struct {
	URLs        []string
	Collector   *colly.Collector
	ScrapedNews []News
}

type IScrapper interface {
	Scrap() ([]News, error)
}
