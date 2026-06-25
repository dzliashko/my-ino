package main

import (
	"fmt"
	"time"
)

func main() {
	goFeed := Feed{
		Name:        "Go Blog",
		URL:         "https://go.dev/blog/feed.atom",
		Subscribers: 500,
		Active:      true,
		Articles:    10_000,
	}
	hnFeed := Feed{
		Name:        "Hacker News",
		URL:         "https://hnrss.org/frontpage",
		Subscribers: 50_000,
		Active:      true,
		Articles:    100_000,
	}

	fmt.Print(goFeed.Info())
	fmt.Print(hnFeed.Info())

	fmt.Printf("%s -> %s\n", goFeed.Name, goFeed.Rating())
	fmt.Printf("%s -> %s\n", hnFeed.Name, hnFeed.Rating())

	fmt.Println(goFeed.IsPopular())
	fmt.Println(hnFeed.IsPopular())

	feeds := []Feed{goFeed, hnFeed}

	for _, feed := range feeds {
		fmt.Print(feed.Info())
	}

	fmt.Println(totalSubscribers(feeds))

	popular := popularFeeds(feeds)

	for _, feed := range popular {
		fmt.Println(feed.Name)
	}

	article1 := Article{
		Title:       "Go 1.25 Released",
		Description: "string",
		Link:        "string",
		PublishedAt: time.Now(),
		Feed:        goFeed,
	}
	article2 := Article{
		Title:       "Understanding Goroutines",
		Description: "string",
		Link:        "string",
		PublishedAt: time.Now(),
		Feed:        goFeed,
	}
	articles := []Article{article1, article2}

	printArticles(articles)

	filteredArticles := (articlesByFeed(goFeed, articles))

	for _, article := range filteredArticles {
		fmt.Println(article.ArticleInfo())
	}
}

type Feed struct {
	Name        string
	URL         string
	Subscribers int
	Active      bool
	Articles    int
}

	"my-ino/internal/domain"
	"my-ino/internal/service"
)

type Article struct {
	Title       string
	Description string
	Link        string
	PublishedAt time.Time
	Feed        Feed
}

func printArticles(articles []Article) {
	for _, article := range articles {
		fmt.Println("Feed:", article.Feed.Name)
		fmt.Println("Title:", article.Title)
	}
}

func articlesByFeed(feed Feed, articles []Article) []Article {
	var result []Article
	for _, article := range articles {
		if article.Feed.Name == feed.Name {
			result = append(result, article)
		}
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
