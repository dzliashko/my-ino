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
	return fmt.Sprintf("Feed: %s\nSubscribers: %d\n", feed.Name, feed.Subscribers)
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
	feedGo := Feed{
		Name:        "Go Blog",
		URL:         "https://go.dev/blog/feed.atom",
		Subscribers: 1_500,
		Active:      true,
		Articles:    10_000,
	}
	feedHacker := Feed{
		Name:        "Hacker News",
		URL:         "https://hnrss.org/frontpage",
		Subscribers: 50_000,
		Active:      true,
		Articles:    100_000,
	}

	fmt.Print(feedInfo(feedGo))
	fmt.Print(feedInfo(feedHacker))

	fmt.Print(feedRating(feedGo))
	fmt.Print(feedRating(feedHacker))

	fmt.Printf("Feed: %s Articles: %d\n", feedGo.Name, feedGo.Articles)
	fmt.Printf("Feed: %s Articles: %d\n", feedHacker.Name, feedHacker.Articles)
}
