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
		feed.Info()
	}

	fmt.Println(totalSubscribers(feeds))

	popular := popularFeeds(feeds)

	for _, feed := range popular {
		fmt.Println(feed.Name)
	}

	arcticle1 := Article{
		Title:       "Go 1.25 Released",
		Description: "string",
		Link:        "string",
		PublishedAt: time.Now(),
		Feed:        goFeed,
	}
	arcticle2 := Article{
		Title:       "Understanding Goroutines",
		Description: "string",
		Link:        "string",
		PublishedAt: time.Now(),
		Feed:        goFeed,
	}
	articles := []Article{arcticle1, arcticle2}

	printArticles(articles)

	articlesByFeed := (articlesByFeed(goFeed, articles))

	for _, arcticle := range articlesByFeed {
		fmt.Println(arcticle.ArticleInfo())
	}
}

type Feed struct {
	Name        string
	URL         string
	Subscribers int
	Active      bool
	Articles    int
}

func channelRating(subscribers int) string {
	switch {
	case subscribers < 100:
		return "New"
	case subscribers < 1_000:
		return "Growing"
	case subscribers < 10_000:
		return "Popular"
	default:
		return "Top"
	}
}

func (f Feed) Info() string {
	return fmt.Sprintf(
		"Feed: %s\nURL: %s\nSubscribers: %d\nArticles: %d\nActive: %t\n",
		f.Name,
		f.URL,
		f.Subscribers,
		f.Articles,
		f.Active,
	)
}

func (f Feed) Rating() string {
	return channelRating(f.Subscribers)
}

func (f Feed) IsPopular() bool {
	return f.Subscribers >= 1000
}

func totalSubscribers(feeds []Feed) int {
	var totalSubscribers int
	for _, feed := range feeds {
		totalSubscribers += feed.Subscribers
	}
	return totalSubscribers
}

func popularFeeds(feeds []Feed) []Feed {
	var result []Feed

	for _, feed := range feeds {
		if feed.IsPopular() {
			result = append(result, feed)
		}
	}
	return result
}

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
		"Title: %s\nDescription: %s\nLink: %d\nPublishedAt: %v\nFeed: %v\n",
		f.Title,
		f.Description,
		f.Link,
		f.PublishedAt,
		f.Feed,
	)
}
