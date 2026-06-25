package service

import (
	"fmt"

	"my-ino/internal/domain"
)

func PrintArticles(articles []domain.Article) {
	for _, article := range articles {
		fmt.Print(article.Info())
		fmt.Println()
	}
}

func ArticlesByFeed(feed domain.Feed, articles []domain.Article) []domain.Article {
	var result []domain.Article

	for _, article := range articles {
		if article.Feed.Name == feed.Name {
			result = append(result, article)
		}
	}

	return result
}
