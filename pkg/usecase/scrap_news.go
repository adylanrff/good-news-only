package usecase

import (
	"log"

	"github.com/adylanrff/good-news-only/pkg/config"
	"github.com/adylanrff/good-news-only/pkg/news"
)

type ScrapNewsUsecase struct {
	NewsScraper news.IScrapper
}

func NewScrapNewsUsecase(config *config.Config) *ScrapNewsUsecase {
	reddit := news.NewRedditScrapper(config)
	return &ScrapNewsUsecase{NewsScraper: reddit}
}

func (t *ScrapNewsUsecase) Scrap() {
	news, err := t.NewsScraper.Scrap()
	if err != nil {
		log.Fatal(err)
	}
	log.Println(news)
}
