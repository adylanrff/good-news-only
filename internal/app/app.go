package app

import (
	"github.com/adylanrff/good-news-only/pkg/config"
	"github.com/adylanrff/good-news-only/pkg/usecase"
)

type App struct {
	ScrapNewsUsecase *usecase.ScrapNewsUsecase
	Config           *config.Config
}

func New(config *config.Config) *App {
	scrapNewsUsecase := usecase.NewScrapNewsUsecase(config)
	return &App{ScrapNewsUsecase: scrapNewsUsecase, Config: config}
}

func (t *App) Execute() {
	t.ScrapNewsUsecase.Scrap()
}
