package main

import "fmt"

func main() {
	name := "Go Blog"
	url := "https://go.dev/blog/feed.atom"
	articles := 15
	active := true
	rating := 4.9
	subscribers := 120000
	averagePublicationsPerDay := 2.0

	fmt.Printf("Feed: %s\n", name)
	fmt.Printf("URL: %s\n", url)
	fmt.Printf("Articles: %d\n", articles)
	fmt.Println("Active:", active)
	fmt.Printf("Rating: %.2f\n", rating)
	fmt.Printf("Subscribers: %d\n", subscribers)
	fmt.Printf("Average publications per day: %.2f\n", averagePublicationsPerDay)
}
