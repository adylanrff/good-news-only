package app

import (
	"log"

	"github.com/adylanrff/good-news-only/pkg/config"
	"github.com/adylanrff/good-news-only/pkg/news"
	"github.com/adylanrff/good-news-only/pkg/twitter"
)

type App struct {
	TwitterClient *twitter.Twitter
	NewsScraper   news.IScrapper
	Config        *config.Config
}

func New(config *config.Config) *App {
	reddit := news.NewRedditScrapper(config)
	log.Println(reddit)
	return &App{NewsScraper: reddit, Config: config}
}

func (t *App) Execute() {
	news, _ := t.NewsScraper.Scrap()
	log.Println(news)
}
