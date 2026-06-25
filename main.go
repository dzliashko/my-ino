package main

import "fmt"

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

func totalSubscibers(feeds []Feed) int {
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
	var totalSubs int

	for _, feed := range feeds {
		totalSubs += feed.Subscribers
		fmt.Println(feed)
	}
	fmt.Println(totalSubs)

	fmt.Println(totalSubscibers(feeds))
	fmt.Println(popularFeeds(feeds))
}
