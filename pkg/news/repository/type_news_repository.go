package repository

import "github.com/adylanrff/good-news-only/pkg/news"

type NewsRepository interface {
	Get() ([]news.News, error)
	Save([]news.News) error
}
