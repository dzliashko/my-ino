package main

import "fmt"

type Feed struct {
	Name        string
	URL         string
	Subscribers int
	Active      bool
	Articles    int
}

func feedInfo(feed Feed) string {
	return fmt.Sprintf(
		"Feed: %s\nURL: %s\nSubscribers: %d\nArticles: %d\nActive: %t\n",
		feed.Name,
		feed.URL,
		feed.Subscribers,
		feed.Articles,
		feed.Active,
	)
}

func feedRating(feed Feed) string {
	return fmt.Sprintf("%s -> %s\n", feed.Name, channelRating(feed.Subscribers))
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

func main() {
	goFeed := Feed{
		Name:        "Go Blog",
		URL:         "https://go.dev/blog/feed.atom",
		Subscribers: 1_500,
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

	fmt.Print(feedInfo(goFeed))
	fmt.Print(feedInfo(hnFeed))

	fmt.Print(feedRating(goFeed))
	fmt.Print(feedRating(hnFeed))
}
