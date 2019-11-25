package repository

import (
	"github.com/adylanrff/good-news-only/pkg/common"
	"github.com/adylanrff/good-news-only/pkg/config"
	"github.com/adylanrff/good-news-only/pkg/news"
)

type PostgreNewsRepository struct {
	Repository common.Repository
}

func NewPostgreNewsRepository(config *config.Config) *PostgreNewsRepository {
	dbURI := config.Postgre.URI
	return &PostgreNewsRepository{Repository: common.Repository{DatabaseURI: dbURI}}
}

func Get() ([]news.News, error) {
	return []news.News{}, nil
}

func Save([]news.News) error {
	return nil
}
