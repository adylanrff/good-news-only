package news

import (
	"fmt"
	"time"

	"github.com/adylanrff/good-news-only/pkg/config"
	"github.com/gocolly/colly"
)

type RedditScrapper struct {
	Scrapper
	maxPages      int
	scrappedPages int
}

func NewRedditScrapper(config *config.Config) *RedditScrapper {
	var urls []string
	for _, subreddit := range config.Reddit.Subreddits {
		urls = append(urls, fmt.Sprintf("%s%s", config.Reddit.URL, subreddit))
	}

	scrapper := initScrapper(config)
	return &RedditScrapper{
		Scrapper: Scrapper{
			URLs:      urls,
			Collector: scrapper,
		},
		maxPages:      config.Reddit.MaxPages,
		scrappedPages: 0}
}

func initScrapper(config *config.Config) *colly.Collector {
	scrapper := colly.NewCollector(
		colly.AllowedDomains(config.Reddit.Domain),
		colly.Async(true),
	)

	scrapper.Limit(&colly.LimitRule{
		Parallelism: config.Scrapper.Paralellism,
		RandomDelay: time.Duration(config.Scrapper.SecondDelay) * time.Second,
	})

	scrapper.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL.String())
	})

	return scrapper
}

func (t *RedditScrapper) Scrap() ([]News, error) {
	t.parseNews()
	for _, url := range t.URLs {
		t.Collector.Visit(url)
	}
	t.Collector.Wait()
	return t.ScrapedNews, nil
}

func (t *RedditScrapper) parseNews() {
	t.Collector.OnHTML(".top-matter", func(e *colly.HTMLElement) {
		news := News{}
		news.title = e.ChildText("a[data-event-action=title]")
		news.url = e.ChildAttr("a[data-event-action=title]", "href")
		news.createdDate = time.Now()
		t.ScrapedNews = append(t.ScrapedNews, news)
	})

	t.Collector.OnHTML("span.next-button", func(h *colly.HTMLElement) {
		next := h.ChildAttr("a", "href")
		if t.scrappedPages < t.maxPages {
			t.Collector.Visit(next)
			t.scrappedPages++
		}
	})
}
