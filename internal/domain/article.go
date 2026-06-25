package domain

import (
	"fmt"
	"time"
)

type Article struct {
	Title       string
	Description string
	Link        string
	PublishedAt time.Time
	Feed        Feed
}

func (a Article) Info() string {
	return fmt.Sprintf(
		"Title: %s\nDescription: %s\nLink: %s\nPublishedAt: %v\nFeed: %s\n",
		a.Title,
		a.Description,
		a.Link,
		a.PublishedAt,
		a.Feed.Name,
	)
}
