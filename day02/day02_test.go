package day02

import (
	"testing"
)

func TestIsPopular(t *testing.T) {
	got := IsPopular(100)
	want := false
	if got != want {
		t.Errorf("isPopular(100) = %t want %t", got, want)
	}
	got = IsPopular(1000)
	want = true
	if got != want {
		t.Errorf("isPopular(100) = %t want %t", got, want)
	}
	got = IsPopular(2500)
	want = true
	if got != want {
		t.Errorf("isPopular(100) = %t want %t", got, want)
	}
}

func TestChannelRating(t *testing.T) {
	got := ChannelRating(50)
	want := "New"
	if got != want {
		t.Errorf("ChannelRating(50) = %s want %s", got, want)
	}
	got = ChannelRating(500)
	want = "Growing"
	if got != want {
		t.Errorf("ChannelRating(500) = %s want %s", got, want)
	}
	got = ChannelRating(5_000)
	want = "Popular"
	if got != want {
		t.Errorf("ChannelRating(5_000) = %s want %s", got, want)
	}
	got = ChannelRating(50_000)
	want = "Top"
	if got != want {
		t.Errorf("ChannelRating(5_0000) = %s want %s", got, want)
	}
}

func TestAverageArticles(t *testing.T) {
	got := AverageArticles(60, 30)
	want := 2.0
	if got != want {
		t.Errorf("AverageArticles(60, 30) = %.f want %.f", got, want)
	}
}
