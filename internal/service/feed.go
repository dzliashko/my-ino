package service

import "my-ino/internal/domain"

func TotalSubscribers(feeds []domain.Feed) int {
	total := 0

	for _, feed := range feeds {
		total += feed.Subscribers
	}

	return total
}

func PopularFeeds(feeds []domain.Feed) []domain.Feed {
	var result []domain.Feed

	for _, feed := range feeds {
		if feed.IsPopular() {
			result = append(result, feed)
		}
	}

	return result
}
