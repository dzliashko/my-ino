package main

import (
	"fmt"
	"time"

	"my-ino/internal/domain"
	"my-ino/internal/service"
)

func main() {
	goFeed := domain.Feed{
		Name:        "Go Blog",
		URL:         "https://go.dev/blog/feed.atom",
		Subscribers: 500,
		Active:      true,
		Articles:    10000,
	}

	hnFeed := domain.Feed{
		Name:        "Hacker News",
		URL:         "https://hnrss.org/frontpage",
		Subscribers: 50000,
		Active:      true,
		Articles:    100000,
	}

	feeds := []domain.Feed{
		goFeed,
		hnFeed,
	}

	for _, feed := range feeds {
		fmt.Print(feed.Info())
		fmt.Println(feed.Rating())
		fmt.Println()
	}

	fmt.Println("Total subscribers:", service.TotalSubscribers(feeds))

	fmt.Println()

	popular := service.PopularFeeds(feeds)

	fmt.Println("Popular feeds:")

	for _, feed := range popular {
		fmt.Println("-", feed.Name)
	}
}

	article1 := domain.Article{
		Title:       "Go 1.25 Released",
		Description: "Release notes",
		Link:        "https://go.dev/blog/go1.25",
		PublishedAt: time.Now(),
		Feed:        goFeed,
	}

	article2 := domain.Article{
		Title:       "Understanding Goroutines",
		Description: "Guide",
		Link:        "https://go.dev/blog/goroutines",
		PublishedAt: time.Now(),
		Feed:        goFeed,
	}

	articles := []domain.Article{
		article1,
		article2,
	}

	service.PrintArticles(articles)

	fmt.Println("Articles of Go Blog:")

	for _, article := range service.ArticlesByFeed(goFeed, articles) {
		fmt.Print(article.Info())
	}
	return result
}

func (f Article) ArticleInfo() string {
	return fmt.Sprintf(
		"Title: %s\nDescription: %s\nLink: %s\nPublishedAt: %v\nFeed: %s\n",
		f.Title,
		f.Description,
		f.Link,
		f.PublishedAt,
		f.Feed.Name,
	)
}
