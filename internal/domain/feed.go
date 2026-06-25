package domain

import "fmt"

type Feed struct {
	Name        string
	URL         string
	Subscribers int
	Active      bool
	Articles    int
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
	switch {
	case f.Subscribers < 100:
		return "New"
	case f.Subscribers < 1000:
		return "Growing"
	case f.Subscribers < 10000:
		return "Popular"
	default:
		return "Top"
	}
}

func (f Feed) IsPopular() bool {
	return f.Subscribers >= 1000
}
