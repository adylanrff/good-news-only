package main

import (
	"github.com/adylanrff/good-news-only/internal/app"
	"github.com/adylanrff/good-news-only/pkg/config"
)

func main() {
	config, _ := config.LoadConfig("config/config.yaml")
	app := app.New(config)
	app.Execute()
}
