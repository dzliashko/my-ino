package day02

func IsPopular(subscribers int) bool {
	return subscribers >= 1000
}

func ChannelRating(subscribers int) string {
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

func AverageArticles(totalArticles int, days int) float64 {
	return float64(totalArticles) / float64(days)
}
